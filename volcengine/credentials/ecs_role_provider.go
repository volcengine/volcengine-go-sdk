package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"os"
	"strings"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

// EcsRoleProviderName provides a name of EcsRole provider.
const EcsRoleProviderName = "EcsRoleProvider"

const (
	// IMDSv2 endpoint and paths
	imdsEndpoint      = "http://100.96.0.96"
	imdsRoleCredsPath = "/volcstack/latest/iam/security_credentials/%s"        // GET
	imdsRoleNamePath  = "/volcstack/latest/iam/security_credentials?type=user" // GET
	imdsTokenPath     = "/latest/api/token"                                    // GET

	// IMDSv2 headers
	imdsTokenTTLHeader  = "X-volc-ecs-metadata-token-ttl-seconds"
	imdsTokenHeader     = "X-volc-ecs-metadata-token"
	imdsTokenTTLSeconds = "21600" // 6 hours

	// Defaults
	imdsDefaultConnectTimeout = 1 * time.Second
	imdsDefaultReadTimeout    = 1 * time.Second
	imdsDefaultRetryInterval  = 1 * time.Second
	imdsDefaultExpiryWindow   = 5 * time.Minute
)

// imdsCredentialResponse represents the JSON response from the IMDS credential endpoint.
type imdsCredentialResponse struct {
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Expiration      string `json:"ExpiredTime"`
}

// EcsRoleProvider retrieves credentials from the ECS Instance Metadata Service (IMDSv2).
//
// Flow:
//  1. GET IMDSv2 token (fresh every time, not cached)
//  2. Resolve roleName: param > env > auto-detect from IMDS
//  3. POST to get STS credentials with token header
type EcsRoleProvider struct {
	Expiry

	// RoleName is the IAM role name. If empty, resolved from env or auto-detected.
	RoleName string

	// ExpiryWindow will allow the credentials to trigger refreshing prior to
	// the credentials actually expiring.
	ExpiryWindow time.Duration

	// MaxRetries is the number of retry attempts after the initial IMDS request.
	// Nil falls back to DefaultRetryerMaxNumRetries (3); 0 disables retries.
	MaxRetries *int

	// RetryInterval is the sleep interval between IMDS retry attempts.
	// If zero or negative, it falls back to DefaultRetryerRetryDelay (1s).
	RetryInterval time.Duration

	httpClient *http.Client
}

// EcsRoleProviderOptions contains optional configuration for EcsRoleProvider.
type EcsRoleProviderOptions struct {
	ExpiryWindow  time.Duration
	MaxRetries    *int
	RetryInterval time.Duration
	HttpTimeout   time.Duration
}

// NewEcsRoleProviderWithOptions constructs an EcsRoleProvider with the role
// name as a required parameter and optional functional options.
func NewEcsRoleProviderWithOptions(roleName string, optFns ...func(*EcsRoleProviderOptions)) *EcsRoleProvider {
	opts := EcsRoleProviderOptions{
		ExpiryWindow: imdsDefaultExpiryWindow,
		HttpTimeout:  imdsDefaultConnectTimeout + imdsDefaultReadTimeout,
	}
	for _, fn := range optFns {
		fn(&opts)
	}
	return &EcsRoleProvider{
		RoleName:      roleName,
		ExpiryWindow:  opts.ExpiryWindow,
		MaxRetries:    opts.MaxRetries,
		RetryInterval: opts.RetryInterval,
		httpClient:    &http.Client{Timeout: opts.HttpTimeout},
	}
}

// NewEcsRoleProvider returns a new EcsRoleProvider.
func NewEcsRoleProvider(roleName string) *EcsRoleProvider {
	return NewEcsRoleProviderWithOptions(roleName)
}

// NewEcsRoleCredentials returns a pointer to a new Credentials object wrapping
// the EcsRoleProvider.
func NewEcsRoleCredentials(roleName string) *Credentials {
	return NewCredentials(NewEcsRoleProvider(roleName))
}

// Retrieve retrieves ECS role credentials from IMDS via IMDSv2.
func (p *EcsRoleProvider) Retrieve() (Value, error) {
	if isIMDSDisabled() {
		return Value{ProviderName: EcsRoleProviderName},
			volcengineerr.New("EcsRoleDisabled",
				"IMDS credential retrieval is disabled via VOLCENGINE_ECS_METADATA_DISABLED=true", nil)
	}

	// Step 1: Get IMDSv2 token (fresh every time)
	token, err := p.getIMDSv2Token()
	if err != nil {
		return Value{ProviderName: EcsRoleProviderName}, err
	}

	// Step 2: Resolve role name
	roleName, err := p.resolveRoleName(token)
	if err != nil {
		return Value{ProviderName: EcsRoleProviderName}, err
	}

	// Step 3: POST to get credentials
	creds, err := p.getCredentials(roleName, token)
	if err != nil {
		return Value{ProviderName: EcsRoleProviderName}, err
	}

	expiration, err := time.Parse(time.RFC3339, creds.Expiration)
	if err != nil {
		return Value{ProviderName: EcsRoleProviderName},
			volcengineerr.New("EcsRoleExpirationParse",
				fmt.Sprintf("failed to parse credential expiration: %s", creds.Expiration), err)
	}

	p.SetExpiration(expiration, p.ExpiryWindow)

	return Value{
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
		ProviderName:    EcsRoleProviderName,
	}, nil
}

// IsExpired returns whether the cached credentials have expired.
func (p *EcsRoleProvider) IsExpired() bool {
	return p.Expiry.IsExpired()
}

func isIMDSDisabled() bool {
	v := os.Getenv("VOLCENGINE_ECS_METADATA_DISABLED")
	return strings.EqualFold(v, "true")
}

// getIMDSv2Token fetches a fresh IMDSv2 token. Not cached.
func (p *EcsRoleProvider) getIMDSv2Token() (string, error) {
	url := imdsEndpoint + imdsTokenPath
	headers := map[string]string{imdsTokenTTLHeader: imdsTokenTTLSeconds}
	body, err := p.doRequestWithRetry(url, "GET", headers)
	if err != nil {
		return "", volcengineerr.New("EcsRoleTokenFailed",
			"failed to get IMDSv2 token", err)
	}
	token := strings.TrimSpace(string(body))
	if token == "" {
		return "", volcengineerr.New("EcsRoleTokenEmpty",
			"IMDSv2 token endpoint returned empty response", nil)
	}
	return token, nil
}

// resolveRoleName determines the role name: param > env > auto-detect from IMDS (every time).
// Not cached — roles can be dynamically unbound/rebound.
func (p *EcsRoleProvider) resolveRoleName(imdsToken string) (string, error) {
	if p.RoleName != "" {
		return p.RoleName, nil
	}

	if envRole := os.Getenv("VOLCENGINE_ECS_METADATA"); envRole != "" {
		return envRole, nil
	}

	// Auto-detect from IMDS
	return p.autoDetectRoleName(imdsToken)
}

// autoDetectRoleName queries IMDS for the role list and returns the first one.
func (p *EcsRoleProvider) autoDetectRoleName(imdsToken string) (string, error) {
	url := imdsEndpoint + imdsRoleNamePath
	headers := map[string]string{imdsTokenHeader: imdsToken}
	body, err := p.doRequestWithRetry(url, "GET", headers)
	if err != nil {
		return "", volcengineerr.New("EcsRoleDetectFailed",
			"failed to auto-detect role name from IMDS", err)
	}

	var roles []string
	if err := json.Unmarshal(body, &roles); err != nil {
		// Fallback: try splitting by newlines
		for _, r := range strings.Split(strings.TrimSpace(string(body)), "\n") {
			if r = strings.TrimSpace(r); r != "" {
				roles = append(roles, r)
			}
		}
	}

	if len(roles) == 0 {
		return "", volcengineerr.New("EcsRoleNotFound",
			"no IAM roles found via IMDS", nil)
	}

	return roles[0], nil
}

// getCredentials GET to IMDS to get STS credentials for the given role.
func (p *EcsRoleProvider) getCredentials(roleName, imdsToken string) (*imdsCredentialResponse, error) {
	url := fmt.Sprintf("%s"+imdsRoleCredsPath, imdsEndpoint, neturl.PathEscape(roleName))
	headers := map[string]string{imdsTokenHeader: imdsToken}
	body, err := p.doRequestWithRetry(url, "GET", headers)
	if err != nil {
		return nil, volcengineerr.New("EcsRoleCredentialsFailed",
			fmt.Sprintf("failed to get credentials for role %s from IMDS", roleName), err)
	}

	var resp imdsCredentialResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, volcengineerr.New("EcsRoleCredentialsParse",
			"failed to parse IMDS credential response", err)
	}

	if resp.AccessKeyID == "" || resp.SecretAccessKey == "" {
		return nil, volcengineerr.New("EcsRoleCredentialsEmpty",
			"IMDS returned empty credentials", nil)
	}

	return &resp, nil
}

// doRequestWithRetry performs an HTTP request with retry logic.
func (p *EcsRoleProvider) doRequestWithRetry(url, method string, headers map[string]string) ([]byte, error) {
	maxRetries, retryInterval := p.retryConfig()

	var lastErr error
	for i := 0; i <= maxRetries; i++ {
		if i > 0 {
			time.Sleep(retryInterval)
		}

		body, err := p.doRequest(url, method, headers)
		if err == nil {
			return body, nil
		}
		lastErr = err
	}
	return nil, lastErr
}

func (p *EcsRoleProvider) retryConfig() (int, time.Duration) {
	maxRetries := resolveCredentialMaxRetries(p.MaxRetries)
	retryInterval := p.RetryInterval
	if retryInterval <= 0 {
		retryInterval = DefaultRetryerRetryDelay
	}
	return maxRetries, retryInterval
}

func (p *EcsRoleProvider) doRequest(url, method string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create IMDS request: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("IMDS request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("IMDS returned HTTP %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read IMDS response body: %w", err)
	}

	return body, nil
}
