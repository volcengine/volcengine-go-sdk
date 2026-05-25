package clicreds

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
)

func writeTempConfig(t *testing.T, dir string, config interface{}) string {
	t.Helper()
	path := filepath.Join(dir, "config.json")
	data, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestRetrieve_AKMode(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "test",
		"profiles": map[string]interface{}{
			"test": map[string]interface{}{
				"mode":       "AK",
				"access-key": "AKID123",
				"secret-key": "SK456",
			},
		},
	})

	p := NewCliProvider(configPath, "test")
	v, err := p.Retrieve()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.AccessKeyID != "AKID123" || v.SecretAccessKey != "SK456" {
		t.Errorf("unexpected credentials: %+v", v)
	}
	if v.SessionToken != "" {
		t.Errorf("expected empty session token, got %q", v.SessionToken)
	}
}

func TestRetrieve_UnsupportedMode(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "custom",
		"profiles": map[string]interface{}{
			"custom": map[string]interface{}{
				"mode":       "SomeUnknownMode",
				"access-key": "AK",
				"secret-key": "SK",
			},
		},
	})

	p := NewCliProvider(configPath, "custom")
	_, err := p.Retrieve()
	if err == nil {
		t.Fatal("expected error for unsupported mode")
	}
	if !strings.Contains(err.Error(), "unsupported mode") {
		t.Errorf("expected 'unsupported mode' in error, got: %v", err)
	}
}

func TestRetrieve_RamRoleArnMode_ConstructsProvider(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "assume",
		"profiles": map[string]interface{}{
			"assume": map[string]interface{}{
				"mode":       "RamRoleArn",
				"access-key": "AK_ASSUME",
				"secret-key": "SK_ASSUME",
				"role-name":  "TestRole",
				"account-id": "12345678",
			},
		},
	})

	p := NewCliProvider(configPath, "assume")
	_, err := p.Retrieve()
	// The call will fail because STS endpoint is not reachable in test,
	// but the error should be from AssumeRole, not "unsupported mode".
	if err == nil {
		t.Fatal("expected error from STS call (no network), but got nil")
	}
	errMsg := err.Error()
	if strings.Contains(errMsg, "unsupported mode") {
		t.Errorf("RamRoleArn should not be unsupported, got: %v", err)
	}
}

func TestRetrieve_RamRoleArnMode_MissingRoleName(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "assume",
		"profiles": map[string]interface{}{
			"assume": map[string]interface{}{
				"mode":       "RamRoleArn",
				"access-key": "AK",
				"secret-key": "SK",
				"account-id": "12345678",
			},
		},
	})

	p := NewCliProvider(configPath, "assume")
	_, err := p.Retrieve()
	if err == nil {
		t.Fatal("expected error for missing role-name")
	}
	if !strings.Contains(err.Error(), "role-name") {
		t.Errorf("expected error about role-name, got: %v", err)
	}
}

func TestRetrieve_EcsRoleMode_MissingRoleName(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "ecs",
		"profiles": map[string]interface{}{
			"ecs": map[string]interface{}{
				"mode": "EcsRole",
			},
		},
	})

	p := NewCliProvider(configPath, "ecs")
	_, err := p.Retrieve()
	if err == nil {
		t.Fatal("expected error for missing role-name in EcsRole mode")
	}
	if !strings.Contains(err.Error(), "role-name") {
		t.Errorf("expected error about role-name, got: %v", err)
	}
}

func TestRetrieve_OIDCMode_MissingTokenFile(t *testing.T) {
	// Clear env vars that might interfere
	t.Setenv("VOLCENGINE_OIDC_TOKEN_FILE", "")
	t.Setenv("VOLCENGINE_OIDC_ROLE_TRN", "")

	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "oidc",
		"profiles": map[string]interface{}{
			"oidc": map[string]interface{}{
				"mode": "OIDC",
			},
		},
	})

	p := NewCliProvider(configPath, "oidc")
	_, err := p.Retrieve()
	if err == nil {
		t.Fatal("expected error for missing oidc-token-file")
	}
	if !strings.Contains(err.Error(), "oidc-token-file") {
		t.Errorf("expected error about oidc-token-file, got: %v", err)
	}
}

func TestRetrieve_RamRoleArnMode_PassesSessionToken(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "ramrole",
		"profiles": map[string]interface{}{
			"ramrole": map[string]interface{}{
				"mode":          "RamRoleArn",
				"access-key":    "STS_AK",
				"secret-key":    "STS_SK",
				"session-token": "STS_TOKEN",
				"role-name":     "TargetRole",
				"account-id":    "12345678",
			},
		},
	})

	p := NewCliProvider(configPath, "ramrole")
	_, _ = p.Retrieve()

	sp, ok := p.delegate.(*credentials.StsProvider)
	if !ok {
		t.Fatalf("expected delegate to be *credentials.StsProvider, got %T", p.delegate)
	}
	if sp.SessionToken != "STS_TOKEN" {
		t.Errorf("expected SessionToken=STS_TOKEN, got %q", sp.SessionToken)
	}
}

func writeLoginCache(t *testing.T, dir, loginSession string, accessToken interface{}, issuedAt time.Time, expiresIn int) string {
	t.Helper()
	cacheDir := filepath.Join(dir, "login", "cache")
	if err := os.MkdirAll(cacheDir, 0700); err != nil {
		t.Fatal(err)
	}
	cachePath := filepath.Join(cacheDir, fmt.Sprintf("%x.json", sha1.Sum([]byte(loginSession))))
	data, err := json.Marshal(map[string]interface{}{
		"login_session": loginSession,
		"access_token":  accessToken,
		"issued_at":     issuedAt.UTC().Format(time.RFC3339),
		"expires_in":    expiresIn,
		"token_type":    "urn:ietf:params:oauth:token-type:access_token_sts",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(cachePath, data, 0600); err != nil {
		t.Fatal(err)
	}
	return cachePath
}

func TestRetrieve_ConsoleLoginMode_ReadsObjectAccessToken(t *testing.T) {
	dir := t.TempDir()
	loginSession := "trn:iam::123456789012:login/session/test"
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "console",
		"profiles": map[string]interface{}{
			"console": map[string]interface{}{
				"mode":          "console-login",
				"login-session": loginSession,
			},
		},
	})
	writeLoginCache(t, dir, loginSession, map[string]interface{}{
		"access_key_id":     "AK_CONSOLE",
		"secret_access_key": "SK_CONSOLE",
		"session_token":     "TOKEN_CONSOLE",
	}, time.Now(), 900)

	p := NewCliProvider(configPath, "console")
	v, err := p.Retrieve()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.AccessKeyID != "AK_CONSOLE" || v.SecretAccessKey != "SK_CONSOLE" || v.SessionToken != "TOKEN_CONSOLE" {
		t.Fatalf("unexpected credentials: %+v", v)
	}
	if p.IsExpired() {
		t.Fatalf("expected console-login credentials to have a live expiration")
	}
}

func TestRetrieve_ConsoleLoginMode_ReadsStringAccessToken(t *testing.T) {
	dir := t.TempDir()
	loginSession := "trn:iam::123456789012:login/session/string"
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "console",
		"profiles": map[string]interface{}{
			"console": map[string]interface{}{
				"mode":          "console-login",
				"login-session": loginSession,
			},
		},
	})
	writeLoginCache(t, dir, loginSession, `{"access_key_id":"AK_STR","secret_access_key":"SK_STR","session_token":"TOKEN_STR"}`, time.Now(), 900)

	p := NewCliProvider(configPath, "console")
	v, err := p.Retrieve()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.AccessKeyID != "AK_STR" || v.SecretAccessKey != "SK_STR" || v.SessionToken != "TOKEN_STR" {
		t.Fatalf("unexpected credentials: %+v", v)
	}
}

func TestRetrieve_ConsoleLoginMode_ExpiredCacheReturnsError(t *testing.T) {
	dir := t.TempDir()
	loginSession := "trn:iam::123456789012:login/session/expired"
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "console",
		"profiles": map[string]interface{}{
			"console": map[string]interface{}{
				"mode":          "console-login",
				"login-session": loginSession,
			},
		},
	})
	writeLoginCache(t, dir, loginSession, map[string]interface{}{
		"access_key_id":     "AK_EXPIRED",
		"secret_access_key": "SK_EXPIRED",
		"session_token":     "TOKEN_EXPIRED",
	}, time.Now().Add(-time.Hour), 900)

	p := NewCliProvider(configPath, "console")
	_, err := p.Retrieve()
	if err == nil {
		t.Fatal("expected error for expired console-login cache")
	}
	if !strings.Contains(err.Error(), "expired") || !strings.Contains(err.Error(), "ve login") {
		t.Fatalf("expected expired ve login error, got: %v", err)
	}
}

func TestRetrieve_ConsoleLoginMode_MissingCacheFileReturnsError(t *testing.T) {
	dir := t.TempDir()
	loginSession := "trn:iam::123456789012:login/session/missing"
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "console",
		"profiles": map[string]interface{}{
			"console": map[string]interface{}{
				"mode":          "console-login",
				"login-session": loginSession,
			},
		},
	})

	p := NewCliProvider(configPath, "console")
	_, err := p.Retrieve()
	if err == nil {
		t.Fatal("expected error for missing console-login cache file")
	}
	if !strings.Contains(err.Error(), "does not exist") || !strings.Contains(err.Error(), "ve login") {
		t.Fatalf("expected missing-cache ve login error, got: %v", err)
	}
}

func TestRetrieve_ConsoleLoginMode_MissingLoginSession(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "console",
		"profiles": map[string]interface{}{
			"console": map[string]interface{}{
				"mode": "console-login",
			},
		},
	})

	p := NewCliProvider(configPath, "console")
	_, err := p.Retrieve()
	if err == nil {
		t.Fatal("expected error for missing login-session")
	}
	if !strings.Contains(err.Error(), "login-session") {
		t.Fatalf("expected login-session error, got: %v", err)
	}
}

func TestRetrieve_ConsoleLoginMode_UsesCustomCacheDirectory(t *testing.T) {
	dir := t.TempDir()
	loginSession := "trn:iam::123456789012:login/session/custom"
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "console",
		"profiles": map[string]interface{}{
			"console": map[string]interface{}{
				"mode":          "console-login",
				"login-session": loginSession,
			},
		},
	})
	customDir := filepath.Join(t.TempDir(), "cache")
	if err := os.MkdirAll(customDir, 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv(loginCacheDirectoryEnv, customDir)

	cachePath := writeLoginCache(t, t.TempDir(), loginSession, map[string]interface{}{
		"access_key_id":     "AK_CUSTOM",
		"secret_access_key": "SK_CUSTOM",
		"session_token":     "TOKEN_CUSTOM",
	}, time.Now(), 900)
	customCachePath := filepath.Join(customDir, filepath.Base(cachePath))
	if err := os.Rename(cachePath, customCachePath); err != nil {
		t.Fatal(err)
	}

	p := NewCliProvider(configPath, "console")
	v, err := p.Retrieve()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.AccessKeyID != "AK_CUSTOM" || v.SecretAccessKey != "SK_CUSTOM" || v.SessionToken != "TOKEN_CUSTOM" {
		t.Fatalf("unexpected credentials: %+v", v)
	}
}
