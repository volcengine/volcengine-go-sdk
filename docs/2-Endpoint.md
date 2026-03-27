[← Credentials](1-Credentials.md) | Endpoint[(中文)](2-Endpoint-zh.md) | [Transport →](3-Transport.md)

---

# Endpoint Configuration

## Custom Endpoint

> - **Default**: if `endpoint` is not specified, the SDK uses [Automatic Endpoint Resolution](#automatic-endpoint-resolution).

```go
func main() {
    region := "cn-beijing"
    config := volcengine.NewConfig().
       WithCredentials(credentials.NewEnvCredentials()).
       WithRegion(region).
       WithEndpoint("<example>.<regionId>.volcengineapi.com")
    sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
}
```

## Custom RegionId

```go
func main() {
    regionId := "cn-beijing"
    config := volcengine.NewConfig().
       WithCredentials(credentials.NewEnvCredentials()).
       WithRegion(regionId)
    sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
}
```

## Automatic Endpoint Resolution

> - **Default**: automatic resolution is enabled; no manual endpoint configuration is required.

Volcengine provides a flexible endpoint resolution mechanism. The SDK automatically builds the endpoint based on service name and region, and supports DualStack.

### Default Endpoint Resolution

**Logic**

1. Whether region is in the bootstrap list.
   - Built-in list implementation: `./volcengine/volcengineutil/url.go#bootstrapRegion`.
   - Only predefined regions (e.g., `cn-beijing-autodriving`, `ap-southeast-2`) or user-configured regions are auto-resolved; others fall back to `open.volcengineapi.com`.
   - You can extend the list via env var `VOLC_BOOTSTRAP_REGION_LIST_CONF` or `customBootstrapRegion`.
2. DualStack support (IPv6)
   - Enable via `useDualStack=true` or env var `VOLC_ENABLE_DUALSTACK=true`.
   - When enabled, the suffix changes from `volcengineapi.com` to `volcengine-api.com`.
3. Construct endpoint:
   - Global services: `<service>.volcengineapi.com`.
   - Regional services: `<service>.<region>.volcengineapi.com`.

**Example**

```go
func main() {
    regionId := "cn-beijing"
    config := volcengine.NewConfig().
        WithCredentials(credentials.NewEnvCredentials()).
        WithRegion(regionId).
        WithUseDualStack(true).
        WithBootstrapRegion(map[string]struct{}{
            "custom_example_region1": {},
            "custom_example_region2": {},
        })
    sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
}
```

### Standard Endpoint Resolution

| Global service | DualStack | Format |
|---|---|---|
| Yes | Yes | `{Service}.volcengine-api.com` |
| Yes | No  | `{Service}.volcengineapi.com` |
| No  | Yes | `{Service}.{region}.volcengine-api.com` |
| No  | No  | `{Service}.{region}.volcengineapi.com` |

Whether a service is global depends on the service itself and cannot be changed. See: `./volcengine/endpoints/standard_resolver.go#ServiceInfos`.

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
    WithEndpointResolver(endpoints.NewStandardEndpointResolver()).
    WithRegion(regionId).
    WithUseDualStack(true)
  sess, err := session.NewSession(config)
  if err != nil {
    panic(err)
  }
}
```

---

[← Credentials](1-Credentials.md) | Endpoint[(中文)](2-Endpoint-zh.md) | [Transport →](3-Transport.md)
