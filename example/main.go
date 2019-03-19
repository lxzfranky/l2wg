package main

import (
	"github.com/lxzfranky/l2wg"
	"github.com/lxzfranky/l2wg/middleware"
)

func BeforeRun(c *l2wg.Context) () {

}

func Run(c *l2wg.Context) () {

}

func AfterRun(c *l2wg.Context) () {

}

func UserMiddleWare(c *l2wg.Context) {

}

func main() {
	// 返回一个路由
	router := l2wg.AppRun()

	// 加一些中间件比如 Auth Log啥的（支持一些框架预定义的和用户自定义的）
	router.AddMiddleWare(middleware.Logger)
	router.AddMiddleWare(UserMiddleWare)

	router.GET("/hello/:name", BeforeRun, Run, AfterRun)

	l2wg.Run(router, "8080")
}
