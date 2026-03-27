[← 访问凭据](1-Credentials-zh.md) | Endpoint 配置[(English)](2-Endpoint.md) | [Transport →](3-Transport-zh.md)

---

# EndPoint配置

## 自定义Endpoint

> - **默认**
>   不指定endpoint时，走[自动化Endpoint寻址](#自动化Endpoint寻址)

用户可以通过在初始化客户端时指定Endpoint

```go
func main() {
    region := "cn-beijing"
    config := volcengine.NewConfig().
       WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：VOLCSTACK_ACCESS_KEY_ID、VOLCSTACK_SECRET_ACCESS_KEY、VOLCSTACK_SESSION_TOKEN
       WithRegion(region).
       // 自定义Endpoint
       WithEndpoint("<example>.<regionId>.volcengineapi.com")
     sess, err := session.NewSession(config)
     if err != nil {
        panic(err)
     }
}
```

## 自定义RegionId

```go
func main() {
    regionId := "cn-beijing"
    config := volcengine.NewConfig().
       WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：VOLCSTACK_ACCESS_KEY_ID、VOLCSTACK_SECRET_ACCESS_KEY、VOLCSTACK_SESSION_TOKEN
       WithRegion(regionId) // 自定义regionId
	sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
}
```

## 自动化Endpoint寻址

> - **默认**
>   默认支持自动寻址，无需手动指定Endpoint

为了简化用户配置，Volcengine 提供了灵活的 Endpoint 自动寻址机制。用户无需手动指定服务地址，SDK 会根据服务名称、区域（Region）等信息自动拼接出合理的访问地址，并支持用户自定义DualStack（双栈）支持。
### Endpoint默认寻址
**Endpoint默认寻址逻辑**
1. 是否自动寻址Region
内置自动寻址Region列表代码:[./volcengine/volcengineutil/url.go#bootstrapRegion](./volcengine/volcengineutil/url.go#L463)
SDK 仅对部分预设区域（如 cn-beijing-autodriving、ap-southeast-2）或用户配置的区域执行自动寻址；其他区域默认返回Endpoint：open.volcengineapi.com。
用户可通过环境变量 VOLC_BOOTSTRAP_REGION_LIST_CONF 或代码中自定义 customBootstrapRegion 来扩展控制区域列表。
2. DualStack 支持（IPv6）
SDK 支持双栈网络（IPv4 + IPv6）访问地址，自动启用条件如下：
显式传入参数 useDualStack = true，或设置环境变量 VOLC_ENABLE_DUALSTACK=true，优先级useDualStack>VOLC_ENABLE_DUALSTACK
启用后，域名后缀将从 volcengineapi.com 切换为 volcengine-api.com。
3. 根据服务名和区域自动构造 Endpoint 地址，规则如下：
**全局服务（如 CDN、IAM）**
使用 <服务名>.volcengineapi.com（或启用双栈时使用 volcengine-api.com）。
示例：cdn.volcengineapi.com
**区域服务（如 ECS、RDS）**
使用 <服务名>.<区域名>.volcengineapi.com 作为默认 Endpoint。
示例：ecs.cn-beijing.volcengineapi.com

**代码示例：**

```go
func main() {
    regionId := "cn-beijing"
    config := volcengine.NewConfig().
        WithCredentials(credentials.NewEnvCredentials()). //环境变量配置：VOLCSTACK_ACCESS_KEY_ID、VOLCSTACK_SECRET_ACCESS_KEY、VOLCSTACK_SESSION_TOKEN
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

### Endpoint标准寻址
**标准寻址规则**

| Global服务 | 双栈 | 格式                                                                                                               |
|----------|----|------------------------------------------------------------------------------------------------------------------|
| 是        | 是  | `{Service}.volcengine-api.com`                                                                                   |
| 是        | 否  | `{Service}.volcengineapi.com`                                                                                    |
| 否        | 是  | `{Service}.{region}.volcengine-api.com`|
| 否        | 否  | `{Service}.{region}.volcengineapi.com` |

**代码示例：**

是否global服务根据具体调用的服务决定的，是否global无法修改的。
可以参考列表：[./volcengine/endpoints/standard_resolver.go#ServiceInfos](./volcengine/endpoints/standard_resolver.go#L69)
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
