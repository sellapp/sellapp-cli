package cli

import (
	"strings"
	"testing"
)

func TestKeychainWriteReplacesWholePairWithoutDelete(t *testing.T) {
	secret := []byte("{\"access_token\":\"new-access\",\"refresh_token\":\"new-refresh\"}")
	expected := string(secret)
	released := false
	updated := false
	api := &keychainBackend{
		find: func(service, entry string) (uintptr, int32) {
			if service != "service" || entry != "profile" {
				t.Fatal("wrong item")
			}
			return 42, 0
		},
		update: func(item uintptr, data []byte) int32 {
			if item != 42 || string(data) != expected {
				t.Fatal("wrong replacement")
			}
			updated = true
			return 0
		},
		add:     func(string, string, []byte) int32 { t.Fatal("existing item added again"); return 0 },
		release: func(item uintptr) { released = item == 42 },
	}
	if err := writeKeychainItem(api, "service", "profile", secret); err != nil {
		t.Fatal(err)
	}
	if !updated || !released {
		t.Fatal("missing update or release")
	}
	for _, v := range secret {
		if v != 0 {
			t.Fatal("secret buffer not cleared")
		}
	}
}
func TestKeychainWriteAddsAndHandlesConcurrentCreation(t *testing.T) {
	calls := 0
	updated := false
	api := &keychainBackend{
		find: func(string, string) (uintptr, int32) {
			calls++
			if calls == 1 {
				return 0, -25300
			}
			return 7, 0
		},
		add:     func(string, string, []byte) int32 { return -25299 },
		update:  func(item uintptr, data []byte) int32 { updated = item == 7 && string(data) == "pair"; return 0 },
		release: func(uintptr) {},
	}
	if err := writeKeychainItem(api, "service", "profile", []byte("pair")); err != nil {
		t.Fatal(err)
	}
	if calls != 2 || !updated {
		t.Fatal("concurrent insertion not updated")
	}
}
func TestKeychainNewItemStoresExactBytes(t *testing.T) {
	expected := "{\"access_token\":\"a\",\"refresh_token\":\"r\"}"
	added := false
	api := &keychainBackend{
		find: func(string, string) (uintptr, int32) { return 0, -25300 },
		add: func(service, entry string, data []byte) int32 {
			added = service == "service" && entry == "entry" && string(data) == expected
			return 0
		},
	}
	if err := writeKeychainItem(api, "service", "entry", []byte(expected)); err != nil {
		t.Fatal(err)
	}
	if !added {
		t.Fatal("new credential changed or omitted")
	}
}
func TestKeychainUpdateFailureReleasesReference(t *testing.T) {
	released := false
	api := &keychainBackend{
		find:    func(string, string) (uintptr, int32) { return 9, 0 },
		update:  func(uintptr, []byte) int32 { return -25308 },
		release: func(item uintptr) { released = item == 9 },
	}
	if err := writeKeychainItem(api, "service", "entry", []byte("pair")); err == nil {
		t.Fatal("failure ignored")
	}
	if !released {
		t.Fatal("reference leaked after failure")
	}
}
func TestKeychainReadAndDeleteLifecycle(t *testing.T) {
	stored := []byte(" exact pair \n")
	released := 0
	deleted := false
	api := &keychainBackend{
		find: func(service, entry string) (uintptr, int32) {
			if service != "service" || entry != "entry" {
				t.Fatal("wrong lookup")
			}
			return 10, 0
		},
		read: func(item uintptr) ([]byte, int32) {
			if item != 10 {
				t.Fatal("wrong item")
			}
			return stored, 0
		},
		delete: func(item uintptr) int32 { deleted = item == 10; return 0 },
		release: func(item uintptr) {
			if item != 10 {
				t.Fatal("wrong release")
			}
			released++
		},
	}
	value, err := readKeychainItem(api, "service", "entry")
	if err != nil || value != " exact pair \n" {
		t.Fatal("read changed bytes or failed")
	}
	for _, v := range stored {
		if v != 0 {
			t.Fatal("read buffer not cleared")
		}
	}
	if err = deleteKeychainItem(api, "service", "entry"); err != nil {
		t.Fatal(err)
	}
	if !deleted || released != 2 {
		t.Fatal("incomplete lifecycle")
	}
}
func TestKeychainReadDeleteErrorsReleaseAndHideDetails(t *testing.T) {
	released := 0
	api := &keychainBackend{
		find:    func(string, string) (uintptr, int32) { return 12, 0 },
		read:    func(uintptr) ([]byte, int32) { return []byte("secret"), -25308 },
		delete:  func(uintptr) int32 { return -25308 },
		release: func(uintptr) { released++ },
	}
	value, err := readKeychainItem(api, "private-service", "private-entry")
	if value != "" || err == nil || strings.Contains(err.Error(), "private") || strings.Contains(err.Error(), "secret") {
		t.Fatal("unsafe read failure")
	}
	err = deleteKeychainItem(api, "private-service", "private-entry")
	if err == nil || strings.Contains(err.Error(), "private") || released != 2 {
		t.Fatal("unsafe delete failure")
	}
}
func TestKeychainFailuresDoNotExposeCredentials(t *testing.T) {
	secret := []byte("never-print-this")
	api := &keychainBackend{find: func(string, string) (uintptr, int32) { return 0, -25308 }}
	err := writeKeychainItem(api, "private-service", "private-profile", secret)
	if err == nil || strings.Contains(err.Error(), "private") || strings.Contains(err.Error(), "never-print") {
		t.Fatal("unsafe error")
	}
}
