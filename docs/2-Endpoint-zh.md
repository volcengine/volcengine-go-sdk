[← 访问凭据](1-Credentials-zh.md) | Endpoint 配置[(English)](2-Endpoint.md) | [Transport →](3-Transport-zh.md)

---

## EndPoint 配置

> **默认**
>
> 不指定 endpoint 时，走 [自动化 Endpoint 寻址](#自动化-endpoint-寻址)。

### 自定义 Endpoint

用户可以通过在初始化客户端时指定 Endpoint：

`volcengine.Config.Endpoint`（`WithEndpoint`）只配置业务服务地址，不会用作获取凭证的 STS 地址。
STS AssumeRole 默认使用 `sts.volcengineapi.com`，可通过 Provider 的 `Host` 或独立的
`credentials.STSConfig.Host` 指定；Provider 的 `Host` 优先，重新绑定时空的 `STSConfig.Host`
保留已有 STS Host。OIDC/SAML 使用各自 Provider 的 `Endpoint`。
在创建 Session、`Session.Copy` 或服务客户端时覆盖业务 Endpoint，不会改变这些 STS 地址。

```go
func main() {
	region := "cn-beijing"
	config := volcengine.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()). // 需要配置环境变量
		WithRegion(region).
		// 自定义Endpoint
		WithEndpoint("<example>.<regionId>.volcengineapi.com")
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

### 自定义 RegionId

```go
func main() {
	regionId := "cn-beijing"
	config := volcengine.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()). // 需要配置环境变量
		WithRegion(regionId)                              // 自定义regionId
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

### 自动化 Endpoint 寻址

> **默认**
>
> 默认支持自动寻址，无需手动指定 Endpoint。

为了简化用户配置，Volcengine 提供了灵活的 Endpoint 自动寻址机制。用户无需手动指定服务地址，SDK 会根据服务名称、区域（Region）等信息自动拼接出合理的访问地址，并支持用户自定义 DualStack（双栈）。

#### Endpoint 默认寻址

##### 寻址逻辑

1. **是否自动寻址 Region**

    内置自动寻址 Region 列表代码：[`./volcengine/volcengineutil/url.go#bootstrapRegion`](./volcengine/volcengineutil/url.go#L463)

    SDK 仅对部分预设区域（如 `cn-beijing-autodriving`、`ap-southeast-2`）或用户配置的区域执行自动寻址；其他区域默认返回 Endpoint：`open.volcengineapi.com`。

    用户可通过环境变量 `VOLC_BOOTSTRAP_REGION_LIST_CONF` 或代码中自定义 `customBootstrapRegion` 来扩展控制区域列表。

2. **DualStack 支持（IPv6）**

    SDK 支持双栈网络（IPv4 + IPv6）访问地址，自动启用条件如下：显式传入参数 `useDualStack = true`，或设置环境变量 `VOLC_ENABLE_DUALSTACK=true`。优先级：`useDualStack` > `VOLC_ENABLE_DUALSTACK`。

    启用后，域名后缀将从 `volcengineapi.com` 切换为 `volcengine-api.com`。

3. **根据服务名和区域自动构造 Endpoint 地址**

    - **全局服务（如 `CDN`、`IAM`）**：使用 `<服务名>.volcengineapi.com`（或启用双栈时使用 `volcengine-api.com`）。示例：`cdn.volcengineapi.com`。
    - **区域服务（如 `ECS`、`RDS`）**：使用 `<服务名>.<区域名>.volcengineapi.com` 作为默认 Endpoint。示例：`ecs.cn-beijing.volcengineapi.com`。

##### 代码示例

```go
func main() {
	regionId := "cn-beijing"
	config := volcengine.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()). // 需要配置环境变量
		WithRegion(regionId).
		WithUseDualStack(true). // 定义是否启用双栈网络（IPv4 + IPv6）访问地址，默认false；也可以使用环境变量VOLC_ENABLE_DUALSTACK=true
		WithBootstrapRegion(map[string]struct{}{
			"custom_example_region1": {},
			"custom_example_region2": {},
		}) // 自定义自动寻址Region列表；也可以使用环境变量VOLC_BOOTSTRAP_REGION_LIST_CONF
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

##### region / service 合法性校验

service 和 region 会被拼接进 Endpoint 的主机名。为防止通过构造 region / service 改写请求主机（SSRF），SDK 在寻址之前会先校验二者都是合法的单个主机名标签：

- 只能包含 ASCII 字母、数字、`-` 和 `_`；
- 必须以字母开头；
- 不能以 `-` 或 `_` 结尾。

例如 `cn-beijing`、`rds_mysql` 合法；`cn-beijing@127.0.0.1`、`cn-beijing:443`、`evil.com`、`cn beijing`、`1cn-beijing`、`cn-beijing-` 非法。先校验 service，再校验 region。空字符串不在此处校验，沿用原有逻辑处理。

校验时机与覆盖范围：

- 未显式设置 Endpoint（`Config.Endpoint` 为 nil）时，Session 在调用任何 Endpoint 寻址逻辑之前执行校验，生成的服务客户端（如 `ecs.New(sess)`）和 `universal` 客户端都会经过这一步。配置了自定义 `EndpointResolver` 时同样先校验，非法值不会传给 Resolver。
- 校验失败时 `ecs.New(sess)` 等构造过程不会 panic，错误在请求的 Validate 阶段返回，早于签名和发送，不会发出网络请求。
- `session.ClientConfigWithError(serviceName, ...)` 会在构造阶段直接返回解析 client config 时产生的错误（如非法 region / service，或自定义 `EndpointResolver` 返回的错误），方便调用方尽早失败。它返回的错误与 `session.ClientConfig(...)`（保留无 error 签名，生成的客户端使用它）携带的错误完全相同；错误原样返回，不加任何消息前缀，可按下方示例用 `errors.As` 识别。
- 标准寻址 `endpoints.NewStandardEndpointResolver()` 的 `EndpointFor` 自身也会执行该校验，直接调用时同样返回错误。
- `feature/rds/auth` 的 `BuildAuthToken` 在生成 token 之前执行该校验，失败时返回错误，不生成 token。
- 显式设置 Endpoint（`WithEndpoint`）时，主机名不由 region / service 拼接，不执行该校验。

校验失败返回 `volcengineerr.Error`，错误码为 `endpoints.ErrCodeInvalidRegion`（`"InvalidRegion"`）或 `endpoints.ErrCodeInvalidService`（`"InvalidService"`），通过 `Code()` 获取。错误信息格式如下，原因部分为 `illegal character at index N`、`must start with a letter` 或 `must not end with '-' or '_'` 之一：

```
InvalidRegion: invalid region "cn-beijing@127.0.0.1": illegal character at index 10, region must start with a letter, end with a letter or digit, and contain only letters, digits, '-' and '_'
```

`BuildAuthToken` 返回的错误会再包装一层 `unable to resolve client config: ...`，建议统一使用 `errors.As` 判断：

```go
package main

import (
	"errors"
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/endpoints"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

func main() {
	config := volcengine.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()). // 需要配置环境变量
		WithRegion("cn-beijing@127.0.0.1")                // 非法 region
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	svc := ecs.New(sess)
	_, err = svc.DescribeInstances(&ecs.DescribeInstancesInput{})

	var vErr volcengineerr.Error
	if errors.As(err, &vErr) {
		switch vErr.Code() {
		case endpoints.ErrCodeInvalidRegion, endpoints.ErrCodeInvalidService:
			fmt.Println(vErr.Code(), vErr.Message())
		}
	}
}
```

也可以使用导出的 `endpoints.ValidateRegion`、`endpoints.ValidateService`、`endpoints.ValidateServiceAndRegion` 提前校验，返回的错误与上面相同。

**与旧版本不同：**

- region 首尾带空白（如 `" cn-beijing "`）不会被 trim，客户端请求直接返回 `InvalidRegion`。旧版本只在判断是否属于自动寻址 Region 列表时忽略首尾空白，签名使用的始终是带空白的原值。
- 上述规则只针对代码中设置的 region（如 `WithRegion`）。从环境变量 `VOLCSTACK_REGION` / `VOLCSTACK_DEFAULT_REGION` 或共享配置文件读取的 region 会在读取时去掉首尾空白（包括空格、Tab、换行等，带引号的值同样处理），纯空白的环境变量视为未设置；去掉空白后的值仍按同样的规则校验。
- 只有开启共享配置加载时，才会从 `VOLCSTACK_DEFAULT_REGION` 和共享配置 / 共享凭证文件中读取 region；否则 region 只取 `VOLCSTACK_REGION`。开启方式：将 `VOLCSTACK_SDK_LOAD_CONFIG` 设为 Go `strconv.ParseBool` 认可的真值（`1`、`t`、`T`、`true`、`True`、`TRUE`，`yes` 等其他值不生效），或在代码中设置 `session.Options.SharedConfigState`（优先于环境变量）。这一条只影响 region：即使未开启，共享凭证文件仍会被读取，用于获取 AK/SK。
- 共享配置文件和共享凭证文件为同一 profile 都定义了 `region` 时，后加载的文件（共享凭证文件）生效。后一个文件里无引号的空值（`region =`）会被忽略，但带引号的纯空白值（`region = "   "`）视为已定义：去掉空白后变为空串，会清掉前一个文件里的值，此时 region 视为未设置。
- 导出函数 `volcengineutil.GetDefaultEndpointByServiceInfo` 没有 error 返回值，遇到非法 service / region 时不报错，而是记录 Warn 日志（传入的 logger 非 nil 时），并回退到默认 Endpoint `open.volcengineapi.com`。直接调用该函数且需要感知错误的调用方，请先调用 `endpoints.ValidateServiceAndRegion`。

#### Endpoint 标准寻址

##### 标准寻址规则

| Global 服务 | 双栈 | 格式 |
|---|---|---|
| 是 | 是 | `{Service}.volcengine-api.com` |
| 是 | 否 | `{Service}.volcengineapi.com` |
| 否 | 是 | `{Service}.{region}.volcengine-api.com` |
| 否 | 否 | `{Service}.{region}.volcengineapi.com` |

服务是否为 Global 由具体服务决定，不可修改。可以参考列表：[`./volcengine/endpoints/standard_resolver.go#ServiceInfos`](./volcengine/endpoints/standard_resolver.go#L69)。

##### 代码示例

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/endpoints"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	regionId := "cn-beijing"
	config := volcengine.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()).
		WithEndpointResolver(endpoints.NewStandardEndpointResolver()). // 配置标准寻址
		WithRegion(regionId).                                          // 配置regionId
		WithUseDualStack(true)                                         // 配置是否双栈
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

---

[← 访问凭据](1-Credentials-zh.md) | Endpoint 配置[(English)](2-Endpoint.md) | [Transport →](3-Transport-zh.md)
