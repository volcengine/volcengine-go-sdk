package credentials

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
)

var _ Provider = &SAMLCredentialsProvider{}

// SAMLResult is the result body returned by AssumeRoleWithSAML.
type SAMLResult struct {
	Credentials     OIDCCredentials `json:"Credentials,omitempty"`
	AssumedRoleUser AssumedRoleUser `json:"AssumedRoleUser,omitempty"`
}

// AssumeRoleWithSAMLResponse is the full STS response for AssumeRoleWithSAML.
type AssumeRoleWithSAMLResponse struct {
	ResponseMetadata response.ResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           SAMLResult                `json:"Result,omitempty"`
}

// SAMLCredentialsProvider obtains temporary credentials from STS via
// AssumeRoleWithSAML. It refreshes credentials automatically before expiration.
type SAMLCredentialsProvider struct {
	RoleTrn         string
	SAMLProviderTrn string
	SAMLAssertion   string
	DurationSeconds int
	Policy          string
	// for sts endpoint
	Schema        string
	Endpoint      string
	MaxRetries    *int          // Retry attempts. Nil falls back to DefaultRetryerMaxNumRetries (3); 0 disables retries.
	RetryInterval time.Duration // Sleep interval between retries. If zero or negative, falls back to DefaultRetryerRetryDelay (1s).

	lastUpdateTimestamp int64
	expirationTimestamp int64
	sessionValue        *Value
	httpOptions         *HttpOptions
}

// SAMLProviderOptions contains optional configuration for SAMLCredentialsProvider.
type SAMLProviderOptions struct {
	DurationSeconds int
	Policy          string
	Schema          string
	Endpoint        string
	MaxRetries      *int
	RetryInterval   time.Duration
	HttpOptions     *HttpOptions
}

// NewSAMLCredentialsProviderWithOptions constructs a SAMLCredentialsProvider
// with required parameters and optional functional options.
func NewSAMLCredentialsProviderWithOptions(roleTrn, samlProviderTrn, samlAssertion string, optFns ...func(*SAMLProviderOptions)) *SAMLCredentialsProvider {
	opts := SAMLProviderOptions{
		DurationSeconds: 3600,
		Schema:          "https",
		HttpOptions:     &HttpOptions{Timeout: 10 * time.Second},
	}
	for _, fn := range optFns {
		fn(&opts)
	}
	return &SAMLCredentialsProvider{
		RoleTrn:         roleTrn,
		SAMLProviderTrn: samlProviderTrn,
		SAMLAssertion:   samlAssertion,
		DurationSeconds: opts.DurationSeconds,
		Policy:          opts.Policy,
		Schema:          opts.Schema,
		Endpoint:        opts.Endpoint,
		MaxRetries:      opts.MaxRetries,
		RetryInterval:   opts.RetryInterval,
		httpOptions:     opts.HttpOptions,
	}
}

// NewSAMLCredentialsProvider constructs a SAMLCredentialsProvider with the
// required role/saml-provider TRNs and the base64-encoded SAML assertion
// returned by your IdP.
func NewSAMLCredentialsProvider(roleTrn, samlProviderTrn, samlAssertion string) *SAMLCredentialsProvider {
	return NewSAMLCredentialsProviderWithOptions(roleTrn, samlProviderTrn, samlAssertion)
}

func (p *SAMLCredentialsProvider) Retrieve() (Value, error) {
	return p.fetchOnce()
}

func (p *SAMLCredentialsProvider) IsExpired() bool {
	if p.expirationTimestamp == 0 {
		return true
	}
	return time.Now().Unix() > p.expirationTimestamp
}

func (p *SAMLCredentialsProvider) fetchOnce() (Value, error) {
	if p.RoleTrn == "" || p.SAMLProviderTrn == "" || p.SAMLAssertion == "" {
		return Value{}, fmt.Errorf("missing required SAML parameters: RoleTrn, SAMLProviderTrn or SAMLAssertion")
	}

	data := url.Values{}
	if p.DurationSeconds > 0 {
		data.Set("DurationSeconds", fmt.Sprintf("%d", p.DurationSeconds))
	} else {
		data.Set("DurationSeconds", "3600")
	}
	data.Set("RoleTrn", p.RoleTrn)
	data.Set("SAMLProviderTrn", p.SAMLProviderTrn)
	data.Set("SAMLResp", p.SAMLAssertion)
	if p.Policy != "" {
		data.Set("Policy", p.Policy)
	}

	endpoint := DefaultEndpoint
	if p.Endpoint != "" {
		endpoint = p.Endpoint
	}
	scheme := "https"
	if p.Schema != "" {
		scheme = p.Schema
	}

	var client *http.Client
	if p.httpOptions != nil {
		client = &http.Client{Timeout: p.httpOptions.Timeout}
	} else {
		client = &http.Client{}
	}

	var stsResp AssumeRoleWithSAMLResponse
	if err := doSTSFormRequest(client, scheme, endpoint, "AssumeRoleWithSAML",
		data, &stsResp, p.MaxRetries, p.RetryInterval,
		"STS AssumeRoleWithSAML returned error"); err != nil {
		return Value{}, err
	}

	creds := Value{
		AccessKeyID:     stsResp.Result.Credentials.AccessKeyId,
		SecretAccessKey: stsResp.Result.Credentials.SecretAccessKey,
		SessionToken:    stsResp.Result.Credentials.SessionToken,
		ProviderName:    "SAML",
	}
	p.lastUpdateTimestamp = time.Now().Unix()
	expirationTime, err := time.Parse(time.RFC3339, stsResp.Result.Credentials.Expiration)
	if err != nil {
		return Value{}, fmt.Errorf("failed to parse expiration time: %v", err)
	}
	p.expirationTimestamp = expirationTime.Add(-60 * time.Second).Unix()
	p.sessionValue = &creds
	return creds, nil
}
