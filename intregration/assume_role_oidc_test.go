package intregration

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
)

func TestOIDCCredentialsProvider_Integration(t *testing.T) {
	// Create a temporary token file
	tokenDir := t.TempDir()
	tokenPath := filepath.Join(tokenDir, "oidc-token")
	tokenContent := "xxx"
	if err := os.WriteFile(tokenPath, []byte(tokenContent), 0644); err != nil {
		t.Fatalf("Failed to write test token file: %v", err)
	}

	// Start a test HTTP server that simulates STS service
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		err := r.ParseForm()
		if err != nil {
			t.Errorf("Failed to parse form: %v", err)
		}

		// Check request parameters
		if r.FormValue("RoleTrn") != "trn:iam::123456789012:oidc-provider/test-provider" {
			t.Errorf("Unexpected RoleTrn: %s", r.FormValue("RoleTrn"))
		}

		// Return a simulated STS response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		now := time.Now()
		expiration := now.Add(time.Hour).Format(time.RFC3339)
		currentTime := now.Format(time.RFC3339)
		
		response := `{
			"ResponseMetadata": {
				"RequestId": "test-request-id",
				"Action": "AssumeRoleWithOIDC",
				"Version": "2021-06-01",
				"Service": "sts",
				"Region": "cn-beijing"
			},
			"Result": {
				"Credentials": {
					"CurrentTime": "` + currentTime + `",
					"Expiration": "` + expiration + `",
					"AccessKeyId": "test-access-key",
					"SecretAccessKey": "test-secret-key",
					"SessionToken": "test-session-token"
				},
				"OIDCTokenInfo": {
					"Subject": "test-subject",
					"Issuer": "test-issuer",
					"ClientIds": ["test-client-id"],
					"ExpirationTime": "` + expiration + `",
					"IssuanceTime": "` + currentTime + `"
				},
				"AssumedRoleUser": {
					"AssumedRoleId": "test-role-id",
					"Trn": "trn:iam::123456789012:role/test-role"
				}
			}
		}`
		_, _ = w.Write([]byte(response))
	}))
	defer server.Close()

	// Set environment variables for the test
	originalEnv := map[string]string{
		"OIDC_PROVIDER_TRN":   os.Getenv("OIDC_PROVIDER_TRN"),
		"OIDC_TOKEN_FILE_PATH": os.Getenv("OIDC_TOKEN_FILE_PATH"),
		"ROLE_ARN":            os.Getenv("ROLE_ARN"),
		"ROLE_SESSION_NAME":   os.Getenv("ROLE_SESSION_NAME"),
		"POLICY":              os.Getenv("POLICY"),
		"STS_REGION":          os.Getenv("STS_REGION"),
		"ENABLE_VPC":          os.Getenv("ENABLE_VPC"),
		"STS_ENDPOINT":        os.Getenv("STS_ENDPOINT"),
	}

	// Restore original environment after test
	defer func() {
		for k, v := range originalEnv {
			if v == "" {
				os.Unsetenv(k)
			} else {
				os.Setenv(k, v)
			}
		}
	}()

	// Set test environment variables
	os.Setenv("OIDC_PROVIDER_TRN", "trn:iam::123456789012:oidc-provider/test-provider")
	os.Setenv("OIDC_TOKEN_FILE_PATH", tokenPath)
	os.Setenv("ROLE_ARN", "trn:iam::123456789012:role/test-role")
	os.Setenv("ROLE_SESSION_NAME", "test-session")
	os.Setenv("STS_ENDPOINT", server.URL)

	// Create provider from environment variables
	provider := credentials.NewOIDCCredentialsProviderFromEnv()
	
	// First retrieval
	creds, err := provider.Retrieve()
	if err != nil {
		t.Fatalf("Failed to retrieve credentials: %v", err)
	}

	// Verify credentials
	if creds.AccessKeyID != "test-access-key" {
		t.Errorf("Unexpected access key: got %s, want test-access-key", creds.AccessKeyID)
	}
	if creds.SecretAccessKey != "test-secret-key" {
		t.Errorf("Unexpected secret key: got %s, want test-secret-key", creds.SecretAccessKey)
	}
	if creds.SessionToken != "test-session-token" {
		t.Errorf("Unexpected session token: got %s, want test-session-token", creds.SessionToken)
	}
	if creds.ProviderName != "OIDC" {
		t.Errorf("Unexpected provider name: got %s, want OIDC", creds.ProviderName)
	}

	// Verify caching works
	secondCreds, err := provider.Retrieve()
	if err != nil {
		t.Fatalf("Failed to retrieve cached credentials: %v", err)
	}
	
	// Should be the same object due to caching
	if secondCreds.AccessKeyID != creds.AccessKeyID {
		t.Errorf("Expected cached credentials but got different ones")
	}

	// Check IsExpired functionality
	if provider.IsExpired() {
		t.Errorf("Credentials shouldn't be expired yet")
	}
}