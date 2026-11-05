package clicreds

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/endpoints"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

type failIfCalledOAuthClient struct{ t *testing.T }

func (c failIfCalledOAuthClient) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error) {
	c.t.Fatalf("OAuth client must not be used with an invalid region")
	return nil, nil
}

type failIfCalledPortalClient struct{ t *testing.T }

func (c failIfCalledPortalClient) GetRoleCredentials(context.Context, *GetRoleCredentialsRequest) (*GetRoleCredentialsResponse, error) {
	c.t.Fatalf("Portal client must not be used with an invalid region")
	return nil, nil
}

func writeSsoCacheForRegionTest(t *testing.T, path string) {
	t.Helper()
	data, err := json.Marshal(&SsoTokenCache{
		AccessToken:           "access",
		ExpiresAt:             time.Now().Add(time.Hour).UTC().Format(time.RFC3339),
		ClientId:              "client-id",
		ClientSecret:          "client-secret",
		ClientSecretExpiresAt: time.Now().Add(time.Hour).Unix(),
		RefreshToken:          "refresh",
	})
	if err != nil {
		t.Fatalf("marshal sso cache: %v", err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatalf("write sso cache: %v", err)
	}
}

func TestSsoInvalidRegionIsRejectedWithoutNetwork(t *testing.T) {
	regions := []string{
		"@127.0.0.1:9/x#",
		"127.0.0.1:9/#",
		"evil.com/#",
		"a.b",
		"cn beijing",
		"1-region",
		"region-",
	}
	for _, region := range regions {
		region := region
		t.Run(region, func(t *testing.T) {
			dir := t.TempDir()
			tokenPath := filepath.Join(dir, "token.json")
			writeSsoCacheForRegionTest(t, tokenPath)

			profile := &cliProfile{AccountId: "12345678", RoleName: "RoleA"}
			p := newSsoRefreshableProvider(tokenPath, profile, "default", filepath.Join(dir, "config.json"), region)
			p.oauthClientFactory = func(string) OAuthClientAPI { return failIfCalledOAuthClient{t} }
			p.portalClientFactory = func(string) PortalClientAPI { return failIfCalledPortalClient{t} }

			_, err := p.Retrieve()
			if err == nil {
				t.Fatalf("Retrieve with region %q: expected error", region)
			}
			ve, ok := err.(volcengineerr.Error)
			if !ok || ve.Code() != endpoints.ErrCodeInvalidRegion {
				t.Fatalf("Retrieve with region %q: error = %v, want code %s", region, err, endpoints.ErrCodeInvalidRegion)
			}
		})
	}
}

func TestSsoValidRegionStillReachesClients(t *testing.T) {
	for _, region := range []string{"", "cn-beijing", " cn-beijing "} {
		region := region
		t.Run(region, func(t *testing.T) {
			dir := t.TempDir()
			tokenPath := filepath.Join(dir, "token.json")
			writeSsoCacheForRegionTest(t, tokenPath)

			profile := &cliProfile{AccountId: "12345678", RoleName: "RoleA"}
			p := newSsoRefreshableProvider(tokenPath, profile, "default", filepath.Join(dir, "config.json"), region)
			portalCalled := false
			p.portalClientFactory = func(string) PortalClientAPI {
				return portalFunc(func(context.Context, *GetRoleCredentialsRequest) (*GetRoleCredentialsResponse, error) {
					portalCalled = true
					return &GetRoleCredentialsResponse{RoleCredentials: RoleCredentials{
						AccessKeyID:     "AK",
						SecretAccessKey: "SK",
						Expiration:      time.Now().Add(time.Hour).Unix(),
					}}, nil
				})
			}

			v, err := p.Retrieve()
			if err != nil {
				t.Fatalf("Retrieve: %v", err)
			}
			if !portalCalled || v.AccessKeyID != "AK" {
				t.Fatalf("portal not reached or wrong creds: called=%v value=%+v", portalCalled, v)
			}
		})
	}
}

type portalFunc func(context.Context, *GetRoleCredentialsRequest) (*GetRoleCredentialsResponse, error)

func (f portalFunc) GetRoleCredentials(ctx context.Context, r *GetRoleCredentialsRequest) (*GetRoleCredentialsResponse, error) {
	return f(ctx, r)
}
