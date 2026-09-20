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

显式传入的内置凭证 Provider 在 `NewSession`、`NewSessionWithOptions` 和已弃用的 `New`
中绑定独立的配置快照、凭证缓存和刷新锁，源 Provider 不会被回写。

`Session.Copy`（包括服务客户端创建）在凭证指针以及 `Region`、`HTTPProxy`、`HTTPSProxy`
的值与上次绑定一致时，复用 Session 已绑定的 `Credentials`、缓存和刷新锁。
无配置覆盖、nil/空配置、使用新指针但值相等的配置，以及仅修改业务 `Endpoint` 等无关字段，
都保留共享。任一绑定字段变化或替换凭证时，创建独立绑定，不回写源 Session。
从新绑定继续派生且默认配置相同的客户端共享该绑定。
不同 Session、直接调用 `Credentials.WithSTSConfig` 仍各自独立；
手工构造且没有绑定记录的 `Session` 在首次 Copy 时执行绑定。

请在创建客户端之前完成 Provider 和 Session 配置，不要在使用过程中并发修改。
之后修改源 Provider 不影响已绑定实例；Session 默认配置变化会在下一次 Copy 或创建客户端时
被识别，但不会重新配置已有客户端。对共享凭证调用 `Expire()` 会使所有共享该绑定的客户端
缓存一同失效，下一次获取凭证时通过共享锁协调刷新。

代理优先级：Provider 的显式 `ProxyURL` > 与 **STS 请求自身协议** 对应的
`WithHTTPProxy` / `WithHTTPSProxy` > 原传输层代理行为。环境变量不会被转存到
`ProxyURL`，标准默认传输层继续处理 `HTTP_PROXY` / `HTTPS_PROXY` / `NO_PROXY`。
显式代理不受环境 `NO_PROXY` 覆盖。未覆盖的协议继续使用原代理规则。
只有当前 STS 请求协议选中的客户端代理才会被解析和校验；另一协议的非法配置
不会阻断当前请求。选中的非法代理仍会报错，不会静默绕过代理。
Provider 的显式 `ProxyURL` 同时覆盖两种协议，仍在构建传输层时校验。
例如 HTTPS STS 通过 HTTP CONNECT 代理时，应配置
`WithHTTPSProxy("http://proxy:8080")`，与业务 API 的 `DisableSSL` 无关。

未设置代理/TLS 选项时，OIDC/SAML 保留 `http.Client.Transport == nil`，
继续使用调用方的 `http.DefaultTransport`；AssumeRole 保留上游客户端传输层。
显式覆盖会克隆支持的 `*http.Transport` 并保留未覆盖的 TLS 设置。
自定义 RoundTripper 无法安全叠加配置时返回错误，不会 panic。
这些绑定规则适用于显式传入的内置 Provider，不改变第三方 Provider 或聚合凭证链的契约。

---

[← Transport](3-Transport-zh.md) | 代理配置[(English)](4-Proxy.md) | [超时配置 →](5-Timeout-zh.md)
