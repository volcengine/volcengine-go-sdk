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
	"sync"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

// ConsoleLoginRefreshableProvider 提供 console-login 模式的凭证解析与续期。
//
// 契约：
//   - 永不写磁盘。读 login cache 文件只发生在 (a) bootstrap 第一次解析；
//     (b) refresh_token 被服务端判定为 invalid_grant 时触发的 fallback。
//   - access_token 内含 STS 三元组，未过期时直接返回；过期时用 cache 中的
//     refresh_token 调 OAuth refresh，更新内存里的 cache + STS。
//   - refresh_token 失效（服务端 400 invalid_grant）时按 fallback 流程：
//     重读磁盘 cache → 若磁盘 RT 与内存不同则用磁盘 RT 重试一次 →
//     仍失败抛错并提示 've login'。
//   - 进程内并发使用 mu 串行化 refresh，避免重复刷新或竞态。
type ConsoleLoginRefreshableProvider struct {
	loginSession string
	cacheDir     string
	// oauthClientFactory 接收 endpoint URL 返回 OAuth 客户端，便于在测试中
	// 注入 httptest.Server。生产路径使用 defaultConsoleLoginOAuthFactory。
	oauthClientFactory func(endpointURL string) ConsoleLoginOAuthClientAPI

	mu          sync.Mutex
	cached      *LoginTokenCache
	creds       credentials.Value
	expiration  time.Time
	initialized bool
}

// newConsoleLoginRefreshableProvider 构造默认配置的 provider。
func newConsoleLoginRefreshableProvider(loginSession, cacheDir string) *ConsoleLoginRefreshableProvider {
	return &ConsoleLoginRefreshableProvider{
		loginSession:       loginSession,
		cacheDir:           cacheDir,
		oauthClientFactory: defaultConsoleLoginOAuthFactory,
	}
}

// defaultConsoleLoginOAuthFactory 生产环境使用，按 cache 中的 endpoint_url
// 实例化 OAuth 客户端；空字符串时回退到 defaultConsoleLoginEndpoint。
func defaultConsoleLoginOAuthFactory(endpointURL string) ConsoleLoginOAuthClientAPI {
	return NewConsoleLoginOAuthClient(&ConsoleLoginOAuthClientConfig{
		EndpointURL: endpointURL,
	})
}

// Retrieve 返回当前可用的凭证；如内存状态过期则触发一次同步 refresh。
func (p *ConsoleLoginRefreshableProvider) Retrieve() (credentials.Value, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.initialized || p.isExpiredLocked() {
		if err := p.refreshLocked(); err != nil {
			return credentials.Value{ProviderName: CliProviderName}, err
		}
	}
	return p.creds, nil
}

// IsExpired 描述"下次 Retrieve 是否需要 refresh"。
func (p *ConsoleLoginRefreshableProvider) IsExpired() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if !p.initialized {
		return true
	}
	return p.isExpiredLocked()
}

func (p *ConsoleLoginRefreshableProvider) isExpiredLocked() bool {
	if p.expiration.IsZero() {
		return true
	}
	// 保留 1 分钟安全余量，避免边界期返回即将过期的 STS。
	return !time.Now().Before(p.expiration.Add(-time.Minute))
}

// refreshLocked 是核心刷新流程；调用方必须已持有 p.mu。
//
// 流程：
//  1. 若内存 cache 不存在，从磁盘 bootstrap 一份；磁盘读取失败直接报错。
//  2. 若内存 cache 中的 access_token 仍未过期，直接解析复用，不调网络。
//  3. 否则用 RT 调 OAuth refresh。
//  4. refresh 因 invalid_grant 失败时触发 fallback：重读磁盘 → 若 RT 不同
//     则用磁盘 RT 重试一次；磁盘 RT 与内存相同则直接报错。
func (p *ConsoleLoginRefreshableProvider) refreshLocked() error {
	if p.cached == nil {
		cache, err := p.loadCacheFromDisk()
		if err != nil {
			return err
		}
		p.cached = cache
	}

	// 快速路径：access_token 尚未过期。
	if exp, ok := tryParseConsoleLoginAccessTokenAndExpiration(p.cached); ok {
		p.applyCredentials(exp)
		return nil
	}

	// 慢路径：调 OAuth refresh。
	err := p.refreshTokenWithCachedRT(context.Background())
	if err == nil {
		return nil
	}

	// fallback：仅在"RT 被服务端拒绝"时触发，避免网络抖动误判。
	if !isRefreshTokenInvalidErr(err) {
		return err
	}
	diskCache, diskErr := p.loadCacheFromDisk()
	if diskErr != nil {
		return wrapConsoleLoginRefreshErr(err, diskErr)
	}
	if strings.TrimSpace(diskCache.RefreshToken) == "" ||
		diskCache.RefreshToken == p.cached.RefreshToken {
		// 磁盘上的 RT 与内存相同（或更糟），CLI 没刷过，确实只能让用户重登。
		return volcengineerr.New(
			"CliConsoleLoginRefreshTokenExpired",
			"console-login refresh token has been rejected by signin service; please run 've login' to re-authenticate",
			err,
		)
	}
	// 用磁盘 RT 整体替换内存 cache，再试一次。
	p.cached = diskCache
	if exp, ok := tryParseConsoleLoginAccessTokenAndExpiration(p.cached); ok {
		p.applyCredentials(exp)
		return nil
	}
	if err2 := p.refreshTokenWithCachedRT(context.Background()); err2 != nil {
		return volcengineerr.New(
			"CliConsoleLoginRefreshTokenExpired",
			"console-login refresh token rejected; reloaded cache from disk but new refresh token also failed; please run 've login'",
			err2,
		)
	}
	return nil
}

// refreshTokenWithCachedRT 持锁调用 OAuth /token 续期；成功时更新内存 cache 与 STS。
func (p *ConsoleLoginRefreshableProvider) refreshTokenWithCachedRT(ctx context.Context) error {
	if p.cached == nil || strings.TrimSpace(p.cached.RefreshToken) == "" {
		return volcengineerr.New(
			"CliConsoleLoginRefreshTokenMissing",
			fmt.Sprintf("console-login cache for login_session %q does not contain refresh_token; please run 've login' first", p.loginSession),
			nil,
		)
	}
	if strings.TrimSpace(p.cached.ClientID) == "" {
		return volcengineerr.New(
			"CliConsoleLoginClientIDMissing",
			fmt.Sprintf("console-login cache for login_session %q does not contain client_id; please run 've login' to regenerate", p.loginSession),
			nil,
		)
	}

	client := p.oauthClientFactory(p.cached.EndpointURL)
	resp, err := client.RefreshToken(ctx, &ConsoleLoginRefreshTokenRequest{
		ClientID:     p.cached.ClientID,
		Scope:        p.cached.Scope,
		RefreshToken: p.cached.RefreshToken,
	})
	if err != nil {
		// Pass invalid_grant through verbatim so refreshLocked can detect it
		// and trigger the disk-reload fallback. Wrap all other errors with a
		// stable error code and the actionable hint.
		if isRefreshTokenInvalidErr(err) {
			return err
		}
		return volcengineerr.New(
			"CliConsoleLoginRefreshTokenFailed",
			"failed to refresh console-login access token; please run 've login' to re-authenticate",
			err,
		)
	}

	// 更新内存 cache：access_token 必更，refresh_token 仅当服务端返回新值时
	// 更新（与 ve cli EnsureValidLoginToken 行为对齐）。
	p.cached.AccessToken = json.RawMessage(resp.AccessToken)
	if strings.TrimSpace(resp.RefreshToken) != "" {
		p.cached.RefreshToken = resp.RefreshToken
	}
	p.cached.IssuedAt = time.Now().UTC().Format(time.RFC3339)
	p.cached.ExpiresIn = int64(resp.ExpiresIn)
	if strings.TrimSpace(resp.TokenType) != "" {
		p.cached.TokenType = resp.TokenType
	}

	exp, ok := tryParseConsoleLoginAccessTokenAndExpiration(p.cached)
	if !ok {
		return volcengineerr.New(
			"CliConsoleLoginAccessTokenInvalid",
			"console-login refresh succeeded but the new access_token could not be parsed into STS credentials",
			nil,
		)
	}
	p.applyCredentials(exp)
	return nil
}

// applyCredentials 把解析后的 STS 写入 p.creds 并设置 expiration。
func (p *ConsoleLoginRefreshableProvider) applyCredentials(exp time.Time) {
	creds, err := parseConsoleLoginAccessToken(p.cached.AccessToken)
	if err != nil {
		// 调用方已经在 tryParseConsoleLoginAccessTokenAndExpiration 校验过，
		// 不应该走到这里；写空值返回让上层报错而非 panic。
		p.creds = credentials.Value{ProviderName: CliProviderName}
		p.expiration = time.Time{}
		p.initialized = true
		return
	}
	p.creds = credentials.Value{
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
		ProviderName:    CliProviderName,
	}
	p.expiration = exp
	p.initialized = true
}

// loadCacheFromDisk 从 cacheDir 读取 login cache 文件并解码。
func (p *ConsoleLoginRefreshableProvider) loadCacheFromDisk() (*LoginTokenCache, error) {
	cachePath, err := p.cachePath()
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(cachePath); err != nil {
		if os.IsNotExist(err) {
			return nil, volcengineerr.New(
				"CliConsoleLoginCacheMissing",
				fmt.Sprintf("console-login token cache file %s does not exist; please run 've login' first", cachePath),
				err,
			)
		}
		return nil, volcengineerr.New(
			"CliConsoleLoginCacheStat",
			fmt.Sprintf("failed to stat console-login token cache file %s", cachePath),
			err,
		)
	}
	data, err := ioutil.ReadFile(cachePath)
	if err != nil {
		return nil, volcengineerr.New(
			"CliConsoleLoginCacheLoad",
			fmt.Sprintf("failed to load console-login token cache file %s", cachePath),
			err,
		)
	}
	if len(strings.TrimSpace(string(data))) == 0 {
		return nil, volcengineerr.New(
			"CliConsoleLoginCacheEmpty",
			fmt.Sprintf("console-login token cache file %s was empty", cachePath),
			nil,
		)
	}
	var cache LoginTokenCache
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, volcengineerr.New(
			"CliConsoleLoginCacheUnmarshal",
			fmt.Sprintf("failed to unmarshal console-login token cache file %s", cachePath),
			err,
		)
	}
	return &cache, nil
}

// cachePath 复用 sha1(login_session) 命名规则，与 ve cli 保持一致。
func (p *ConsoleLoginRefreshableProvider) cachePath() (string, error) {
	if strings.TrimSpace(p.cacheDir) == "" {
		return "", volcengineerr.New(
			"CliConsoleLoginCacheDirMissing",
			fmt.Sprintf("console-login provider for login_session %q has no cache directory", p.loginSession),
			nil,
		)
	}
	hash := sha1.Sum([]byte(p.loginSession))
	return filepath.Join(p.cacheDir, fmt.Sprintf("%x.json", hash)), nil
}

// tryParseConsoleLoginAccessTokenAndExpiration 尝试解析 cache 的 access_token
// 并按 issued_at + expires_in 计算到期时间；任一缺失或过期则返回 false。
func tryParseConsoleLoginAccessTokenAndExpiration(cache *LoginTokenCache) (time.Time, bool) {
	if cache == nil {
		return time.Time{}, false
	}
	exp, err := consoleLoginCacheExpiration(cache)
	if err != nil {
		return time.Time{}, false
	}
	if !time.Now().Before(exp.Add(-time.Minute)) {
		return time.Time{}, false
	}
	if _, err := parseConsoleLoginAccessToken(cache.AccessToken); err != nil {
		return time.Time{}, false
	}
	return exp, true
}

// isRefreshTokenInvalidErr 抽出 unwrap 逻辑，便于上层只判定语义层错误。
func isRefreshTokenInvalidErr(err error) bool {
	if err == nil {
		return false
	}
	type unwrapper interface {
		Unwrap() error
	}
	cur := err
	for cur != nil {
		if apiErr, ok := cur.(*ConsoleLoginOAuthAPIError); ok {
			return apiErr.IsRefreshTokenInvalid()
		}
		uw, ok := cur.(unwrapper)
		if !ok {
			break
		}
		cur = uw.Unwrap()
	}
	return false
}

// wrapConsoleLoginRefreshErr 把"refresh 失败 + 重读磁盘也失败"的双重错误整合成
// 一条可观测的提示。
func wrapConsoleLoginRefreshErr(refreshErr, diskErr error) error {
	return volcengineerr.New(
		"CliConsoleLoginRefreshTokenExpired",
		fmt.Sprintf("console-login refresh token rejected (%v); failed to reload cache from disk: %v; please run 've login'", refreshErr, diskErr),
		refreshErr,
	)
}
