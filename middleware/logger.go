package middleware

import (
	"fmt"
	"github.com/lxzfranky/l2wg"
	"io"
	"os"
	"time"
)

var DefaultWriter io.Writer = os.Stdout

func Logger(c *l2wg.Context) {
	// 按照一定的格式打印下输入的内容
	out := DefaultWriter

	// 开始时间
	start := time.Now()
	path := c.Request.URL.Path
	c.Next()
	// 结束时间
	end := time.Now()
	// 耗费时间
	last := end.Sub(start)
	// 客户端IP
	clientIp := c.ClientIP()
	// 请求方法
	method := c.Request.Method
	// 状态码
	state := c.State()

	fmt.Fprintf(out, "[GIN] %v | %13v | %d |%s  %s %s\n",
		end.Format("2006/01/02 - 15:04:05"),
		last,
		state,
		clientIp,
		method,
		path,
	)
}
