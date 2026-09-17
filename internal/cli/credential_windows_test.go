//go:build windows

package cli

import (
	"errors"
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"unsafe"
)

func credentialTestSetACL(path, sddl string) error {
	value, err := syscall.UTF16PtrFromString(sddl)
	if err != nil {
		return err
	}
	var descriptor uintptr
	ok, _, callErr := credentialConvertSD.Call(uintptr(unsafe.Pointer(value)), 1, uintptr(unsafe.Pointer(&descriptor)), 0)
	if ok == 0 {
		return callErr
	}
	defer credentialLocalFree.Call(descriptor)
	name, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	set := credentialAdvapi.NewProc("SetFileSecurityW")
	ok, _, callErr = set.Call(uintptr(unsafe.Pointer(name)), 0x80000004, descriptor)
	if ok == 0 {
		return callErr
	}
	return nil
}
func TestCredentialWindowsCurrentUserACLAndReplacement(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	record := credentialTestRecord("windows-entry")
	for _, secret := range []string{"first-secret", "second-secret"} {
		if err := credentialWrite(record, secret); err != nil {
			t.Fatal(err)
		}
		if value, err := credentialRead(record); err != nil || value != secret {
			t.Fatal(err)
		}
	}
	path, _ := credentialFile(record.VaultEntry)
	for _, target := range []string{path, filepath.Dir(path)} {
		info, err := os.Lstat(target)
		if err != nil {
			t.Fatal(err)
		}
		if err = platformPrivateCheck(target, info); err != nil {
			t.Fatal(err)
		}
	}
}
func TestCredentialWindowsRejectsAdditionalPrincipal(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	record := credentialTestRecord("acl")
	if err := credentialWrite(record, "fake-secret"); err != nil {
		t.Fatal(err)
	}
	path, _ := credentialFile(record.VaultEntry)
	sid, err := credentialCurrentSID()
	if err != nil {
		t.Fatal(err)
	}
	if err = credentialTestSetACL(path, "D:P(A;;FA;;;"+sid+")(A;;FR;;;WD)"); err != nil {
		t.Fatal(err)
	}
	if _, err = credentialRead(record); err == nil {
		t.Fatal("additional principal accepted")
	}
	if err = credentialWrite(record, "replacement"); err == nil {
		t.Fatal("unsafe ACL overwritten")
	}
}
func TestCredentialWindowsRejectsHardLinks(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	record := credentialTestRecord("hardlink")
	if err := credentialWrite(record, "fake-secret"); err != nil {
		t.Fatal(err)
	}
	path, _ := credentialFile(record.VaultEntry)
	if err := os.Link(path, path+".link"); err != nil {
		t.Fatal(err)
	}
	if _, err := credentialRead(record); err == nil {
		t.Fatal("hard link read accepted")
	}
	if err := credentialWrite(record, "replacement"); err == nil {
		t.Fatal("hard link replacement accepted")
	}
}
func TestCredentialWindowsRejectsReparsePoints(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	record := credentialTestRecord("reparse")
	if err := credentialWrite(record, "fake-secret"); err != nil {
		t.Fatal(err)
	}
	path, _ := credentialFile(record.VaultEntry)
	source := path + ".source"
	if err := os.Rename(path, source); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(source, path); err != nil {
		var system syscall.Errno
		if errors.As(err, &system) && system == 1314 {
			t.Skip("native symlink creation privilege is unavailable")
		}
		t.Fatal(err)
	}
	if _, err := credentialRead(record); err == nil {
		t.Fatal("reparse point accepted")
	}
}
