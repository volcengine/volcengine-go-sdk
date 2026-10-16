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
	// 为空时保留原 Transport 的代理行为；标准默认 Transport 会处理
	// HTTP_PROXY / HTTPS_PROXY / NO_PROXY。
	ProxyURL string
	// CAFile 自定义 CA 证书文件路径（PEM 格式），用于校验 STS 服务端证书。
	// 为空时保留原 Transport 的 TLS 根配置；未自定义时使用系统根证书。
	CAFile string
	// InsecureSkipVerify 为 true 时跳过 TLS 证书链与主机名校验。
	// 仅限联调，生产环境禁用。
	InsecureSkipVerify bool
	// Client fallback and the original transport are kept separate from the
	// explicitly configured ProxyURL. Environment proxies are never copied here.
	httpProxy  string
	httpsProxy string
	base       http.RoundTripper
}

// IsZero 判断显式代理/TLS 参数和客户端代理兜底是否都为空。
//
// 为 true 时表示调用方未显式配置任何代理/TLS 选项，此时应保持底层
// http.Client 原样（例如上游 base.Client 复用的 _GlobalClient 及其
// VOLC_ 环境变量代理），避免无意义地覆盖 Transport。
func (cfg ProxyConfig) IsZero() bool {
	return cfg.ProxyURL == "" && cfg.CAFile == "" && !cfg.InsecureSkipVerify &&
		cfg.httpProxy == "" && cfg.httpsProxy == ""
}

// newDefaultProxyTransport 克隆 http.DefaultTransport 并叠加 SDK 的连接池调优。
//
// 直接 new 一个仅设置少量字段的 &http.Transport{} 会丢失标准库的关键默认值
// （DialContext 拨号超时、TLSHandshakeTimeout、ExpectContinueTimeout、
// ForceAttemptHTTP2 等）。特别是一旦显式设置了非 nil 的 TLSClientConfig，
// 若 ForceAttemptHTTP2 未置真，net/http 将不再自动协商 HTTP/2 而静默降级到
// HTTP/1.1。因此这里以 Clone() 为基准，既保留标准库默认，又只覆盖需要定制的
// 连接池参数（代理与 TLS 由调用方在此基础上覆盖）。
func newDefaultProxyTransport(base http.RoundTripper) (*http.Transport, error) {
	if base == nil {
		base = http.DefaultTransport
	}
	source, ok := base.(*http.Transport)
	if !ok || source == nil {
		return nil, fmt.Errorf("cannot apply STS proxy/TLS options to transport %T", base)
	}
	transport := source.Clone()
	// Clone may copy the h2 ALPN protocol initialized on the source while
	// leaving TLSNextProto to be configured on the clone. Opt in explicitly:
	// otherwise its non-nil TLSClientConfig can disable automatic HTTP/2,
	// causing HTTP/1 requests to be sent over an h2-negotiated connection.
	transport.ForceAttemptHTTP2 = true
	transport.MaxIdleConns = 1000
	transport.MaxIdleConnsPerHost = 100
	transport.IdleConnTimeout = 10 * time.Second
	return transport, nil
}

// buildProxyTransport 根据代理与 TLS 配置构建 *http.Transport。
//
// 仅当 cfg 中配置了代理/TLS 选项或客户端代理兜底时才返回非 nil 的
// transport；cfg.IsZero() 时返回 (nil, nil)，由调用方保留原有传输层配置。
//
// transport 以 http.DefaultTransport.Clone() 为基准，保留标准库的超时与
// 协议默认值（TLSHandshakeTimeout / ExpectContinueTimeout / ForceAttemptHTTP2
// 等），再覆盖代理与 TLS，避免出现握手挂起或 HTTP/2 静默降级。
//
// 代理优先级：
//  1. 若 cfg.ProxyURL 非空，则使用该地址作为 HTTP/HTTPS 代理；
//  2. 否则按实际请求 scheme 选择并校验客户端显式代理；
//  3. 未命中时保留原 transport.Proxy（标准默认值支持环境变量和 NO_PROXY）。
//
// TLS 处理：
//   - 若 cfg.CAFile 非空，则以该 CA 校验 STS 服务端证书链；
//   - 若 cfg.InsecureSkipVerify 为 true，则跳过证书校验（仅限联调场景）。
func buildProxyTransport(cfg ProxyConfig) (*http.Transport, error) {
	if cfg.IsZero() {
		return nil, nil
	}

	transport, err := newDefaultProxyTransport(cfg.base)
	if err != nil {
		return nil, err
	}
	if cfg.ProxyURL != "" {
		parsed, err := parseSTSProxy(cfg.ProxyURL)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(parsed)
	} else if cfg.httpProxy != "" || cfg.httpsProxy != "" {
		baseProxy := transport.Proxy
		transport.Proxy = func(req *http.Request) (*url.URL, error) {
			var proxy string
			switch req.URL.Scheme {
			case "http":
				proxy = cfg.httpProxy
			case "https":
				proxy = cfg.httpsProxy
			}
			if proxy != "" {
				return parseSTSProxy(proxy)
			}
			if baseProxy != nil {
				return baseProxy(req)
			}
			return nil, nil
		}
	}

	// Clone preserves caller TLS settings, including mTLS and ServerName.
	if cfg.CAFile == "" && !cfg.InsecureSkipVerify {
		return transport, nil
	}
	tlsCfg := transport.TLSClientConfig
	if tlsCfg == nil {
		tlsCfg = &tls.Config{}
	}
	tlsCfg.InsecureSkipVerify = cfg.InsecureSkipVerify
	if cfg.CAFile != "" {
		pemBytes, err := ioutil.ReadFile(cfg.CAFile)
		if err != nil {
			return nil, fmt.Errorf("read CAFile %q: %w", cfg.CAFile, err)
		}
		// 显式指定 CAFile 时以该 CA 单独构建信任池（不追加系统根证书）；
		// 留空时保留原 Transport 的 TLS 根配置。
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pemBytes) {
			return nil, fmt.Errorf("append CA cert from %q failed", cfg.CAFile)
		}
		tlsCfg.RootCAs = pool
	}

	transport.TLSClientConfig = tlsCfg
	return transport, nil
}

func parseSTSProxy(value string) (*url.URL, error) {
	parsed, err := url.Parse(value)
	if err != nil || parsed.Hostname() == "" {
		// Parser errors may contain proxy passwords.
		return nil, fmt.Errorf("invalid STS proxy URL")
	}
	switch parsed.Scheme {
	case "http", "https", "socks5", "socks5h":
		return parsed, nil
	default:
		return nil, fmt.Errorf("unsupported STS proxy scheme %q", parsed.Scheme)
	}
}

// buildProxyHTTPClient 根据代理与 TLS 配置构建独立的 *http.Client。
//
// 适用于没有可复用底层 client 的独立请求路径（如 OIDC 的
// doSTSFormRequest）。cfg 为空时保留 nil Transport，沿用原 http.DefaultTransport。
//
// 当 timeout > 0 时，会设置到返回的 http.Client 上。
func buildProxyHTTPClient(cfg ProxyConfig, timeout time.Duration) (*http.Client, error) {
	transport, err := buildProxyTransport(cfg)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	if transport != nil {
		client.Transport = transport
	}
	if timeout > 0 {
		client.Timeout = timeout
	}
	return client, nil
}
