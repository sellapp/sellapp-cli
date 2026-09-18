package cli

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const credentialContext = "SellApp CLI encrypted credential v1"

type secretStore interface {
	Read(string) (string, error)
	Write(string, string) error
	Delete(string) error
}
type fileSecretStore struct{}
type keyringSecretStore struct{}

// These adapters keep native-vault tests independent of the user's actual vault.
var credentialVaultRead = vaultRead
var credentialVaultWrite = vaultWrite
var credentialVaultDelete = vaultDelete
var credentialMachineID = platformMachineID

func (keyringSecretStore) Read(entry string) (string, error) { return credentialVaultRead(entry) }
func (keyringSecretStore) Write(entry, secret string) error {
	return credentialVaultWrite(entry, secret)
}
func (keyringSecretStore) Delete(entry string) error { return credentialVaultDelete(entry) }

func credentialBackend(record profileRecord) string {
	if record.SecretStore == "" {
		return "keyring"
	}
	return record.SecretStore
}
func requestedSecretStore() (string, error) {
	value := os.Getenv("SELLAPP_CLI_SECRET_STORE")
	if value != "" && value != "file" && value != "keyring" {
		return "", &UsageError{Message: "SELLAPP_CLI_SECRET_STORE must be file or keyring"}
	}
	return value, nil
}
func newSecretStore() (string, error) {
	value, err := requestedSecretStore()
	if err != nil {
		return "", err
	}
	if value == "" {
		value = "file"
	}
	if value == "keyring" && !vaultAvailable() {
		return "", errors.New("keyring credential storage was requested but is unavailable; enable your OS vault (Secret Service and secret-tool on Linux/WSL), or choose SELLAPP_CLI_SECRET_STORE=file")
	}
	return value, nil
}
func storeFor(record profileRecord) (secretStore, error) {
	switch credentialBackend(record) {
	case "file":
		return fileSecretStore{}, nil
	case "keyring":
		return keyringSecretStore{}, nil
	default:
		return nil, errors.New("unsupported credential storage backend; restore valid profile metadata")
	}
}
func credentialRead(record profileRecord) (string, error) {
	store, err := storeFor(record)
	if err != nil {
		return "", err
	}
	return store.Read(record.VaultEntry)
}
func credentialWrite(record profileRecord, secret string) error {
	store, err := storeFor(record)
	if err != nil {
		return err
	}
	return store.Write(record.VaultEntry, secret)
}
func credentialDelete(record profileRecord) error {
	store, err := storeFor(record)
	if err != nil {
		return err
	}
	return store.Delete(record.VaultEntry)
}

func credentialConfigDir() (string, error) {
	root, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	if !filepath.IsAbs(root) {
		return "", errors.New("configuration directory must be an absolute path")
	}
	return filepath.Join(root, "SellApp", "cli"), nil
}
func credentialDir() (string, error) {
	root, err := credentialConfigDir()
	if err != nil {
		return "", err
	}
	if err = privateDirectory(filepath.Dir(root), true); err != nil {
		return "", err
	}
	if err = privateDirectory(root, true); err != nil {
		return "", err
	}
	dir := filepath.Join(root, "credentials")
	if err = privateDirectory(dir, true); err != nil {
		return "", err
	}
	return dir, nil
}
func credentialFile(entry string) (string, error) {
	if entry == "" || len(entry) > 512 {
		return "", errors.New("invalid credential entry identity")
	}
	dir, err := credentialDir()
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256([]byte(entry))
	return filepath.Join(dir, hex.EncodeToString(sum[:])+".json"), nil
}

// The HKDF-SHA256 extract and single-block expand follow RFC 5869. A 32-byte
// AES key requires exactly one expansion block; identity is bound as GCM AAD.
func credentialKey(identity string, salt []byte) []byte {
	extract := hmac.New(sha256.New, salt)
	_, _ = extract.Write([]byte(identity))
	prk := extract.Sum(nil)
	expand := hmac.New(sha256.New, prk)
	_, _ = expand.Write([]byte(credentialContext))
	_, _ = expand.Write([]byte{1})
	key := expand.Sum(nil)
	for i := range prk {
		prk[i] = 0
	}
	return key
}

type encryptedCredential struct {
	Version        int    `json:"version"`
	IdentitySource string `json:"identity_source"`
	Salt           []byte `json:"salt"`
	Nonce          []byte `json:"nonce"`
	Ciphertext     []byte `json:"ciphertext"`
}

func credentialRecoveryError() error {
	return errors.New("cannot decrypt saved credential: this machine's identity changed or the credential file is damaged; the saved file was retained. Restore the original machine identity and file, or run sellapp login to replace this profile")
}
func credentialIdentity(source string, create bool) (string, string, error) {
	if source == "machine" || source == "" {
		if id, err := credentialMachineID(); err == nil && strings.TrimSpace(id) != "" {
			return "machine:" + strings.TrimSpace(id), "machine", nil
		}
		if source == "machine" {
			return "", "", errors.New("this credential requires the original machine identifier, which is currently unavailable; restore it and retry. The saved credential was retained")
		}
	}
	if source != "" && source != "installation" {
		return "", "", credentialRecoveryError()
	}
	dir, err := credentialDir()
	if err != nil {
		return "", "", err
	}
	path := filepath.Join(dir, "installation-id")
	data, err := privateRead(path)
	if err == nil {
		if len(data) != 64 {
			return "", "", errors.New("invalid private installation identity; restore it before opening existing credentials")
		}
		if _, err = hex.DecodeString(string(data)); err != nil {
			return "", "", errors.New("invalid private installation identity; restore it before opening existing credentials")
		}
		return "installation:" + string(data), "installation", nil
	}
	if !os.IsNotExist(err) {
		return "", "", err
	}
	if !create || source == "installation" {
		return "", "", errors.New("private installation identity is missing; restore the original installation-id file. The saved credential was retained")
	}
	unlock, err := lockCredentialPath(filepath.Join(dir, "identity.lock"), true)
	if err != nil {
		return "", "", err
	}
	defer unlock()
	data, err = privateRead(path)
	if err == nil {
		if len(data) != 64 {
			return "", "", errors.New("invalid private installation identity")
		}
		if _, err = hex.DecodeString(string(data)); err != nil {
			return "", "", errors.New("invalid private installation identity")
		}
		return "installation:" + string(data), "installation", nil
	}
	if !os.IsNotExist(err) {
		return "", "", err
	}
	random := make([]byte, 32)
	if _, err = rand.Read(random); err != nil {
		return "", "", err
	}
	data = []byte(hex.EncodeToString(random))
	if err = privateAtomicWrite(path, data); err != nil {
		return "", "", err
	}
	return "installation:" + string(data), "installation", nil
}
func validPlatformIdentifier(value string, uuid bool) bool {
	if uuid {
		if len(value) != 36 || value[8] != '-' || value[13] != '-' || value[18] != '-' || value[23] != '-' {
			return false
		}
		value = strings.ReplaceAll(value, "-", "")
	}
	if len(value) != 32 {
		return false
	}
	raw, err := hex.DecodeString(value)
	if err != nil {
		return false
	}
	for _, b := range raw {
		if b != 0 {
			return true
		}
	}
	return false
}
func credentialAEAD(identity string, salt []byte) (cipher.AEAD, error) {
	key := credentialKey(identity, salt)
	defer func() {
		for i := range key {
			key[i] = 0
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}
func credentialAAD(entry, source string) []byte {
	return []byte(credentialContext + "\x00" + entry + "\x00" + source)
}
func (fileSecretStore) Read(entry string) (string, error) {
	path, err := credentialFile(entry)
	if err != nil {
		return "", err
	}
	data, err := privateRead(path)
	if err != nil {
		return "", err
	}
	var record encryptedCredential
	if len(data) > 2<<20 || json.Unmarshal(data, &record) != nil || record.Version != 1 || len(record.Salt) != 32 || len(record.Nonce) != 12 || len(record.Ciphertext) < 16 {
		return "", credentialRecoveryError()
	}
	if record.IdentitySource != "machine" && record.IdentitySource != "installation" {
		return "", credentialRecoveryError()
	}
	identity, _, err := credentialIdentity(record.IdentitySource, false)
	if err != nil {
		return "", err
	}
	aead, err := credentialAEAD(identity, record.Salt)
	if err != nil {
		return "", err
	}
	plaintext, err := aead.Open(nil, record.Nonce, record.Ciphertext, credentialAAD(entry, record.IdentitySource))
	if err != nil {
		return "", credentialRecoveryError()
	}
	defer func() {
		for i := range plaintext {
			plaintext[i] = 0
		}
	}()
	return string(plaintext), nil
}
func (fileSecretStore) Write(entry, secret string) error {
	if secret == "" || len(secret) > 1<<20 {
		return errors.New("credential must be nonempty and at most 1 MiB")
	}
	path, err := credentialFile(entry)
	if err != nil {
		return err
	}
	source := ""
	existing, err := privateRead(path)
	if err == nil {
		var previous encryptedCredential
		if json.Unmarshal(existing, &previous) != nil || previous.Version != 1 || (previous.IdentitySource != "machine" && previous.IdentitySource != "installation") {
			return credentialRecoveryError()
		}
		source = previous.IdentitySource
	} else if !os.IsNotExist(err) {
		return err
	}
	identity, source, err := credentialIdentity(source, true)
	if err != nil {
		return err
	}
	record := encryptedCredential{Version: 1, IdentitySource: source, Salt: make([]byte, 32), Nonce: make([]byte, 12)}
	if _, err = rand.Read(record.Salt); err != nil {
		return err
	}
	if _, err = rand.Read(record.Nonce); err != nil {
		return err
	}
	aead, err := credentialAEAD(identity, record.Salt)
	if err != nil {
		return err
	}
	plaintext := []byte(secret)
	record.Ciphertext = aead.Seal(nil, record.Nonce, plaintext, credentialAAD(entry, source))
	for i := range plaintext {
		plaintext[i] = 0
	}
	data, err := json.Marshal(record)
	if err != nil {
		return err
	}
	return privateAtomicWrite(path, append(data, '\n'))
}
func (fileSecretStore) Delete(entry string) error {
	path, err := credentialFile(entry)
	if err != nil {
		return err
	}
	return privateRemove(path)
}

func checkPathParents(path string) error {
	if !filepath.IsAbs(path) {
		return errors.New("credential path must be absolute")
	}
	for current := filepath.Clean(path); ; current = filepath.Dir(current) {
		info, err := os.Lstat(current)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			if err = platformPathSafe(current, info); err != nil {
				return err
			}
			if current != path && !info.IsDir() {
				return errors.New("credential path parent is not a directory")
			}
		}
		if filepath.Dir(current) == current {
			break
		}
	}
	return nil
}
func privateDirectory(path string, create bool) error {
	if err := checkPathParents(path); err != nil {
		return err
	}
	info, err := os.Lstat(path)
	if os.IsNotExist(err) && create {
		if _, parentErr := os.Lstat(filepath.Dir(path)); os.IsNotExist(parentErr) {
			if err = privateDirectory(filepath.Dir(path), true); err != nil {
				return err
			}
		}
		if err = platformPrivateMkdir(path); err != nil && !os.IsExist(err) {
			return err
		}
		info, err = os.Lstat(path)
	}
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.New("credential directory is not a directory")
	}
	return platformPrivateCheck(path, info)
}
func privateRead(path string) ([]byte, error) {
	if err := checkPathParents(path); err != nil {
		return nil, err
	}
	before, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	if err = platformPrivateCheck(path, before); err != nil {
		return nil, err
	}
	if !before.Mode().IsRegular() {
		return nil, errors.New("credential file is not a regular file")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	after, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if !os.SameFile(before, after) {
		return nil, errors.New("credential file changed while opening")
	}
	return io.ReadAll(io.LimitReader(file, 2<<20+1))
}
func privateAtomicWrite(path string, data []byte) error {
	if err := privateDirectory(filepath.Dir(path), false); err != nil {
		return err
	}
	if err := checkPrivateTarget(path); err != nil {
		return err
	}
	random := make([]byte, 16)
	if _, err := rand.Read(random); err != nil {
		return err
	}
	tmp := filepath.Join(filepath.Dir(path), ".new-"+hex.EncodeToString(random))
	file, err := platformPrivateCreate(tmp)
	if err != nil {
		return err
	}
	defer os.Remove(tmp)
	if _, err = file.Write(data); err == nil {
		err = file.Sync()
	}
	closeErr := file.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	if err = checkPrivateTarget(path); err != nil {
		return err
	}
	if err = platformReplace(tmp, path); err != nil {
		return err
	}
	return platformSyncDirectory(filepath.Dir(path))
}
func checkPrivateTarget(path string) error {
	if err := checkPathParents(path); err != nil {
		return err
	}
	info, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() {
		return errors.New("credential target is not a regular file")
	}
	return platformPrivateCheck(path, info)
}
func privateRemove(path string) error {
	if err := checkPrivateTarget(path); err != nil {
		return err
	}
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return platformSyncDirectory(filepath.Dir(path))
}
func privateCreateMarker(path string) error {
	if err := privateDirectory(filepath.Dir(path), false); err != nil {
		return err
	}
	file, err := platformPrivateCreate(path)
	if err != nil {
		return err
	}
	_, err = file.WriteString("refresh pending\n")
	if err == nil {
		err = file.Sync()
	}
	closeErr := file.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	return platformSyncDirectory(filepath.Dir(path))
}

// Never automatically remove an existing credential lock: its process might
// still be exchanging a rotating refresh token with the server.
var credentialLockMkdir = platformPrivateMkdir
var credentialLockCheck = privateDirectory

func lockCredentialPath(path string, wait bool) (func(), error) {
	if err := privateDirectory(filepath.Dir(path), false); err != nil {
		return nil, err
	}
	deadline := time.Now().Add(5 * time.Second)
	for {
		err := credentialLockMkdir(path)
		if err == nil {
			return func() { _ = os.Remove(path) }, nil
		}
		if !os.IsExist(err) && !platformLockTransient(err) {
			return nil, fmt.Errorf("create credential lock: %w", err)
		}
		if os.IsExist(err) {
			if checkErr := credentialLockCheck(path, false); checkErr != nil && !os.IsNotExist(checkErr) && !platformLockTransient(checkErr) {
				return nil, fmt.Errorf("inspect credential lock: %w", checkErr)
			}
		}
		// Windows may retain a deleted lock directory until its last inspection
		// handle closes. Retry only acquiring the lock, never credential access.
		// Successful private directory creation remains the sole ownership signal.
		if !wait || !time.Now().Before(deadline) {
			return nil, fmt.Errorf("credentials are locked by another process; wait for it to finish. After a crashed refresh, run sellapp login with a new profile: %w", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
func updateProfiles(change func(*profileDocument) error) error {
	root, err := credentialConfigDir()
	if err != nil {
		return err
	}
	if _, err = credentialDir(); err != nil {
		return err
	}
	unlock, err := lockCredentialPath(filepath.Join(root, "profiles.lock"), true)
	if err != nil {
		return err
	}
	defer unlock()
	doc, err := loadProfiles()
	if err != nil {
		return err
	}
	if err = change(&doc); err != nil {
		return err
	}
	return saveProfiles(doc)
}
func saveNewProfile(name string, record profileRecord, secret string) error {
	unlock, _, err := lockOAuth(record.VaultEntry)
	if err != nil {
		return err
	}
	defer unlock()
	if err = credentialWrite(record, secret); err != nil {
		return err
	}
	actual, err := credentialRead(record)
	if err != nil {
		return err
	}
	if actual != secret {
		return errors.New("saved credential verification failed; previous profile was retained")
	}
	return updateProfiles(func(doc *profileDocument) error { doc.Profiles[name] = record; doc.Active = name; return nil })
}

// The caller holds the entry's credential lock. Metadata is committed only
// after the destination verifies, and the source is deleted only after commit.
func prepareCredentialRecord(record profileRecord) (profileRecord, error) {
	// Re-read after taking the entry lock so a concurrent completed migration or
	// logout cannot leave a long-running command using stale backend metadata.
	doc, err := loadProfiles()
	if err != nil {
		return record, err
	}
	for _, current := range doc.Profiles {
		if current.VaultEntry == record.VaultEntry {
			record = current
			break
		}
	}
	requested, err := requestedSecretStore()
	if err != nil {
		return record, err
	}
	if requested == "" || requested == credentialBackend(record) {
		if cleanupPreviousCredential(record) == nil {
			record.PreviousSecretStore = ""
		}
		return record, nil
	}
	if requested == "keyring" && !vaultAvailable() {
		return record, errors.New("keyring credential storage was requested but is unavailable; enable your OS vault or choose SELLAPP_CLI_SECRET_STORE=file")
	}
	source := record
	secret, err := credentialRead(source)
	if err != nil {
		return record, err
	}
	replacement := record
	replacement.SecretStore = requested
	replacement.PreviousSecretStore = credentialBackend(source)
	if err = credentialWrite(replacement, secret); err != nil {
		return record, err
	}
	verified, err := credentialRead(replacement)
	if err != nil {
		return record, err
	}
	if verified != secret {
		return record, errors.New("credential migration verification failed; source credential was retained")
	}
	if err = updateProfiles(func(doc *profileDocument) error {
		found := false
		for name, current := range doc.Profiles {
			if current.VaultEntry == record.VaultEntry {
				if credentialBackend(current) != credentialBackend(source) {
					return errors.New("profile changed during credential migration; retry")
				}
				doc.Profiles[name] = replacement
				found = true
			}
		}
		if !found {
			return errors.New("profile was removed during credential migration; source credential retained")
		}
		return nil
	}); err != nil {
		return record, err
	}
	if cleanupPreviousCredential(replacement) == nil {
		replacement.PreviousSecretStore = ""
	}
	return replacement, nil
}

// Logout retains the tracked profile until every local backend is cleared.
// OAuth callers revoke the server connection before reaching this helper.
func deleteProfileCredentials(record profileRecord) error {
	if err := cleanupPreviousCredential(record); err != nil {
		return err
	}
	return credentialDelete(record)
}
func cleanupPreviousCredential(record profileRecord) error {
	if record.PreviousSecretStore == "" {
		return nil
	}
	source := record
	source.SecretStore = record.PreviousSecretStore
	if credentialBackend(source) == credentialBackend(record) {
		return errors.New("invalid credential migration metadata")
	}
	if err := credentialDelete(source); err != nil {
		return fmt.Errorf("credential migration was saved, but the old backend could not be cleared; retry with the same storage setting: %w", err)
	}
	return updateProfiles(func(doc *profileDocument) error {
		for name, current := range doc.Profiles {
			if current.VaultEntry == record.VaultEntry && credentialBackend(current) == credentialBackend(record) {
				current.PreviousSecretStore = ""
				doc.Profiles[name] = current
			}
		}
		return nil
	})
}
