English | [中文](README.MD)

### Volcengine SDK for Go

#### Overview

1. Volcengine Cloud Go SDK.
2. Minimum Go version: **1.14+**. If using Ark service (`service/arkruntime`), **1.18+** is required.
3. Files in the `service` directory are auto-generated; do not modify them.
4. It is recommended to use `go mod` for dependency management.

#### Endpoint Configuration

If you want to customize the SDK endpoint, use the following code:

```go
config = volcengine.NewConfig().
        WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
        WithRegion(region).WithEndpoint("ecs.cn-beijing-autodriving.volcengineapi.com")
```

Standard endpoint rules:

| Regional Service | Global Service |
|---|---|
| `{service}.{region}.volcengineapi.com` <br> Example: `ecs.cn-beijing-autodriving.volcengineapi.com` | `{service}.volcengineapi.com` <br> Example: `iam.volcengineapi.com` |

Note:

- If the service name contains `_`, it should be converted to `-` in the endpoint. Use lowercase for all characters.

#### Code Example

Examples are available in the `example` directory.

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
    var (
        ak     string
        sk     string
        region string
        config *volcengine.Config
        sess   *session.Session
        client *vpc.VPC
        resp   *vpc.DescribeVpcsOutput
        err    error
    )
    ak = "your ak"
    sk = "your sk"
    region = "cn-beijing"
    config = volcengine.NewConfig().
        WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
        WithRegion(region)

    sess, err = session.NewSession(config)
    if err != nil {
        panic(err)
    }

    client = vpc.New(sess)

    resp, err = client.DescribeVpcs(&vpc.DescribeVpcsInput{
        PageNumber: volcengine.Int64(1),
        PageSize:   volcengine.Int64(10),
    })

    if err != nil {
        panic(err)
    }

    fmt.Println(&resp)
}
```

For more examples, see: [SDK Integration Guide](./SDK_Integration.md)
