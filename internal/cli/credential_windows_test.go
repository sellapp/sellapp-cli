//go:build windows

package cli

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

func TestCredentialWindowsLockRetriesTurnoverFailures(t *testing.T) {
	for _, stage := range []string{"create", "inspect"} {
		for _, failure := range []error{syscall.ERROR_ACCESS_DENIED, syscall.Errno(32)} {
			t.Run(stage+failure.Error(), func(t *testing.T) {
				credentialTestHome(t)
				dir, err := credentialDir()
				if err != nil {
					t.Fatal(err)
				}
				path := filepath.Join(dir, "turnover.lock")
				oldMake, oldCheck := credentialLockMkdir, credentialLockCheck
				t.Cleanup(func() { credentialLockMkdir = oldMake; credentialLockCheck = oldCheck })
				failed := false
				if stage == "inspect" {
					if err = platformPrivateMkdir(path); err != nil {
						t.Fatal(err)
					}
				}
				credentialLockMkdir = func(name string) error {
					if stage == "create" && !failed {
						failed = true
						return failure
					}
					if stage == "inspect" && failed {
						if err := os.Remove(name); err != nil {
							return err
						}
					}
					return platformPrivateMkdir(name)
				}
				credentialLockCheck = func(name string, create bool) error {
					if stage == "inspect" && !failed {
						failed = true
						return failure
					}
					return privateDirectory(name, create)
				}
				unlock, err := lockCredentialPath(path, true)
				if err != nil {
					t.Fatal(err)
				}
				defer unlock()
				if !failed {
					t.Fatal("turnover failure was not exercised")
				}
				if err = privateDirectory(path, false); err != nil {
					t.Fatal("acquired lock is not private", err)
				}
			})
		}
	}
}
func TestCredentialWindowsLockDenialRemainsBounded(t *testing.T) {
	credentialTestHome(t)
	dir, err := credentialDir()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "denied.lock")
	old := credentialLockMkdir
	t.Cleanup(func() { credentialLockMkdir = old })
	credentialLockMkdir = func(string) error { return syscall.ERROR_ACCESS_DENIED }
	for _, wait := range []bool{false, true} {
		started := time.Now()
		unlock, err := lockCredentialPath(path, wait)
		elapsed := time.Since(started)
		if unlock != nil || !errors.Is(err, syscall.ERROR_ACCESS_DENIED) {
			t.Fatal("permission denial lost or lock granted", err)
		}
		if !wait && elapsed > time.Second {
			t.Fatal("non-waiting acquisition blocked")
		}
		if wait && (elapsed < 5*time.Second || elapsed > 10*time.Second) {
			t.Fatal("lock retry exceeded its bound", elapsed)
		}
	}
	if _, err = os.Stat(path); !os.IsNotExist(err) {
		t.Fatal("denied acquisition created a lock", err)
	}
}
func TestCredentialWindowsLockRetainsUnsafeExistingDirectory(t *testing.T) {
	credentialTestHome(t)
	dir, err := credentialDir()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "unsafe.lock")
	if err = platformPrivateMkdir(path); err != nil {
		t.Fatal(err)
	}
	sid, err := credentialCurrentSID()
	if err != nil {
		t.Fatal(err)
	}
	if err = credentialTestSetACL(path, "D:P(A;;FA;;;"+sid+")(A;;FR;;;WD)"); err != nil {
		t.Fatal(err)
	}
	unlock, err := lockCredentialPath(path, true)
	if err == nil || unlock != nil || !strings.Contains(err.Error(), "credential ACL") {
		t.Fatal("unsafe lock was not rejected", err)
	}
	if _, err = os.Stat(path); err != nil {
		t.Fatal("existing lock was removed", err)
	}
}
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
