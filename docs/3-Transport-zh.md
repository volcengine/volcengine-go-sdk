[← Endpoint 配置](2-Endpoint-zh.md) | Transport[(English)](3-Transport.md) | [超时配置 →](4-Timeout-zh.md)

---

## HTTP 连接池配置

> **默认**
>
> - 最大空闲连接数（`MaxIdleConns`）：100
> - 空闲连接存活时间（`IdleConnTimeout`）：90s
> - 每个路由最大连接数（`MaxIdleConnsPerHost`）：2

最大空闲连接数、空闲连接存活时间、每个路由最大连接数没有直接提供参数设置，可以通过自定义 HTTPClient 实现。

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
		MaxIdleConns:          100,              // 所有host加起来的最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接最大存活时间
		MaxIdleConnsPerHost:   10,               // 每个host（路由）最大空闲连接数
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second, // 设置ReadTimeout
	}
	config := volcengine.NewConfig().
		WithRegion(region).
		WithHTTPClient(client).
		WithCredentials(credentials.NewEnvCredentials()) //需要配置环境变量

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

## HTTPS 请求配置

### 指定 scheme

> **默认**
>
> - HTTPS

scheme，为 `true` 表示 scheme 为 http，为 `false` 表示 scheme 为 https；建议使用 HTTPS，这样可以提升数据传输的安全性。若不设置，则使用默认支持的请求协议类型（HTTPS）。

```go
func main() {
	region := "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithDisableSSL(true).                            //true 表示scheme为http，false表示为https，默认为false
		WithCredentials(credentials.NewEnvCredentials()) // 需要配置环境变量

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

### 忽略 SSL 验证

> **默认**
>
> - 不忽略（开启 SSL 认证）

没有直接提供参数设置，可以通过自定义 HTTPClient 实现。

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
		MaxIdleConns:          100,              // 所有host加起来的最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接最大存活时间
		MaxIdleConnsPerHost:   10,               // 每个host（路由）最大空闲连接数
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, //跳过服务器证书校验，即使证书无效也继续访问
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second, // 设置ReadTimeout
	}
	config := volcengine.NewConfig().
		WithRegion(region).
		WithHTTPClient(client).
		WithCredentials(credentials.NewEnvCredentials()) //需要配置环境变量

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

### 指定 TLS 协议版本

> **默认**
>
> - ≥ TLS 1.2

目前只支持自定义 HTTPClient 的方式实现；如果没有特殊要求，建议不要修改。

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
		MaxIdleConns:          100,              // 所有host加起来的最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接最大存活时间
		MaxIdleConnsPerHost:   10,               // 每个host（路由）最大空闲连接数
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12, // 最小 TLS1.2
			MaxVersion: tls.VersionTLS13, // 最大 TLS1.3
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second, // 设置ReadTimeout
	}
	config := volcengine.NewConfig().
		WithRegion(region).
		WithHTTPClient(client).
		WithCredentials(credentials.NewEnvCredentials()) //需要配置环境变量

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

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

---

[← Endpoint 配置](2-Endpoint-zh.md) | Transport[(English)](3-Transport.md) | [超时配置 →](4-Timeout-zh.md)
