[← Transport](3-Transport.md) | Proxy[(中文)](4-Proxy-zh.md) | [Timeout →](5-Timeout.md)

---

## HTTP(S) Proxy

> **Default**
>
> - No proxy

### Configure HTTP(S) Proxy

```go
var ak, sk, region string
config = volcengine.NewConfig().
	WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
	WithRegion(region).WithHTTPProxy("http://your_proxy:8080").WithHTTPSProxy("http://your_proxy:8080")

sess, _ = session.NewSession(config)
client = ecs.New(sess)
```

### Notes

Supported environment variables:

- `http_proxy` / `HTTP_PROXY`
- `https_proxy` / `HTTPS_PROXY`

Priority: code > environment variables.

### STS / OIDC / SAML credential requests

`NewSession`, `NewSessionWithOptions`, and `Session.Copy` (including service-client
construction) bind built-in providers to independent configuration snapshots,
credential caches and refresh locks. The source provider is not modified. Finish
configuration before constructing the session/client; subsequent source-provider
changes do not affect existing bindings.

Priority: explicit provider `ProxyURL` > `WithHTTPProxy` / `WithHTTPSProxy` matching
the **STS request's scheme** > the original transport's proxy behavior. Environment
values are never stored in `ProxyURL`; the standard default transport continues
to honor `HTTP_PROXY`, `HTTPS_PROXY` and `NO_PROXY`. Explicit proxies take precedence
over environment `NO_PROXY`; unspecified schemes retain the original proxy rule.
Only the client proxy selected for the current STS request's scheme is parsed
and validated. An invalid proxy for the other scheme does not block the request.
An invalid selected proxy still returns an error instead of silently bypassing
the proxy. An explicit provider `ProxyURL` covers both schemes and is still
validated when constructing the transport.
For HTTPS STS through an HTTP CONNECT proxy, use
`WithHTTPSProxy("http://proxy:8080")`, independently of business API `DisableSSL`.

With no overrides, OIDC/SAML retain a nil `http.Client.Transport` and use the caller's
`http.DefaultTransport`; AssumeRole retains its upstream client transport. Explicit
overrides clone supported `*http.Transport` values, preserving unspecified TLS settings.
Unsupported custom RoundTripper overrides return an error instead of panicking.
These rules cover explicitly supplied built-in providers, not arbitrary providers
or children of aggregate credential chains.

---

[← Transport](3-Transport.md) | Proxy[(中文)](4-Proxy-zh.md) | [Timeout →](5-Timeout.md)
