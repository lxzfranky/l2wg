package l2wg

import "net/http"

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