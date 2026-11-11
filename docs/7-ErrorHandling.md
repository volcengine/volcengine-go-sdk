[← Retry](6-Retry.md) | Error Handling[(中文)](7-ErrorHandling-zh.md) | [Debugging →](8-Debugging.md)

---

## Error Handling

When calling APIs, different types of errors may be returned. You can adopt targeted handling strategies based on the specific error type and error code. For example, network errors can be retried, while business logic errors should be addressed by adjusting parameters or fixing business logic, thereby improving system robustness and user experience.

Error classification:


| Error Type | Description | Returned Error Type | Common Properties | Private Properties |
|---|---|---|---|---|
| 1. Client error | Request did not reach the server; parameter validation failed | volcengineerr.Error or native error | Code(): error code;  <br>Message(): error description;  <br>Error(): detailed error info;  <br>OrigErr(): original error | None |
| 2. Server error | Request successfully reached the server; business logic error returned | volcengineerr.RequestFailure | Same as above | RequestID() to obtain the request ID for server-side troubleshooting |
| 3. Network/timeout error | DNS resolution error or request timeout | volcengineerr.Error | Same as above | None |
| 4. Other errors | Other errors not covered by the above categories | volcengineerr.Error or native error | Same as above | None |

### Error diagnostics for non-standard responses

When the HTTP status code is non-2xx (e.g. 504), the response body may not be a standard TOP gateway JSON structure. The default error handler returns `volcengineerr.RequestFailure` according to the following rules to make troubleshooting easier:

- **Body cannot be decoded into the expected response structure** (e.g. HTML, plain text, an empty response, a top-level array, or incompatible types for fields such as `ResponseMetadata`): returns `RequestFailure` with `Code()` set to `SerializationError`, the raw response body attached to `Message()`, and the HTTP status code and request ID exposed via `StatusCode()` and `RequestID()`.
- **Body can be decoded but has no valid error code** (e.g. missing `ResponseMetadata` with no non-empty simple error code, or metadata with no error code): returns `RequestFailure` with `Code()` set to `UnknownError`, also preserving the HTTP status code, request ID, and response body.
- **Body uses the simple error format** (missing `ResponseMetadata`, but matching `{"errorcode":"error code","message":"error description"}` with a non-empty `errorcode`): returns `RequestFailure` with `Code()` preserving the actual `errorcode`, `Message()` including the description and raw body, and the HTTP status code and request ID preserved.
- **Body is a standard TOP JSON error**: returns `RequestFailure`, preserving the server-side `Code()` and `Message()`; if the body has no non-empty request ID, the fallback below applies.

For the standard and non-standard responses above, including standard business errors under HTTP 200, the request ID is resolved in order from a non-empty `ResponseMetadata.RequestId` in the body, an existing `r.RequestID` on the request object, then the `X-Top-Requestid` / `X-Top-Request-Id` / `X-Tt-Logid` response headers. `RequestID()` returns an empty string if none of these sources has a value. Standard business errors under HTTP 200 also apply `SimpleError`, which still affects only the display.

In addition, when an HTTP 200 response has a JSON `null` body, decoding into a struct or `*map[string]interface{}` with output data returns `RequestFailure` with `Code()` set to `UnknownError` and `StatusCode()` set to 200, preserving the request ID and raw body. This applies to both normal decoding and `ForceJsonNumberDecode`. Other malformed HTTP 200 bodies (e.g. HTML) are outside the non-2xx diagnostics described above and may still return a native JSON decoding error.

When map decoding rejects JSON `null`, it preserves the original output pointer. If a custom retryer retries `UnknownError`, subsequent responses are decoded into the original output object, which remains available to `After` interceptors. Exhausted retries still return an error. The default retryer does not retry an HTTP 200 response solely because of `UnknownError`.

Compatibility: some native JSON decoding errors that were previously returned directly are now wrapped in `RequestFailure`; callers can obtain the underlying error through `OrigErr()`. The error classification also affects default retries: with `MaxRetries=3`, an HTTP 400 response with an HTML body changes from at most 4 attempts to 1 attempt. HTTP 500, 502, 503, and 504 responses still retry under the existing status-code rules. `SimpleError` changes only the display of `Error()`, leaving `StatusCode()`, `RequestID()`, and `Message()` unchanged. Custom error handlers and retryers follow their own implementations.

### Error handling for STS credential requests

`StsAssumeRole`, `StsAssumeRoleWithOptions`, `StsProvider` (including `NewStsCredentials` / `NewStsCredentialsWithOptions`), OIDC, and SAML credential requests return `volcengineerr.RequestFailure`, accessible through `errors.As`:

- Standard TOP errors, including `ResponseMetadata.Error` under HTTP 200, preserve the server-side `Code()` and `Message()`. `StatusCode()` always uses the actual HTTP status, ignoring `HTTPCode` in the body.
- Request IDs come first from a non-empty `ResponseMetadata.RequestId` in the body, then from the `X-Top-Requestid`, `X-Top-Request-Id`, and `X-Tt-Logid` response headers in that order.
- Non-standard error bodies follow the `SerializationError`, `UnknownError`, or simple error-code rules above. Read failures return `ReadError`, preserving any partially read body. `OrigErr()` exposes underlying read and decoding errors.
- Network and timeout errors return `RequestError`. If no HTTP response was received, `StatusCode()` is 0 and `RequestID()` is empty; no HTTP 500 is fabricated.
- A JSON `null` body, missing result, or incomplete credentials under a successful HTTP status returns `UnknownError`. Invalid OIDC/SAML credential expiration times return `SerializationError` with the HTTP status, request ID, and underlying parsing error, without updating the cache.

STS credential requests continue to retry network failures, HTTP 429/502/503/504, and existing throttling/transient error codes. The final error carries the last attempt's diagnostics. HTTP 500 without a retryable error code is not retried; terminal 4xx responses are not retried. OIDC/SAML default to 3 retries when `MaxRetries` is `nil` and disable retries when it points to 0. AssumeRole retains the default when its integer `MaxRetries` is less than or equal to 0.

Compatibility: these credential errors now use the SDK's default `Code: Message` format, followed by `status code` and `request id` on the next line, and `caused by` when an underlying error exists. Use accessor methods instead of parsing previous error strings. Credential Providers have no `SimpleError` option; Session's `SimpleError` applies only to API requests through `service/sts` and does not change credential-refresh error displays.

```go
var failure volcengineerr.RequestFailure
if errors.As(err, &failure) {
    fmt.Println(failure.StatusCode(), failure.Code(), failure.Message(), failure.RequestID())
    // Underlying network, read, or decoding cause: failure.OrigErr()
}
```

Local parameter validation, token-file reads, and proxy/TLS configuration errors retain each Provider's existing local error types.

### Retry classification of wrapped errors

Wrapping an SDK error with `%w` preserves default throttling, credential-expiration, and custom error-code classification, along with recursive retry decisions for `OrigErr()` causes. Unknown causes under `ReadError` or `SerializationError` remain non-retryable unless existing rules such as HTTP status apply; see [Wrapped errors and underlying causes](6-Retry.md#wrapped-errors-and-underlying-causes).

Access underlying SDK error causes through `OrigErr()`. This change does not add `Unwrap()` to public error types.

### Full error details

`volcengineerr.PrintError(writer, err)` explicitly writes indented JSON with a trailing newline. Its return value reports output failures. `volcengineerr.GetErrorDetails(err)` provides the same read-only extraction for use with your own logger.

```go
// Import fmt, os and volcengineerr.
if err != nil {
    if printErr := volcengineerr.PrintError(os.Stderr, err); printErr != nil {
        fmt.Fprintf(os.Stderr, "failed to print error: %v\n", printErr)
    }
}
```

Call `PrintError` directly when you only need output; no type assertion is required. To read the HTTP response separately, optionally match the new extension:

```go
var failure volcengineerr.DetailedRequestFailure
if errors.As(err, &failure) {
    if response := failure.ResponseDetails(); response != nil {
        _ = response // A response snapshot copy with status, headers and body.
    }
}
```

The new `DetailedRequestFailure` extends `RequestFailure` with an optional `ResponseDetails() *ResponseDetails` method. It returns only a captured HTTP snapshot copy, or nil when no snapshot exists. Continue using `Code()` / `Message()` for error properties; use the package-level `volcengineerr.GetErrorDetails(err)` for the complete error tree. Existing `Error` / `RequestFailure` interfaces and constructor signatures are unchanged. Customer-defined errors require no new methods. Errors returned by customer parsing callbacks preserve their identity, concrete type, properties and `Error()` text, including errors created through SDK constructors. The SDK does not attach details to arbitrary final errors.

`ErrorDetails` exposes `Type`, `Error`, `Code`, `Message`, `Response` and `Causes`. `ResponseDetails` exposes `StatusCode`, `Status`, `Proto`, `RequestID`, multi-value `Headers` / `Trailers`, original `Body` bytes, `BodyCaptured` and `BodyComplete`. Body request IDs still take precedence over headers, and HTTP 200 business errors retain status 200. `SimpleError` does not suppress structured fields; existing `Error()` / `String()` displays are unchanged.

Full snapshots cover request errors constructed by the default TOP parser and STS / AssumeRole / OIDC / SAML credential requests, including OIDC/SAML expiration parsing failures. Capture reuses existing reads without additional requests or reading closed bodies, and retains only the attempt corresponding to the error. Interrupted reads preserve partial bytes with `body_complete: false` and the read cause in `causes`.

Default TOP response-read timeouts triggered by `request.WithResponseReadTimeout` also retain the partial body, HTTP status, headers and request ID. The existing `ResponseTimeout` type and text remain unchanged, with a nil `OrigErr()` and no additional `causes`. Each request owns its snapshot without modifying the shared timeout error.

Without an HTTP response, `response` is `null`. Legacy interfaces exposing only status and request ID yield `headers: null` and `body_captured: false`. A captured empty body has `body_captured: true` and `body: ""`. JSON renders UTF-8 bodies as text with `body_encoding: "utf-8"` and other bytes as reversible Base64. Returned details are deep copies. Extraction supports `%w`, `OrigErr()`, `OrigErrs()` and multi-cause standard wrapping, with cycle protection and bounded depth.

Native read/decoding errors returned unchanged, customer errors, and paths without capture such as ECS/IMDS, endpoint credentials and CLI OAuth/Portal provide only available information. Package functions fall back to legacy fields when optional customer `ResponseDetails()` methods return nil. `GetErrorDetails(nil)` returns nil; `PrintError(writer, nil)` writes nothing. Write failures and short writes return wrapped output errors. Full bodies may contain credentials; callers can modify the returned copy before logging.

Error-tree deduplication supports custom value errors containing slice or map fields. Distinct pointer errors remain separate even when their contents match.

### Code Example

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
	"net"
)

func main() {
	region := "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewEnvCredentials())
	sess, err := session.NewSession(config)
	var be volcengineerr.Error
	if err != nil {
		if errors.As(err, &be) {
			fmt.Println("1. Client error (failed to create session)", be.Code(), be.Message(), be.Error())
		} else {
			fmt.Println("4. Other error", err.Error())
		}
		panic(err)
	}
	svc := ecs.New(sess)

	tags := make([]*ecs.TagForCreateKeyPairInput, 0, 2)
	tags = append(tags, &ecs.TagForCreateKeyPairInput{Key: volcengine.String("testTag")})
	createKeyPairInput := &ecs.CreateKeyPairInput{
		KeyPairName: volcengine.String(("testKeyPairName")),
		Tags:        tags,
	}

	_, err = svc.CreateKeyPair(createKeyPairInput)
	if err != nil {
		var requestFailure volcengineerr.RequestFailure // Server-side error
		var errInvalidParam request.ErrInvalidParam     // Parameter validation error
		// Pre-request parameter validation
		if errors.As(err, &errInvalidParam) {
			fmt.Println("1. Client error (parameter validation error):", errInvalidParam.Code(), errInvalidParam.Field(), errInvalidParam.Error())
			// Request reached the server; server returned an error
		} else if errors.As(err, &requestFailure) {
			fmt.Println("2. Server error:", requestFailure.RequestID(), requestFailure.Code(), requestFailure.StatusCode(), requestFailure.Error())
		} else if errors.As(err, &be) {
			// Request was sent but did not reach the backend service
			switch be.Code() {
			case "RequestCanceled":
				fmt.Println("3. Network/timeout error: context timeout passed to the API call")
			case "RequestError":
				if be.OrigErr() != nil {
					var netErr net.Error
					var dnsError *net.DNSError
					if errors.As(be.OrigErr(), &dnsError) {
						fmt.Println("3. Network/timeout error: DNS resolution error handling")
					} else if errors.As(be.OrigErr(), &netErr) && netErr.Timeout() {
						var oPError *net.OpError
						if errors.Is(be.OrigErr(), context.DeadlineExceeded) {
							fmt.Println("3. Network/timeout error: http.Client Timeout (ReadTimeout)....", be.Code(), be.Error())
						} else if errors.As(be.OrigErr(), &oPError) && oPError.Op == "dial" {
							fmt.Println("3. Network/timeout error: http.Client Transport.Dialer Timeout (ConnectTimeout)....", be.Code(), be.Error())
						} else {
							fmt.Println("3. Network/timeout error: other timeout handling", be.Code(), be.Message(), be.Error())
						}
					}
				}
			default:
				fmt.Println("4. Other error", be.Code(), be.Message(), be.Error())
			}
		} else {
			fmt.Println("4. Other error", err.Error())
		}

	}

}
```

---

[← Retry](6-Retry.md) | Error Handling[(中文)](7-ErrorHandling-zh.md) | [Debugging →](8-Debugging.md)
