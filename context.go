package l2wg

import (
	"net"
	"net/http"
	"strings"
)

type Handler func(*Context)

type Context struct {
	setState bool
	state    int
	index    int
	Writer   http.ResponseWriter
	Request  *http.Request
	Params   map[string]string
	Data     map[string]interface{}
	handlers []Handler
}

func (n *node) addRoute(path string, handler ...Handler) {

}

// State 获取State
func (c *Context) State() int {
	return c.state
}

// Next 跳过中间件
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) ClientIP() string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}