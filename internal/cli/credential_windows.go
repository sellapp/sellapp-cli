//go:build windows

package cli

import (
	"errors"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

var credentialAdvapi = syscall.NewLazyDLL("advapi32.dll")
var credentialKernel = syscall.NewLazyDLL("kernel32.dll")
var credentialConvertSD = credentialAdvapi.NewProc("ConvertStringSecurityDescriptorToSecurityDescriptorW")
var credentialGetSecurity = credentialAdvapi.NewProc("GetNamedSecurityInfoW")
var credentialGetControl = credentialAdvapi.NewProc("GetSecurityDescriptorControl")
var credentialGetACE = credentialAdvapi.NewProc("GetAce")
var credentialMoveFile = credentialKernel.NewProc("MoveFileExW")
var credentialLocalFree = credentialKernel.NewProc("LocalFree")

func credentialCurrentSID() (string, error) {
	token, err := syscall.OpenCurrentProcessToken()
	if err != nil {
		return "", err
	}
	defer token.Close()
	user, err := token.GetTokenUser()
	if err != nil {
		return "", err
	}
	return user.User.Sid.String()
}
func credentialSecurityAttributes(directory bool) (*syscall.SecurityAttributes, func(), error) {
	sid, err := credentialCurrentSID()
	if err != nil {
		return nil, nil, err
	}
	flags := ""
	if directory {
		flags = "OICI"
	}
	text, err := syscall.UTF16PtrFromString("O:" + sid + "D:P(A;" + flags + ";FA;;;" + sid + ")")
	if err != nil {
		return nil, nil, err
	}
	var descriptor uintptr
	success, _, callErr := credentialConvertSD.Call(uintptr(unsafe.Pointer(text)), 1, uintptr(unsafe.Pointer(&descriptor)), 0)
	if success == 0 {
		return nil, nil, callErr
	}
	return &syscall.SecurityAttributes{Length: uint32(unsafe.Sizeof(syscall.SecurityAttributes{})), SecurityDescriptor: descriptor}, func() { _, _, _ = credentialLocalFree.Call(descriptor) }, nil
}
func platformPathSafe(path string, info os.FileInfo) error {
	if info.Mode()&os.ModeSymlink != 0 {
		return errors.New("credential paths must not contain links or junctions")
	}
	p, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	attributes, err := syscall.GetFileAttributes(p)
	if err != nil {
		return err
	}
	if attributes&syscall.FILE_ATTRIBUTE_REPARSE_POINT != 0 {
		return errors.New("credential paths must not contain reparse points")
	}
	return nil
}

type credentialACL struct {
	Revision byte
	Reserved byte
	Size     uint16
	Count    uint16
	Unused   uint16
}
type credentialACE struct {
	Type  byte
	Flags byte
	Size  uint16
	Mask  uint32
	SID   uint32
}

func platformPrivateCheck(path string, info os.FileInfo) error {
	if err := platformPathSafe(path, info); err != nil {
		return err
	}
	sid, err := credentialCurrentSID()
	if err != nil {
		return err
	}
	p, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	var owner *syscall.SID
	var dacl *credentialACL
	var descriptor uintptr
	status, _, _ := credentialGetSecurity.Call(uintptr(unsafe.Pointer(p)), 1, 5, uintptr(unsafe.Pointer(&owner)), 0, uintptr(unsafe.Pointer(&dacl)), 0, uintptr(unsafe.Pointer(&descriptor)))
	if status != 0 {
		return syscall.Errno(status)
	}
	defer credentialLocalFree.Call(descriptor)
	if owner == nil || dacl == nil {
		return errors.New("credential security descriptor has no owner or DACL")
	}
	ownerSID, err := owner.String()
	if err != nil || ownerSID != sid || dacl == nil {
		return errors.New("credential ACL must be owned by and restricted to the current user")
	}
	var control uint16
	var revision uint32
	ok, _, callErr := credentialGetControl.Call(descriptor, uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&revision)))
	if ok == 0 {
		return callErr
	}
	if control&0x1000 == 0 {
		return errors.New("credential ACL must disable inherited permissions")
	}
	if dacl.Count == 0 {
		return errors.New("credential ACL does not grant access to the current user")
	}
	for i := uint16(0); i < dacl.Count; i++ {
		var ace *credentialACE
		ok, _, callErr = credentialGetACE.Call(uintptr(unsafe.Pointer(dacl)), uintptr(i), uintptr(unsafe.Pointer(&ace)))
		if ok == 0 {
			return callErr
		}
		if ace.Type != 0 || ace.Size < 12 {
			return errors.New("credential ACL contains unsupported access rules")
		}
		aceSID, err := (*syscall.SID)(unsafe.Pointer(&ace.SID)).String()
		if err != nil || aceSID != sid || ace.Mask&0x1f01ff != 0x1f01ff {
			return errors.New("credential ACL must grant full control only to the current user")
		}
	}
	if !info.IsDir() {
		handle, err := syscall.CreateFile(p, 0x80, syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE|syscall.FILE_SHARE_DELETE, nil, syscall.OPEN_EXISTING, syscall.FILE_FLAG_OPEN_REPARSE_POINT, 0)
		if err != nil {
			return err
		}
		defer syscall.CloseHandle(handle)
		var details syscall.ByHandleFileInformation
		if err = syscall.GetFileInformationByHandle(handle, &details); err != nil {
			return err
		}
		if details.NumberOfLinks != 1 {
			return errors.New("credential files must not have hard links")
		}
	}
	return nil
}
func platformLockTransient(err error) bool {
	return errors.Is(err, syscall.ERROR_ACCESS_DENIED) || errors.Is(err, syscall.Errno(32))
}
func platformPrivateMkdir(path string) error {
	p, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	attributes, free, err := credentialSecurityAttributes(true)
	if err != nil {
		return err
	}
	defer free()
	return syscall.CreateDirectory(p, attributes)
}
func platformPrivateCreate(path string) (*os.File, error) {
	p, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}
	attributes, free, err := credentialSecurityAttributes(false)
	if err != nil {
		return nil, err
	}
	defer free()
	handle, err := syscall.CreateFile(p, syscall.GENERIC_WRITE, 0, attributes, syscall.CREATE_NEW, syscall.FILE_ATTRIBUTE_NORMAL, 0)
	if err != nil {
		return nil, err
	}
	return os.NewFile(uintptr(handle), path), nil
}
func platformReplace(from, to string) error {
	source, err := syscall.UTF16PtrFromString(from)
	if err != nil {
		return err
	}
	destination, err := syscall.UTF16PtrFromString(to)
	if err != nil {
		return err
	}
	ok, _, callErr := credentialMoveFile.Call(uintptr(unsafe.Pointer(source)), uintptr(unsafe.Pointer(destination)), 0x1|0x8)
	if ok == 0 {
		return callErr
	}
	return nil
}

// MoveFileEx uses MOVEFILE_WRITE_THROUGH and private file writes use FlushFileBuffers.
func platformSyncDirectory(string) error { return nil }
func platformMachineID() (string, error) {
	path, err := syscall.UTF16PtrFromString("SOFTWARE\\Microsoft\\Cryptography")
	if err != nil {
		return "", err
	}
	var key syscall.Handle
	if err = syscall.RegOpenKeyEx(syscall.HKEY_LOCAL_MACHINE, path, 0, syscall.KEY_READ|syscall.KEY_WOW64_64KEY, &key); err != nil {
		return "", err
	}
	defer syscall.RegCloseKey(key)
	name, _ := syscall.UTF16PtrFromString("MachineGuid")
	var kind uint32
	var size uint32
	if err = syscall.RegQueryValueEx(key, name, nil, &kind, nil, &size); err != nil {
		return "", err
	}
	if kind != syscall.REG_SZ || size < 2 || size > 4096 {
		return "", errors.New("platform machine identifier unavailable")
	}
	buffer := make([]uint16, (size+1)/2)
	if err = syscall.RegQueryValueEx(key, name, nil, &kind, (*byte)(unsafe.Pointer(&buffer[0])), &size); err != nil {
		return "", err
	}
	value := strings.TrimSpace(syscall.UTF16ToString(buffer))
	if !validPlatformIdentifier(value, true) {
		return "", errors.New("platform machine identifier unavailable")
	}
	return "windows:" + value, nil
}
