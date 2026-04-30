[← Retry](5-Retry.md) | Error Handling[(中文)](6-ErrorHandling-zh.md) | [Debugging →](7-Debugging.md)

---

# Error Handling

When calling APIs, different types of errors may be returned. You can adopt targeted handling strategies based on the specific error type and error code. For example, network errors can be retried, while business logic errors should be addressed by adjusting parameters or fixing business logic, thereby improving system robustness and user experience.

Error classification:


| Error Type | Description | Returned Error Type | Common Properties | Private Properties |
|---|---|---|---|---|
| 1. Client error | Request did not reach the server; parameter validation failed | volcengineerr.Error or native error | Code(): error code;  <br>Message(): error description;  <br>Error(): detailed error info;  <br>OrigErr(): original error | None |
| 2. Server error | Request successfully reached the server; business logic error returned | volcengineerr.RequestFailure | Same as above | RequestID() to obtain the request ID for server-side troubleshooting |
| 3. Network/timeout error | DNS resolution error or request timeout | volcengineerr.Error | Same as above | None |
| 4. Other errors | Other errors not covered by the above categories | volcengineerr.Error or native error | Same as above | None |

**Code example:**

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

[← Retry](5-Retry.md) | Error Handling[(中文)](6-ErrorHandling-zh.md) | [Debugging →](7-Debugging.md)
