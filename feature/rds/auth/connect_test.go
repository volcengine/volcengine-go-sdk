package auth

import (
	"context"
	"strings"
	"testing"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func newTestSession(t *testing.T, ak, sk, region string, disableSSL *bool, endpoint ...string) *session.Session {
	t.Helper()
	cfg := volcengine.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithRegion(region)
	if disableSSL != nil {
		cfg = cfg.WithDisableSSL(*disableSSL)
	}
	if len(endpoint) > 0 {
		cfg = cfg.WithEndpoint(endpoint[0])
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		t.Fatalf("failed to create session: %v", err)
	}
	return sess
}

func TestBuildAuthToken(t *testing.T) {
	tests := []struct {
		name         string
		sess         *session.Session
		dbUser       string
		instanceId   string
		expires      int
		wantErr      bool
		errContains  string
		disableSSL   bool
		region       string
		expectedHost string
	}{
		{
			name:       "valid config and parameters with custom expires",
			sess:       newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    1800,
			wantErr:    false,
			region:     "cn-beijing",
		},
		{
			name:       "valid config with default expires",
			sess:       newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    0,
			wantErr:    false,
			region:     "cn-beijing",
		},
		{
			name:       "valid config with negative expires",
			sess:       newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    -1,
			wantErr:    false,
			region:     "cn-beijing",
		},
		{
			name:       "valid config with DisableSSL true",
			sess:       newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", volcengine.Bool(true)),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    900,
			wantErr:    false,
			disableSSL: true,
			region:     "cn-beijing",
		},
		{
			name:         "valid config with custom endpoint",
			sess:         newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil, "custom-rds.example.com"),
			dbUser:       "testuser",
			instanceId:   "mysql-instance-123",
			expires:      900,
			wantErr:      false,
			region:       "cn-beijing",
			expectedHost: "custom-rds.example.com",
		},
		{
			name:        "nil session",
			sess:        nil,
			dbUser:      "testuser",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "session must not be nil",
		},
		{
			name:        "empty dbUser",
			sess:        newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil),
			dbUser:      "",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "dbUser or instanceId must not be empty",
		},
		{
			name:        "empty instanceId",
			sess:        newTestSession(t, "test-access-key-id", "test-secret-access-key", "cn-beijing", nil),
			dbUser:      "testuser",
			instanceId:  "",
			expires:     900,
			wantErr:     true,
			errContains: "dbUser or instanceId must not be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			token, err := BuildAuthToken(ctx, tt.sess, tt.dbUser, tt.instanceId, tt.expires)

			if tt.wantErr {
				if err == nil {
					t.Errorf("BuildAuthToken() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("BuildAuthToken() error = %v, should contain %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("BuildAuthToken() unexpected error = %v", err)
				return
			}

			if token == "" {
				t.Error("BuildAuthToken() returned empty token")
			}

			// Verify token contains expected parameters
			if !strings.Contains(token, "Action=ConnectDatabase") {
				t.Error("BuildAuthToken() token should contain Action=ConnectDatabase")
			}
			if !strings.Contains(token, "DBUser="+tt.dbUser) {
				t.Errorf("BuildAuthToken() token should contain DBUser=%s", tt.dbUser)
			}
			if tt.expires > 0 && !strings.Contains(token, "X-Expires=") {
				t.Error("BuildAuthToken() token should contain X-Expires parameter when expires > 0")
			}
			if tt.expires <= 0 && !strings.Contains(token, "X-Expires=900") {
				t.Error("BuildAuthToken() token should contain X-Expires=900 when expires <= 0")
			}

			// Verify scheme based on DisableSSL
			if tt.disableSSL {
				if !strings.HasPrefix(token, "http://") {
					t.Error("BuildAuthToken() token should start with http:// when DisableSSL is true")
				}
			} else {
				if !strings.HasPrefix(token, "https://") {
					t.Error("BuildAuthToken() token should start with https:// when DisableSSL is false")
				}
			}

			// Verify endpoint host
			expectedHost := tt.expectedHost
			if expectedHost == "" && tt.region != "" {
				expectedHost = "rds-mysql." + tt.region + ".volcengineapi.com"
			}
			if expectedHost != "" && !strings.Contains(token, expectedHost) {
				t.Errorf("BuildAuthToken() token should contain endpoint host %s, got %s", expectedHost, token)
			}
		})
	}
}
