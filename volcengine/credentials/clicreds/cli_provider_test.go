package clicreds

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
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

func TestRetrieve_StsTokenMode_Success(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "sts",
		"profiles": map[string]interface{}{
			"sts": map[string]interface{}{
				"mode":          "StsToken",
				"access-key":    "AKID_STS",
				"secret-key":    "SK_STS",
				"session-token": "TOKEN_STS",
			},
		},
	})

	p := NewCliProvider(configPath, "sts")
	v, err := p.Retrieve()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.AccessKeyID != "AKID_STS" || v.SecretAccessKey != "SK_STS" || v.SessionToken != "TOKEN_STS" {
		t.Errorf("unexpected credentials: %+v", v)
	}
}

func TestRetrieve_StsTokenMode_MissingToken(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "sts",
		"profiles": map[string]interface{}{
			"sts": map[string]interface{}{
				"mode":       "StsToken",
				"access-key": "AKID_STS",
				"secret-key": "SK_STS",
				// no session-token
			},
		},
	})

	p := NewCliProvider(configPath, "sts")
	_, err := p.Retrieve()
	if err == nil {
		t.Fatal("expected error for StsToken mode without session-token")
	}
}

func TestRetrieve_UnsupportedMode(t *testing.T) {
	dir := t.TempDir()
	configPath := writeTempConfig(t, dir, map[string]interface{}{
		"current": "custom",
		"profiles": map[string]interface{}{
			"custom": map[string]interface{}{
				"mode":       "RamRoleArn",
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
}
