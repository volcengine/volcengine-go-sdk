[← 重试机制](6-Retry-zh.md) | 异常处理[(English)](7-ErrorHandling.md) | [Debug 机制 →](8-Debugging-zh.md)

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

### 非标准响应的错误诊断

当 HTTP 状态码为非 2xx（如 504 等）时，响应体可能不是标准的 TOP 网关 JSON 结构。默认错误处理按以下规则返回 `volcengineerr.RequestFailure`，便于排查问题：

- **响应体无法解码为预期响应结构**（如 HTML、纯文本、空响应、顶层数组，或 `ResponseMetadata` 等字段的类型不兼容）：返回 `RequestFailure`，`Code()` 为 `SerializationError`，`Message()` 中附带原始响应正文，并通过 `StatusCode()`、`RequestID()` 暴露 HTTP 状态码与请求 ID。
- **响应体可解码但缺少有效错误码**（如缺少 `ResponseMetadata`，且没有非空的简易错误码，或元数据中缺少错误码）：返回 `RequestFailure`，`Code()` 为 `UnknownError`，同样保留 HTTP 状态码、请求 ID 与响应正文。
- **响应体为简易错误格式**（缺少 `ResponseMetadata`，但符合 `{"errorcode":"错误码","message":"错误描述"}` 且 `errorcode` 非空）：返回 `RequestFailure`，`Code()` 保留真实的 `errorcode`，`Message()` 附带错误描述与原始正文，同时保留 HTTP 状态码与请求 ID。
- **响应体为标准 TOP JSON 错误**：返回 `RequestFailure`，保留服务端的 `Code()`、`Message()`；正文没有非空请求 ID 时，按下述规则兜底。

上述标准及非标准响应（包括 HTTP 200 下的标准业务错误）的请求 ID 依次取自响应体中非空的 `ResponseMetadata.RequestId`、请求对象中已有的 `r.RequestID`、以及 `X-Top-Requestid` / `X-Top-Request-Id` / `X-Tt-Logid` 响应头；所有来源均无值时，`RequestID()` 返回空字符串。标准业务错误在 HTTP 200 下也应用 `SimpleError`，该选项仍只影响显示。

此外，有输出数据的结构体和 `*map[string]interface{}` 解码路径收到 HTTP 200、正文为 JSON `null` 时，均返回 `RequestFailure`，`Code()` 为 `UnknownError`，`StatusCode()` 为 200，并保留请求 ID 与原始正文。普通解码和 `ForceJsonNumberDecode` 均适用。

成功状态码下，默认 TOP 解析器处理输出数据时也统一保存异常诊断：

- **读取正文失败**：返回 `RequestFailure`，`Code()` 为 `ReadError`，通过 `OrigErr()` 保留原始读取错误；快照保存已读字节，`BodyCaptured` 为 true，`BodyComplete` 为 false。
- **正文解码失败**（如 HTML、空正文、JSON 结构不兼容）：返回 `RequestFailure`，`Code()` 为 `SerializationError`，底层解码错误保留在 `OrigErr()` 中。
- **结果序列化、输出对象解码或 map 字段访问失败**：同样返回 `SerializationError`，保留底层原因和原始 HTTP 正文，而非转换后的结果字节。客户 `CustomerUnmarshalData` 返回的数据在 SDK 编解码时出错也适用；客户 `CustomerUnmarshalError` 返回的错误对象仍原样返回，客户回调中的 panic 沿用原行为。

上述 SDK 构造的错误通过 `StatusCode()`、`RequestID()` 和 `PrintError` 暴露 HTTP 状态码、请求 ID、响应头及正文快照；HTTP 200 仍记录为 200。即使其它字段类型不兼容，也会尝试从已读取的有效 JSON 中单独提取 `ResponseMetadata.RequestId`，再按既有顺序兜底，不重新读取响应流。普通解码和 `ForceJsonNumberDecode` 均适用。

map 解码拒绝 JSON `null` 时会保留原输出指针。若自定义重试器重试 `UnknownError`，后续响应会继续解析到原输出对象，`After` 拦截器也能获取该对象；重试耗尽仍返回错误。默认重试器不会仅因 `UnknownError` 而重试 HTTP 200 响应。

兼容性说明：此前直接返回的部分原生 JSON 解码错误现在封装为 `RequestFailure`，调用方可通过 `OrigErr()` 获取底层错误。默认重试行为也会受错误分类影响：以 `MaxRetries=3` 为例，HTTP 400、正文为 HTML 的请求从原先最多 4 次尝试变为 1 次；HTTP 500、502、503、504 仍按既有状态码规则重试。`SimpleError` 仅改变 `Error()` 的显示形式，不影响 `StatusCode()`、`RequestID()` 和 `Message()`。自定义错误处理或重试器以其实现为准。

HTTP 200 下的读取、解码及结果转换错误也由原生错误改为上述 SDK 错误类型。HTML 或普通类型不匹配等不可重试原因不再走“未知错误可重试”的兜底，`MaxRetries=3` 时只尝试 1 次；具有明确可重试底层原因的读取错误仍按既有规则重试。通过 `request.WithResponseReadTimeout` 触发的读取超时最终仍归一化为原有 `ResponseTimeout`，保留快照，详见下文。

### STS 凭证请求的错误处理

`StsAssumeRole`、`StsAssumeRoleWithOptions`、`StsProvider`（含 `NewStsCredentials` / `NewStsCredentialsWithOptions`）、OIDC 和 SAML 凭证请求统一返回 `volcengineerr.RequestFailure`，支持通过 `errors.As` 获取错误：

- 标准 TOP 错误（包括 HTTP 200 下的 `ResponseMetadata.Error`）保留服务端 `Code()` 和 `Message()`；`StatusCode()` 始终取真实 HTTP 状态码，不采用正文中的 `HTTPCode`。
- 请求 ID 优先取正文中非空的 `ResponseMetadata.RequestId`，再依次取 `X-Top-Requestid`、`X-Top-Request-Id`、`X-Tt-Logid` 响应头。
- 非标准错误正文采用上述 `SerializationError`、`UnknownError` 或简易错误码规则。读取失败返回 `ReadError`，保留已读取的正文；读取和解码错误可通过 `OrigErr()` 获取底层原因。
- 网络或超时错误返回 `RequestError`；没有收到 HTTP 响应时，`StatusCode()` 为 0、`RequestID()` 为空，不伪造 HTTP 500。
- 成功状态码下的 JSON `null`、缺失结果或不完整凭证返回 `UnknownError`；OIDC/SAML 的凭证过期时间无法解析时返回 `SerializationError`，保留 HTTP 状态码、请求 ID 和底层解析错误，且不更新缓存。

STS 凭证请求继续按网络错误、HTTP 429/502/503/504 和已有的限流/瞬时错误码重试，最终错误携带最后一次尝试的信息。普通 HTTP 500 且没有可重试错误码时不重试；终止性 4xx 不重试。OIDC/SAML 的 `MaxRetries` 为 `nil` 时默认重试 3 次，指向 0 时不重试；AssumeRole 的整型 `MaxRetries` 小于等于 0 时仍采用默认值。

兼容性说明：上述凭证请求的错误字符串改为 SDK 默认格式 `Code: Message`，下一行显示 `status code` 和 `request id`，有底层错误时追加 `caused by`。调用方应使用属性方法而非解析旧错误字符串。凭证 Provider 没有 `SimpleError` 选项；Session 的 `SimpleError` 仅影响通过 `service/sts` 发起的 API 请求，不改变凭证刷新错误的显示。

```go
var failure volcengineerr.RequestFailure
if errors.As(err, &failure) {
    fmt.Println(failure.StatusCode(), failure.Code(), failure.Message(), failure.RequestID())
    // 底层网络、读取或解码原因：failure.OrigErr()
}
```

本地参数验证、令牌文件读取及代理/TLS 配置错误仍沿用各 Provider 的本地错误类型。

### 包装错误的重试识别

使用 `%w` 包装 SDK 错误后，默认重试判断仍识别限流、凭证过期及自定义错误码，并递归判断 `OrigErr()` 中的底层原因。`ReadError`、`SerializationError` 的未知底层原因仍不可重试，HTTP 状态码等既有规则除外；详见[包装错误与底层原因](6-Retry-zh.md#包装错误与底层原因)。

SDK 错误的底层原因仍通过 `OrigErr()` 访问；本次调整没有为公共错误类型增加 `Unwrap()`。

### 全量异常详情

`volcengineerr.PrintError(writer, err)` 主动输出带缩进的 JSON，末尾追加换行；返回值表示打印是否成功。`volcengineerr.GetErrorDetails(err)` 提供相同的只读提取，可交给自己的日志系统。

```go
// 需要导入 fmt、os 和 volcengineerr。
if err != nil {
    if printErr := volcengineerr.PrintError(os.Stderr, err); printErr != nil {
        fmt.Fprintf(os.Stderr, "打印异常失败：%v\n", printErr)
    }
}
```

只需要打印时直接调用 `PrintError`，不需要先断言。需要单独读取 HTTP 响应时，可选择匹配新增的扩展接口：

```go
var failure volcengineerr.DetailedRequestFailure
if errors.As(err, &failure) {
    if response := failure.ResponseDetails(); response != nil {
        _ = response // 响应快照副本，可读取状态码、header 和 body。
    }
}
```

新增的 `DetailedRequestFailure` 扩展原有 `RequestFailure`，增加可选的 `ResponseDetails() *ResponseDetails` 方法。该方法只返回已采集的 HTTP 响应快照副本，没有快照时返回 nil；错误码和描述仍通过 `Code()` / `Message()` 获取。需要完整错误树时使用包函数 `volcengineerr.GetErrorDetails(err)`。旧 `Error` / `RequestFailure` 接口和构造函数签名保持不变，老客户的自定义错误无需增加方法。客户解析回调返回的错误保持原对象、具体类型、属性和 `Error()` 文本，包括客户自己通过 SDK 构造函数创建的错误；SDK 不在最终返回前统一追加详情。

`ErrorDetails` 提供 `Type`、`Error`、`Code`、`Message`、`Response` 和 `Causes`。`ResponseDetails` 提供 `StatusCode`、`Status`、`Proto`、`RequestID`、多值 `Headers` / `Trailers`、原始字节 `Body`、`BodyCaptured` 和 `BodyComplete`。正文中的 RequestID 仍优先于响应头；HTTP 200 业务错误的状态码仍为 200。结构化字段不受 `SimpleError` 影响，旧 `Error()` / `String()` 显示不变。

完整快照覆盖默认 TOP 解析器自己构造的请求错误，以及 STS / AssumeRole / OIDC / SAML 凭证请求错误，包含 OIDC/SAML 过期时间解析失败。复用原有正文读取，不额外发请求或读取关闭后的 body，仅保存错误对应的尝试。读取中断保留部分字节，`body_complete` 为 false，原因位于 `causes`。

`Request.SetResponseBody` 补充已有快照的正文时，保留状态码、协议、RequestID 和 headers，只更新正文、`BodyCaptured` / `BodyComplete` 以及 trailers。更新使用快照副本，不修改调用方传入的旧快照。正文复制保存；有 `HTTPResponse` 时复制其 trailers，没有时保留已有 trailers。已明确初始化为 nil 的快照不会从占位响应重新创建。

通过 `request.WithResponseReadTimeout` 触发的默认 TOP 响应读取超时，也保留已读取的部分正文、HTTP 状态码、响应头及请求 ID；仍返回原有 `ResponseTimeout` 类型和文本，`OrigErr()` 为 nil，此类错误不额外生成 `causes`。每个请求独立保存快照，不修改共享超时错误。

没有 HTTP 响应时，`response` 为 `null`；旧接口只能提供状态码和 RequestID 时，`headers` 为 `null`、`body_captured` 为 false。采集到空正文时，`body_captured` 为 true、`body` 为 `""`。JSON 中 UTF-8 正文使用文本和 `body_encoding: "utf-8"`，其它字节使用可还原的 Base64。`GetErrorDetails` 返回深拷贝；支持 `%w`、`OrigErr()`、`OrigErrs()` 及多分支标准包装，遍历具有循环保护和深度上限。

错误树去重支持字段中含 slice 或 map 的自定义值类型错误；不同指针错误保持独立，不会仅因内容相同而合并。

客户自定义错误，以及 ECS/IMDS、Endpoint credentials 和 CLI OAuth/Portal 等未接入路径中仍原样返回的读取/解码错误，仅输出实际已有信息，不承诺完整快照。客户实现可选 `ResponseDetails()` 方法且返回 nil 时，包函数自动退回旧接口字段。`GetErrorDetails(nil)` 返回 nil，`PrintError(writer, nil)` 不输出；写失败和短写作为可解包的打印错误返回。完整正文可能包含凭证等内容，调用方可修改详情副本后自行输出。

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

[← 重试机制](6-Retry-zh.md) | 异常处理[(English)](7-ErrorHandling.md) | [Debug 机制 →](8-Debugging-zh.md)
