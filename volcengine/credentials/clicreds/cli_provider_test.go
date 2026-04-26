package clicreds

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
)

func writeTempConfig(t *testing.T, dir string, config interface{}) string {
	t.Helper()
	path := filepath.Join(dir, "config.json")
	data, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
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
