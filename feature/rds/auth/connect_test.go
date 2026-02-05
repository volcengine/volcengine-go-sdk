package auth

import (
	"context"
	"strings"
	"testing"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
)

func TestBuildAuthToken(t *testing.T) {
	tests := []struct {
		name        string
		config      *volcengine.Config
		dbUser      string
		instanceId  string
		expires     int
		wantErr     bool
		errContains string
	}{
		{
			name: "valid config and parameters with custom expires",
			config: volcengine.NewConfig().
				WithCredentials(credentials.NewStaticCredentials("test-access-key-id", "test-secret-access-key", "")).
				WithRegion("cn-beijing"),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    1800,
			wantErr:    false,
		},
		{
			name: "valid config with default expires",
			config: volcengine.NewConfig().
				WithCredentials(credentials.NewStaticCredentials("test-access-key-id", "test-secret-access-key", "")).
				WithRegion("cn-beijing"),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    0,
			wantErr:    false,
		},
		{
			name: "valid config with negative expires",
			config: volcengine.NewConfig().
				WithCredentials(credentials.NewStaticCredentials("test-access-key-id", "test-secret-access-key", "")).
				WithRegion("cn-beijing"),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    -1,
			wantErr:    false,
		},
		{
			name: "valid config with DisableSSL true",
			config: volcengine.NewConfig().
				WithCredentials(credentials.NewStaticCredentials("test-access-key-id", "test-secret-access-key", "")).
				WithRegion("cn-beijing").
				WithDisableSSL(true),
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    900,
			wantErr:    false,
		},
		{
			name:        "nil config",
			config:      nil,
			dbUser:      "testuser",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "config must not be nil",
		},
		{
			name:        "nil credentials",
			config:      volcengine.NewConfig().WithRegion("cn-beijing"),
			dbUser:      "testuser",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "credentials must not be nil",
		},
		{
			name: "nil region",
			config: volcengine.NewConfig().
				WithCredentials(credentials.NewStaticCredentials("test-access-key-id", "test-secret-access-key", "")),
			dbUser:      "testuser",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "region must not be empty",
		},
		{
			name: "empty region",
			config: volcengine.NewConfig().
				WithCredentials(credentials.NewStaticCredentials("test-access-key-id", "test-secret-access-key", "")).
				WithRegion(""),
			dbUser:      "testuser",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "region must not be empty",
		},
		{
			name: "empty dbUser",
			config: volcengine.NewConfig().
				WithCredentials(credentials.NewStaticCredentials("test-access-key-id", "test-secret-access-key", "")).
				WithRegion("cn-beijing"),
			dbUser:      "",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "dbUser or instanceId must not be empty",
		},
		{
			name: "empty instanceId",
			config: volcengine.NewConfig().
				WithCredentials(credentials.NewStaticCredentials("test-access-key-id", "test-secret-access-key", "")).
				WithRegion("cn-beijing"),
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
			token, err := BuildAuthToken(ctx, tt.config, tt.dbUser, tt.instanceId, tt.expires)

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
			if tt.expires > 0 && !strings.Contains(token, "Expires=") {
				t.Error("BuildAuthToken() token should contain Expires parameter when expires > 0")
			}
			if tt.expires <= 0 && !strings.Contains(token, "X-Expires=900") {
				t.Error("BuildAuthToken() token should contain X-Expires=900 when expires <= 0")
			}

			// Verify scheme based on DisableSSL
			if tt.config != nil && tt.config.DisableSSL != nil && *tt.config.DisableSSL {
				if !strings.HasPrefix(token, "http://") {
					t.Error("BuildAuthToken() token should start with http:// when DisableSSL is true")
				}
			} else {
				if !strings.HasPrefix(token, "https://") {
					t.Error("BuildAuthToken() token should start with https:// when DisableSSL is false")
				}
			}

			// Verify regional endpoint format
			if tt.config != nil && tt.config.Region != nil {
				expectedHost := "rds-mysql." + *tt.config.Region + ".volcengineapi.com"
				if !strings.Contains(token, expectedHost) {
					t.Errorf("BuildAuthToken() token should contain regional endpoint %s", expectedHost)
				}
			}
		})
	}
}
