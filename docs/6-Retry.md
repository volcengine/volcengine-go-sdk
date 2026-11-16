[← Timeout](5-Timeout.md) | Retry[(中文)](6-Retry-zh.md) | [Error Handling →](7-ErrorHandling.md)

---

## Retries

The SDK includes retry logic for network errors and throttling. Business logic errors (invalid params, resource not found) are not retried.

### Wrapped errors and underlying causes

Retry decisions preserve recursive results from SDK errors' `OrigErr()` and recognize standard wrappers such as `fmt.Errorf("...: %w", err)`. Throttling, credential expiration, and custom `RetryErrorCodes` / `ThrottleErrorCodes` use `errors.As` to obtain the first SDK error code in the wrapping chain. Wrapping alone does not make a terminal business error retryable.

The default error branch checks the exact legacy cancellation texts `net/http: request canceled` and `net/http: request canceled while waiting for connection` before following `Unwrap()`. A wrapper reporting either text remains non-retryable even when its cause is temporary or has a retryable code.

`ReadError` and `SerializationError` retry only when an underlying cause is known to be retryable, such as a temporary network error, a timeout error code, or an existing connection-reset/broken-pipe match. Syntax errors and unrecognized read failures still stop retries. Existing HTTP status, explicit `Retryable`, and custom error-code rules continue to apply.

Batch causes are traversed through `OrigErrs()` within a fixed budget, without repeatedly calling an `OrigErr()` that may manufacture the same batch. When `ReadError`, `SerializationError`, or `RequestError` depends on its causes, every non-nil cause must be retryable; any non-retryable cause stops the decision. Each decision checks at most 256 errors or causes and at most 64 wrapping levels. Causes left unresolved when the budget is exhausted are treated as non-retryable. Existing HTTP-status and explicit-configuration priorities remain unchanged.

### STS credential retries

AssumeRole, OIDC, and SAML validate request construction before entering the retry loop. Invalid request URLs and missing AssumeRole API configuration return immediately without sending or waiting. Network failures after an HTTP attempt still follow the existing STS retry policy. Each AssumeRole attempt signs an independent request with its own timeout context; each OIDC/SAML attempt sends the complete form body again.

### Enable/Disable Retries

> **Default**
>
> - Enabled (3 times)

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

### Retry Count

> **Default**
>
> - 3

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

### Custom Retry Error Codes

Configure retryable server error codes per request.

```go
resp, err := svc.DescribeAvailableResourceWithContext(ctx, describeAvailableResourceInput, func(request *request.Request) {
	request.RetryErrorCodes = []string{"InvalidAccessKey"}
})
```

---

[← Timeout](5-Timeout.md) | Retry[(中文)](6-Retry-zh.md) | [Error Handling →](7-ErrorHandling.md)
