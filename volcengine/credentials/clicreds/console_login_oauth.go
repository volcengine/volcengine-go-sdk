package clicreds

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// defaultConsoleLoginEndpoint 与 volcengine-cli 端 ve login 默认的 endpoint
	// 保持一致；当 login cache 里没有 endpoint_url 字段时使用该值。
	defaultConsoleLoginEndpoint = "https://signin.volcengine.com"
	// consoleLoginTokenPath 是 OAuth token 端点的固定路径。
	consoleLoginTokenPath = "/authorize/oauth/token"
	// consoleLoginRequestTimeout 是单次 HTTP 请求的超时。
	consoleLoginRequestTimeout = 30 * time.Second
	// consoleLoginRetryAttempts 是 token 请求的最大重试次数（含首发）。
	consoleLoginRetryAttempts = 3
)

// ConsoleLoginOAuthClientConfig 用于配置 ConsoleLoginOAuthClient。
type ConsoleLoginOAuthClientConfig struct {
	// EndpointURL 为 signin 服务的基础地址。空字符串时使用
	// defaultConsoleLoginEndpoint。允许测试时注入 httptest.Server URL。
	EndpointURL string
	// HTTPClient 允许注入自定义 HTTP 客户端（例如代理、超时）。
	HTTPClient *http.Client
}

// ConsoleLoginOAuthClient 封装对 signin OAuth 端点的调用，仅实现
// refresh_token grant，因为 SDK 不参与交互式授权流程。
type ConsoleLoginOAuthClient struct {
	tokenURL   string
	httpClient *http.Client
}

// ConsoleLoginOAuthClientAPI 抽离接口便于在测试中替换实现。
type ConsoleLoginOAuthClientAPI interface {
	RefreshToken(ctx context.Context, req *ConsoleLoginRefreshTokenRequest) (*ConsoleLoginTokenResponse, error)
}

var _ ConsoleLoginOAuthClientAPI = (*ConsoleLoginOAuthClient)(nil)

// ConsoleLoginRefreshTokenRequest 表示 refresh_token grant 的请求参数。
// 与 ve login 写入的 LoginTokenCache 字段一一对应。
type ConsoleLoginRefreshTokenRequest struct {
	ClientID     string
	Scope        string
	RefreshToken string
}

// ConsoleLoginTokenResponse 表示 token 端点的成功响应；字段与
// volcengine-cli ConsoleTokenResponse 对齐，以便 SDK 解析行为完全一致。
type ConsoleLoginTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
}

// consoleLoginOAuthErrorBody 对应 signin 服务返回的标准 OAuth 错误结构。
type consoleLoginOAuthErrorBody struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorURI         string `json:"error_uri,omitempty"`
}

// ConsoleLoginOAuthAPIError 包装 signin OAuth 端点的非 2xx 响应。
// 暴露 StatusCode 和结构化 error code，便于上层判断是否属于"refresh
// token 失效"类不可重试错误（典型为 400 invalid_grant）。
type ConsoleLoginOAuthAPIError struct {
	StatusCode int
	ErrorCode  string
	ErrorDesc  string
	RawBody    string
	RequestID  string
}

func (e *ConsoleLoginOAuthAPIError) Error() string {
	if e == nil {
		return ""
	}
	var parts []string
	if e.ErrorCode != "" {
		parts = append(parts, e.ErrorCode)
	}
	if e.ErrorDesc != "" {
		parts = append(parts, e.ErrorDesc)
	}
	msg := strings.Join(parts, ": ")
	if msg == "" {
		if e.RawBody != "" {
			msg = e.RawBody
		} else {
			msg = "unknown error"
		}
	}
	suffix := fmt.Sprintf("[status %d", e.StatusCode)
	if e.RequestID != "" {
		suffix += ", requestId: " + e.RequestID
	}
	suffix += "]"
	return fmt.Sprintf("console login oauth request failed: %s %s", msg, suffix)
}

// IsRefreshTokenInvalid 判断错误是否来自"refresh token 不可再用"语义。
// 服务端在 RT 失效 / 不存在 / 已撤销时按 OAuth 规范返回 400 + invalid_grant；
// 上层据此触发"重读磁盘 cache fallback"。
func (e *ConsoleLoginOAuthAPIError) IsRefreshTokenInvalid() bool {
	if e == nil {
		return false
	}
	return e.StatusCode == http.StatusBadRequest && e.ErrorCode == "invalid_grant"
}

// NewConsoleLoginOAuthClient 根据配置创建客户端。空配置/空字段时使用默认值。
func NewConsoleLoginOAuthClient(cfg *ConsoleLoginOAuthClientConfig) *ConsoleLoginOAuthClient {
	endpoint := defaultConsoleLoginEndpoint
	if cfg != nil && strings.TrimSpace(cfg.EndpointURL) != "" {
		endpoint = strings.TrimSpace(cfg.EndpointURL)
	}
	endpoint = strings.TrimRight(endpoint, "/")

	client := &http.Client{Timeout: consoleLoginRequestTimeout}
	if cfg != nil && cfg.HTTPClient != nil {
		client = cfg.HTTPClient
	}

	return &ConsoleLoginOAuthClient{
		tokenURL:   endpoint + consoleLoginTokenPath,
		httpClient: client,
	}
}

// RefreshToken 使用 refresh_token grant 换取新的 access_token。
//
// 字段校验失败时返回普通 error；HTTP 非 2xx 时返回 *ConsoleLoginOAuthAPIError，
// 调用方可通过 IsRefreshTokenInvalid 判断是否需触发 fallback。
func (c *ConsoleLoginOAuthClient) RefreshToken(ctx context.Context, req *ConsoleLoginRefreshTokenRequest) (*ConsoleLoginTokenResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if strings.TrimSpace(req.ClientID) == "" {
		return nil, fmt.Errorf("client_id is required")
	}
	if strings.TrimSpace(req.RefreshToken) == "" {
		return nil, fmt.Errorf("refresh_token is required")
	}

	form := url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("client_id", req.ClientID)
	form.Set("refresh_token", req.RefreshToken)
	if strings.TrimSpace(req.Scope) != "" {
		form.Set("scope", req.Scope)
	}
	body := form.Encode()

	var resp ConsoleLoginTokenResponse
	err := doWithRetry(ctx, retryOptions{maxAttempts: consoleLoginRetryAttempts}, func() error {
		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.tokenURL, strings.NewReader(body))
		if err != nil {
			return fmt.Errorf("failed to build request: %w", err)
		}
		httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		httpResp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return fmt.Errorf("request failed: %w", err)
		}
		defer httpResp.Body.Close()

		respBytes, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response: %w", err)
		}
		requestID := httpResp.Header.Get("X-Tt-Logid")

		if httpResp.StatusCode/100 != 2 {
			apiErr := &ConsoleLoginOAuthAPIError{
				StatusCode: httpResp.StatusCode,
				RawBody:    string(respBytes),
				RequestID:  requestID,
			}
			if len(respBytes) > 0 {
				var errBody consoleLoginOAuthErrorBody
				if json.Unmarshal(respBytes, &errBody) == nil {
					apiErr.ErrorCode = errBody.Error
					apiErr.ErrorDesc = errBody.ErrorDescription
				}
			}
			return apiErr
		}

		if len(respBytes) > 0 {
			if err := json.Unmarshal(respBytes, &resp); err != nil {
				return fmt.Errorf("failed to decode token response (status %d, requestId: %s): %w", httpResp.StatusCode, requestID, err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(resp.AccessToken) == "" {
		return nil, fmt.Errorf("RefreshToken succeeded but access_token was empty")
	}
	if resp.ExpiresIn <= 0 {
		return nil, fmt.Errorf("RefreshToken succeeded but expires_in was missing or invalid")
	}
	return &resp, nil
}
