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

`NewSession`, `NewSessionWithOptions`, and the deprecated `New` bind explicitly
supplied built-in providers to independent configuration snapshots, credential
caches and refresh locks. The source provider is not modified.

`Session.Copy` (including service-client construction) reuses the session's bound
`Credentials`, cache and refresh lock when the credentials pointer and the values
of `Region`, `HTTPProxy`, and `HTTPSProxy` match the last binding. No overrides,
nil/empty configs, equal values in new pointers, and business-only changes such as
`Endpoint` retain this sharing. Changing any binding field or replacing the
credentials creates an independent binding without modifying the source session.
Clients subsequently derived from that binding with the same defaults share it.
Separate sessions and direct `Credentials.WithSTSConfig` calls remain independent;
a manually constructed `Session` without a binding record binds on its first copy.

Finish provider and session configuration before creating clients, and do not
modify it concurrently with use. Later source-provider changes do not affect
existing bindings; changes to session defaults are detected on the next copy or
client construction, but do not reconfigure existing clients. Calling `Expire()`
on shared credentials invalidates the cache for all clients sharing that binding;
the next retrieval refreshes it under the shared lock.

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
