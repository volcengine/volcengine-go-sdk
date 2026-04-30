[← Timeout](4-Timeout.md) | Retry[(中文)](5-Retry-zh.md) | [Error Handling →](6-ErrorHandling.md)

---

# Retries

The SDK includes retry logic for network errors and throttling. Business logic errors (invalid params, resource not found) are not retried.

## Enable/Disable Retries

> - **Default**: enabled (3 times)

Set max retries to `0` to disable.

```go
func main() {
	region := "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithDisableSSL(true).
		WithCredentials(credentials.NewEnvCredentials()).
		WithMaxRetries(0)

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

## Retry Count

> - **Default**: 3

```go
func main() {
	region := "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithDisableSSL(true).
		WithCredentials(credentials.NewEnvCredentials()).
		WithMaxRetries(4)

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
}
```

## Custom Retry Error Codes

Configure retryable server error codes per request.

```go
resp, err := svc.DescribeAvailableResourceWithContext(ctx, describeAvailableResourceInput, func(request *request.Request) {
	request.RetryErrorCodes = []string{"InvalidAccessKey"}
})
```

---

[← Timeout](4-Timeout.md) | Retry[(中文)](5-Retry-zh.md) | [Error Handling →](6-ErrorHandling.md)
