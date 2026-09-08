package credentials

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// ProxyConfig 汇总访问 STS 端点时的代理与 TLS 相关配置。
//
// 该结构用于把 ProxyURL、CAFile、InsecureSkipVerify 三个对外参数统一
// 传递给 buildProxyTransport / buildProxyHTTPClient，便于各 Provider
// 复用同一套构建逻辑。
type ProxyConfig struct {
	// ProxyURL 正向代理地址，格式 scheme://[user:password@]host:port。
	// 为空时回退到环境变量 HTTP_PROXY / HTTPS_PROXY / NO_PROXY。
	ProxyURL string
	// CAFile 自定义 CA 证书文件路径（PEM 格式），用于校验 STS 服务端证书。
	// 为空时使用操作系统默认根证书信任库。
	CAFile string
	// InsecureSkipVerify 为 true 时跳过 TLS 证书链与主机名校验。
	// 仅限联调，生产环境禁用。
	InsecureSkipVerify bool
}

// IsZero 判断三个对外参数是否都为空。
//
// 为 true 时表示调用方未显式配置任何代理/TLS 选项，此时应保持底层
// http.Client 原样（例如上游 base.Client 复用的 _GlobalClient 及其
// VOLC_ 环境变量代理），避免无意义地覆盖 Transport。
func (cfg ProxyConfig) IsZero() bool {
	return cfg.ProxyURL == "" && cfg.CAFile == "" && !cfg.InsecureSkipVerify
}

// buildProxyTransport 根据代理与 TLS 配置构建 *http.Transport。
//
// 仅当 cfg 中至少显式配置了一项（ProxyURL / CAFile / InsecureSkipVerify）
// 时才返回非 nil 的 transport；三者全空时返回 (nil, nil)，由调用方保留
// 原有传输层配置。
//
// 代理优先级：
//  1. 若 cfg.ProxyURL 非空，则使用该地址作为 HTTP/HTTPS 代理；
//  2. 否则回退到环境变量（HTTP_PROXY / HTTPS_PROXY / NO_PROXY）。
//
// TLS 处理：
//   - 若 cfg.CAFile 非空，则以该 CA 校验 STS 服务端证书链；
//   - 若 cfg.InsecureSkipVerify 为 true，则跳过证书校验（仅限联调场景）。
func buildProxyTransport(cfg ProxyConfig) (*http.Transport, error) {
	if cfg.IsZero() {
		return nil, nil
	}

	proxyFunc := http.ProxyFromEnvironment
	if cfg.ProxyURL != "" {
		parsed, err := url.Parse(cfg.ProxyURL)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy url %q: %w", cfg.ProxyURL, err)
		}
		proxyFunc = http.ProxyURL(parsed)
	}

	tlsCfg := &tls.Config{
		InsecureSkipVerify: cfg.InsecureSkipVerify,
	}
	if cfg.CAFile != "" {
		pemBytes, err := ioutil.ReadFile(cfg.CAFile)
		if err != nil {
			return nil, fmt.Errorf("read CAFile %q: %w", cfg.CAFile, err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pemBytes) {
			return nil, fmt.Errorf("append CA cert from %q failed", cfg.CAFile)
		}
		tlsCfg.RootCAs = pool
	}

	return &http.Transport{
		Proxy:               proxyFunc,
		TLSClientConfig:     tlsCfg,
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     10 * time.Second,
	}, nil
}

// buildProxyHTTPClient 根据代理与 TLS 配置构建独立的 *http.Client。
//
// 适用于没有可复用底层 client 的独立请求路径（如 OIDC 的
// doSTSFormRequest）。即便 cfg 为空也会返回一个带默认 transport 的
// client，其代理回退到标准环境变量（HTTP_PROXY / HTTPS_PROXY / NO_PROXY）。
//
// 当 timeout > 0 时，会设置到返回的 http.Client 上。
func buildProxyHTTPClient(cfg ProxyConfig, timeout time.Duration) (*http.Client, error) {
	transport, err := buildProxyTransport(cfg)
	if err != nil {
		return nil, err
	}
	if transport == nil {
		// cfg 为空时也需要一个可用 transport（含连接池与环境变量代理回退），
		// 并保留一个默认 TLS 配置以维持原有行为。
		transport = &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			TLSClientConfig:     &tls.Config{},
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     10 * time.Second,
		}
	}

	client := &http.Client{Transport: transport}
	if timeout > 0 {
		client.Timeout = timeout
	}
	return client, nil
}
