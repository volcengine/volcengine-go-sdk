[← Transport](3-Transport-zh.md) | 超时配置[(English)](4-Timeout.md) | [重试机制 →](5-Retry-zh.md)

---

## 超时配置

### 全局超时设置（Client 级别）

> **默认**
>
> - 默认 HTTP Client：`http.DefaultClient`（由 `defaults/defaults.go:WithHTTPClient` 设置）
> - `ConnectTimeout` - 30s
> - `ReadTimeout` / 整体请求超时 - 不限制（`http.DefaultClient.Timeout == 0`）

暂不支持直接设置 `ConnectTimeOut` 和 `ReadTimeout` 配置，可以通过自定义 HttpClient 来实现。

```go
func main() {
	region := "cn-beijing"
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 设置ConnectTimeOut
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
		Timeout:   60 * time.Second, // 设置ReadTimeout
	}
	config := volcengine.NewConfig().
		WithRegion(region).
		WithHTTPClient(client).
		WithCredentials(credentials.NewEnvCredentials()) // 需要配置环境变量

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

### 单接口指定超时设置

单个接口超时配置需要调用以 **WithContext** 结尾的接口。

```go
func main() {
	region := "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewEnvCredentials()) //需要配置环境变量
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	associateInstancesIamRoleInput := &ecs.AssociateInstancesIamRoleInput{
		IamRoleName: volcengine.String("EcsTestRole"),
		InstanceIds: volcengine.StringSlice([]string{"i-3tiefmkskq3vj0******"}),
	}

	// 创建带5秒超时的上下文；注意：如果已经存在上下文context，请把这里context.Background()替换为已有的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 防止资源泄漏
	defer cancel()

	// 调用WithContext结尾的接口
	resp, err := svc.AssociateInstancesIamRoleWithContext(ctx, associateInstancesIamRoleInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

---

[← Transport](3-Transport-zh.md) | 超时配置[(English)](4-Timeout.md) | [重试机制 →](5-Retry-zh.md)
