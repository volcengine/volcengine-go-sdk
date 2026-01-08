English | [中文](SDK_Integration_zh.md)

# Table of Contents

- [SDK Integration](#sdk-integration)
- [Requirements](#requirements)
- [Securely Configure Access Credentials](#securely-configure-access-credentials)
  - [Environment Variable Setup](#environment-variable-setup)
    - [Linux](#linux)
    - [Windows](#windows)
      - [GUI Setup](#gui-setup)
      - [Command Line Setup](#command-line-setup)
- [Credentials](#credentials)
  - [AK/SK](#aksk)
  - [STS Token](#sts-token)
  - [AssumeRole](#assumerole)
- [Endpoint Configuration](#endpoint-configuration)
  - [Custom Endpoint](#custom-endpoint)
  - [Custom RegionId](#custom-regionid)
  - [Automatic Endpoint Resolution](#automatic-endpoint-resolution)
    - [Default Endpoint Resolution](#default-endpoint-resolution)
    - [Standard Endpoint Resolution](#standard-endpoint-resolution)
- [HTTP Connection Pool](#http-connection-pool)
- [HTTPS Request Configuration](#https-request-configuration)
  - [Specify Scheme](#specify-scheme)
  - [Ignore SSL Verification](#ignore-ssl-verification)
  - [Specify TLS Version](#specify-tls-version)
- [HTTP(S) Proxy](#https-proxy)
  - [Configure HTTP(S) Proxy](#configure-https-proxy)
  - [Notes](#notes)
- [Timeouts](#timeouts)
  - [Global Timeouts (Client Level)](#global-timeouts-client-level)
  - [Per-API Timeout](#per-api-timeout)
- [Retries](#retries)
  - [Enable/Disable Retries](#enabledisable-retries)
  - [Retry Count](#retry-count)
  - [Custom Retry Error Codes](#custom-retry-error-codes)
- [Error Handling](#error-handling)
- [Debugging](#debugging)
  - [Enable Debug Mode](#enable-debug-mode)
- [Log Output](#log-output)

# SDK Integration

When calling APIs, it is recommended to integrate the SDK in your project. Using the SDK simplifies development, speeds up integration, and reduces long-term maintenance costs. Volcengine SDK integration typically includes three steps: importing the SDK, configuring access credentials, and writing API call code. This document explains common scenarios and best practices.

# Requirements

1. Go version **>= 1.14**.
2. If you use Ark service (`service/arkruntime`), Go version **>= 1.18** is required.
3. It is recommended to use `go mod` for dependency management.

# Securely Configure Access Credentials

To prevent credential leakage, do not hardcode credentials in plaintext in your code. Volcengine provides multiple secure ways to load credentials, such as environment variables.

## Environment Variable Setup

### Linux

> ⚠️ **Note**
>
> Environment variables configured with `export` are only effective for the current session. They will be lost when the session ends. To persist them, add the exports to your shell startup scripts.

| Key | Command |
| --- | --- |
| `VOLCSTACK_ACCESS_KEY_ID` (or `VOLCSTACK_ACCESS_KEY`) | `export VOLCSTACK_ACCESS_KEY_ID=yourAccessKeyID` |
| `VOLCSTACK_SECRET_ACCESS_KEY` (or `VOLCSTACK_SECRET_KEY`) | `export VOLCSTACK_SECRET_ACCESS_KEY=yourSecretAccessKey` |
| `VOLCSTACK_SESSION_TOKEN` | `export VOLCSTACK_SESSION_TOKEN=yourSessionToken` |

**Verify**: run `echo $VOLCSTACK_ACCESS_KEY_ID`. If it returns the correct Access Key ID, the configuration is effective.

### Windows

Two options are provided: **GUI setup** and **command line setup**.

**Verify**: open Command Prompt and run `echo %VOLCSTACK_ACCESS_KEY_ID%`, `echo %VOLCSTACK_SECRET_ACCESS_KEY%`, `echo %VOLCSTACK_SESSION_TOKEN%`. If the returned values are correct, the configuration is effective.

#### GUI Setup

On Windows 10, right-click **This PC** -> **Properties** -> **Advanced system settings** -> **Environment Variables** -> **System variables/User variables** -> **New**, then set:

| Variable | Example |
| --- | --- |
| AccessKey Id | Name: `VOLCSTACK_ACCESS_KEY_ID` <br/> Value: `*****` |
| AccessKey Secret | Name: `VOLCSTACK_SECRET_ACCESS_KEY` <br/> Value: `*****` |
| Session Token | Name: `VOLCSTACK_SESSION_TOKEN` <br/> Value: `*****` |

#### Command Line Setup

Run Command Prompt as Administrator, and add environment variables:

```
setx VOLCSTACK_ACCESS_KEY_ID yourAccessKeyID /M
setx VOLCSTACK_SECRET_ACCESS_KEY yourAccessKeySecret /M
setx VOLCSTACK_SESSION_TOKEN yourSessionToken /M
```

> ⚠️ Note
>
> `/M` means system-level variables. You may omit it for user-level variables.

# Credentials

Volcengine SDK supports three common authentication mechanisms: **AK/SK**, **STS temporary credentials**, and **AssumeRole**. Choose the one that best matches your scenario.

You can refer to: [Environment Variable Setup](#environment-variable-setup)

## AK/SK

AK/SK is a pair of permanent access keys created in the Volcengine console. The SDK signs each request to authenticate.

> ⚠️ Notes
>
> 1. Do not embed or expose AK/SK in client-side applications.
> 2. Use a configuration center or environment variables.
> 3. Follow least privilege principles.

```go
func main() {
    ak, sk, region := "Your AK", "Your SK", "cn-beijing"
    config := volcengine.NewConfig().
       WithRegion(region).
       // 1. NewStaticCredentials may leak credentials; do not use in production.
       WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
       // 2. Recommended in production: read from env vars.
       // WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
}
```

## STS Token

STS (Security Token Service) provides temporary credentials (temporary AK/SK and Token). You can configure validity duration. It is recommended for high-security scenarios.

> ⚠️ Notes
>
> 1. Least privilege: only grant required permissions.
> 2. Use a reasonable TTL. Shorter is safer; avoid exceeding 1 hour.

```go
func main() {
    ak, sk, token, region := "Your AK", "Your SK", "Your token", "cn-beijing"
    config := volcengine.NewConfig().
       WithRegion(region).
       WithCredentials(credentials.NewStaticCredentials(ak, sk, token))
       // WithCredentials(credentials.NewEnvCredentials())
    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
}
```

## AssumeRole

AssumeRole supports dynamic credentials with automatic refresh. The SDK refreshes before STS token expiry (buffer 60s) to avoid failures at the boundary.

> ⚠️ Notes
>
> 1. Least privilege.
> 2. Choose a reasonable TTL; maximum is 12 hours.
> 3. Use fine-grained roles and policies.

```go
func main() {
    ak, sk, region := "Your AK", "Your SK", "cn-beijing"
    config := volcengine.NewConfig().
        WithRegion(region).
        WithCredentials(credentials.NewStsCredentials(credentials.StsValue{
            AccessKey:  ak,
            SecurityKey: sk,
            RoleName:   "RoleName",
            Host:       "Host",
            Region:     "Region",
            AccountId:  "123456",
            Schema:     "Schema",
            Timeout:    5 * time.Second,
            DurationSeconds: 900,
        }))

    sess, err := session.NewSession(config)
    if err != nil {
        panic(err)
    }
}
```

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

# HTTP Connection Pool

> - **Default**
>   - `MaxIdleConns`: 100
>   - `IdleConnTimeout`: 90s
>   - `MaxIdleConnsPerHost`: 2

You can customize an `http.Client` to adjust these settings.

```go
func main() {
    region := "cn-beijing"
    transport := &http.Transport{
       Proxy: http.ProxyFromEnvironment,
       DialContext: (&net.Dialer{
          Timeout:   30 * time.Second,
          KeepAlive: 30 * time.Second,
          DualStack: true,
       }).DialContext,
       MaxIdleConns:          100,
       IdleConnTimeout:       90 * time.Second,
       MaxIdleConnsPerHost:   10,
       TLSHandshakeTimeout:   10 * time.Second,
       ExpectContinueTimeout: 1 * time.Second,
    }

    client := &http.Client{
       Transport: transport,
       Timeout:   60 * time.Second,
    }
    config := volcengine.NewConfig().
       WithRegion(region).
       WithHTTPClient(client).
       WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
    svc := ecs.New(sess)
    _ = svc
}
```

# HTTPS Request Configuration

## Specify Scheme

> - **Default**: `https`

In the SDK, `disableSSL=true` means using `http`, and `disableSSL=false` means using `https`.

```go
func main() {
    region := "cn-beijing"
    config := volcengine.NewConfig().
       WithRegion(region).
       WithDisableSSL(true).
       WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
    svc := ecs.New(sess)
    _ = svc
}
```

## Ignore SSL Verification

You can customize `http.Client` to skip certificate verification.

```go
func main() {
    region :=  "cn-beijing"
    transport := &http.Transport{
       Proxy: http.ProxyFromEnvironment,
       DialContext: (&net.Dialer{
          Timeout:   30 * time.Second,
          KeepAlive: 30 * time.Second,
          DualStack: true,
       }).DialContext,
       MaxIdleConns:          100,
       IdleConnTimeout:       90 * time.Second,
       MaxIdleConnsPerHost:   10,
       TLSHandshakeTimeout:   10 * time.Second,
       ExpectContinueTimeout: 1 * time.Second,
       TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    client := &http.Client{
       Transport: transport,
       Timeout:   60 * time.Second,
    }
    config := volcengine.NewConfig().
       WithRegion(region).
       WithHTTPClient(client).
       WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
    svc := ecs.New(sess)
    _ = svc
}
```

## Specify TLS Version

You can customize TLS min/max versions via `TLSClientConfig`.

```go
func main() {
    region := "cn-beijing"
    transport := &http.Transport{
       Proxy: http.ProxyFromEnvironment,
       DialContext: (&net.Dialer{
          Timeout:   30 * time.Second,
          KeepAlive: 30 * time.Second,
          DualStack: true,
       }).DialContext,
       MaxIdleConns:          100,
       IdleConnTimeout:       90 * time.Second,
       MaxIdleConnsPerHost:   10,
       TLSHandshakeTimeout:   10 * time.Second,
       ExpectContinueTimeout: 1 * time.Second,
       TLSClientConfig: &tls.Config{
           MinVersion: tls.VersionTLS12,
           MaxVersion: tls.VersionTLS13,
       },
    }

    client := &http.Client{
       Transport: transport,
       Timeout:   60 * time.Second,
    }
    config := volcengine.NewConfig().
       WithRegion(region).
       WithHTTPClient(client).
       WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
    svc := ecs.New(sess)
    _ = svc
}
```

# HTTP(S) Proxy

> - **Default**: no proxy

## Configure HTTP(S) Proxy

```go
var ak, sk, region string
config = volcengine.NewConfig().
    WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
    WithRegion(region).WithHTTPProxy("http://your_proxy:8080").WithHTTPSProxy("http://your_proxy:8080")

sess, _ = session.NewSession(config)
client = ecsops.New(sess)
```

## Notes

Supported environment variables:

- `http_proxy`/`HTTP_PROXY`, `https_proxy`/`HTTPS_PROXY`

Priority: code > environment variables.

# Timeouts

## Global Timeouts (Client Level)

> - **Default**
>   - `ConnectTimeout`: 30s
>   - `ReadTimeout`: unlimited
>   - Default client: `http.DefaultClient`

Configure timeouts via a custom `http.Client`.

```go
func main() {
    region := "cn-beijing"
    transport := &http.Transport{
       Proxy: http.ProxyFromEnvironment,
       DialContext: (&net.Dialer{
          Timeout:   30 * time.Second,
          KeepAlive: 30 * time.Second,
          DualStack: true,
       }).DialContext,
       MaxIdleConns:          100,
       IdleConnTimeout:       90 * time.Second,
       TLSHandshakeTimeout:   10 * time.Second,
       ExpectContinueTimeout: 1 * time.Second,
    }

    client := &http.Client{
       Transport: transport,
       Timeout:   60 * time.Second,
    }
    config := volcengine.NewConfig().
       WithRegion(region).
       WithHTTPClient(client).
       WithCredentials(credentials.NewEnvCredentials())

    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
    svc := ecs.New(sess)
    _ = svc
}
```

## Per-API Timeout

Use API methods with suffix `WithContext`.

```go
func main() {
    region :=  "cn-beijing"
    config := volcengine.NewConfig().
       WithRegion(region).
       WithCredentials(credentials.NewEnvCredentials())
    sess, err := session.NewSession(config)
    if err != nil {
       panic(err)
    }
    svc := ecs.New(sess)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    resp, err := svc.AssociateInstancesIamRoleWithContext(ctx, associateInstancesIamRoleInput)
    if err != nil {
       panic(err)
    }
    fmt.Println(resp)
}
```

# Retries

The SDK includes retry logic for network errors and throttling. Business logic errors (invalid params, resource not found) are not retried.

## Enable/Disable Retries

> - **Default**: enabled (3 times)

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
    _ = svc
}
```

## Retry Count

> - **Default**: 3

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
    _ = svc
}
```

## Custom Retry Error Codes

Configure retryable server error codes per request.

```go
resp, err := svc.DescribeAvailableResourceWithContext(ctx, describeAvailableResourceInput, func(request *request.Request) {
    request.RetryErrorCodes = []string{"InvalidAccessKey"}
})
```

# Error Handling

Errors are categorized as:

| Type | Description | Error type |
|---|---|---|
| 1. Client error | Validation failures before reaching server | `volcengineerr.Error` or native `error` |
| 2. Server error | Request reaches server; business error returned | `volcengineerr.RequestFailure` |
| 3. Network/timeout | DNS or timeout | `volcengineerr.Error` |
| 4. Others | Other errors | `volcengineerr.Error` or native `error` |

(See the original Chinese guide for the full code sample.)

# Debugging

## Enable Debug Mode

Debug logs are disabled by default. Enable via `WithDebug(true)`.

```go
config := volcengine.NewConfig().
    WithRegion(region).
    WithDebug(true).
    WithCredentials(credentials.NewEnvCredentials())
```

# Log Output

By default, logs are written to `os.Stdout`. Use `WithLogWriter` to write to a file or other writer.

```go
file, _ := os.Create("sdk.log")
config := volcengine.NewConfig().
  WithRegion(region).
  WithDebug(true).
  WithLogWriter(file).
  WithCredentials(credentials.NewEnvCredentials())
```
