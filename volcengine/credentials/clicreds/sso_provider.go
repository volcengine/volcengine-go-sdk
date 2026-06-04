package clicreds

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

// SsoRefreshableProvider handles SSO credential retrieval and in-memory refresh.
//
// Contract:
//   - Never writes to disk. The SSO token cache file is read on bootstrap and
//     once more as a fallback when the OAuth server rejects the in-memory
//     refresh token with invalid_grant.
//   - Fast path: if the in-memory access_token is still within TTL, the OAuth
//     refresh call is skipped; Portal GetRoleCredentials is still called to
//     obtain the STS triple.
//   - Slow path: calls OAuth refresh then Portal GetRoleCredentials; updates
//     in-memory snapshot only.
//   - On refresh_token rejection (invalid_grant), re-reads the disk cache once;
//     if the disk RT differs, retries once with the disk state; otherwise errors
//     with a "please run 've sso login'" hint.
//   - Concurrent Retrieve calls are serialized with sync.Mutex.
type SsoRefreshableProvider struct {
	tokenPath   string
	profile     *cliProfile
	profileName string
	configPath  string
	region      string

	mu          sync.Mutex
	cached      *SsoTokenCache
	creds       credentials.Value
	expiration  time.Time
	initialized bool

	// oauthClientFactory and portalClientFactory allow injection for testing.
	oauthClientFactory  func(region string) OAuthClientAPI
	portalClientFactory func(region string) PortalClientAPI
}

func defaultSsoOAuthFactory(region string) OAuthClientAPI {
	return NewOAuthClient(&OAuthClientConfig{Region: region})
}

func defaultSsoPortalFactory(region string) PortalClientAPI {
	return NewPortalClient(&PortalClientConfig{Region: region})
}

// newSsoRefreshableProvider constructs an SsoRefreshableProvider with production defaults.
func newSsoRefreshableProvider(tokenPath string, profile *cliProfile, profileName, configPath, region string) *SsoRefreshableProvider {
	return &SsoRefreshableProvider{
		tokenPath:           tokenPath,
		profile:             profile,
		profileName:         profileName,
		configPath:          configPath,
		region:              region,
		oauthClientFactory:  defaultSsoOAuthFactory,
		portalClientFactory: defaultSsoPortalFactory,
	}
}

// Retrieve returns current valid credentials, refreshing if needed.
func (p *SsoRefreshableProvider) Retrieve() (credentials.Value, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.initialized || p.isExpiredLocked() {
		if err := p.refreshLocked(); err != nil {
			return credentials.Value{ProviderName: CliProviderName}, err
		}
	}
	return p.creds, nil
}

// IsExpired reports whether the next Retrieve will require a refresh.
func (p *SsoRefreshableProvider) IsExpired() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if !p.initialized {
		return true
	}
	return p.isExpiredLocked()
}

func (p *SsoRefreshableProvider) isExpiredLocked() bool {
	if p.expiration.IsZero() {
		return true
	}
	// 60-second safety buffer to avoid handing out credentials near expiry.
	return !time.Now().Before(p.expiration.Add(-time.Minute))
}

// refreshLocked is the core refresh flow; caller must hold p.mu.
//
// Flow:
//  1. Bootstrap from disk on first call.
//  2. Fast path: if the in-memory access_token is still valid, use it.
//  3. Slow path: OAuth refresh + Portal GetRoleCredentials.
//  4. On invalid_grant, re-read disk once; if disk RT differs, retry once.
func (p *SsoRefreshableProvider) refreshLocked() error {
	if p.cached == nil {
		cache, err := loadSsoTokenCache(p.tokenPath)
		if err != nil {
			return err
		}
		p.cached = cache
	}

	// Fast path: access_token still within TTL.
	if !p.isAccessTokenExpiredLocked() {
		return p.fetchRoleCredentialsLocked()
	}

	// Slow path: refresh the access token.
	err := p.refreshAccessTokenLocked(context.Background())
	if err == nil {
		return p.fetchRoleCredentialsLocked()
	}

	// Fallback: only for explicit invalid_grant rejections.
	if !isSsoRefreshTokenInvalidErr(err) {
		return err
	}

	diskCache, diskErr := loadSsoTokenCache(p.tokenPath)
	if diskErr != nil {
		return wrapSsoRefreshErr(err, diskErr)
	}
	if strings.TrimSpace(diskCache.RefreshToken) == "" ||
		diskCache.RefreshToken == p.cached.RefreshToken {
		return volcengineerr.New(
			"CliSsoTokenRefreshExpired",
			"SSO refresh token has been rejected by the authorization server; please run 've sso login' to re-authenticate",
			err,
		)
	}
	// Retry once with disk state.
	p.cached = diskCache
	if !p.isAccessTokenExpiredLocked() {
		return p.fetchRoleCredentialsLocked()
	}
	if err2 := p.refreshAccessTokenLocked(context.Background()); err2 != nil {
		return volcengineerr.New(
			"CliSsoTokenRefreshExpired",
			"SSO refresh token rejected; reloaded cache from disk but new refresh token also failed; please run 've sso login'",
			err2,
		)
	}
	return p.fetchRoleCredentialsLocked()
}

// isAccessTokenExpiredLocked returns true when the in-memory access_token is
// past its expiry buffer (60 seconds). Caller must hold p.mu.
func (p *SsoRefreshableProvider) isAccessTokenExpiredLocked() bool {
	if p.cached == nil || strings.TrimSpace(p.cached.AccessToken) == "" {
		return true
	}
	exp, err := parseTokenExpiration(p.cached.ExpiresAt)
	if err != nil {
		return true
	}
	// Apply 60-second safety buffer, matching the plan document and
	// the console-login provider behaviour.
	return !time.Now().Before(exp.Add(-time.Minute))
}

// refreshAccessTokenLocked calls OAuth /token with the in-memory refresh_token.
// Caller must hold p.mu.
func (p *SsoRefreshableProvider) refreshAccessTokenLocked(ctx context.Context) error {
	if p.cached == nil || strings.TrimSpace(p.cached.RefreshToken) == "" {
		return volcengineerr.New(
			"CliSsoTokenRefreshMissing",
			fmt.Sprintf("SSO token cache %s did not contain refresh token; please run 've sso login' to re-authenticate", p.tokenPath),
			nil,
		)
	}
	refreshExpired, err := isRefreshTokenExpired(p.cached)
	if err != nil {
		return volcengineerr.New(
			"CliSsoTokenRefreshExpiresParse",
			fmt.Sprintf("failed to parse refresh token expiration in %s; please run 've sso login' to re-authenticate", p.tokenPath),
			err,
		)
	}
	if refreshExpired {
		return volcengineerr.New(
			"CliSsoTokenRefreshExpired",
			fmt.Sprintf("SSO refresh token in %s has expired; please run 've sso login' to re-authenticate", p.tokenPath),
			nil,
		)
	}
	if strings.TrimSpace(p.cached.ClientId) == "" || strings.TrimSpace(p.cached.ClientSecret) == "" {
		return volcengineerr.New(
			"CliSsoTokenClientMissing",
			fmt.Sprintf("SSO token cache %s did not contain client id/secret; please run 've sso login' to re-authenticate", p.tokenPath),
			nil,
		)
	}

	oauthClient := p.oauthClientFactory(p.region)
	tokenResp, err := oauthClient.CreateToken(ctx, &CreateTokenRequest{
		GrantType:    "refresh_token",
		ClientID:     p.cached.ClientId,
		ClientSecret: p.cached.ClientSecret,
		RefreshToken: p.cached.RefreshToken,
	})
	if err != nil {
		// Pass invalid_grant through verbatim so refreshLocked can detect it
		// and trigger the disk-reload fallback.  All other errors get wrapped
		// with the stable error code and actionable hint.
		if isSsoRefreshTokenInvalidErr(err) {
			return err
		}
		return volcengineerr.New(
			"CliSsoTokenRefreshFailed",
			"failed to refresh SSO access token; please run 've sso login' to re-authenticate",
			err,
		)
	}
	if tokenResp == nil || strings.TrimSpace(tokenResp.AccessToken) == "" {
		return volcengineerr.New(
			"CliSsoTokenRefreshEmpty",
			"SSO refresh token response did not include access token; please run 've sso login' to re-authenticate",
			nil,
		)
	}
	if tokenResp.ExpiresIn <= 0 {
		return volcengineerr.New(
			"CliSsoTokenRefreshExpiresIn",
			"SSO refresh token response did not include expires_in; please run 've sso login' to re-authenticate",
			nil,
		)
	}

	p.cached.AccessToken = tokenResp.AccessToken
	if strings.TrimSpace(tokenResp.RefreshToken) != "" {
		p.cached.RefreshToken = tokenResp.RefreshToken
	}
	p.cached.ExpiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second).UTC().Format(time.RFC3339)
	return nil
}

// fetchRoleCredentialsLocked calls Portal GetRoleCredentials and stores results.
// Caller must hold p.mu.
func (p *SsoRefreshableProvider) fetchRoleCredentialsLocked() error {
	accountID := strings.TrimSpace(p.profile.AccountId)
	if accountID == "" {
		return volcengineerr.New(
			"CliConfigAccountIDMissing",
			fmt.Sprintf("cli config profile %s in %s did not contain account-id", p.profileName, p.configPath),
			nil,
		)
	}
	roleName := strings.TrimSpace(p.profile.RoleName)
	if roleName == "" {
		return volcengineerr.New(
			"CliConfigRoleNameMissing",
			fmt.Sprintf("cli config profile %s in %s did not contain role-name", p.profileName, p.configPath),
			nil,
		)
	}

	portalClient := p.portalClientFactory(p.region)
	resp, err := portalClient.GetRoleCredentials(context.Background(), &GetRoleCredentialsRequest{
		AccessToken: p.cached.AccessToken,
		AccountID:   accountID,
		RoleName:    roleName,
	})
	if err != nil {
		return volcengineerr.New(
			"CliSsoPortalCredentials",
			"failed to get SSO role credentials; please run 've sso login' to re-authenticate",
			err,
		)
	}

	creds := resp.RoleCredentials
	if strings.TrimSpace(creds.AccessKeyID) == "" || strings.TrimSpace(creds.SecretAccessKey) == "" {
		return volcengineerr.New(
			"CliSsoPortalCredentialsEmpty",
			"SSO portal credentials response did not include access key or secret key",
			nil,
		)
	}

	exp := time.Time{}
	if creds.Expiration > 0 {
		exp = UnixTimestampToTime(creds.Expiration)
	}

	p.creds = credentials.Value{
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
		ProviderName:    CliProviderName,
	}
	p.expiration = exp
	p.initialized = true
	return nil
}

// isSsoRefreshTokenInvalidErr checks whether err is an OAuth invalid_grant response.
func isSsoRefreshTokenInvalidErr(err error) bool {
	if err == nil {
		return false
	}
	type unwrapper interface {
		Unwrap() error
	}
	cur := err
	for cur != nil {
		if apiErr, ok := cur.(*OAuthAPIError); ok {
			return apiErr.StatusCode == 400 &&
				strings.EqualFold(apiErr.Response.Error, "invalid_grant")
		}
		uw, ok := cur.(unwrapper)
		if !ok {
			break
		}
		cur = uw.Unwrap()
	}
	return false
}

// wrapSsoRefreshErr combines a refresh error and a disk-reload error into one.
func wrapSsoRefreshErr(refreshErr, diskErr error) error {
	return volcengineerr.New(
		"CliSsoTokenRefreshFailed",
		fmt.Sprintf("SSO refresh token rejected (%v); failed to reload cache from disk: %v; please run 've sso login'", refreshErr, diskErr),
		refreshErr,
	)
}
