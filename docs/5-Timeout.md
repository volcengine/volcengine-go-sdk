[← Proxy](4-Proxy.md) | Timeout[(中文)](5-Timeout-zh.md) | [Retry →](6-Retry.md)

---

## Timeouts

### Global Timeouts (Client Level)

> **Default**
>
> - Default client: `http.DefaultClient` (set by `defaults/defaults.go:WithHTTPClient`)
> - `ConnectTimeout`: 30s (which uses `net.Dialer{Timeout: 30s}`)
> - `ReadTimeout` / overall request timeout: unlimited (`http.DefaultClient.Timeout == 0`)

Configure timeouts via a custom `http.Client`.

```go
func main() {
	region := "cn-beijing"
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second,
	}
	config := volcengine.NewConfig().
		WithRegion(region).
		WithHTTPClient(client).
		WithCredentials(credentials.NewEnvCredentials())

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

### Per-API Timeout

Use API methods with suffix `WithContext`.

```go
func main() {
	region := "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewEnvCredentials())
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := svc.AssociateInstancesIamRoleWithContext(ctx, associateInstancesIamRoleInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

---

[← Proxy](4-Proxy.md) | Timeout[(中文)](5-Timeout-zh.md) | [Retry →](6-Retry.md)
