package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func credentialTestHome(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", root)
	t.Setenv("APPDATA", root)
	t.Setenv("HOME", root)
	t.Setenv("SELLAPP_CLI_SECRET_STORE", "")
	t.Setenv("SELLAPP_API_KEY", "")
	t.Setenv("SELLAPP_STORE", "")
	t.Setenv("SELLAPP_ACCESS_TOKEN", "")
	t.Setenv("SELLAPP_PROFILE", "")
	return root
}
func credentialTestIdentity(t *testing.T) {
	t.Helper()
	old := credentialMachineID
	credentialMachineID = func() (string, error) { return "test-machine-identity", nil }
	t.Cleanup(func() { credentialMachineID = old })
}
func credentialTestRecord(entry string) profileRecord {
	return profileRecord{Method: "api-key", VaultEntry: entry, SecretStore: "file", Store: "launch-lab"}
}
func TestCredentialBackendSelection(t *testing.T) {
	credentialTestHome(t)
	if actual, err := newSecretStore(); err != nil || actual != "file" {
		t.Fatal(actual, err)
	}
	if credentialBackend(profileRecord{}) != "keyring" {
		t.Fatal("missing metadata must denote keyring")
	}
	for _, value := range []string{"file", "keyring"} {
		t.Setenv("SELLAPP_CLI_SECRET_STORE", value)
		if actual, err := requestedSecretStore(); err != nil || actual != value {
			t.Fatal(actual, err)
		}
	}
	t.Setenv("SELLAPP_CLI_SECRET_STORE", "invalid")
	if _, err := newSecretStore(); err == nil {
		t.Fatal("invalid backend accepted")
	}
}
func TestCredentialForcedUnavailableKeyring(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("macOS native availability is covered by native vault tests")
	}
	credentialTestHome(t)
	t.Setenv("PATH", t.TempDir())
	t.Setenv("SELLAPP_CLI_SECRET_STORE", "keyring")
	if _, err := newSecretStore(); err == nil || !strings.Contains(err.Error(), "requested") {
		t.Fatal("forced keyring silently fell back", err)
	}
}
func TestCredentialEncryptedRoundTripAndRandomizedRecords(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	record := credentialTestRecord("entry")
	secret := "private-access private-refresh\n exact bytes "
	if err := credentialWrite(record, secret); err != nil {
		t.Fatal(err)
	}
	path, _ := credentialFile(record.VaultEntry)
	first, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(first, []byte("private")) {
		t.Fatal("plaintext persisted")
	}
	actual, err := credentialRead(record)
	if err != nil || actual != secret {
		t.Fatal(actual, err)
	}
	if err = credentialWrite(record, secret); err != nil {
		t.Fatal(err)
	}
	second, _ := os.ReadFile(path)
	if bytes.Equal(first, second) {
		t.Fatal("salt or nonce reused")
	}
	var envelope encryptedCredential
	if err = json.Unmarshal(second, &envelope); err != nil || envelope.Version != 1 || len(envelope.Salt) != 32 || len(envelope.Nonce) != 12 || envelope.IdentitySource != "machine" {
		t.Fatal("invalid format", err)
	}
}
func TestCredentialEntryIdentityAuthenticated(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	store := fileSecretStore{}
	if err := store.Write("first", "secret"); err != nil {
		t.Fatal(err)
	}
	first, _ := credentialFile("first")
	second, _ := credentialFile("second")
	data, _ := privateRead(first)
	if err := privateAtomicWrite(second, data); err != nil {
		t.Fatal(err)
	}
	if _, err := store.Read("second"); err == nil {
		t.Fatal("credential identity substitution accepted")
	}
	after, _ := os.ReadFile(second)
	if !bytes.Equal(data, after) {
		t.Fatal("failed authentication deleted credential")
	}
}
func TestCredentialCorruptAndChangedIdentityRetained(t *testing.T) {
	for _, kind := range []string{"ciphertext", "nonce", "version", "json", "machine", "source"} {
		t.Run(kind, func(t *testing.T) {
			credentialTestHome(t)
			credentialTestIdentity(t)
			record := credentialTestRecord("entry")
			if err := credentialWrite(record, "sensitive"); err != nil {
				t.Fatal(err)
			}
			path, _ := credentialFile("entry")
			data, _ := privateRead(path)
			var encrypted encryptedCredential
			_ = json.Unmarshal(data, &encrypted)
			switch kind {
			case "ciphertext":
				encrypted.Ciphertext[0] ^= 1
			case "nonce":
				encrypted.Nonce = []byte{1}
			case "version":
				encrypted.Version = 99
			case "machine":
				credentialMachineID = func() (string, error) { return "different-machine", nil }
			case "source":
				encrypted.IdentitySource = "installation"
			}
			if kind == "json" {
				data = []byte("{broken")
			} else {
				data, _ = json.Marshal(encrypted)
			}
			if err := privateAtomicWrite(path, data); err != nil {
				t.Fatal(err)
			}
			if value, err := credentialRead(record); err == nil || value != "" || strings.Contains(err.Error(), "sensitive") {
				t.Fatal("unsafe decryption result", err)
			}
			after, _ := privateRead(path)
			if !bytes.Equal(data, after) {
				t.Fatal("credential modified after failed authentication")
			}
		})
	}
}
func TestCredentialFallbackIdentityPersistsAndDoesNotChangeSources(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	credentialMachineID = func() (string, error) { return "", errors.New("unavailable") }
	record := credentialTestRecord("entry")
	if err := credentialWrite(record, "secret"); err != nil {
		t.Fatal(err)
	}
	dir, _ := credentialDir()
	path := filepath.Join(dir, "installation-id")
	original, err := privateRead(path)
	if err != nil || len(original) != 64 {
		t.Fatal(err)
	}
	credentialMachineID = func() (string, error) { return "newly-available-machine", nil }
	if actual, err := credentialRead(record); err != nil || actual != "secret" {
		t.Fatal("fallback became unreadable", err)
	}
	if err := credentialWrite(record, "updated"); err != nil {
		t.Fatal(err)
	}
	if actual, err := credentialRead(record); err != nil || actual != "updated" {
		t.Fatal(err)
	}
	if after, _ := privateRead(path); !bytes.Equal(after, original) {
		t.Fatal("fallback identity replaced")
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	if _, err := credentialRead(record); err == nil {
		t.Fatal("missing identity accepted")
	}
	if _, err := os.Lstat(path); !os.IsNotExist(err) {
		t.Fatal("read regenerated fallback identity")
	}
}
func TestCredentialMachineIdentityTemporaryFailureDoesNotCreateFallback(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	record := credentialTestRecord("entry")
	if err := credentialWrite(record, "secret"); err != nil {
		t.Fatal(err)
	}
	credentialMachineID = func() (string, error) { return "", errors.New("unavailable") }
	if _, err := credentialRead(record); err == nil {
		t.Fatal("machine failure ignored")
	}
	dir, _ := credentialDir()
	if _, err := os.Lstat(filepath.Join(dir, "installation-id")); !os.IsNotExist(err) {
		t.Fatal("read generated substitute identity")
	}
	credentialMachineID = func() (string, error) { return "test-machine-identity", nil }
	if actual, err := credentialRead(record); err != nil || actual != "secret" {
		t.Fatal(err)
	}
}
func TestCredentialPlatformIdentifiersRejectInvalidDefaults(t *testing.T) {
	for _, value := range []string{"", "uninitialized", "00000000000000000000000000000000", "secret", "gggggggggggggggggggggggggggggggg"} {
		if validPlatformIdentifier(value, false) {
			t.Fatal("invalid machine identity accepted")
		}
	}
	if !validPlatformIdentifier("fb718f7442e04743bd309d0b6e54171b", false) || !validPlatformIdentifier("fb718f74-42e0-4743-bd30-9d0b6e54171b", true) || validPlatformIdentifier("00000000-0000-0000-0000-000000000000", true) {
		t.Fatal("UUID validation failed")
	}
}
func TestCredentialPrivatePermissionsAndFreshConfigurationRoot(t *testing.T) {
	root := credentialTestHome(t)
	credentialTestIdentity(t)
	nested := filepath.Join(root, "not-created", "configuration")
	t.Setenv("XDG_CONFIG_HOME", nested)
	t.Setenv("APPDATA", nested)
	record := credentialTestRecord("entry")
	if err := credentialWrite(record, "secret"); err != nil {
		t.Fatal(err)
	}
	file, _ := credentialFile("entry")
	for _, path := range []string{file, filepath.Dir(file), filepath.Dir(filepath.Dir(file))} {
		info, err := os.Lstat(path)
		if err != nil {
			t.Fatal(err)
		}
		if err = platformPrivateCheck(path, info); err != nil {
			t.Fatal("permissions were not private", err)
		}
	}
}
func TestCredentialRejectsUnsafeModesAndHardLinks(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows ACL checks run in dedicated native tests")
	}
	for _, kind := range []string{"file", "directory", "hardlink", "symlink", "parent"} {
		t.Run(kind, func(t *testing.T) {
			credentialTestHome(t)
			credentialTestIdentity(t)
			record := credentialTestRecord("entry")
			if err := credentialWrite(record, "secret"); err != nil {
				t.Fatal(err)
			}
			path, _ := credentialFile("entry")
			switch kind {
			case "file":
				_ = os.Chmod(path, 0644)
			case "directory":
				_ = os.Chmod(filepath.Dir(path), 0755)
			case "hardlink":
				if err := os.Link(path, path+".link"); err != nil {
					t.Fatal(err)
				}
			case "symlink":
				if err := os.Rename(path, path+".source"); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(path+".source", path); err != nil {
					t.Fatal(err)
				}
			case "parent":
				dir := filepath.Dir(path)
				if err := os.Rename(dir, dir+".source"); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(dir+".source", dir); err != nil {
					t.Fatal(err)
				}
			}
			if _, err := credentialRead(record); err == nil {
				t.Fatal("unsafe read accepted")
			}
			if err := credentialWrite(record, "replacement"); err == nil {
				t.Fatal("unsafe replacement accepted")
			}
		})
	}
}
func TestCredentialAtomicReplacementPreservesOldFileOnUnsafeTarget(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	record := credentialTestRecord("entry")
	if err := credentialWrite(record, "secret"); err != nil {
		t.Fatal(err)
	}
	path, _ := credentialFile("entry")
	source := path + ".original"
	if err := os.Rename(path, source); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(path, 0700); err != nil {
		t.Fatal(err)
	}
	if err := credentialWrite(record, "replacement"); err == nil {
		t.Fatal("directory target replaced")
	}
	if info, err := os.Stat(path); err != nil || !info.IsDir() {
		t.Fatal("target was destroyed")
	}
	if _, err := os.Stat(source); err != nil {
		t.Fatal("old credential lost")
	}
	entries, _ := os.ReadDir(filepath.Dir(path))
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".new-") {
			t.Fatal("temporary credential left behind")
		}
	}
}
func TestCredentialConcurrentProfilesRemainIndependent(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	var wg sync.WaitGroup
	failures := make(chan error, 12)
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			name := fmt.Sprintf("profile-%d", i)
			failures <- saveNewProfile(name, credentialTestRecord(name), "secret-"+name)
		}(i)
	}
	wg.Wait()
	close(failures)
	for err := range failures {
		if err != nil {
			t.Fatal(err)
		}
	}
	doc, err := loadProfiles()
	if err != nil || len(doc.Profiles) != 12 {
		t.Fatal("profiles overwritten", err, len(doc.Profiles))
	}
	for name, record := range doc.Profiles {
		actual, err := credentialRead(record)
		if err != nil || actual != "secret-"+name {
			t.Fatal("credential overwritten", err)
		}
	}
}
func TestCredentialLegacyMigrationAndCleanupRetry(t *testing.T) {
	for _, cleanupFails := range []bool{false, true} {
		t.Run(fmt.Sprint(cleanupFails), func(t *testing.T) {
			credentialTestHome(t)
			credentialTestIdentity(t)
			oldRead, oldDelete := credentialVaultRead, credentialVaultDelete
			t.Cleanup(func() { credentialVaultRead = oldRead; credentialVaultDelete = oldDelete })
			sourceExists := true
			credentialVaultRead = func(string) (string, error) {
				if !sourceExists {
					return "", errors.New("source missing")
				}
				return "old-secret", nil
			}
			credentialVaultDelete = func(string) error {
				doc, err := loadProfiles()
				if err != nil {
					t.Fatal(err)
				}
				if doc.Profiles["legacy"].SecretStore != "file" {
					t.Fatal("source deleted before metadata commit")
				}
				if cleanupFails {
					return errors.New("vault locked")
				}
				sourceExists = false
				return nil
			}
			record := profileRecord{VaultEntry: "legacy-entry", Store: "launch-lab"}
			if err := saveProfiles(profileDocument{SchemaVersion: 1, Active: "legacy", Profiles: map[string]profileRecord{"legacy": record}}); err != nil {
				t.Fatal(err)
			}
			t.Setenv("SELLAPP_CLI_SECRET_STORE", "file")
			cfg := &Config{}
			if err := resolveProfileConfig(cfg); err != nil || cfg.APIKey != "old-secret" {
				t.Fatal("migration made credential unusable", err)
			}
			doc, _ := loadProfiles()
			if doc.Profiles["legacy"].SecretStore != "file" || sourceExists != cleanupFails {
				t.Fatal("wrong migration state")
			}
			if cleanupFails {
				if doc.Profiles["legacy"].PreviousSecretStore != "keyring" {
					t.Fatal("cleanup was not retained for retry")
				}
				cleanupFails = false
				if err := resolveProfileConfig(&Config{}); err != nil {
					t.Fatal(err)
				}
				if sourceExists {
					t.Fatal("cleanup retry did not finish")
				}
			}
		})
	}
}
func TestCredentialMigrationWriteFailureRetainsSourceAndMetadata(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	old := credentialVaultRead
	t.Cleanup(func() { credentialVaultRead = old })
	credentialVaultRead = func(string) (string, error) { return "source-secret", nil }
	record := profileRecord{VaultEntry: "legacy", Store: "launch-lab"}
	if err := saveProfiles(profileDocument{SchemaVersion: 1, Active: "legacy", Profiles: map[string]profileRecord{"legacy": record}}); err != nil {
		t.Fatal(err)
	}
	target, _ := credentialFile("legacy")
	if err := os.Mkdir(target, 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("SELLAPP_CLI_SECRET_STORE", "file")
	if err := resolveProfileConfig(&Config{}); err == nil {
		t.Fatal("migration succeeded with unsafe destination")
	}
	doc, _ := loadProfiles()
	if doc.Profiles["legacy"].SecretStore != "" {
		t.Fatal("metadata changed before verified save")
	}
	t.Setenv("SELLAPP_CLI_SECRET_STORE", "")
	if value, err := credentialRead(doc.Profiles["legacy"]); err != nil || value != "source-secret" {
		t.Fatal("source lost", err)
	}
}
func TestCredentialMigrationMetadataFailureRetainsSource(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	oldRead, oldDelete := credentialVaultRead, credentialVaultDelete
	t.Cleanup(func() { credentialVaultRead = oldRead; credentialVaultDelete = oldDelete })
	record := profileRecord{VaultEntry: "legacy", Store: "launch-lab"}
	if err := saveProfiles(profileDocument{SchemaVersion: 1, Active: "legacy", Profiles: map[string]profileRecord{"legacy": record}}); err != nil {
		t.Fatal(err)
	}
	path, _ := profilePath()
	original, _ := privateRead(path)
	credentialVaultDelete = func(string) error { t.Fatal("source deleted before metadata commit"); return nil }
	credentialVaultRead = func(string) (string, error) {
		if err := os.Rename(path, path+".saved"); err != nil {
			return "", err
		}
		if err := os.Mkdir(path, 0700); err != nil {
			return "", err
		}
		return "source-secret", nil
	}
	t.Setenv("SELLAPP_CLI_SECRET_STORE", "file")
	if err := resolveProfileConfig(&Config{}); err == nil {
		t.Fatal("unsafe metadata commit succeeded")
	}
	retained, _ := privateRead(path + ".saved")
	if !bytes.Equal(retained, original) {
		t.Fatal("original metadata lost")
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(path+".saved", path); err != nil {
		t.Fatal(err)
	}
	doc, _ := loadProfiles()
	if doc.Profiles["legacy"].SecretStore != "" {
		t.Fatal("source metadata changed")
	}
}
func TestCredentialInterruptedReplacementPreservesReadableRecord(t *testing.T) {
	if os.Getenv("SELLAPP_TEST_CREDENTIAL_REPLACE_CHILD") == "1" {
		credentialMachineID = func() (string, error) { return "test-machine-identity", nil }
		record := credentialTestRecord("interrupted")
		if err := credentialWrite(record, "replacement"); err != nil {
			t.Fatal(err)
		}
		_, _ = os.Stdout.Write([]byte("."))
		for {
			if err := credentialWrite(record, "replacement"); err != nil {
				t.Fatal(err)
			}
		}
	}
	credentialTestHome(t)
	credentialTestIdentity(t)
	record := credentialTestRecord("interrupted")
	if err := credentialWrite(record, "original"); err != nil {
		t.Fatal(err)
	}
	child := exec.Command(os.Args[0], "-test.run=^TestCredentialInterruptedReplacementPreservesReadableRecord$")
	child.Env = append(os.Environ(), "SELLAPP_TEST_CREDENTIAL_REPLACE_CHILD=1")
	output, err := child.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err = child.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = child.Process.Kill() })
	ready := make([]byte, 1)
	if _, err = io.ReadFull(output, ready); err != nil {
		t.Fatal(err)
	}
	if err = child.Process.Kill(); err != nil {
		t.Fatal(err)
	}
	_ = child.Wait()
	actual, err := credentialRead(record)
	if err != nil || (actual != "original" && actual != "replacement") {
		t.Fatal("interrupted atomic write damaged credential", err)
	}
}
func TestCredentialAPIKeyPrecedenceAndStatusRedaction(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	if err := saveNewProfile("file", credentialTestRecord("entry"), "saved-private-key"); err != nil {
		t.Fatal(err)
	}
	cfg := &Config{APIKey: "explicit", Store: "chosen"}
	if err := resolveProfileConfig(cfg); err != nil || cfg.APIKey != "explicit" {
		t.Fatal(err)
	}
	t.Setenv("SELLAPP_API_KEY", "environment")
	cfg = &Config{}
	if err := resolveProfileConfig(cfg); err != nil || cfg.APIKey != "environment" {
		t.Fatal(err)
	}
	t.Setenv("SELLAPP_API_KEY", "")
	output, err := execute(t, "auth", "status")
	if err != nil || !strings.Contains(output, "storage_backend") || !strings.Contains(output, "file") || strings.Contains(output, "saved-private-key") {
		t.Fatal(output, err)
	}
}
func TestCredentialManualKeyLoginWithoutVault(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	t.Setenv("PATH", t.TempDir())
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if r.Header.Get("Authorization") != "Bearer manual-secret" {
			t.Error("wrong key")
		}
		if r.Method != "GET" || r.URL.Path != "/v2/products" || r.URL.Query().Get("limit") != "1" || r.Header.Get("X-STORE") != "launch-lab" {
			t.Error("manual key login must retain its bounded store read", r.Method, r.URL.String())
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, responseFixture("GET /v2/products"))
	}))
	defer server.Close()
	output, err := execute(t, "auth", "login", "--method", "api-key", "--api-key", "manual-secret", "--store", "launch-lab", "--base-url", server.URL, "--no-interactive")
	if err != nil || requests != 1 || strings.Contains(output, "manual-secret") {
		t.Fatal(output, err, requests)
	}
	doc, _ := loadProfiles()
	record := doc.Profiles["default"]
	if record.SecretStore != "file" {
		t.Fatal("manual login did not default to file")
	}
	if value, err := credentialRead(record); err != nil || value != "manual-secret" {
		t.Fatal(err)
	}
}
func TestCredentialRejectedManualKeyDoesNotSaveProfile(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		_, _ = io.WriteString(w, "{}")
	}))
	defer server.Close()
	if _, err := execute(t, "login", "--method", "api-key", "--api-key", "bad-secret", "--store", "launch-lab", "--base-url", server.URL); err == nil {
		t.Fatal("invalid key accepted")
	}
	doc, _ := loadProfiles()
	if len(doc.Profiles) != 0 {
		t.Fatal("invalid key saved")
	}
}
func TestCredentialBrowserlessOAuthLoginWithoutVault(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	t.Setenv("PATH", t.TempDir())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/oauth/token":
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "login-access", "refresh_token": "login-refresh", "token_type": "Bearer", "expires_in": 3600})
		case "/api/v2/stores":
			_, _ = io.WriteString(w, "{\"data\":[{\"id\":\"1\",\"slug\":\"launch-lab\",\"name\":\"Launch Lab\",\"permissions\":[]}]}")
		default:
			t.Error("unexpected path", r.URL.Path)
			w.WriteHeader(404)
		}
	}))
	defer server.Close()
	reader, writer := io.Pipe()
	defer reader.Close()
	var stderr bytes.Buffer
	root := NewRootCommand(writer, &stderr)
	root.SetArgs([]string{"login", "--client-id", "test-native", "--base-url", server.URL + "/api", "--no-browser", "--no-interactive"})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	done := make(chan error, 1)
	go func() { err := root.ExecuteContext(ctx); _ = writer.Close(); done <- err }()
	decoder := json.NewDecoder(reader)
	var event map[string]any
	if err := decoder.Decode(&event); err != nil {
		t.Fatal(err)
	}
	address, _ := url.Parse(event["url"].(string))
	callback := address.Query().Get("redirect_uri") + "?state=" + address.Query().Get("state") + "&code=single-use"
	response, err := http.Get(callback)
	if err != nil {
		t.Fatal(err)
	}
	response.Body.Close()
	if err = decoder.Decode(&event); err != nil {
		t.Fatal(err)
	}
	if err = <-done; err != nil {
		t.Fatal(err)
	}
	if event["type"] != "connected" || event["store"] != "launch-lab" || event["storage_backend"] != "file" {
		t.Fatal(event)
	}
	if stderr.Len() != 0 {
		t.Fatal("noninteractive output was decorated", stderr.String())
	}
	doc, _ := loadProfiles()
	record := doc.Profiles["default"]
	pair, err := loadOAuthPair(record)
	if err != nil || pair.AccessToken != "login-access" || record.SecretStore != "file" {
		t.Fatal(err)
	}
}
func TestCredentialFileRefreshConcurrentAndUncertainOutcomes(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	var requests atomic.Int32
	entered := make(chan struct{})
	release := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		close(entered)
		<-release
		_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "new-access", "refresh_token": "new-refresh", "token_type": "Bearer", "expires_in": 3600})
	}))
	defer server.Close()
	record := profileRecord{Method: "oauth", SecretStore: "file", VaultEntry: "refresh-entry", Origin: server.URL, ClientID: "client"}
	pair := oauthPair{AccessToken: "old-access", RefreshToken: "old-refresh", ExpiresAt: time.Now().Add(-time.Hour)}
	data, _ := json.Marshal(pair)
	if err := saveNewProfile("refresh", record, string(data)); err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { done <- resolveOAuth(record, &Config{}) }()
	<-entered
	if err := resolveOAuth(record, &Config{}); err == nil || !strings.Contains(err.Error(), "locked") {
		t.Fatal("simultaneous refresh accepted", err)
	}
	close(release)
	if err := <-done; err != nil {
		t.Fatal(err)
	}
	cfg := &Config{}
	if err := resolveOAuth(record, cfg); err != nil || cfg.AccessToken != "new-access" || requests.Load() != 1 {
		t.Fatal("rotation was not persisted", err)
	}
	persisted, err := loadOAuthPair(record)
	if err != nil || persisted.RefreshToken != "new-refresh" {
		t.Fatal(err)
	}
}
func TestCredentialRefreshReplayRejectedAndMarkerRetained(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.WriteHeader(400)
		_, _ = io.WriteString(w, "{\"error\":\"invalid_grant\"}")
	}))
	defer server.Close()
	record := profileRecord{Method: "oauth", SecretStore: "file", VaultEntry: "refresh", Origin: server.URL, ClientID: "client"}
	pair := oauthPair{AccessToken: "old-access", RefreshToken: "used-refresh", ExpiresAt: time.Now().Add(-time.Hour)}
	data, _ := json.Marshal(pair)
	if err := saveNewProfile("refresh", record, string(data)); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		if err := resolveOAuth(record, &Config{}); err == nil {
			t.Fatal("replayed refresh accepted")
		}
	}
	if requests != 1 {
		t.Fatal("refresh replay retried")
	}
	if actual, err := credentialRead(record); err != nil || actual != string(data) {
		t.Fatal("uncertain credentials deleted", err)
	}
}
func TestCredentialLogoutRevokesBeforeDeletingAndRetainsOnFailure(t *testing.T) {
	for _, success := range []bool{false, true} {
		t.Run(fmt.Sprint(success), func(t *testing.T) {
			credentialTestHome(t)
			credentialTestIdentity(t)
			var record profileRecord
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/oauth/revoke" {
					t.Error("wrong revocation path")
				}
				if _, err := credentialRead(record); err != nil {
					t.Error("credential removed before revocation")
				}
				if !success {
					w.WriteHeader(503)
					return
				}
				w.WriteHeader(204)
			}))
			defer server.Close()
			record = profileRecord{Method: "oauth", SecretStore: "file", VaultEntry: "revoke", Origin: server.URL, ClientID: "client"}
			data, _ := json.Marshal(oauthPair{AccessToken: "access", RefreshToken: "refresh", ExpiresAt: time.Now().Add(time.Hour)})
			if err := saveNewProfile("logout", record, string(data)); err != nil {
				t.Fatal(err)
			}
			_, err := execute(t, "auth", "logout")
			if (err == nil) != success {
				t.Fatal(err)
			}
			doc, _ := loadProfiles()
			_, exists := doc.Profiles["logout"]
			if exists == success {
				t.Fatal("wrong profile deletion state")
			}
			_, err = credentialRead(record)
			if success && !os.IsNotExist(err) {
				t.Fatal("credential remains after revocation", err)
			}
			if !success && err != nil {
				t.Fatal("failed revoke lost credentials", err)
			}
		})
	}
}
func TestCredentialLogoutTracksFailedMigrationCleanup(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	oldDelete := credentialVaultDelete
	t.Cleanup(func() { credentialVaultDelete = oldDelete })
	record := credentialTestRecord("pending-logout")
	record.PreviousSecretStore = "keyring"
	if err := saveNewProfile("pending", record, "still-valid-key"); err != nil {
		t.Fatal(err)
	}
	blocked := true
	deletions := 0
	credentialVaultDelete = func(string) error {
		if blocked {
			return errors.New("keyring unavailable")
		}
		deletions++
		if deletions > 1 {
			return errors.New("keyring entry already removed")
		}
		return nil
	}
	if _, err := execute(t, "auth", "logout"); err == nil {
		t.Fatal("logout hid failed source cleanup")
	}
	doc, err := loadProfiles()
	if err != nil || doc.Profiles["pending"].PreviousSecretStore != "keyring" {
		t.Fatal("pending source no longer tracked", err)
	}
	if secret, err := credentialRead(record); err != nil || secret != "still-valid-key" {
		t.Fatal("working destination removed on cleanup failure", err)
	}
	blocked = false
	if _, err = execute(t, "auth", "logout"); err != nil {
		t.Fatal("logout cleanup retry failed", err)
	}
	doc, err = loadProfiles()
	if err != nil || len(doc.Profiles) != 0 || deletions != 1 {
		t.Fatal("logout did not clear both tracked backends", err, deletions)
	}
	if _, err = credentialRead(record); !os.IsNotExist(err) {
		t.Fatal("credential remains after logout", err)
	}
}
func TestCredentialAuditBeforeLoginKeepsPrivateStorageDirectories(t *testing.T) {
	credentialTestHome(t)
	credentialTestIdentity(t)
	t.Setenv("SELLAPP_AUDIT_DISABLED", "")
	writeAudit(&Config{Store: "launch-lab"}, "listProducts", time.Now(), nil)
	if err := saveNewProfile("after-audit", credentialTestRecord("after-audit"), "key"); err != nil {
		t.Fatal("audit made later login storage unavailable", err)
	}
}
func TestCredentialMissingLinuxVaultDeleteIsRetryableButErrorsAreNotHidden(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("secret-tool exit contract is Linux-only")
	}
	credentialTestHome(t)
	bin := t.TempDir()
	t.Setenv("PATH", bin)
	tool := filepath.Join(bin, "secret-tool")
	if err := os.WriteFile(tool, []byte("#!/bin/sh\nexit 1\n"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := vaultDelete("already-removed"); err != nil {
		t.Fatal("missing entry cannot be retried", err)
	}
	if err := os.WriteFile(tool, []byte("#!/bin/sh\nprintf 'service unavailable' >&2\nexit 1\n"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := vaultDelete("locked-vault"); err == nil {
		t.Fatal("vault failure hidden as a missing entry")
	}
	if err := os.WriteFile(tool, []byte("#!/bin/sh\nexit 2\n"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := vaultDelete("unknown-failure"); err == nil {
		t.Fatal("unexpected exit hidden")
	}
}
