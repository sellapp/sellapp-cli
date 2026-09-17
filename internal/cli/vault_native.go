package cli

import (
	"errors"
	"fmt"
	"sync"
)

type keychainBackend struct {
	find    func(string, string) (uintptr, int32)
	add     func(string, string, []byte) int32
	update  func(uintptr, []byte) int32
	read    func(uintptr) ([]byte, int32)
	delete  func(uintptr) int32
	release func(uintptr)
}

var keychainOnce sync.Once
var keychainAPI *keychainBackend
var keychainLoadError error

func nativeVaultAvailable() bool { _, err := nativeKeychain(); return err == nil }
func nativeKeychain() (*keychainBackend, error) {
	keychainOnce.Do(func() { keychainAPI, keychainLoadError = loadNativeKeychain() })
	return keychainAPI, keychainLoadError
}
func nativeVaultWrite(entry, secret string) error {
	api, err := nativeKeychain()
	if err != nil {
		return errors.New("native credential vault unavailable")
	}
	return writeKeychainItem(api, vaultService, entry, []byte(secret))
}
func nativeVaultRead(entry string) (string, error) {
	api, err := nativeKeychain()
	if err != nil {
		return "", errors.New("native credential vault unavailable")
	}
	return readKeychainItem(api, vaultService, entry)
}
func nativeVaultDelete(entry string) error {
	api, err := nativeKeychain()
	if err != nil {
		return errors.New("native credential vault unavailable")
	}
	return deleteKeychainItem(api, vaultService, entry)
}
func readKeychainItem(api *keychainBackend, service, entry string) (string, error) {
	item, status := api.find(service, entry)
	if status != 0 {
		return "", fmt.Errorf("credential vault entry unavailable (OSStatus %d)", status)
	}
	defer api.release(item)
	data, status := api.read(item)
	defer func() {
		for i := range data {
			data[i] = 0
		}
	}()
	if status != 0 {
		return "", fmt.Errorf("credential vault read failed (OSStatus %d)", status)
	}
	if len(data) == 0 {
		return "", errors.New("credential vault returned an empty key")
	}
	return string(data), nil
}
func deleteKeychainItem(api *keychainBackend, service, entry string) error {
	item, status := api.find(service, entry)
	if status != 0 {
		return fmt.Errorf("credential vault entry unavailable (OSStatus %d)", status)
	}
	defer api.release(item)
	if status = api.delete(item); status != 0 {
		return fmt.Errorf("credential vault delete failed (OSStatus %d)", status)
	}
	return nil
}

// Replace the complete pair atomically; never delete the previous item first.
func writeKeychainItem(api *keychainBackend, service, entry string, secret []byte) error {
	defer func() {
		for i := range secret {
			secret[i] = 0
		}
	}()
	item, status := api.find(service, entry)
	if status == 0 {
		defer api.release(item)
		status = api.update(item, secret)
	} else if status == -25300 {
		status = api.add(service, entry, secret)
		if status == -25299 {
			item, status = api.find(service, entry)
			if status == 0 {
				defer api.release(item)
				status = api.update(item, secret)
			}
		}
	}
	if status != 0 {
		return fmt.Errorf("credential vault write failed (OSStatus %d)", status)
	}
	return nil
}
