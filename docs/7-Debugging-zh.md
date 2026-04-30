[← 异常处理](6-ErrorHandling-zh.md) | Debug 机制[(English)](7-Debugging.md) | [概览 →](0-Overview-zh.md)

---

# Debug机制

为便于客户在处理请求时进行问题排查和调试，SDK 支持日志功能，并提供多种日志级别设置。客户可根据实际需求配置日志级别，获取详细的请求与响应信息，以提升排障效率和系统可 observability（可观测性）。

## 开启Debug模式

debug日志默认是关闭的，开启可以用`WithDebug`方法开启

> **默认**
> * `debug` - `False`

**代码示例：**
```go
import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	region := "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		// 开启debug日志
		WithDebug(true).
		WithCredentials(credentials.NewEnvCredentials()) // 环境变量配置：VOLCSTACK_ACCESS_KEY_ID、VOLCSTACK_SECRET_ACCESS_KEY、VOLCSTACK_SESSION_TOKEN
}
```

# 指定日志输出位置
> **默认**
> * `LogWriter` - `os.Stdout`

默认情况下，SDK 会将日志输出到标准输出（stdout）。如果需要将日志输出到文件或其他位置，可以使用 `WithLogWriter` 方法指定日志输出位置。
如果你用的其它日志库，传入其它日志库的Writer即可。

**代码示例：**
```go
package main

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"os"
)

func main() {
	region := "cn-beijing"
	file, _ := os.Create("sdk.log")
	config := volcengine.NewConfig().
		WithRegion(region).
		WithDebug(true).
		WithLogWriter(file).
		WithCredentials(credentials.NewEnvCredentials()) // 环境变量配置：VOLCSTACK_ACCESS_KEY_ID、VOLCSTACK_SECRET_ACCESS_KEY、VOLCSTACK_SESSION_TOKEN
}
```

---

[← 异常处理](6-ErrorHandling-zh.md) | Debug 机制[(English)](7-Debugging.md) | [概览 →](0-Overview-zh.md)
