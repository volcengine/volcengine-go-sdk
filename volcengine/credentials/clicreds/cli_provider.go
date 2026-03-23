package clicreds

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/volcengine/volcengine-go-sdk/internal/shareddefaults"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

const (
	// CliProviderName provides a name of CLI config provider.
	CliProviderName = "CliProvider"
	modeSSO         = "sso"
	modeAK          = "ak"
	defaultRegion   = "cn-beijing"
)

type cliConfigure struct {
	Current     string                 `json:"current"`
	Profiles    map[string]*cliProfile `json:"profiles"`
	SsoSessions map[string]*SsoSession `json:"sso-session"`
}

type cliProfile struct {
	Mode           string `json:"mode"`
	AccessKey      string `json:"access-key"`
	SecretKey      string `json:"secret-key"`
	SessionToken   string `json:"session-token"`
	StsExpiration  int64  `json:"sts-expiration"`
	SsoSessionName string `json:"sso-session-name"`
	AccountId      string `json:"account-id"`
	RoleName       string `json:"role-name"`
}

type SsoSession struct {
	Name               string   `json:"name"`
	StartURL           string   `json:"start-url"`
	Region             string   `json:"region"`
	RegistrationScopes []string `json:"registration-scopes,omitempty"`
}

type SsoTokenCache struct {
	StartURL              string `json:"start_url"`
	SessionName           string `json:"session_name"`
	AccessToken           string `json:"access_token"`
	ExpiresAt             string `json:"expires_at"`
	ClientId              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	ClientIdIssuedAt      int64  `json:"client_id_issued_at,omitempty"`
	ClientSecretExpiresAt int64  `json:"client_secret_expires_at,omitempty"`
	RefreshToken          string `json:"refresh_token,omitempty"`
	Region                string `json:"region"`
}

// CliProvider retrieves credentials from volcengine-cli's config file
// (~/.volcengine/config.json).
//
// It supports reading access-key/secret-key/session-token, and will treat
// sts-expiration (Unix timestamp in seconds or milliseconds) as the credential's expiration time if present.
type CliProvider struct {
	credentials.Expiry

	// filePath is path to ~/.volcengine/config.json. If empty, it will use the
	// default path derived from the current user's home directory.
	configPath string

	// Directory for caching SSO token information
	cacheDir string

	// Profile is the profile name in the config. If empty, it will first try
	// VOLCSTACK_PROFILE, then use config.current.
	profile string

	retrieved     bool
	hasExpiration bool
}

// NewCliCredentials returns a pointer to a new Credentials object wrapping the
// volcengine-cli config provider.
func NewCliCredentials(configPath, profile string) *credentials.Credentials {
	cacheDir := ""
	if configPath == "" {
		home := shareddefaults.UserHomeDir()
		if home != "" {
			configPath = filepath.Join(home, ".volcengine", "config.json")
			cacheDir = filepath.Join(home, ".volcengine", "sso", "cache")
		}
	} else {
		cacheDir = filepath.Join(filepath.Dir(configPath), "sso", "cache")
	}
	return credentials.NewExpireAbleCredentials(&CliProvider{
		configPath: configPath,
		cacheDir:   cacheDir,
		profile:    profile,
		Expiry:     credentials.Expiry{},
	})
}

func (p *CliProvider) Retrieve() (credentials.Value, error) {
	p.retrieved = false
	p.hasExpiration = false
	p.SetExpiration(time.Time{}, 0)

	configPath, err := p.getConfigPath()
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, err
	}

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliConfigLoad",
			fmt.Sprintf("failed to load cli config file %s", configPath),
			err,
		)
	}

	cfg := &cliConfigure{}
	if len(strings.TrimSpace(string(b))) != 0 {
		if err := json.Unmarshal(b, cfg); err != nil {
			return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
				"CliConfigUnmarshal",
				fmt.Sprintf("failed to unmarshal cli config file %s", configPath),
				err,
			)
		}
	}

	profileName := p.getProfile(cfg)
	if cfg.Profiles == nil || len(cfg.Profiles) == 0 {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliConfigNoProfiles",
			fmt.Sprintf("cli config file %s did not contain any profiles", configPath),
			nil,
		)
	}

	profile, ok := cfg.Profiles[profileName]
	if !ok || profile == nil {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliConfigProfileNotFound",
			fmt.Sprintf("cli config file %s did not contain profile %s", configPath, profileName),
			nil,
		)
	}

	mode := strings.ToLower(strings.TrimSpace(profile.Mode))
	switch mode {
	case "", modeAK:
		return p.retrieveAK(profile, profileName, configPath)
	case modeSSO:
		return p.retrieveSSO(profile, profileName, configPath, cfg)
	default:
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliConfigModeInvalid",
			fmt.Sprintf("cli config profile %s in %s contained unsupported mode %q", profileName, configPath, profile.Mode),
			nil,
		)
	}
}

func (p *CliProvider) retrieveAK(profile *cliProfile, profileName, configPath string) (credentials.Value, error) {
	if len(profile.AccessKey) == 0 {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliConfigAccessKey",
			fmt.Sprintf("cli config profile %s in %s did not contain access-key", profileName, configPath),
			nil,
		)
	}
	if len(profile.SecretKey) == 0 {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliConfigSecretKey",
			fmt.Sprintf("cli config profile %s in %s did not contain secret-key", profileName, configPath),
			nil,
		)
	}

	p.retrieved = true
	return credentials.Value{
		AccessKeyID:     profile.AccessKey,
		SecretAccessKey: profile.SecretKey,
		SessionToken:    profile.SessionToken,
		ProviderName:    CliProviderName,
	}, nil
}

func (p *CliProvider) retrieveSSO(profile *cliProfile, profileName, configPath string, cfg *cliConfigure) (credentials.Value, error) {
	if value, ok, err := p.useStsCredentialsIfValid(profile, profileName, configPath); err != nil {
		return value, err
	} else if ok {
		return value, nil
	}

	sessionName, startURL, region, err := p.resolveSsoSession(profile, profileName, configPath, cfg)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, err
	}
	tokenPath, err := p.resolveTokenCachePath(startURL, sessionName, profileName, configPath)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, err
	}

	tokenCache, err := loadSsoTokenCache(tokenPath)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, err
	}
	accessToken := strings.TrimSpace(tokenCache.AccessToken)
	if accessToken == "" {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliSsoTokenAccessMissing",
			fmt.Sprintf("sso token cache file %s did not contain access token", tokenPath),
			nil,
		)
	}
	expired, err := isTokenExpired(tokenCache.ExpiresAt)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliSsoTokenExpiresParse",
			fmt.Sprintf("failed to parse access token expiration in %s", tokenPath),
			err,
		)
	}
	if !expired {
		return p.getRoleCredentials(accessToken, profile, profileName, configPath, region)
	}

	refreshedToken, err := p.refreshAccessToken(tokenCache, tokenPath, region)
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, err
	}

	return p.getRoleCredentials(refreshedToken, profile, profileName, configPath, region)
}

func (p *CliProvider) useStsCredentialsIfValid(profile *cliProfile, profileName, configPath string) (credentials.Value, bool, error) {
	expTimestamp := profile.StsExpiration
	if expTimestamp == 0 {
		return credentials.Value{ProviderName: CliProviderName}, false, nil
	}
	if expTimestamp < 0 {
		return credentials.Value{ProviderName: CliProviderName}, false, volcengineerr.New(
			"CliConfigExpiration",
			fmt.Sprintf("cli config profile %s in %s contained invalid sts-expiration", profileName, configPath),
			nil,
		)
	}

	exp := UnixTimestampToTime(expTimestamp)
	if time.Now().After(exp) {
		return credentials.Value{ProviderName: CliProviderName}, false, nil
	}
	if len(profile.AccessKey) == 0 {
		return credentials.Value{ProviderName: CliProviderName}, false, volcengineerr.New(
			"CliConfigAccessKey",
			fmt.Sprintf("cli config profile %s in %s did not contain access-key", profileName, configPath),
			nil,
		)
	}
	if len(profile.SecretKey) == 0 {
		return credentials.Value{ProviderName: CliProviderName}, false, volcengineerr.New(
			"CliConfigSecretKey",
			fmt.Sprintf("cli config profile %s in %s did not contain secret-key", profileName, configPath),
			nil,
		)
	}

	p.hasExpiration = true
	p.SetExpiration(exp, time.Minute)
	p.retrieved = true
	return credentials.Value{
		AccessKeyID:     profile.AccessKey,
		SecretAccessKey: profile.SecretKey,
		SessionToken:    profile.SessionToken,
		ProviderName:    CliProviderName,
	}, true, nil
}

func (p *CliProvider) resolveSsoSession(profile *cliProfile, profileName, configPath string, cfg *cliConfigure) (string, string, string, error) {
	sessionName := strings.TrimSpace(profile.SsoSessionName)
	if sessionName == "" {
		return "", "", "", volcengineerr.New(
			"CliConfigSsoSessionNameMissing",
			fmt.Sprintf("cli config profile %s in %s did not contain sso-session-name", profileName, configPath),
			nil,
		)
	}
	if cfg == nil || len(cfg.SsoSessions) == 0 {
		return "", "", "", volcengineerr.New(
			"CliConfigSsoSessionsMissing",
			fmt.Sprintf("cli config file %s did not contain any sso-sessions", configPath),
			nil,
		)
	}
	session := cfg.SsoSessions[sessionName]
	if session == nil {
		return "", "", "", volcengineerr.New(
			"CliConfigSsoSessionNotFound",
			fmt.Sprintf("cli config file %s did not contain sso session %s", configPath, sessionName),
			nil,
		)
	}
	startURL := strings.TrimSpace(session.StartURL)
	if startURL == "" {
		return "", "", "", volcengineerr.New(
			"CliConfigSsoStartURLMissing",
			fmt.Sprintf("sso session %s in %s did not contain start-url", sessionName, configPath),
			nil,
		)
	}

	region := strings.TrimSpace(session.Region)
	if region == "" {
		region = defaultRegion
	}
	return sessionName, startURL, region, nil
}

func (p *CliProvider) resolveTokenCachePath(startURL, sessionName, profileName, configPath string) (string, error) {
	if p.cacheDir == "" {
		return "", volcengineerr.New(
			"CliSsoCacheDirMissing",
			fmt.Sprintf("cli config profile %s in %s did not resolve cache directory", profileName, configPath),
			nil,
		)
	}

	tokenPath := filepath.Join(p.cacheDir, tokenCacheFileName(startURL, sessionName))
	if _, err := os.Stat(tokenPath); err != nil {
		if os.IsNotExist(err) {
			return "", volcengineerr.New(
				"CliSsoTokenCacheMissing",
				fmt.Sprintf("sso token cache file %s does not exist", tokenPath),
				err,
			)
		}
		return "", volcengineerr.New(
			"CliSsoTokenCacheStat",
			fmt.Sprintf("failed to stat sso token cache file %s", tokenPath),
			err,
		)
	}

	return tokenPath, nil
}

func (p *CliProvider) refreshAccessToken(tokenCache *SsoTokenCache, tokenPath, region string) (string, error) {
	if strings.TrimSpace(tokenCache.RefreshToken) == "" {
		return "", volcengineerr.New(
			"CliSsoTokenRefreshMissing",
			fmt.Sprintf("sso token cache file %s did not contain refresh token", tokenPath),
			nil,
		)
	}
	refreshExpired, err := isRefreshTokenExpired(tokenCache)
	if err != nil {
		return "", volcengineerr.New(
			"CliSsoTokenRefreshExpiresParse",
			fmt.Sprintf("failed to parse refresh token expiration in %s", tokenPath),
			err,
		)
	}
	if refreshExpired {
		return "", volcengineerr.New(
			"CliSsoTokenRefreshExpired",
			fmt.Sprintf("refresh token in %s has expired", tokenPath),
			nil,
		)
	}
	if strings.TrimSpace(tokenCache.ClientId) == "" || strings.TrimSpace(tokenCache.ClientSecret) == "" {
		return "", volcengineerr.New(
			"CliSsoTokenClientMissing",
			fmt.Sprintf("sso token cache file %s did not contain client id/secret", tokenPath),
			nil,
		)
	}

	oauthClient := NewOAuthClient(&OAuthClientConfig{Region: region})
	tokenResp, err := oauthClient.CreateToken(context.Background(), &CreateTokenRequest{
		GrantType:    "refresh_token",
		ClientID:     tokenCache.ClientId,
		ClientSecret: tokenCache.ClientSecret,
		RefreshToken: tokenCache.RefreshToken,
	})
	if err != nil {
		return "", volcengineerr.New(
			"CliSsoTokenRefreshFailed",
			"failed to refresh access token",
			err,
		)
	}
	if tokenResp == nil || strings.TrimSpace(tokenResp.AccessToken) == "" {
		return "", volcengineerr.New(
			"CliSsoTokenRefreshEmpty",
			"refresh token response did not include access token",
			nil,
		)
	}
	if tokenResp.ExpiresIn <= 0 {
		return "", volcengineerr.New(
			"CliSsoTokenRefreshExpiresIn",
			"refresh token response did not include expires_in",
			nil,
		)
	}

	tokenCache.AccessToken = tokenResp.AccessToken
	if strings.TrimSpace(tokenResp.RefreshToken) != "" {
		tokenCache.RefreshToken = tokenResp.RefreshToken
	}
	tokenCache.ExpiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second).UTC().Format(time.RFC3339)
	if err := saveSsoTokenCache(tokenPath, tokenCache); err != nil {
		return "", err
	}

	return tokenCache.AccessToken, nil
}

func (p *CliProvider) getRoleCredentials(accessToken string, profile *cliProfile, profileName, configPath, region string) (credentials.Value, error) {
	accountID := strings.TrimSpace(profile.AccountId)
	if accountID == "" {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliConfigAccountIDMissing",
			fmt.Sprintf("cli config profile %s in %s did not contain account-id", profileName, configPath),
			nil,
		)
	}
	roleName := strings.TrimSpace(profile.RoleName)
	if roleName == "" {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliConfigRoleNameMissing",
			fmt.Sprintf("cli config profile %s in %s did not contain role-name", profileName, configPath),
			nil,
		)
	}
	portalClient := NewPortalClient(&PortalClientConfig{Region: region})
	resp, err := portalClient.GetRoleCredentials(context.Background(), &GetRoleCredentialsRequest{
		AccessToken: accessToken,
		AccountID:   accountID,
		RoleName:    roleName,
	})
	if err != nil {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliSsoPortalCredentials",
			"failed to get role credentials",
			err,
		)
	}

	creds := resp.RoleCredentials
	if strings.TrimSpace(creds.AccessKeyID) == "" || strings.TrimSpace(creds.SecretAccessKey) == "" {
		return credentials.Value{ProviderName: CliProviderName}, volcengineerr.New(
			"CliSsoPortalCredentialsEmpty",
			"portal credentials response did not include access key or secret key",
			nil,
		)
	}
	if creds.Expiration > 0 {
		exp := UnixTimestampToTime(creds.Expiration)
		p.hasExpiration = true
		p.SetExpiration(exp, time.Minute)
	}

	p.retrieved = true
	return credentials.Value{
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
		ProviderName:    CliProviderName,
	}, nil
}

func loadSsoTokenCache(path string) (*SsoTokenCache, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, volcengineerr.New(
			"CliSsoTokenCacheLoad",
			fmt.Sprintf("failed to load sso token cache file %s", path),
			err,
		)
	}
	if len(strings.TrimSpace(string(data))) == 0 {
		return nil, volcengineerr.New(
			"CliSsoTokenCacheEmpty",
			fmt.Sprintf("sso token cache file %s was empty", path),
			nil,
		)
	}

	cache := &SsoTokenCache{}
	if err := json.Unmarshal(data, cache); err != nil {
		return nil, volcengineerr.New(
			"CliSsoTokenCacheUnmarshal",
			fmt.Sprintf("failed to unmarshal sso token cache file %s", path),
			err,
		)
	}
	return cache, nil
}

func saveSsoTokenCache(path string, cache *SsoTokenCache) error {
	if err := writeJSONFileAtomic(path, 0600, cache); err != nil {
		return volcengineerr.New(
			"CliSsoTokenCacheWrite",
			fmt.Sprintf("failed to write sso token cache file %s", path),
			err,
		)
	}
	return nil
}

func isTokenExpired(expiresAt string) (bool, error) {
	exp, err := parseTokenExpiration(expiresAt)
	if err != nil {
		return true, err
	}
	return time.Now().After(exp), nil
}

func parseTokenExpiration(expiresAt string) (time.Time, error) {
	value := strings.TrimSpace(expiresAt)
	if value == "" {
		return time.Time{}, fmt.Errorf("expires_at is empty")
	}
	exp, err := time.Parse(time.RFC3339, value)
	if err == nil {
		return exp, nil
	}
	exp, err = time.Parse(time.RFC3339Nano, value)
	if err == nil {
		return exp, nil
	}
	return time.Time{}, fmt.Errorf("expires_at is invalid: %s", value)
}

func isRefreshTokenExpired(cache *SsoTokenCache) (bool, error) {
	if cache == nil {
		return true, fmt.Errorf("token cache is nil")
	}
	if cache.ClientSecretExpiresAt <= 0 {
		return true, fmt.Errorf("refresh token expiration is missing")
	}
	exp := UnixTimestampToTime(cache.ClientSecretExpiresAt)
	if exp.IsZero() {
		return true, fmt.Errorf("refresh token expiration is invalid")
	}
	return time.Now().After(exp), nil
}

func (p *CliProvider) IsExpired() bool {
	if !p.retrieved {
		return true
	}
	if p.hasExpiration {
		return p.Expiry.IsExpired()
	}
	return false
}

func (p *CliProvider) getConfigPath() (string, error) {
	if len(p.configPath) != 0 {
		return p.configPath, nil
	}

	if env := os.Getenv("VOLCENGINE_CLI_CONFIG_FILE"); env != "" {
		p.configPath = env
		return p.configPath, nil
	}

	home := shareddefaults.UserHomeDir()
	if len(home) == 0 {
		return "", credentials.ErrSharedCredentialsHomeNotFound
	}

	p.configPath = filepath.Join(home, ".volcengine", "config.json")
	return p.configPath, nil
}

func (p *CliProvider) getProfile(cfg *cliConfigure) string {
	if p.profile == "" {
		p.profile = os.Getenv("VOLCENGINE_CLI_PROFILE")
	}
	if p.profile == "" && cfg != nil && cfg.Current != "" {
		p.profile = cfg.Current
	}
	if p.profile == "" {
		p.profile = "default"
	}
	return p.profile
}

func UnixTimestampToTime(ts int64) time.Time {
	switch {
	case ts >= 1e18: // 纳秒
		return time.Unix(0, ts)
	case ts >= 1e15: // 微秒
		return time.Unix(0, ts*int64(time.Microsecond))
	case ts >= 1e12: // 毫秒
		sec := ts / 1000
		nsec := (ts % 1000) * int64(time.Millisecond)
		return time.Unix(sec, nsec)
	default: // 秒
		return time.Unix(ts, 0)
	}
}

func tokenCacheFileName(startURL, sessionName string) string {
	payload := struct {
		StartURL    string `json:"start_url"`
		SessionName string `json:"session_name"`
	}{
		StartURL:    startURL,
		SessionName: sessionName,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		data = []byte(startURL + "\n" + sessionName)
	}
	hash := sha1.Sum(data)
	return fmt.Sprintf("%x.json", hash)
}

// 使用临时文件写入后原子替换，避免中断导致缓存损坏。
func writeJSONFileAtomic(path string, perm os.FileMode, payload interface{}) (retErr error) {
	dir := filepath.Dir(path)
	tempFile, err := os.CreateTemp(dir, ".tmp-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tempName := tempFile.Name()
	defer func() {
		if retErr != nil {
			_ = tempFile.Close()
			_ = os.Remove(tempName)
		}
	}()

	if err := tempFile.Chmod(perm); err != nil {
		retErr = fmt.Errorf("failed to set cache file permissions: %w", err)
		return retErr
	}

	encoder := json.NewEncoder(tempFile)
	if err := encoder.Encode(payload); err != nil {
		retErr = fmt.Errorf("failed to write cache file: %w", err)
		return retErr
	}

	if err := tempFile.Close(); err != nil {
		retErr = fmt.Errorf("failed to close cache file: %w", err)
		return retErr
	}

	if err := os.Rename(tempName, path); err != nil {
		removeErr := os.Remove(path)
		if removeErr == nil || os.IsNotExist(removeErr) {
			if err2 := os.Rename(tempName, path); err2 == nil {
				return nil
			}
		}
		retErr = fmt.Errorf("failed to replace cache file: %w", err)
		return retErr
	}

	return nil
}
