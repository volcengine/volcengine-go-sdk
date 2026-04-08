package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	MaxRetries    int           // Retry Times. Default is 0.
	RetryInterval time.Duration // the sleep interval between retries. Defaults to 1s.

	lastUpdateTimestamp int64
	expirationTimestamp int64
	sessionValue        *Value
	httpOptions         *HttpOptions
}

// NewSAMLCredentialsProvider constructs a SAMLCredentialsProvider with the
// required role/saml-provider TRNs and the base64-encoded SAML assertion
// returned by your IdP.
func NewSAMLCredentialsProvider(roleTrn, samlProviderTrn, samlAssertion string) *SAMLCredentialsProvider {
	return &SAMLCredentialsProvider{
		RoleTrn:         roleTrn,
		SAMLProviderTrn: samlProviderTrn,
		SAMLAssertion:   samlAssertion,
		DurationSeconds: 3600,
		Schema:          "https",
		httpOptions:     &HttpOptions{Timeout: 10 * time.Second},
	}
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

	maxRetries := p.MaxRetries
	if maxRetries < 0 {
		maxRetries = 0
	}
	retryInterval := p.RetryInterval
	if retryInterval <= 0 {
		retryInterval = time.Second
	}

	requestURL := scheme + "://" + endpoint + "/?Action=AssumeRoleWithSAML&Version=2018-01-01"
	var stsResp AssumeRoleWithSAMLResponse
	var lastErr error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		stsResp, lastErr = doSAMLAssumeRoleOnce(client, requestURL, data.Encode())
		if lastErr == nil {
			break
		}
		if attempt < maxRetries {
			time.Sleep(retryInterval)
		}
	}
	if lastErr != nil {
		return Value{}, lastErr
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

// doSAMLAssumeRoleOnce performs one AssumeRoleWithSAML HTTP call and decodes the response. It is intended to be invoked inside a retry loop.
func doSAMLAssumeRoleOnce(client *http.Client, requestURL, body string) (AssumeRoleWithSAMLResponse, error) {
	var stsResp AssumeRoleWithSAMLResponse
	resp, err := client.Post(requestURL, "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		return stsResp, fmt.Errorf("failed to request STS service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return stsResp, fmt.Errorf("STS service returned non-OK status: %d, body: %s", resp.StatusCode, string(respBody))
	}
	if err := json.NewDecoder(resp.Body).Decode(&stsResp); err != nil {
		return AssumeRoleWithSAMLResponse{}, fmt.Errorf("failed to decode STS response: %v", err)
	}
	if stsResp.ResponseMetadata.Error != nil {
		return AssumeRoleWithSAMLResponse{}, fmt.Errorf("STS AssumeRoleWithSAML returned error: %v", stsResp.ResponseMetadata)
	}
	return stsResp, nil
}
