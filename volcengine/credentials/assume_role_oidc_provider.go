package credentials

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var _ Provider = &OIDCCredentialsProvider{}

const (
	defaultSessionName = "volcengine-go-sdk-oidc-session"
)

type HttpOptions struct {
	// The maximum number of retries to attempt.
	MaxRetries int
	// The maximum time to wait for a connection.
	ConnectTimeout time.Duration
}

// OIDCReAssumeRoleWithOIDCsponse represents the response from AssumeRoleWithOIDC API call
type AssumeRoleWithOIDCResponse struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId,omitempty"`
		Action    string `json:"Action,omitempty"`
		Version   string `json:"Version,omitempty"`
		Service   string `json:"Service,omitempty"`
		Region    string `json:"Region,omitempty"`
	} `json:"ResponseMetadata,omitempty"`
	Result struct {
		Credentials struct {
			CurrentTime     string `json:"CurrentTime,omitempty"`
			Expiration      string `json:"Expiration,omitempty"`
			AccessKeyId     string `json:"AccessKeyId,omitempty"`
			SecretAccessKey string `json:"SecretAccessKey,omitempty"`
			SessionToken    string `json:"SessionToken,omitempty"`
		} `json:"Credentials,omitempty"`
		OIDCTokenInfo struct {
			Subject        string   `json:"Subject,omitempty"`
			Issuer         string   `json:"Issuer,omitempty"`
			ClientIds      []string `json:"ClientIds,omitempty"`
			ExpirationTime string   `json:"ExpirationTime,omitempty"`
			IssuanceTime   string   `json:"IssuanceTime,omitempty"`
		} `json:"OIDCTokenInfo,omitempty"`
		AssumedRoleUser struct {
			AssumedRoleId string `json:"AssumedRoleId,omitempty"`
			Trn           string `json:"Trn,omitempty"`
		} `json:"AssumedRoleUser,omitempty"`
	} `json:"Result,omitempty"`
}

type OIDCCredentialsProvider struct {
	// mutex protects the credentials from being updated concurrently
	mutex             sync.RWMutex
	oidcProviderTRN   string
	oidcTokenFilePath string
	roleTrn           string
	roleSessionName   string
	durationSeconds   int
	policy            string
	// for sts endpoint
	stsRegion           string
	enableVpc           bool
	stsEndpoint         string
	lastUpdateTimestamp int64
	expirationTimestamp int64
	sessionValue        *Value
	// for http options
	httpOptions *HttpOptions
}

func NewOIDCCredentialsProviderFromEnv() *OIDCCredentialsProvider {
	// get all from env

	return &OIDCCredentialsProvider{
		oidcProviderTRN:   os.Getenv("VOLCENGINE_OIDC_PROVIDER_TRN"),
		oidcTokenFilePath: os.Getenv("VOLCENGINE_OIDC_TOKEN_FILE"),
		roleTrn:           os.Getenv("VOLCENGINE_ROLE_TRN"),
		roleSessionName:   os.Getenv("VOLCENGINE_ROLE_SESSION_NAME"),
		durationSeconds:   3600, // default duration
		policy:            os.Getenv("VOLCENGINE_ROLE_POLICY"),
		stsRegion:         os.Getenv("VOLCENGINE_STS_REGION"),
		enableVpc:         os.Getenv("ENABLE_VPC") == "true",
		stsEndpoint:       os.Getenv("STS_ENDPOINT"),
		httpOptions:       &HttpOptions{MaxRetries: 3, ConnectTimeout: 10 * time.Second},
	}
}
func (p *OIDCCredentialsProvider) Retrieve() (Value, error) {
	if p.IsExpired() {
		return p.fetchOnce()
	}
	// 如果没有过期，直接返回缓存的凭证
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return *p.sessionValue, nil
}

func (p *OIDCCredentialsProvider) fetchOnce() (Value, error) {
	// 检查必要参数
	if p.oidcTokenFilePath == "" || p.roleTrn == "" {
		return Value{}, fmt.Errorf("missing required environment variables: VOLCENGINE_OIDC_TOKEN_FILE or VOLCENGINE_ROLE_TRN")
	}

	// 从文件中读取OIDC令牌
	tokenBytes, err := os.ReadFile(p.oidcTokenFilePath)
	if err != nil {
		return Value{}, fmt.Errorf("failed to read OIDC token file: %v", err)
	}
	oidcToken := string(tokenBytes)

	// 构建请求参数
	data := url.Values{}
	data.Set("RoleTrn", p.roleTrn)
	data.Set("OIDCToken", oidcToken)
	if p.roleSessionName != "" {
		data.Set("RoleSessionName", p.roleSessionName)
	} else {
		data.Set("RoleSessionName", defaultSessionName)
	}
	data.Set("DurationSeconds", fmt.Sprintf("%d", p.durationSeconds))
	if p.policy != "" {
		data.Set("Policy", p.policy)
	}

	endpoint := "https://sts.volcengineapi.com"
	if p.stsEndpoint != "" {
		endpoint = p.stsEndpoint
	}
	// 发送请求到STS服务
	client := &http.Client{Timeout: p.httpOptions.ConnectTimeout}
	resp, err := client.Post(
		endpoint+"/?Action=AssumeRoleWithOIDC&Version=2018-01-01",
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return Value{}, fmt.Errorf("failed to request STS service: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return Value{}, fmt.Errorf("STS service returned non-OK status: %d, body: %s", resp.StatusCode, string(body))
	}
	stsResp := AssumeRoleWithOIDCResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&stsResp); err != nil {
		return Value{}, fmt.Errorf("failed to decode STS response: %v", err)
	}

	// 创建凭证对象
	creds := Value{
		AccessKeyID:     stsResp.Result.Credentials.AccessKeyId,
		SecretAccessKey: stsResp.Result.Credentials.SecretAccessKey,
		SessionToken:    stsResp.Result.Credentials.SessionToken,
		ProviderName:    "OIDC",
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.lastUpdateTimestamp = time.Now().Unix()
	// parse time 2021-04-12T11:57:09+08:00 to unix timestamp
	expirationTime, err := time.Parse(time.RFC3339, stsResp.Result.Credentials.Expiration)
	if err != nil {
		return Value{}, fmt.Errorf("failed to parse expiration time: %v", err)
	}
	p.expirationTimestamp = expirationTime.Unix()
	p.sessionValue = &creds
	return creds, nil
}

func (p *OIDCCredentialsProvider) IsExpired() bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	if p.expirationTimestamp == 0 {
		return true
	}
	return time.Now().Unix() > p.expirationTimestamp
}
