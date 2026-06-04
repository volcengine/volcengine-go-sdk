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

---

[← Transport](3-Transport.md) | Proxy[(中文)](4-Proxy-zh.md) | [Timeout →](5-Timeout.md)
