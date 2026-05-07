package credentials

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
)

var _ Provider = &OIDCCredentialsProvider{}

const DefaultEndpoint = "sts.volcengineapi.com"

const (
	defaultSessionName = "volcengine-go-sdk-oidc-session"
)

type HttpOptions struct {
	// The maximum time to wait for a connection.
	Timeout time.Duration
}

// OIDCCredentials 包含访问凭证信息
type OIDCCredentials struct {
	CurrentTime     string `json:"CurrentTime,omitempty"`
	Expiration      string `json:"Expiration,omitempty"`
	AccessKeyId     string `json:"AccessKeyId,omitempty"`
	SecretAccessKey string `json:"SecretAccessKey,omitempty"`
	SessionToken    string `json:"SessionToken,omitempty"`
}

// OIDCTokenInfo 包含 OIDC 令牌的信息
type OIDCTokenInfo struct {
	Subject        string   `json:"Subject,omitempty"`
	Issuer         string   `json:"Issuer,omitempty"`
	ClientIds      []string `json:"ClientIds,omitempty"`
	ExpirationTime string   `json:"ExpirationTime,omitempty"`
	IssuanceTime   string   `json:"IssuanceTime,omitempty"`
}

// AssumedRoleUser 包含假定角色用户的信息
type AssumedRoleUser struct {
	AssumedRoleId string `json:"AssumedRoleId,omitempty"`
	Trn           string `json:"Trn,omitempty"`
}

// Result 包含请求结果的详细信息
type Result struct {
	Credentials     OIDCCredentials `json:"Credentials,omitempty"`
	OIDCTokenInfo   OIDCTokenInfo   `json:"OIDCTokenInfo,omitempty"`
	AssumedRoleUser AssumedRoleUser `json:"AssumedRoleUser,omitempty"`
}

// AssumeRoleWithOIDCResponse 表示使用 OIDC 进行角色假定的响应结构体
type AssumeRoleWithOIDCResponse struct {
	ResponseMetadata response.ResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           Result                    `json:"Result,omitempty"`
}

type OIDCCredentialsProvider struct {
	OIDCTokenFilePath string
	RoleTrn           string
	RoleSessionName   string
	DurationSeconds   int
	Policy            string
	// for sts endpoint
	Schema              string
	Endpoint            string
	MaxRetries          *int          // Retry attempts. Nil falls back to DefaultRetryerMaxNumRetries (3); 0 disables retries.
	RetryInterval       time.Duration // Sleep interval between retries. If zero or negative, falls back to DefaultRetryerRetryDelay (1s).
	lastUpdateTimestamp int64
	expirationTimestamp int64
	sessionValue        *Value
	// for http options
	httpOptions *HttpOptions
}

// OIDCProviderOptions contains optional configuration for OIDCCredentialsProvider.
type OIDCProviderOptions struct {
	RoleSessionName string
	DurationSeconds int
	Policy          string
	Schema          string
	Endpoint        string
	MaxRetries      *int
	RetryInterval   time.Duration
	HttpOptions     *HttpOptions
}

// NewOIDCCredentialsProviderWithOptions constructs an OIDCCredentialsProvider
// with required parameters and optional functional options.
func NewOIDCCredentialsProviderWithOptions(tokenFilePath, roleTrn string, optFns ...func(*OIDCProviderOptions)) *OIDCCredentialsProvider {
	opts := OIDCProviderOptions{
		DurationSeconds: 3600,
		Schema:          "https",
		HttpOptions:     &HttpOptions{Timeout: 10 * time.Second},
	}
	for _, fn := range optFns {
		fn(&opts)
	}
	return &OIDCCredentialsProvider{
		OIDCTokenFilePath: tokenFilePath,
		RoleTrn:           roleTrn,
		RoleSessionName:   opts.RoleSessionName,
		DurationSeconds:   opts.DurationSeconds,
		Policy:            opts.Policy,
		Schema:            opts.Schema,
		Endpoint:          opts.Endpoint,
		MaxRetries:        opts.MaxRetries,
		RetryInterval:     opts.RetryInterval,
		httpOptions:       opts.HttpOptions,
	}
}

func NewOIDCCredentialsProviderFromEnv() *OIDCCredentialsProvider {
	sessionName := os.Getenv("VOLCENGINE_OIDC_ROLE_SESSION_NAME")
	if sessionName == "" {
		sessionName = "credentials-go-" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
	}
	return NewOIDCCredentialsProviderWithOptions(
		os.Getenv("VOLCENGINE_OIDC_TOKEN_FILE"),
		os.Getenv("VOLCENGINE_OIDC_ROLE_TRN"),
		func(o *OIDCProviderOptions) {
			o.RoleSessionName = sessionName
			o.Policy = os.Getenv("VOLCENGINE_OIDC_ROLE_POLICY")
			o.Endpoint = os.Getenv("VOLCENGINE_OIDC_STS_ENDPOINT")
		},
	)
}

func (p *OIDCCredentialsProvider) Retrieve() (Value, error) {
	return p.fetchOnce()
}

func (p *OIDCCredentialsProvider) fetchOnce() (Value, error) {
	// 检查必要参数
	if p.OIDCTokenFilePath == "" || p.RoleTrn == "" {
		return Value{}, fmt.Errorf("missing required environment variables: VOLCENGINE_OIDC_TOKEN_FILE or VOLCENGINE_OIDC_ROLE_TRN")
	}

	// 从文件中读取OIDC令牌
	tokenBytes, err := ioutil.ReadFile(p.OIDCTokenFilePath)
	if err != nil {
		return Value{}, fmt.Errorf("failed to read OIDC token file: %v", err)
	}
	oidcToken := string(tokenBytes)

	// 构建请求参数
	data := url.Values{}
	data.Set("RoleTrn", p.RoleTrn)
	data.Set("OIDCToken", oidcToken)
	if p.RoleSessionName != "" {
		data.Set("RoleSessionName", p.RoleSessionName)
	} else {
		data.Set("RoleSessionName", defaultSessionName)
	}
	data.Set("DurationSeconds", fmt.Sprintf("%d", p.DurationSeconds+60))
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

	var stsResp AssumeRoleWithOIDCResponse
	if err := doSTSFormRequest(client, scheme, endpoint, "AssumeRoleWithOIDC",
		data, &stsResp, p.MaxRetries, p.RetryInterval,
		"STS AssumeRoleWithOIDC service error"); err != nil {
		return Value{}, err
	}

	creds := Value{
		AccessKeyID:     stsResp.Result.Credentials.AccessKeyId,
		SecretAccessKey: stsResp.Result.Credentials.SecretAccessKey,
		SessionToken:    stsResp.Result.Credentials.SessionToken,
		ProviderName:    "OIDC",
	}
	p.lastUpdateTimestamp = time.Now().Unix()
	expirationTime, err := time.Parse(time.RFC3339, stsResp.Result.Credentials.Expiration)
	if err != nil {
		return Value{}, fmt.Errorf("failed to parse expiration time: %v", err)
	}
	earlyExpirationTime := expirationTime.Add(time.Duration(-60) * time.Second)
	p.expirationTimestamp = earlyExpirationTime.Unix()
	p.sessionValue = &creds
	return creds, nil
}

func (p *OIDCCredentialsProvider) IsExpired() bool {
	if p.expirationTimestamp == 0 {
		return true
	}
	return time.Now().Unix() > p.expirationTimestamp
}
