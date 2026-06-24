[← Overview](0-Overview.md) | Credentials[(中文)](1-Credentials-zh.md) | [Endpoint →](2-Endpoint.md)

---

## Credentials

Volcengine Go SDK supports multiple authentication mechanisms. Choose the one that best matches your scenario.

### Credential Providers Overview

| Provider | Purpose | Refresh Support | Typical Scenario |
| --- | --- | --- | --- |
| `StaticCredentials` | Static AK/SK(/Token) | No | Long-lived server-side credentials |
| `EnvCredentials` | Read from env vars | No | CI/CD and container env injection |
| `StsCredentials` | STS AssumeRole | Yes | Role-based temporary credentials |
| `OIDCCredentialsProvider` | STS AssumeRoleWithOIDC | Yes | OIDC federation |
| `SAMLCredentialsProvider` | STS AssumeRoleWithSAML | Yes | SAML federation |
| `CliProvider` | Read from `~/.volcengine/config.json` | Depends on mode | Reuse CLI login/profile |
| `EcsRoleProvider` | Read from ECS IMDS (IMDSv2) | Yes | ECS instance role credentials |
| `DefaultCredentialProvider` | 4-step chain wrapper | Depends on delegated provider | No AK/SK in application code |

You can refer to: [Environment Variable Setup](0-Overview.md)

### AK/SK

AK/SK is a pair of permanent access keys created in the Volcengine console. The SDK signs each request to authenticate.

> ⚠️ **Notes**
>
> 1. Do not embed or expose AK/SK in client-side applications.
> 2. Use a configuration center or environment variables.
> 3. Follow least privilege principles.

```go
func main() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		// 1. credentials.NewStaticCredentials takes static AK/SK and may leak credentials; not recommended in production.
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
		// 2. credentials.NewEnvCredentials() takes no arguments and reads from env vars:
		//    VOLCENGINE_ACCESS_KEY, VOLCENGINE_SECRET_KEY, VOLCENGINE_SESSION_TOKEN. Recommended in production.
		// WithCredentials(credentials.NewEnvCredentials())

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

### STS Token

STS (Security Token Service) provides temporary credentials (temporary AK/SK and Token). You can configure validity duration. It is recommended for high-security scenarios.

> ⚠️ **Notes**
>
> 1. Least privilege: only grant required permissions.
> 2. Use a reasonable TTL. Shorter is safer; avoid exceeding 1 hour.

```go
func main() {
	ak, sk, token, region := "Your AK", "Your SK", "Your token", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		// 1. credentials.NewStaticCredentials takes static AK/SK(/Token) and may leak credentials; not recommended in production.
		WithCredentials(credentials.NewStaticCredentials(ak, sk, token))
		// 2. credentials.NewEnvCredentials() takes no arguments and reads from env vars:
		//    VOLCENGINE_ACCESS_KEY, VOLCENGINE_SECRET_KEY, VOLCENGINE_SESSION_TOKEN. Recommended in production.
		// WithCredentials(credentials.NewEnvCredentials())
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

### AssumeRole

AssumeRole supports dynamic credentials with automatic refresh. The SDK refreshes before STS token expiry (buffer 60s) to avoid failures at the boundary.

> ⚠️ **Notes**
>
> 1. Least privilege.
> 2. Choose a reasonable TTL; maximum is 12 hours.
> 3. Use fine-grained roles and policies.

#### Option 1: Using WithOptions (recommended)

```go
package main

import (
	"os"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewStsCredentialsWithOptions(
			os.Getenv("VOLCENGINE_ACCESS_KEY"), // Sub-account AK (required)
			os.Getenv("VOLCENGINE_SECRET_KEY"), // Sub-account SK (required)
			"RoleName",                         // Name of the role to assume (required)
			"123456",                           // Main account ID (required)
			// All options below are optional; omit the func to use defaults
			// func(o *credentials.StsAssumeRoleOptions) {
			//     o.Host = "open.volcengineapi.com" // STS host
			//     o.Region = "cn-beijing"           // STS region
			//     o.Schema = "https"                // STS schema
			//     o.Timeout = 5 * time.Second       // Request timeout
			//     o.DurationSeconds = 900           // TTL in seconds, default 3600
			//     o.Policy = `{"Statement":[...]}`  // Session policy JSON
			//     o.MaxRetries = 3                  // Retry attempts; 0 or negative falls back to DefaultRetryerMaxNumRetries (3)
			//     o.RetryInterval = 1 * time.Second // Sleep between retries; <= 0 falls back to 1s
			// },
		))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

#### Option 2: Using StsValue struct

```go
package main

import (
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStsCredentials(credentials.StsValue{
			AccessKey:       ak,              // Sub-account AK, preferably read from env: os.Getenv("VOLCENGINE_ACCESS_KEY")
			SecurityKey:     sk,              // Sub-account SK, preferably read from env: os.Getenv("VOLCENGINE_SECRET_KEY")
			RoleName:        "RoleName",      // Name of the role to assume
			Host:            "Host",          // STS host
			Region:          "Region",        // STS region
			AccountId:       "123456",        // Main account ID that owns the role
			Schema:          "Schema",        // STS schema (http/https)
			Timeout:         5 * time.Second, // STS request timeout
			DurationSeconds: 900,             // TTL of the temporary credentials, in seconds
			// Policy: optional session policy JSON, e.g. `{"Statement":[{"Effect":"Allow","Action":["vpc:DescribeVpcs"],"Resource":["*"]}]}`
			MaxRetries:    3,                 // optional extra retry attempts; 0 or negative falls back to DefaultRetryerMaxNumRetries (3)
			RetryInterval: 1 * time.Second,   // optional sleep between retries; <= 0 falls back to 1s
		}))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

### STS AssumeRoleWithOIDC Example

STS AssumeRoleOIDC (Security Token Service) is a temporary access credential mechanism provided by Volcengine. Developers use an `oidc_token` to call the STS interface on the server side to obtain temporary credentials (temporary AK, SK, and Token). The validity period is configurable, making it suitable for scenarios with high security requirements.

> ⚠️ **Notes**
>
> 1. **Least Privilege**: Grant only the minimum permissions required for the caller to access resources, avoiding the use of `*` wildcards to grant full resource and operation permissions.
> 2. **Reasonable Validity Period**: Set a reasonable validity period based on actual conditions. Shorter periods are safer; it is recommended not to exceed 1 hour.
> 3. **OIDC Token Storage**: In the Go SDK, the OIDC Token must be stored in a file.

#### Option 1: Using WithOptions (recommended)

```go
package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	p := credentials.NewOIDCCredentialsProviderWithOptions(
		"/path/to/oidc_token_file", // OIDC Token file path (required)
		"Your Role Trn",            // Role TRN (required)
		func(o *credentials.OIDCProviderOptions) {
			// o.RoleSessionName = ""                         // env: VOLCENGINE_OIDC_ROLE_SESSION_NAME (optional)
			// o.Policy = ""                                  // env: VOLCENGINE_OIDC_ROLE_POLICY (optional)
			// o.Endpoint = ""                                // env: VOLCENGINE_OIDC_STS_ENDPOINT (optional)
			// o.DurationSeconds = 3600                       // Validity period, default 3600
			// o.MaxRetries = volcengine.Int(3)               // optional: retry attempts; nil defaults to 3, 0 disables
			// o.RetryInterval = 1 * time.Second               // optional: sleep between retries; <= 0 falls back to 1s
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

> Tip: `credentials.NewOIDCCredentialsProviderFromEnv()` builds the provider from environment variables without any arguments.

#### Option 2: Using struct literal

```go
package main

import (
	"fmt"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	p := &credentials.OIDCCredentialsProvider{
		OIDCTokenFilePath: "/path/to/oidc_token_file", // env: VOLCENGINE_OIDC_TOKEN_FILE (required)
		RoleTrn:           "Your Role Trn",            // env: VOLCENGINE_OIDC_ROLE_TRN  (required)
		RoleSessionName:   "",                         // env: VOLCENGINE_OIDC_ROLE_SESSION_NAME (optional)
		Policy:            "",                         // env: VOLCENGINE_OIDC_ROLE_POLICY (optional)
		Endpoint:          "",                         // env: VOLCENGINE_OIDC_STS_ENDPOINT (optional)
		DurationSeconds:   3600,                       // Validity period
		MaxRetries:        volcengine.Int(3),          // optional extra retry attempts; nil defaults to 3, 0 disables retries
		RetryInterval:     1 * time.Second,            // optional sleep between retries; <= 0 falls back to 1s
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

### STS AssumeRoleWithSAML Example

`SAMLCredentialsProvider` exchanges a SAML assertion (returned by your SAML 2.0 IdP) for temporary STS credentials via `AssumeRoleWithSAML`. Credentials are auto-refreshed before expiration.

> ⚠️ **Notes**
>
> 1. **Least Privilege**: Grant only the minimum permissions required.
> 2. **Reasonable Validity Period**: Recommended not to exceed 1 hour.
> 3. The `SAMLAssertion` is the base64-encoded SAML Response returned by your IdP.

#### Option 1: Using WithOptions (recommended)

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	p := credentials.NewSAMLCredentialsProviderWithOptions(
		"trn:iam::1234567890:role/saml-role",      // Role TRN (required)
		"trn:iam::1234567890:saml-provider/MyIdp", // SAML provider TRN (required)
		"BASE64_ENCODED_SAML_RESPONSE_FROM_IDP",   // SAML assertion (required)
		func(o *credentials.SAMLProviderOptions) {
			// o.DurationSeconds = 3600                                 // Validity period, default 3600
			// o.MaxRetries = volcengine.Int(3)                         // optional: retry attempts; nil defaults to 3, 0 disables
			// o.RetryInterval = 1 * time.Second                        // optional: sleep between retries; <= 0 falls back to 1s
		},
	)

	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewCredentials(p))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

#### Option 2: Using convenience constructor

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
		"trn:iam::1234567890:role/saml-role",      // RoleTrn
		"trn:iam::1234567890:saml-provider/MyIdp", // SAMLProviderTrn
		"BASE64_ENCODED_SAML_RESPONSE_FROM_IDP",   // SAMLAssertion
	)
	p.DurationSeconds = 3600
	p.MaxRetries = volcengine.Int(3)  // optional extra retry attempts; nil defaults to 3, 0 disables retries
	p.RetryInterval = 1 * time.Second // optional sleep between retries; <= 0 falls back to 1s

	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewCredentials(p))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

### Environment Variable Credential Provider

`EnvProvider` reads credentials from environment variables. Priority order:

- Access Key: `VOLCENGINE_ACCESS_KEY` > `VOLCSTACK_ACCESS_KEY_ID` > `VOLCSTACK_ACCESS_KEY`
- Secret Key: `VOLCENGINE_SECRET_KEY` > `VOLCSTACK_SECRET_ACCESS_KEY` > `VOLCSTACK_SECRET_KEY`
- Session Token: `VOLCENGINE_SESSION_TOKEN` > `VOLCSTACK_SESSION_TOKEN` (optional)

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
}
```

### CLI Config Credential Provider

`CliProvider` reads credentials from the volcengine-cli config file (`~/.volcengine/config.json`).

- Config path priority: constructor param > `VOLCENGINE_CLI_CONFIG_FILE` > `~/.volcengine/config.json`
- Profile priority: constructor param > `VOLCENGINE_PROFILE` > `VOLCSTACK_PROFILE` > `current` in config > `default`

Supported modes in profile (case-insensitive):

| Mode | Description                                                     |
| --- |-----------------------------------------------------------------|
| `ak` / empty | Static AK/SK from profile                                       |
| `sso` | Reads STS credentials from the CLI sso cache (SDK refreshes access token in-memory, never writes the cache file) |
| `ramrolearn` | STS AssumeRole (delegates to `StsProvider`)                     |
| `oidc` | STS AssumeRoleWithOIDC (delegates to `OIDCCredentialsProvider`) |
| `ecsrole` | ECS IMDS (delegates to `EcsRoleProvider`)                       |
| `console-login` | Reads STS credentials from the CLI console-login cache (SDK refreshes via OAuth `refresh_token` in-memory, never writes the cache file) |

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials/clicreds"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	// Use default config path and profile
	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(clicreds.NewCliCredentials("", ""))

	// Or specify a profile
	// WithCredentials(clicreds.NewCliCredentials("", "prod"))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

#### Runtime Refresh Behavior (sso / console-login)

For `sso` and `console-login` modes the SDK now owns refresh in-memory and
never writes any local file. Key invariants:

- **Read-only on disk**: `config.json`, `~/.volcengine/sso/cache/*.json` and
  `~/.volcengine/login/cache/*.json` are read on bootstrap and once more if
  the authorization service rejects the in-memory refresh token (`invalid_grant`
  fallback). They are never written by the SDK.
- **In-memory refresh**: when the cached `access_token` is past its expiry
  buffer (60 seconds), the SDK exchanges the cached `refresh_token` at the
  OAuth `/token` endpoint and updates its in-memory state. SSO then calls
  the Portal `GetRoleCredentials` API for the STS triple.
- **Invalid-grant fallback** (both sso and console-login): on HTTP 400
  `invalid_grant`, the SDK re-reads the cache file once. If the disk
  `refresh_token` differs from the in-memory one (i.e. `ve login` /
  `ve sso login` rotated it under the SDK), the SDK retries with the disk RT;
  otherwise it reports an actionable error pointing at `ve login` or
  `ve sso login`.
- **Refresh-token expiry**: when the SDK exhausts both the in-memory and
  disk refresh tokens, it raises a clear error instructing the user to run
  `ve login` (console-login) or `ve sso login` (SSO).
- **Concurrency**: a per-process lock serializes refreshes so concurrent
  callers share a single in-flight refresh.

See [`cli-console-login-credential-plan.md`](./cli-console-login-credential-plan.md)
for the full contract.

### ECS Role Credential Provider

`EcsRoleProvider` retrieves temporary credentials from the ECS Instance Metadata Service (IMDSv2).

- Role name priority: constructor param > `VOLCENGINE_ECS_METADATA` env var > auto-detect from IMDS
- Disable switch: `VOLCENGINE_ECS_METADATA_DISABLED=true`
- IMDS endpoint: `http://100.96.0.96` (IMDSv2 with token-based authentication)
- Credentials are automatically refreshed before expiration (5-minute buffer)

> ⚠️ **Notes**
>
> 1. Only works on ECS instances with an IAM role attached.
> 2. Auto-detection queries the IMDS role list and uses the first role found.

#### Option 1: Using WithOptions (recommended, supports custom retry config)

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	p := credentials.NewEcsRoleProviderWithOptions(
		"your-ecs-role-name", // Role name (leave empty to auto-detect from IMDS)
		func(o *credentials.EcsRoleProviderOptions) {
			// o.MaxRetries = volcengine.Int(3)                // optional: retry attempts; nil defaults to 3, 0 disables
			// o.RetryInterval = 1 * time.Second               // optional: sleep between retries; <= 0 falls back to 1s
		},
	)

	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewCredentials(p))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

#### Option 2: Using convenience constructor

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	// Specify role name explicitly
	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewEcsRoleCredentials("your-ecs-role-name"))

	// Or auto-detect role name from IMDS
	// WithCredentials(credentials.NewEcsRoleCredentials(""))

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

### Default Credential Provider (Credential Chain)

When no credentials are explicitly configured, the SDK automatically uses `DefaultCredentialProvider` — a 4-step chain that tries each provider in order until one succeeds:

1. **EnvProvider** — environment variables (`VOLCENGINE_ACCESS_KEY` / `VOLCSTACK_ACCESS_KEY_ID`)
2. **OIDCCredentialsProvider** — OIDC from env vars (`VOLCENGINE_OIDC_TOKEN_FILE`, `VOLCENGINE_OIDC_ROLE_TRN`)
3. **CliProvider** — CLI config file (`~/.volcengine/config.json`)
4. **EcsRoleProvider** — ECS IMDS (instance metadata)

By default, `reuseLastProviderEnabled=true`: after the first successful resolution, subsequent calls reuse the cached provider for performance, falling back to full chain traversal only if the cached provider fails.

**Implicit usage** (no credentials configured — the SDK uses the default chain automatically):

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	// No WithCredentials — SDK automatically uses DefaultCredentialProvider
	config := volcengine.NewConfig().
		WithRegion("cn-beijing")

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
}
```

**Explicit usage** (customize the default chain, e.g., specify an ECS role name):

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
}
```

### Shared credentials file (Deprecated)

> ⚠️ **This mechanism is deprecated and may be removed in a future release.** It is kept only for backward compatibility. New code should use one of the following instead:
> 1. **Environment variables**: `VOLCENGINE_ACCESS_KEY` / `VOLCENGINE_SECRET_KEY` (see [EnvironmentVariables.md](EnvironmentVariables.md))
> 2. **CLI config file**: `~/.volcengine/config.json` (see *CLI Config Credential Provider* above)
> 3. **Default credential chain**: automatic discovery when no credentials are configured (see *Default Credential Provider* above)

The `~/.volcengine/credentials` shared credentials file (INI format, AWS-style) is still supported. Its path and field semantics are unchanged. In the new default credential chain, this file is no longer surfaced as a standalone provider; instead, it is handled by the `sharedConfig` resolution path inside the `session` layer — existing user code requires no changes to keep working.

Note: this file and the CLI credentials file `~/.volcengine/config.json` (JSON format, written by volcengine-cli) are **two independent systems**.

```ini
# ~/.volcengine/credentials
[default]
volcstack_access_key_id = AK_DEFAULT
volcstack_secret_access_key = SK_DEFAULT

[prod]
volcstack_access_key_id = AK_PROD
volcstack_secret_access_key = SK_PROD
```

```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	// Explicitly pick the [prod] profile
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "prod",
	})
	if err != nil {
		panic(err)
	}
}
```

Loading behavior:

- When `Options.Profile` is empty, if the file exists and contains a usable profile, it is loaded automatically (parity with master). If not, the default credential chain is used as a fallback.
- When `Options{Profile: "prod"}` is set explicitly but the `prod` section is missing in the file, `creds.Get()` returns `failed to load profile, prod.` — it does NOT silently fall back to the default chain, preventing accidental use of a different identity.

---

[← Overview](0-Overview.md) | Credentials[(中文)](1-Credentials-zh.md) | [Endpoint →](2-Endpoint.md)
