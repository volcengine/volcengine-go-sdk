[← Credentials](1-Credentials.md) | Endpoint[(中文)](2-Endpoint-zh.md) | [Transport →](3-Transport.md)

---

## Endpoint Configuration

> **Default**
>
> If `Endpoint` is not specified, the SDK uses [Automatic Endpoint Resolution](#automatic-endpoint-resolution).

### Custom Endpoint

`volcengine.Config.Endpoint` (`WithEndpoint`) configures the business service address;
it does not configure the STS address used to obtain credentials. STS AssumeRole
defaults to `sts.volcengineapi.com`. Set the provider's `Host` or an independent
`credentials.STSConfig.Host` to override it. The provider's `Host` takes precedence,
and an empty `STSConfig.Host` preserves the previously bound STS host.
OIDC/SAML use their own provider's `Endpoint`. Overriding the business endpoint
when creating a session, copying it, or creating a service client does not change
these STS addresses.

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

### Custom RegionId

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

### Automatic Endpoint Resolution

Volcengine provides a flexible endpoint resolution mechanism. The SDK automatically builds the endpoint based on service name and region, and supports DualStack.

#### Default Endpoint Resolution

##### Resolution Logic

1. **Whether the region is in the bootstrap list**

    Built-in list implementation: [`./volcengine/volcengineutil/url.go#bootstrapRegion`](./volcengine/volcengineutil/url.go#L463).

    Only predefined regions (e.g., `cn-beijing-autodriving`, `ap-southeast-2`) or user-configured regions are auto-resolved; others fall back to `open.volcengineapi.com`.

    You can extend the list via env var `VOLC_BOOTSTRAP_REGION_LIST_CONF` or `customBootstrapRegion`.

2. **DualStack support (IPv6)**

    Enable via `useDualStack=true` or env var `VOLC_ENABLE_DUALSTACK=true`. Priority: `useDualStack` > `VOLC_ENABLE_DUALSTACK`.

    When enabled, the suffix changes from `volcengineapi.com` to `volcengine-api.com`.

3. **Construct endpoint based on service name and region**

    - **Global services (e.g., `CDN`, `IAM`)**: `<service>.volcengineapi.com` (or `volcengine-api.com` when DualStack is enabled). Example: `cdn.volcengineapi.com`.
    - **Regional services (e.g., `ECS`, `RDS`)**: `<service>.<region>.volcengineapi.com` is used as the default endpoint. Example: `ecs.cn-beijing.volcengineapi.com`.

##### Code Example

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

##### Region and service validation

The service and region are concatenated into the endpoint host name. To prevent a crafted region or service from rewriting the request host (SSRF), the SDK checks that both are a valid single host name label before resolving the endpoint:

- only ASCII letters, digits, `-` and `_`;
- starts with a letter;
- does not end with `-` or `_`.

For example `cn-beijing` and `rds_mysql` are valid; `cn-beijing@127.0.0.1`, `cn-beijing:443`, `evil.com`, `cn beijing`, `1cn-beijing` and `cn-beijing-` are not. The service is checked first, then the region. An empty value is not checked here and is handled by the existing logic.

When the check runs and what it covers:

- When no endpoint is set explicitly (`Config.Endpoint` is nil), the session runs the check before any endpoint resolution, so generated service clients (e.g. `ecs.New(sess)`) and the `universal` client are covered. With a custom `EndpointResolver` the check still runs first, and an invalid value is never passed to the resolver.
- On failure `ecs.New(sess)` and similar constructors do not panic; the error is returned in the request Validate stage, before signing and sending, so no network request is made.
- `EndpointFor` of the standard resolver `endpoints.NewStandardEndpointResolver()` runs the same check itself, so calling it directly also returns the error.
- `BuildAuthToken` in `feature/rds/auth` runs the check before building the token and returns the error without building a token.
- When an endpoint is set explicitly (`WithEndpoint`), the host is not built from the region or service and the check is skipped.

A failed check returns a `volcengineerr.Error` with the code `endpoints.ErrCodeInvalidRegion` (`"InvalidRegion"`) or `endpoints.ErrCodeInvalidService` (`"InvalidService"`), available via `Code()`. The message has the following format, where the reason is one of `illegal character at index N`, `must start with a letter` or `must not end with '-' or '_'`:

```
InvalidRegion: invalid region "cn-beijing@127.0.0.1": illegal character at index 10, region must start with a letter, end with a letter or digit, and contain only letters, digits, '-' and '_'
```

The error returned by `BuildAuthToken` is additionally wrapped as `unable to resolve client config: ...`, so use `errors.As` to check it:

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
		WithCredentials(credentials.NewEnvCredentials()).
		WithRegion("cn-beijing@127.0.0.1") // invalid region
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

The exported functions `endpoints.ValidateRegion`, `endpoints.ValidateService` and `endpoints.ValidateServiceAndRegion` can be used to check values up front; they return the same errors.

**Behavior change:**

- A region with leading or trailing whitespace (e.g. `" cn-beijing "`) is no longer trimmed; the client request returns `InvalidRegion`. Previous versions only ignored the whitespace when matching the bootstrap region list, and always signed with the raw value including the whitespace.
- This only applies to regions set in code (such as `WithRegion`). A region read from the `VOLCSTACK_REGION` / `VOLCSTACK_DEFAULT_REGION` environment variables or from the shared config file is trimmed when it is read (leading and trailing whitespace such as spaces, tabs and line breaks are removed, quoted values included), and a whitespace-only environment variable is treated as not set. The trimmed value is then validated like any other region.
- The exported function `volcengineutil.GetDefaultEndpointByServiceInfo` has no error return. For an invalid service or region it does not fail; it logs a Warn message (when the given logger is not nil) and falls back to the default endpoint `open.volcengineapi.com`. Callers of this function that need to detect the error should call `endpoints.ValidateServiceAndRegion` first.

#### Standard Endpoint Resolution

##### Resolution Rules

| Global service | DualStack | Format |
|---|---|---|
| Yes | Yes | `{Service}.volcengine-api.com` |
| Yes | No  | `{Service}.volcengineapi.com` |
| No  | Yes | `{Service}.{region}.volcengine-api.com` |
| No  | No  | `{Service}.{region}.volcengineapi.com` |

Whether a service is global depends on the service itself and cannot be changed. See: [`./volcengine/endpoints/standard_resolver.go#ServiceInfos`](./volcengine/endpoints/standard_resolver.go#L69).

##### Code Example

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
