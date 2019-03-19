package middleware

import (
	"fmt"
	"github.com/lxzfranky/l2wg"
	"io"
	"net"
	"os"
	"time"
)

var DefaultWriter io.Writer = os.Stdout

func Logger(c *l2wg.Context)  {
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
	clientIp := c.ClientIP()
	method


	fmt.Sprintf()
}
