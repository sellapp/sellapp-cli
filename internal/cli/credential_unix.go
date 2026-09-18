//go:build !windows

package cli

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func platformPathSafe(_ string, info os.FileInfo) error {
	if info.Mode()&os.ModeSymlink != 0 {
		return errors.New("credential paths must not contain symbolic links")
	}
	if info.IsDir() && info.Mode().Perm()&0022 != 0 {
		stat, ok := info.Sys().(*syscall.Stat_t)
		if !ok || info.Mode()&os.ModeSticky == 0 || stat.Uid != 0 {
			return errors.New("credential path parent is writable by another user")
		}
	}
	return nil
}
func platformPrivateCheck(_ string, info os.FileInfo) error {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || stat.Uid != uint32(os.Geteuid()) || info.Mode().Perm()&0077 != 0 {
		return errors.New("credential path must be owned by the current user with owner-only permissions (0700 directories, 0600 files)")
	}
	if !info.IsDir() && (!info.Mode().IsRegular() || stat.Nlink != 1) {
		return errors.New("credential files must be regular files without hard links")
	}
	return nil
}
func platformLockTransient(error) bool       { return false }
func platformPrivateMkdir(path string) error { return os.Mkdir(path, 0700) }
func platformPrivateCreate(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
}
func platformReplace(from, to string) error { return os.Rename(from, to) }
func platformSyncDirectory(path string) error {
	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	defer dir.Close()
	return dir.Sync()
}
func platformMachineID() (string, error) {
	if runtime.GOOS == "darwin" {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		data, err := exec.CommandContext(ctx, "/usr/sbin/ioreg", "-rd1", "-c", "IOPlatformExpertDevice").Output()
		if err != nil {
			return "", errors.New("platform machine identifier unavailable")
		}
		for _, line := range strings.Split(string(data), "\n") {
			key, value, ok := strings.Cut(line, "=")
			if ok && strings.TrimSpace(key) == "\"IOPlatformUUID\"" {
				value = strings.Trim(strings.TrimSpace(value), "\"")
				if validPlatformIdentifier(value, true) {
					return "darwin:" + value, nil
				}
			}
		}
	} else {
		for _, path := range []string{"/etc/machine-id", "/var/lib/dbus/machine-id"} {
			if data, err := os.ReadFile(path); err == nil {
				value := strings.TrimSpace(string(data))
				if validPlatformIdentifier(value, false) {
					return "linux:" + value, nil
				}
			}
		}
	}
	return "", errors.New("platform machine identifier unavailable")
}
