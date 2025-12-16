package clicreds

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// OAuthClientConfig 用于配置 OAuth 客户端的可选项。
type OAuthClientConfig struct {
	// Region 控制使用的区域（默认：cn-beijing）。
	Region string
	// HTTPClient 允许注入自定义 HTTP 客户端（例如代理、超时）。
	HTTPClient *http.Client
}

const (
	// defaultOAuthRegion 为默认区域标识（当前未启用拼接区域模板的逻辑）。
	defaultOAuthRegion = "cn-beijing"
	// OAuth API 各路径常量。
	defaultTokenPath = "/token"
	// defaultRequestTimeout 为默认请求超时。
	defaultRequestTimeout = 10 * time.Second
	// deviceCodeGrantType 为设备码授权的 grant_type 标识。
	deviceCodeGrantType = "urn:ietf:params:oauth:grant-type:device_code"
	// oAuthBaseURLTemplate  = "https://cloudidentity-ssoauth-%s.volces.com"
	// oAuthBaseURLTemplate 为 OAuth 服务基础地址（已固定为 BOE 环境）。
	oAuthBaseURLTemplate = "https://cloudidentity-oauth.bytedance.net"
)

// OAuthClient 缓存拼好的 URL 和 HTTP 客户端，避免每次调用重新计算。
type OAuthClient struct {
	baseURL    string
	tokenURL   string
	httpClient *http.Client
}

// OAuthClientAPI 定义 OAuth 客户端对外暴露的方法集合，便于测试或替换实现。
type OAuthClientAPI interface {
	CreateToken(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error)
}

// 编译期断言：确保 *OAuthClient 实现了 OAuthClientAPI 接口（缺方法会直接编译失败）。
var _ OAuthClientAPI = (*OAuthClient)(nil)

// CreateTokenRequest 对应 CreateToken API 的请求参数。
type CreateTokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token,omitempty"`
	DeviceCode   string `json:"device_code,omitempty"`
}

// CreateTokenResponse 表示获取 Token 成功后的返回结构。
type CreateTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int    `json:"expires_in"`
}

// oauthErrorResponse 对应 OAuth 错误响应结构。
type oauthErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// OAuthAPIError 用于承载 OAuth API 非 2xx 响应时的结构化错误信息。
type OAuthAPIError struct {
	StatusCode int
	Response   oauthErrorResponse
	RawBody    string
}

func (e *OAuthAPIError) Error() string {
	if e == nil {
		return ""
	}
	if e.Response.ErrorDescription != "" {
		return fmt.Sprintf("request failed: %s (%s) [status %d]", e.Response.Error, e.Response.ErrorDescription, e.StatusCode)
	}
	if e.Response.Error != "" {
		return fmt.Sprintf("request failed: %s [status %d]", e.Response.Error, e.StatusCode)
	}
	if e.RawBody != "" {
		return fmt.Sprintf("request failed with status %d: %s", e.StatusCode, e.RawBody)
	}
	return fmt.Sprintf("request failed with status %d", e.StatusCode)
}

// NewOAuthClient 根据配置创建 OAuthClient，包含默认值和可选覆盖项。
func NewOAuthClient(cfg *OAuthClientConfig) *OAuthClient {
	/**
	region := defaultOAuthRegion
	if cfg != nil && strings.TrimSpace(cfg.Region) != "" {
		region = strings.TrimSpace(cfg.Region)
	}**/

	// base := fmt.Sprintf(oAuthBaseURLTemplate, region)
	base := oAuthBaseURLTemplate
	client := &http.Client{Timeout: defaultRequestTimeout}
	if cfg != nil && cfg.HTTPClient != nil {
		client = cfg.HTTPClient
	}

	return &OAuthClient{
		baseURL:    strings.TrimRight(base, "/"),
		tokenURL:   strings.TrimRight(base, "/") + defaultTokenPath,
		httpClient: client,
	}
}

// CreateToken 调用 CreateToken API，获取 access/refresh token。
func (c *OAuthClient) CreateToken(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if strings.TrimSpace(req.GrantType) == "" {
		return nil, fmt.Errorf("grantType is required")
	}
	if strings.TrimSpace(req.ClientID) == "" || strings.TrimSpace(req.ClientSecret) == "" {
		return nil, fmt.Errorf("clientId and clientSecret are required")
	}
	switch strings.ToLower(req.GrantType) {
	case "refresh_token":
		if strings.TrimSpace(req.RefreshToken) == "" {
			return nil, fmt.Errorf("refreshToken is required for refresh_token grant")
		}
	case deviceCodeGrantType, "device_code":
		if strings.TrimSpace(req.DeviceCode) == "" {
			return nil, fmt.Errorf("deviceCode is required for device_code grant")
		}
	default:
		return nil, fmt.Errorf("grantType %s is not supported", req.GrantType)
	}

	var apiResp CreateTokenResponse
	if err := doOAuthPost(ctx, c.httpClient, c.tokenURL, req, &apiResp); err != nil {
		return nil, err
	}
	if apiResp.AccessToken == "" && apiResp.TokenType == "" && apiResp.RefreshToken == "" && apiResp.ExpiresIn == 0 {
		return nil, fmt.Errorf("CreateToken succeeded but response was empty")
	}
	return &apiResp, nil
}

// doOAuthPost 负责发起 OAuth POST 请求并统一处理错误与响应解析。
func doOAuthPost(ctx context.Context, client *http.Client, url string, payload interface{}, out interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-tt-env", "boe_cli_cli")

	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode/100 != 2 {
		var errResp oauthErrorResponse
		if len(respBytes) > 0 && json.Unmarshal(respBytes, &errResp) == nil && (errResp.Error != "" || errResp.ErrorDescription != "") {
			return &OAuthAPIError{
				StatusCode: resp.StatusCode,
				Response:   errResp,
				RawBody:    string(respBytes),
			}
		}
		if len(respBytes) == 0 {
			return &OAuthAPIError{
				StatusCode: resp.StatusCode,
			}
		}
		return &OAuthAPIError{
			StatusCode: resp.StatusCode,
			RawBody:    string(respBytes),
		}
	}

	if len(respBytes) > 0 && out != nil {
		if err := json.Unmarshal(respBytes, out); err != nil {
			return fmt.Errorf("failed to decode response (status %d): %w", resp.StatusCode, err)
		}
	}

	return nil
}
