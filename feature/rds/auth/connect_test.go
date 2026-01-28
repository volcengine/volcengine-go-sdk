package auth

import (
	"context"
	"strings"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
)

func TestBuildAuthToken(t *testing.T) {
	tests := []struct {
		name        string
		credentials *base.Credentials
		dbUser      string
		instanceId  string
		expires     int
		wantErr     bool
		errContains string
	}{
		{
			name: "valid credentials and parameters with custom expires",
			credentials: &base.Credentials{
				AccessKeyID:     "test-access-key-id",
				SecretAccessKey: "test-secret-access-key",
				Region:          "cn-beijing",
			},
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    1800,
			wantErr:    false,
		},
		{
			name: "valid credentials with default expires",
			credentials: &base.Credentials{
				AccessKeyID:     "test-access-key-id",
				SecretAccessKey: "test-secret-access-key",
				Region:          "cn-beijing",
			},
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    0,
			wantErr:    false,
		},
		{
			name: "valid credentials with negative expires",
			credentials: &base.Credentials{
				AccessKeyID:     "test-access-key-id",
				SecretAccessKey: "test-secret-access-key",
				Region:          "cn-beijing",
			},
			dbUser:     "testuser",
			instanceId: "mysql-instance-123",
			expires:    -1,
			wantErr:    false,
		},
		{
			name:        "nil credentials",
			credentials: nil,
			dbUser:      "testuser",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "credentials provider must not be nil",
		},
		{
			name: "empty AccessKeyID",
			credentials: &base.Credentials{
				AccessKeyID:     "",
				SecretAccessKey: "test-secret-access-key",
				Region:          "cn-beijing",
			},
			dbUser:      "testuser",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "credentials provider must not be nil",
		},
		{
			name: "empty SecretAccessKey",
			credentials: &base.Credentials{
				AccessKeyID:     "test-access-key-id",
				SecretAccessKey: "",
				Region:          "cn-beijing",
			},
			dbUser:      "testuser",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "credentials provider must not be nil",
		},
		{
			name: "empty Region",
			credentials: &base.Credentials{
				AccessKeyID:     "test-access-key-id",
				SecretAccessKey: "test-secret-access-key",
				Region:          "",
			},
			dbUser:      "testuser",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "credentials provider must not be nil",
		},
		{
			name: "empty dbUser",
			credentials: &base.Credentials{
				AccessKeyID:     "test-access-key-id",
				SecretAccessKey: "test-secret-access-key",
				Region:          "cn-beijing",
			},
			dbUser:      "",
			instanceId:  "mysql-instance-123",
			expires:     900,
			wantErr:     true,
			errContains: "dbUser or instanceId must not be empty",
		},
		{
			name: "empty instanceId",
			credentials: &base.Credentials{
				AccessKeyID:     "test-access-key-id",
				SecretAccessKey: "test-secret-access-key",
				Region:          "cn-beijing",
			},
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
			token, err := BuildAuthToken(ctx, tt.credentials, tt.dbUser, tt.instanceId, tt.expires)

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
		})
	}
}
