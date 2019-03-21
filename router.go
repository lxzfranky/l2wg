package l2wg

import (
	"net/http"
)

type Router struct {
	trees      map[string]*node
	Middleware []Handler
}

//给路由增加中间件
func (r *Router) AddMiddleWare(handler ...Handler) {
	r.Middleware = handler
}

//GET方法
func (r *Router) GET(path string, handler ...Handler) {
	r.Handle("GET", path, handler...)
}

//POST方法
func (r *Router) POST(path string, handler ...Handler) () {
	r.Handle("POST", path, handler...)
}

//PUT方法
func (r *Router) PUT(path string, handler ...Handler) () {
	r.Handle("PUT", path, handler...)
}

//DELETE方法
func (r *Router) DELETE(path string, handler ...Handler) {
	r.Handle("DELETE", path, handler...)
}

//HEAD
func (r *Router) HEAD(path string, handler ...Handler) {
	r.Handle("HEAD", path, handler...)
}

//异常判断
func (r *Router) Handle(method, path string, handler ...Handler) {
	if path[0] != '/' {
		panic("Path must begin with '/' in path'" + path + "' ")
	}
	if r.trees == nil {
		r.trees = make(map[string]*node)
	}
	root := r.trees[method]
	if root == nil {
		root = new(node)
		r.trees[method] = root
	}
	root.addRoute(path, handler...)
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request)  {

}