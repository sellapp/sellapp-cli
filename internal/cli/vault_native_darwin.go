//go:build darwin

package cli

import (
	"errors"
	"github.com/ebitengine/purego"
	"runtime"
	"unsafe"
)

func loadNativeKeychain() (*keychainBackend, error) {
	security, err := purego.Dlopen("/System/Library/Frameworks/Security.framework/Security", purego.RTLD_NOW|purego.RTLD_LOCAL)
	if err != nil {
		return nil, errors.New("Security framework unavailable")
	}
	core, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_NOW|purego.RTLD_LOCAL)
	if err != nil {
		return nil, errors.New("CoreFoundation framework unavailable")
	}
	var interaction func(bool) int32
	var find func(uintptr, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uintptr, uintptr, *uintptr) int32
	var add func(uintptr, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uint32, unsafe.Pointer, uintptr) int32
	var update func(uintptr, uintptr, uint32, unsafe.Pointer) int32
	var copyData func(uintptr, uintptr, uintptr, uintptr, *uint32, *unsafe.Pointer) int32
	var freeData func(uintptr, unsafe.Pointer) int32
	var deleteItem func(uintptr) int32
	var release func(uintptr)
	bindings := []struct {
		lib    uintptr
		name   string
		target any
	}{
		{security, "SecKeychainSetUserInteractionAllowed", &interaction},
		{security, "SecKeychainFindGenericPassword", &find},
		{security, "SecKeychainAddGenericPassword", &add},
		{security, "SecKeychainItemModifyAttributesAndData", &update},
		{security, "SecKeychainItemCopyAttributesAndData", &copyData},
		{security, "SecKeychainItemFreeAttributesAndData", &freeData},
		{security, "SecKeychainItemDelete", &deleteItem},
		{core, "CFRelease", &release},
	}
	for _, binding := range bindings {
		symbol, err := purego.Dlsym(binding.lib, binding.name)
		if err != nil {
			return nil, errors.New("native credential vault API unavailable")
		}
		purego.RegisterFunc(binding.target, symbol)
	}
	// Locked/inaccessible keychains fail without a prompt during token rotation.
	if interaction(false) != 0 {
		return nil, errors.New("cannot disable keychain interaction")
	}
	pointer := func(data []byte) unsafe.Pointer {
		if len(data) == 0 {
			return nil
		}
		return unsafe.Pointer(&data[0])
	}
	return &keychainBackend{
		find: func(service, entry string) (uintptr, int32) {
			s, e := []byte(service), []byte(entry)
			var item uintptr
			status := find(0, uint32(len(s)), pointer(s), uint32(len(e)), pointer(e), 0, 0, &item)
			runtime.KeepAlive(s)
			runtime.KeepAlive(e)
			return item, status
		},
		add: func(service, entry string, secret []byte) int32 {
			s, e := []byte(service), []byte(entry)
			status := add(0, uint32(len(s)), pointer(s), uint32(len(e)), pointer(e), uint32(len(secret)), pointer(secret), 0)
			runtime.KeepAlive(s)
			runtime.KeepAlive(e)
			runtime.KeepAlive(secret)
			return status
		},
		update: func(item uintptr, secret []byte) int32 {
			status := update(item, 0, uint32(len(secret)), pointer(secret))
			runtime.KeepAlive(secret)
			return status
		},
		read: func(item uintptr) ([]byte, int32) {
			var length uint32
			var data unsafe.Pointer
			status := copyData(item, 0, 0, 0, &length, &data)
			if data != nil {
				defer freeData(0, data)
			}
			if status != 0 {
				return nil, status
			}
			if length == 0 {
				return nil, 0
			}
			if data == nil {
				return nil, -50
			}
			result := append([]byte(nil), unsafe.Slice((*byte)(data), int(length))...)
			return result, 0
		},
		delete:  deleteItem,
		release: release,
	}, nil
}
