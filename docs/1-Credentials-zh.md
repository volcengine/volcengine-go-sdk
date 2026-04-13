[← 概览](0-Overview-zh.md) | 访问凭据[(English)](1-Credentials.md) | [Endpoint 配置 →](2-Endpoint-zh.md)

---

# 访问凭据

火山引擎 Go SDK 支持多种认证方式，开发者可根据业务需求选择合适的方式接入。

## 凭证提供者概览

| 提供者 | 用途 | 自动刷新 | 典型场景 |
| --- | --- | --- | --- |
| `StaticCredentials` | 静态 AK/SK(/Token) | 否 | 长期服务端凭证 |
| `EnvCredentials` | 从环境变量读取 | 否 | CI/CD 和容器环境注入 |
| `StsCredentials` | STS AssumeRole | 是 | 基于角色的临时凭证 |
| `OIDCCredentialsProvider` | STS AssumeRoleWithOIDC | 是 | OIDC 联合身份认证 |
| `SAMLCredentialsProvider` | STS AssumeRoleWithSAML | 是 | SAML 联合身份认证 |
| `CliProvider` | 从 `~/.volcengine/config.json` 读取 | 取决于 mode | 复用 CLI 登录/配置 |
| `EcsRoleProvider` | 从 ECS IMDS (IMDSv2) 读取 | 是 | ECS 实例角色凭证 |
| `DefaultCredentialProvider` | 4 步凭证链 | 取决于委托的提供者 | 应用代码中无需 AK/SK |

环境变量设置可以参考这里:[**环境变量设置**](../SDK_Integration_zh.md#环境变量设置)

## AK、SK设置

AK/SK 是由火山引擎用户在控制台创建的一对永久访问密钥。SDK 使用该密钥对每次请求进行签名，从而完成身份验证。

> ⚠️ 注意事项
>
> 1. 不得在客户端嵌入或暴露 AK/SK。
> 2. 推荐使用配置中心或环境变量存储密钥。
> 3. 配置合理的最小权限访问策略。

```go
func main() {
    ak, sk, region := "Your AK", "Your SK", "cn-beijing"
    config := volcengine.NewConfig().
       WithRegion(region).
       // 1. credentials.NewStaticCredentials 是输入静态ak和sk可能泄漏会导致AK/SK泄漏，生产环境不能这样使用
       WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
       // 2. credentials.NewEnvCredentials() 不用传入任何参数，会从环境变量中读取：VOLCENGINE_ACCESS_KEY、VOLCENGINE_SECRET_KEY、VOLCENGINE_SESSION_TOKEN，生产环境建议使用这个
       // WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
}
```

## STS Token设置

STS（Security Token Service）是火山引擎提供的临时访问凭证机制。开发者通过服务端调用 STS 接口获取临时凭证（临时 AK、SK 和 Token），有效期可配置，适用于安全要求较高的场景。

> ⚠️ 注意事项
>
> 1. 最小权限： 仅授予调用方访问所需资源的最小权限，避免使用 * 通配符授予全资源、全操作权限。
> 2. 设置合理的有效期: 请根据实际情况设置合理有效期，越短越安全，建议不要超过1小时。

```Go
func main() {
    ak, sk,token,region := "Your AK", "Your SK", "Your token", "cn-beijing"
    config := volcengine.NewConfig().
       WithRegion(region).
       // 1. credentials.NewStaticCredentials 是输入静态ak和sk可能泄漏会导致AK/SK泄漏，生产环境不推荐这样使用
       WithCredentials(credentials.NewStaticCredentials(ak, sk, token))
       // 2. credentials.NewEnvCredentials() 不用传入任何参数，会从环境变量中读取：VOLCENGINE_ACCESS_KEY、VOLCENGINE_SECRET_KEY、VOLCENGINE_SESSION_TOKEN，生产环境建议使用这个
       // WithCredentials(credentials.NewEnvCredentials())
    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
}
```

## AssumeRole

动态访问凭证信息，支持动态刷新，在STS临时Token过期前60S会进行自动的刷新，避免临界时间点Token过期

> ⚠️ 注意事项
>
> 1. 最小权限： 仅授予调用方访问所需资源的最小权限，避免使用 * 通配符授予全资源、全操作权限。
> 2. 设置合理的有效期: 请根据实际情况设置合理有效期，越短越安全，最长不能超过12小时。
> 3. 细粒度角色: 角色应绑定精细的访问控制策略，仅允许访问特定服务、资源、操作，防止角色滥用。

**方式一：使用 WithOptions（推荐）**

```go
func main() {
    config := volcengine.NewConfig().
        WithRegion("cn-beijing").
        WithCredentials(credentials.NewStsCredentialsWithOptions(
            os.Getenv("VOLCENGINE_ACCESS_KEY"),  // 子账号AK（必填）
            os.Getenv("VOLCENGINE_SECRET_KEY"),  // 子账号SK（必填）
            "RoleName",                          // 扮演角色名称（必填）
            "123456",                            // 被扮演的主账号ID（必填）
            // 以下为可选配置，不传则使用默认值
            // func(o *credentials.StsAssumeRoleOptions) {
            //     o.Host = "open.volcengineapi.com" // STS 域名
            //     o.Region = "cn-beijing"           // STS region
            //     o.Schema = "https"                // STS schema
            //     o.Timeout = 5 * time.Second       // 请求超时时间
            //     o.DurationSeconds = 900           // 临时凭证有效期（秒），默认 3600
            //     o.Policy = `{"Statement":[...]}`  // session policy JSON
            //     o.MaxRetries = volcengine.Int(3)   // 重试次数；nil 默认 3，0 关闭重试
            //     o.RetryInterval = 1 * time.Second // 重试间隔；<= 0 回退到 1s
            // },
        ))

    sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
}
```

**方式二：使用 StsValue 结构体**

```go
func main() {
    ak, sk, region := "Your AK", "Your SK", "cn-beijing"
    config := volcengine.NewConfig().
        WithRegion(region).
        WithCredentials(credentials.NewStsCredentials(credentials.StsValue{
            AccessKey:  ak,         // 子账号AK,最好从环境变量获取：os.Getenv("VOLCENGINE_ACCESS_KEY")
            SecurityKey: sk,        // 子账号SK，最好从环境变量获取：os.Getenv("VOLCENGINE_SECRET_KEY")
            RoleName:   "RoleName", // 扮演角色名称
            Host:       "Host",     // 请求的sts域名
            Region:     "Region",   // 请求sts的region信息
            AccountId:  "123456",   // 被扮演的主账号ID，即角色所属的主账号ID
            Schema:     "Schema",   // 请求sts的schema信息
            Timeout:    5 * time.Second, // 请求sts的超时时间
            DurationSeconds: 900,        // STS临时凭证过期时长，单位为秒
            // Policy: 可选的 session policy JSON，用于进一步收窄临时凭证的权限，例如：`{"Statement":[{"Effect":"Allow","Action":["vpc:DescribeVpcs"],"Resource":["*"]}]}`,
            MaxRetries:    volcengine.Int(3),    // 可选：AssumeRole 失败时的额外重试次数；nil 默认 3，0 表示关闭重试
            RetryInterval: 1 * time.Second, // 可选：重试间隔；<= 0 时回退到 1s
        }))

    sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
}
```

## AssumeRoleOIDC

STS AssumeRoleOIDC（Security Token Service）是火山引擎提供的临时访问凭证机制。开发者通过oidc_token在服务端调用 STS 接口获取临时凭证（临时 AK、SK 和 Token），有效期可配置，适用于安全要求较高的场景。

> ⚠️ 注意事项
>
> 1. 最小权限： 仅授予调用方访问所需资源的最小权限，避免使用 * 通配符授予全资源、全操作权限。
> 2. 设置合理的有效期: 请根据实际情况设置合理有效期，越短越安全，建议不要超过1小时。
> 3. Go SDK 中 OIDC Token 需要存储在文件中。

**方式一：使用 WithOptions（推荐）**

```go
func main() {
	p := credentials.NewOIDCCredentialsProviderWithOptions(
		"/path/to/oidc_token_file",              // OIDC Token 文件路径（必填）
		"Your Role Trn",                         // 角色 TRN（必填）
		func(o *credentials.OIDCProviderOptions) {
			// o.RoleSessionName = ""                         // env: VOLCENGINE_OIDC_ROLE_SESSION_NAME（可选）
			// o.Policy = ""                                  // env: VOLCENGINE_OIDC_ROLE_POLICY（可选）
			// o.Endpoint = ""                                // env: VOLCENGINE_OIDC_STS_ENDPOINT（可选）
			// o.DurationSeconds = 3600                       // 有效期，默认 3600
			// o.MaxRetries = volcengine.Int(3)               // 可选：失败时额外重试次数；nil 默认 3，0 表示关闭重试
			// o.RetryInterval = 1 * time.Second              // 可选：重试间隔；<= 0 时回退到 1s
		},
	)

	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewCredentials(p))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	svc := vpc.New(sess)
	resp, err := svc.DescribeVpcs(&vpc.DescribeVpcsInput{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

> 提示：`credentials.NewOIDCCredentialsProviderFromEnv()` 可从环境变量直接构造，无需传入任何参数。

**方式二：使用结构体字面量**

```go
func main() {
	p := &credentials.OIDCCredentialsProvider{
		OIDCTokenFilePath: "/path/to/oidc_token_file", // env: VOLCENGINE_OIDC_TOKEN_FILE（必填）
		RoleTrn:           "Your Role Trn",            // env: VOLCENGINE_OIDC_ROLE_TRN（必填）
		RoleSessionName:   "",              // env: VOLCENGINE_OIDC_ROLE_SESSION_NAME（可选）
		Policy:            "",              // env: VOLCENGINE_OIDC_ROLE_POLICY（可选）
		Endpoint:          "",              // env: VOLCENGINE_OIDC_STS_ENDPOINT（可选）
		DurationSeconds:   3600,            // 有效期
		MaxRetries:        volcengine.Int(3),    // 可选：失败时额外重试次数；nil 默认 3，0 表示关闭重试
		RetryInterval:     1 * time.Second, // 可选：重试间隔；<= 0 时回退到 1s
	}

	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewCredentials(p))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	svc := vpc.New(sess)
	resp, err := svc.DescribeVpcs(&vpc.DescribeVpcsInput{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

## AssumeRoleWithSAML

`SAMLCredentialsProvider` 通过 SAML 2.0 IdP 返回的 SAML 断言调用 STS `AssumeRoleWithSAML` 接口换取临时凭证，并在到期前自动刷新。

> ⚠️ 注意事项
>
> 1. **最小权限原则**：仅授予必要的权限。
> 2. **合理的有效期**：建议不超过 1 小时。
> 3. `SAMLAssertion` 为 IdP 返回的 base64 编码的 SAML Response。

**方式一：使用 WithOptions（推荐）**

```go
package main

import (
    "github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
    "github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
    p := credentials.NewSAMLCredentialsProviderWithOptions(
        "trn:iam::1234567890:role/saml-role",                // 角色 TRN（必填）
        "trn:iam::1234567890:saml-provider/MyIdp",           // SAML 提供者 TRN（必填）
        "BASE64_ENCODED_SAML_RESPONSE_FROM_IDP",             // SAML 断言（必填）
        func(o *credentials.SAMLProviderOptions) {
            // o.DurationSeconds = 3600                                 // 有效期，默认 3600
            // o.MaxRetries = volcengine.Int(3)                         // 可选：失败时额外重试次数；nil 默认 3，0 表示关闭重试
            // o.RetryInterval = 1 * time.Second                        // 可选：重试间隔；<= 0 时回退到 1s
        },
    )

    config := volcengine.NewConfig().
        WithRegion("cn-beijing").
        WithCredentials(credentials.NewCredentials(p))

    sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
    _ = sess
}
```

**方式二：使用便捷构造函数**

```go
package main

import (
    "time"

    "github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
    "github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
    p := credentials.NewSAMLCredentialsProvider(
        "trn:iam::1234567890:role/saml-role",                // RoleTrn
        "trn:iam::1234567890:saml-provider/MyIdp",           // SAMLProviderTrn
        "BASE64_ENCODED_SAML_RESPONSE_FROM_IDP",             // SAMLAssertion
    )
    p.DurationSeconds = 3600
    p.MaxRetries = volcengine.Int(3)         // 可选：失败时额外重试次数；nil 默认 3，0 表示关闭重试
    p.RetryInterval = 1 * time.Second  // 可选：重试间隔；<= 0 时回退到 1s

    config := volcengine.NewConfig().
        WithRegion("cn-beijing").
        WithCredentials(credentials.NewCredentials(p))

    sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
    _ = sess
}
```

## 环境变量凭证提供者

`EnvProvider` 从环境变量中读取凭证。优先级顺序：

- Access Key：`VOLCENGINE_ACCESS_KEY` > `VOLCSTACK_ACCESS_KEY_ID` > `VOLCSTACK_ACCESS_KEY`
- Secret Key：`VOLCENGINE_SECRET_KEY` > `VOLCSTACK_SECRET_ACCESS_KEY` > `VOLCSTACK_SECRET_KEY`
- Session Token：`VOLCENGINE_SESSION_TOKEN` > `VOLCSTACK_SESSION_TOKEN`（可选）

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewEnvCredentials())

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	_ = sess
}
```

## CLI 配置文件凭证提供者

`CliProvider` 从 volcengine-cli 配置文件 (`~/.volcengine/config.json`) 读取凭证。

- 配置文件路径优先级：构造参数 > `VOLCENGINE_CLI_CONFIG_FILE` 环境变量 > `~/.volcengine/config.json`
- Profile 优先级：构造参数 > `VOLCENGINE_PROFILE` > `VOLCSTACK_PROFILE` > 配置中的 `current` > `default`

支持的 Profile 模式（不区分大小写）：

| 模式 | 说明 |
| --- | --- |
| `ak` / 空 | 从 profile 中读取静态 AK/SK |
| `ststoken` | 静态 AK/SK + SessionToken |
| `sso` | SSO 登录（OIDC Device Authorization） |
| `ramrolearn` | STS AssumeRole（委托给 `StsProvider`） |
| `oidc` | STS AssumeRoleWithOIDC（委托给 `OIDCCredentialsProvider`） |
| `ecsrole` | ECS IMDS（委托给 `EcsRoleProvider`） |

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials/clicreds"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	// 使用默认配置路径和 profile
	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(clicreds.NewCliCredentials("", ""))

	// 或者指定 profile
	// WithCredentials(clicreds.NewCliCredentials("", "prod"))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	_ = sess
}
```

## ECS 实例角色凭证提供者

`EcsRoleProvider` 通过 ECS 实例元数据服务 (IMDSv2) 获取临时凭证。

- 角色名优先级：构造参数 > `VOLCENGINE_ECS_METADATA` 环境变量 > 从 IMDS 自动检测
- 禁用开关：`VOLCENGINE_ECS_METADATA_DISABLED=true`
- IMDS 端点：`http://100.96.0.96`（IMDSv2 基于 token 的认证）
- 凭证在过期前自动刷新（5 分钟缓冲窗口）

> ⚠️ 注意事项
>
> 1. 仅在绑定了 IAM 角色的 ECS 实例上可用。
> 2. 自动检测会查询 IMDS 角色列表并使用找到的第一个角色。

**方式一：使用 WithOptions（推荐，支持自定义重试等配置）**

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	p := credentials.NewEcsRoleProviderWithOptions(
		"your-ecs-role-name", // 角色名（为空时从 IMDS 自动检测）
		func(o *credentials.EcsRoleProviderOptions) {
			// o.MaxRetries = volcengine.Int(3)                // 可选：失败时额外重试次数；nil 默认 3，0 表示关闭重试
			// o.RetryInterval = 1 * time.Second               // 可选：重试间隔；<= 0 时回退到 1s
		},
	)

	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewCredentials(p))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	_ = sess
}
```

**方式二：使用便捷构造函数**

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	// 显式指定角色名
	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewEcsRoleCredentials("your-ecs-role-name"))

	// 或者从 IMDS 自动检测角色名
	// WithCredentials(credentials.NewEcsRoleCredentials(""))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	_ = sess
}
```

## 默认凭证提供者（凭证链）

当未显式配置凭证时，SDK 会自动使用 `DefaultCredentialProvider` —— 一个 4 步凭证链，按顺序尝试每个提供者，直到成功：

1. **EnvProvider** — 环境变量（`VOLCENGINE_ACCESS_KEY` / `VOLCSTACK_ACCESS_KEY_ID`）
2. **OIDCCredentialsProvider** — 从环境变量读取 OIDC（`VOLCENGINE_OIDC_TOKEN_FILE`、`VOLCENGINE_OIDC_ROLE_TRN`）
3. **CliProvider** — CLI 配置文件（`~/.volcengine/config.json`）
4. **EcsRoleProvider** — ECS IMDS（实例元数据）

默认启用 `reuseLastProviderEnabled=true`：首次成功解析后，后续调用复用缓存的提供者以提升性能，仅在缓存提供者失败时才回退到完整链遍历。

**隐式使用**（不配置凭证 —— SDK 自动使用默认凭证链）：

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	// 不设置 WithCredentials —— SDK 自动使用 DefaultCredentialProvider
	config := volcengine.NewConfig().
		WithRegion("cn-beijing")

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	_ = sess
}
```

**显式使用**（自定义默认凭证链，例如指定 ECS 角色名）：

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/defaults"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(defaults.NewDefaultCredentialProvider(
			func(o *credentials.DefaultCredentialProviderOptions) {
				o.RoleName = "my-ecs-role"
			},
		))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	_ = sess
}
```

---

[← 概览](0-Overview-zh.md) | 访问凭据[(English)](1-Credentials.md) | [Endpoint 配置 →](2-Endpoint-zh.md)
