package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenRequestId() string {
	const length = 34

	sb := strings.Builder{}
	sb.Grow(length)
	sb.WriteString(time.Now().Format("20060102150405"))
	sb.WriteString(fmt.Sprintf("%020X", rand.Uint64()))
	return sb.String()
}
