[← 超时配置](4-Timeout-zh.md) | 重试机制[(English)](5-Retry.md) | [异常处理 →](6-ErrorHandling-zh.md)

---

# 重试机制

请求的处理逻辑内置了网络异常重试逻辑，即当遇到网络异常问题或限流错误时，系统会自动尝试重新发起请求，以确保服务的稳定性和可靠性。若请求因业务逻辑错误而报错，例如参数错误、资源不存在等情况，SDK将不会执行重试操作，这是因为业务层面的错误通常需要应用程序根据具体的错误信息做出相应的处理或调整，而非简单地重复尝试。

> **重试延迟说明：**
> SDK为了防止请求风暴，提高系统稳定性，并有效缓解服务雪崩风险，增加重试间的延迟，实现对服务端压力的自适应控制。
> **延迟时间公式：**
> delay = min(MaxDelay, 2ⁿ × minDelay × (1 + Rand[0, 1)) + Retry-After
>
>
> | 参数            | 说明                                                    |
> | --------------- | ------------------------------------------------------- |
> | maxDelay        | 最大延迟时间，计算的最大延迟不会超过maxDelay，默认500ms |
> | 2ⁿ             | 纯指数增长                                              |
> | (1 + Rand[0, 1) | 把结果随机放大 1 ~ 2 倍，避免「惊群效应」               |
> | min(...)        | 防止无限增长，超过MaxDelay                              |
> | Retry-After     | 服务端如果显式告知休眠时长，则先按它要求静默            |
>
> **举例说明：**
> MaxDelay=500ms，minDelay=30ms
>
>
> | 重试次数（从第0次开始） | 指数退避×抖动区间(ms) | Retry-After(ms) | 最终延迟时间（指数退避×抖动区间+Retry-After） |
> | ----------------------- | ---------------------- | --------------- | ---------------------------------------------- |
> | 0                       | [30,60]                | 10              | [40,70]                                        |
> | 1                       | [60,120]               | 20              | [80,140]                                       |
> | 3                       | [120,240]              | 30              | [150,270]                                      |
> | ...                     | ...<=500ms             | 10              | <=510                                          |

## 开启重试机制

> - **默认**
>   开启（3次）

如果想关闭，可以把最大尝试次数改为0

```go
func main() {
	region := "cn-beijing"
	// Configure retry settings
	config := volcengine.NewConfig().
		WithRegion(region).
		WithDisableSSL(true).
		WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：VOLCSTACK_ACCESS_KEY_ID、VOLCSTACK_SECRET_ACCESS_KEY、VOLCSTACK_SESSION_TOKEN
		// 关闭重试
		WithMaxRetries(0)

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

## 重试次数

> - **默认**
>   3次

1. 设置默认次数

```go
func main() {
	region := "cn-beijing"
	// Configure retry settings
	config := volcengine.NewConfig().
		WithRegion(region).
		WithDisableSSL(true).
		WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：VOLCSTACK_ACCESS_KEY_ID、VOLCSTACK_SECRET_ACCESS_KEY、VOLCSTACK_SESSION_TOKEN
		// 设置最大重试次数 (default is DefaultRetryerMaxNumRetries)
		WithMaxRetries(4)

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

## 自定义重试错误码

在调用接口的时候可以根据业务需求，自定义重试的错误码(服务端返回的错误码)。公共错误码可以参考：[公共错误码](https://www.volcengine.com/docs/6369/68677?lang=zh)

```go
func main() {
	region := "cn-beijing"
	// Configure retry settings
	config := volcengine.NewConfig().
		WithRegion(region).
		WithDisableSSL(true).
		WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：VOLCSTACK_ACCESS_KEY_ID、VOLCSTACK_SECRET_ACCESS_KEY、VOLCSTACK_SESSION_TOKEN
		// 设置最大重试次数 (default is DefaultRetryerMaxNumRetries)
		WithMaxRetries(4)

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	describeAvailableResourceInput := &ecs.DescribeAvailableResourceInput{
		DestinationResource: volcengine.String("InstanceType"),
		InstanceTypeId:      volcengine.String("ecs.g2i.large"),
		ZoneId:              volcengine.String("cn-*****"),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	resp, err := svc.DescribeAvailableResourceWithContext(ctx, describeAvailableResourceInput, func(request *request.Request) {
		// 自定义重试错误码
		request.RetryErrorCodes = []string{"InvalidAccessKey"}
	})
	if err != nil {
		panic(err)
	}
}
```

---

[← 超时配置](4-Timeout-zh.md) | 重试机制[(English)](5-Retry.md) | [异常处理 →](6-ErrorHandling-zh.md)
