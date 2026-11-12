package rdsmysql

import (
	"fmt"
	"os"
	"sync"
)

// deprecationNoticeOnce 保证下线过渡提示在同一进程内最多打印一次，
// 避免重复创建客户端时刷屏。
var deprecationNoticeOnce sync.Once

// deprecationNotice 是客户端被构造时输出到 stderr 的过渡期提示。
// 仅提示，不改变请求版本、参数、签名和返回值。
const deprecationNotice = `Warning: RDS MySQL API 2018-01-01 is deprecated and will be removed in a future release.
Please migrate to service/rdsmysqlv2 with API version 2022-01-01.
The current request continues unchanged using API 2018-01-01.
`

// warnDeprecated 在旧版客户端被构造时输出一次过渡提示。
func warnDeprecated() {
	deprecationNoticeOnce.Do(func() {
		fmt.Fprint(os.Stderr, deprecationNotice)
	})
}
