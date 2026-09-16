[← Transport](3-Transport-zh.md) | 代理配置[(English)](4-Proxy.md) | [超时配置 →](5-Timeout-zh.md)

---

## HTTP(S) 代理配置

> **默认**
>
> - 无代理

### 配置 HTTP(S) 代理

```go
var ak, sk, region string
config = volcengine.NewConfig().
	WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
	WithRegion(region).WithHTTPProxy("http://your_proxy:8080").WithHTTPSProxy("http://your_proxy:8080")

sess, _ = session.NewSession(config)
client = ecs.New(sess)
```

### 注意事项

支持通过以下环境变量配置代理：

- `http_proxy` / `HTTP_PROXY`
- `https_proxy` / `HTTPS_PROXY`

优先级：代码指定 > 环境变量。

### STS / OIDC / SAML 凭证请求

内置凭证 Provider 在 `NewSession`、`NewSessionWithOptions` 和 `Session.Copy`
（包括服务客户端创建）时绑定独立的配置快照、凭证缓存和刷新锁，源 Provider 不会被回写。
请在创建 session 或服务客户端之前完成配置；之后修改源 Provider 不影响已经绑定的实例。

代理优先级：Provider 的显式 `ProxyURL` > 与 **STS 请求自身协议** 对应的
`WithHTTPProxy` / `WithHTTPSProxy` > 原传输层代理行为。环境变量不会被转存到
`ProxyURL`，标准默认传输层继续处理 `HTTP_PROXY` / `HTTPS_PROXY` / `NO_PROXY`。
显式代理不受环境 `NO_PROXY` 覆盖。未覆盖的协议继续使用原代理规则。
例如 HTTPS STS 通过 HTTP CONNECT 代理时，应配置
`WithHTTPSProxy("http://proxy:8080")`，与业务 API 的 `DisableSSL` 无关。

未设置代理/TLS 选项时，OIDC/SAML 保留 `http.Client.Transport == nil`，
继续使用调用方的 `http.DefaultTransport`；AssumeRole 保留上游客户端传输层。
显式覆盖会克隆支持的 `*http.Transport` 并保留未覆盖的 TLS 设置。
自定义 RoundTripper 无法安全叠加配置时返回错误，不会 panic。
这些绑定规则适用于显式传入的内置 Provider，不改变第三方 Provider 或聚合凭证链的契约。

---

[← Transport](3-Transport-zh.md) | 代理配置[(English)](4-Proxy.md) | [超时配置 →](5-Timeout-zh.md)
