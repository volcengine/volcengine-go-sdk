[← 重试机制](5-Retry-zh.md) | 异常处理[(English)](6-ErrorHandling.md) | [Debug 机制 →](7-Debugging-zh.md)

---

## 异常处理

在调用接口时，可能会返回不同类型的错误。客户可根据具体的错误类型和错误码，采取有针对性的处理策略。例如，对于网络异常可选择重试，对于业务逻辑错误则应根据错误信息进行参数调整或业务逻辑修正，从而提升系统的健壮性和用户体验。

错误分类：


| 错误类型       | 错误描述                             | 返回错误类型                   | 公共属性                                                                         | 私有属性                                          |
|------------| ------------------------------------ | ------------------------------ |------------------------------------------------------------------------------| ------------------------------------------------- |
| 1. 客户端错误   | 请求未到达服务端，对参数作验证       | volcengineerr.Error或原生error | Code()：错误码;  <br>Message():错误描述信息;  <br>Error()：详细错误信息;  <br>OrigErr(): 原始错误 | 无                                                |
| 2. 服务端错误   | 请求成功到达服务器，返回业务逻辑错误 | volcengineerr.RequestFailure   | 同上                                                                           | 可以通过RequestID()获取请求id，方便服务端问题排查 |
| 3. 网络/超时错误 | DNS解析错误或请求超时                | volcengineerr.Error            | 同上                                                                           | 无                                                |
| 4. 其它错误    | 未包含在前 3 类错误中的其它错误      | volcengineerr.Error或原生error | 同上                                                                           | 无                                                |

### 代码示例

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
			fmt.Println("1. 客户端错误(创建session失败)", be.Code(), be.Message(), be.Error())
		} else {
			fmt.Println("4. 其它错误", err.Error())
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
		var requestFailure volcengineerr.RequestFailure // 服务端返回的错误
		var errInvalidParam request.ErrInvalidParam     // 参数验证错误
		// 请求未达到服务前参数验证
		if errors.As(err, &errInvalidParam) {
			fmt.Println("1. 客户端错误(参数验证错误)：", errInvalidParam.Code(), errInvalidParam.Field(), errInvalidParam.Error())
			// 请求到达服务端，服务端返回错误
		} else if errors.As(err, &requestFailure) {
			fmt.Println("2. 服务端错误：", requestFailure.RequestID(), requestFailure.Code(), requestFailure.StatusCode(), requestFailure.Error())
		} else if errors.As(err, &be) {
			// 发送请求，但没有到达后端服务
			switch be.Code() {
			case "RequestCanceled":
				fmt.Println("3. 网络/超时错误：这里是请求接口传入context上下文超时")
			case "RequestError":
				if be.OrigErr() != nil {
					var netErr net.Error
					var dnsError *net.DNSError
					if errors.As(be.OrigErr(), &dnsError) {
						fmt.Println("3. 网络/超时错误：DNS解析错误处理")
					} else if errors.As(be.OrigErr(), &netErr) && netErr.Timeout() {
						var oPError *net.OpError
						if errors.Is(be.OrigErr(), context.DeadlineExceeded) {
							fmt.Println("3. 网络/超时错误：http.Client Timeout(ReadTimeout)....", be.Code(), be.Error())
						} else if errors.As(be.OrigErr(), &oPError) && oPError.Op == "dial" {
							fmt.Println("3. 网络/超时错误：http.Client Transport.Dialer Timeout(ConnectTimeout)....", be.Code(), be.Error())
						} else {
							fmt.Println("3. 网络/超时错误：其它超时处理", be.Code(), be.Message(), be.Error())
						}
					}
				}
			default:
				fmt.Println("4. 其它错误", be.Code(), be.Message(), be.Error())
			}
		} else {
			fmt.Println("4. 其它错误", err.Error())
		}

	}

}
```

---

[← 重试机制](5-Retry-zh.md) | 异常处理[(English)](6-ErrorHandling.md) | [Debug 机制 →](7-Debugging-zh.md)
