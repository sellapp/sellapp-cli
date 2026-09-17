//go:build !darwin

package cli

import "errors"

func loadNativeKeychain() (*keychainBackend, error) {
	return nil, errors.New("native keychain is available only on macOS")
}
