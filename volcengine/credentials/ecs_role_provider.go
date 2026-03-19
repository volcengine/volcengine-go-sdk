package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

// EcsRoleProviderName provides a name of EcsRole provider.
const EcsRoleProviderName = "EcsRoleProvider"

// TODO: Confirm with ECS team and replace placeholders.
const (
	// imdsEndpoint is the IMDS endpoint address. Pending confirmation from ECS team.
	imdsEndpoint = "http://100.96.0.96"

	// imdsRoleCredsPath is the path template to get STS credentials for a role.
	// Use fmt.Sprintf with roleName.
	imdsRoleCredsPath = "/volcstack/latest/iam/security_credentials/%s"

	// imdsDefaultConnectTimeout is the default connect timeout for IMDS requests.
	imdsDefaultConnectTimeout = 1 * time.Second

	// imdsDefaultReadTimeout is the default read timeout for IMDS requests.
	imdsDefaultReadTimeout = 1 * time.Second

	// imdsDefaultMaxRetries is the default max retries for IMDS requests.
	imdsDefaultMaxRetries = 3

	// imdsDefaultRetryInterval is the default retry interval.
	imdsDefaultRetryInterval = 1 * time.Second

	// imdsDefaultExpiryWindow is the default expiry window.
	imdsDefaultExpiryWindow = 5 * time.Minute
)

// imdsCredentialResponse represents the JSON response from the IMDS credential endpoint.
// TODO: Confirm exact field names with ECS team.
type imdsCredentialResponse struct {
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Expiration      string `json:"ExpiredTime"`
}

// EcsRoleProvider retrieves credentials from the ECS Instance Metadata Service (IMDS).
type EcsRoleProvider struct {
	Expiry

	// RoleName is the IAM role name. If empty, it is resolved from the env var
	// VOLCENGINE_ECS_METADATA. If neither is set, an error is returned.
	RoleName string

	// ExpiryWindow will allow the credentials to trigger refreshing prior to
	// the credentials actually expiring.
	ExpiryWindow time.Duration

	httpClient *http.Client
	maxRetries int
	retryInterval time.Duration
}

// NewEcsRoleProvider returns a new EcsRoleProvider.
func NewEcsRoleProvider(roleName string) *EcsRoleProvider {
	return &EcsRoleProvider{
		RoleName:     roleName,
		ExpiryWindow: imdsDefaultExpiryWindow,
		httpClient: &http.Client{
			Timeout: imdsDefaultConnectTimeout + imdsDefaultReadTimeout,
		},
		maxRetries:    imdsDefaultMaxRetries,
		retryInterval: imdsDefaultRetryInterval,
	}
}

// NewEcsRoleCredentials returns a pointer to a new Credentials object wrapping
// the EcsRoleProvider.
func NewEcsRoleCredentials(roleName string) *Credentials {
	return NewCredentials(NewEcsRoleProvider(roleName))
}

// Retrieve retrieves ECS role credentials from IMDS.
func (p *EcsRoleProvider) Retrieve() (Value, error) {
	if isIMDSDisabled() {
		return Value{ProviderName: EcsRoleProviderName},
			volcengineerr.New("EcsRoleDisabled",
				"IMDS credential retrieval is disabled via VOLCENGINE_ECS_METADATA_DISABLED=true", nil)
	}

	roleName, err := p.resolveRoleName()
	if err != nil {
		return Value{ProviderName: EcsRoleProviderName}, err
	}

	creds, err := p.getCredentials(roleName)
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

// resolveRoleName determines the role name from: parameter > env var.
// If neither is set, returns an error (no auto-detect from IMDS).
func (p *EcsRoleProvider) resolveRoleName() (string, error) {
	if p.RoleName != "" {
		return p.RoleName, nil
	}

	if envRole := os.Getenv("VOLCENGINE_ECS_METADATA"); envRole != "" {
		return envRole, nil
	}

	return "", volcengineerr.New("EcsRoleNameNotSet",
		"ECS role name not set: specify via RoleName parameter or VOLCENGINE_ECS_METADATA environment variable", nil)
}

// getCredentials fetches STS credentials for the given role from IMDS.
func (p *EcsRoleProvider) getCredentials(roleName string) (*imdsCredentialResponse, error) {
	url := fmt.Sprintf("%s"+imdsRoleCredsPath, imdsEndpoint, roleName)
	body, err := p.doRequestWithRetry(url)
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

// doRequestWithRetry performs an HTTP GET with retry logic.
func (p *EcsRoleProvider) doRequestWithRetry(url string) ([]byte, error) {
	var lastErr error
	for i := 0; i < p.maxRetries; i++ {
		if i > 0 {
			time.Sleep(p.retryInterval)
		}

		body, err := p.doRequest(url)
		if err == nil {
			return body, nil
		}
		lastErr = err
	}
	return nil, lastErr
}

func (p *EcsRoleProvider) doRequest(url string) ([]byte, error) {
	resp, err := p.httpClient.Get(url)
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
