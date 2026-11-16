[← 超时配置](5-Timeout-zh.md) | 重试机制[(English)](6-Retry.md) | [异常处理 →](7-ErrorHandling-zh.md)

---

## 重试机制

请求的处理逻辑内置了网络异常重试逻辑，即当遇到网络异常问题或限流错误时，系统会自动尝试重新发起请求，以确保服务的稳定性和可靠性。若请求因业务逻辑错误而报错，例如参数错误、资源不存在等情况，SDK 将不会执行重试操作，这是因为业务层面的错误通常需要应用程序根据具体的错误信息做出相应的处理或调整，而非简单地重复尝试。

### 包装错误与底层原因

重试判断会递归保留 SDK 错误 `OrigErr()` 的结果，并识别 `fmt.Errorf("...: %w", err)` 等标准包装。限流、凭证过期及自定义 `RetryErrorCodes` / `ThrottleErrorCodes` 通过 `errors.As` 获取包装链中的首个 SDK 错误码；包装本身不会使终止性业务错误变为可重试。

默认错误分支会在调用 `Unwrap()` 前，先检查 `net/http: request canceled` 和 `net/http: request canceled while waiting for connection` 两个精确的旧取消文本。包装错误返回任一文本时，即使底层原因是临时错误或具有可重试错误码，仍按不可重试处理。

`ReadError` 和 `SerializationError` 仅在底层存在已知可重试原因（如临时网络错误、超时错误码或既有连接重置/断管识别）时重试。语法错误及无法识别的读取原因仍停止重试；HTTP 状态码、显式 `Retryable` 和自定义错误码的既有规则继续生效。

批量底层原因通过 `OrigErrs()` 有限遍历，不反复调用可能生成相同批量错误的 `OrigErr()`。当 `ReadError`、`SerializationError` 或 `RequestError` 需要依据底层原因判定时，非空原因必须全部可重试；任一原因不可重试即停止。单次判定最多检查 256 个错误或原因，包装深度最多 64 层；预算耗尽时尚未判定的底层原因按不可重试处理。HTTP 状态码及显式配置的既有优先级保持不变。

### STS 凭证重试

AssumeRole、OIDC 和 SAML 会在进入重试循环前完成请求构造检查。请求 URL 无效或 AssumeRole API 配置缺失时立即返回，不发送请求，也不等待重试。进入 HTTP 调用后的网络失败仍按既有 STS 策略重试；每轮 AssumeRole 使用独立请求重新签名并创建独立超时上下文，每轮 OIDC/SAML 都重新发送完整表单正文。

### 重试延迟说明

SDK 为了防止请求风暴，提高系统稳定性，并有效缓解服务雪崩风险，增加重试间的延迟，实现对服务端压力的自适应控制。

#### 延迟时间公式

```
delay = min(MaxDelay, 2ⁿ × minDelay × (1 + Rand[0, 1)) + Retry-After
```

| 参数 | 说明 |
|---|---|
| `maxDelay` | 最大延迟时间，计算的最大延迟不会超过 `maxDelay`，默认 300s（`DefaultRetryerMaxRetryDelay`） |
| `minDelay` | 最小延迟时间基线，默认 30ms（`DefaultRetryerMinRetryDelay`） |
| `2ⁿ` | 纯指数增长 |
| `(1 + Rand[0, 1))` | 把结果随机放大 1 ~ 2 倍，避免「惊群效应」 |
| `min(...)` | 防止无限增长，超过 `MaxDelay` |
| `Retry-After` | 服务端如果显式告知休眠时长，则先按它要求静默 |

#### 举例说明（演示用：把 `MaxDelay` 缩小到 500ms 便于观察封顶行为）

`MaxDelay=500ms`，`minDelay=30ms`：

| 重试次数（从第 0 次开始） | 指数退避×抖动区间（ms） | Retry-After（ms） | 最终延迟时间（指数退避×抖动区间 + Retry-After） |
|---|---|---|---|
| 0 | [30, 60] | 10 | [40, 70] |
| 1 | [60, 120] | 20 | [80, 140] |
| 3 | [120, 240] | 30 | [150, 270] |
| ... | ... ≤ 500ms | 10 | ≤ 510 |

### 开启重试机制

> **默认**
>
> - 开启（3 次）

如果想关闭，可以把最大尝试次数改为 0。

```go
func main() {
	region := "cn-beijing"
	// Configure retry settings
	config := volcengine.NewConfig().
		WithRegion(region).
		WithDisableSSL(true).
		WithCredentials(credentials.NewEnvCredentials()). //需要配置环境变量
		// 关闭重试
		WithMaxRetries(0)

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

### 重试次数

> **默认**
>
> - 3 次

设置最大重试次数：

```go
func main() {
	region := "cn-beijing"
	// Configure retry settings
	config := volcengine.NewConfig().
		WithRegion(region).
		WithDisableSSL(true).
		WithCredentials(credentials.NewEnvCredentials()). //需要配置环境变量
		// 设置最大重试次数 (default is DefaultRetryerMaxNumRetries)
		WithMaxRetries(4)

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

### 自定义重试错误码

在调用接口的时候可以根据业务需求，自定义重试的错误码（服务端返回的错误码）。公共错误码可以参考：[公共错误码](https://www.volcengine.com/docs/6369/68677?lang=zh)。

```go
func main() {
	region := "cn-beijing"
	// Configure retry settings
	config := volcengine.NewConfig().
		WithRegion(region).
		WithDisableSSL(true).
		WithCredentials(credentials.NewEnvCredentials()). //需要配置环境变量
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

[← 超时配置](5-Timeout-zh.md) | 重试机制[(English)](6-Retry.md) | [异常处理 →](7-ErrorHandling-zh.md)
