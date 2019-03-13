package main

import "github.com/lxzfranky/l2wg"

func BeforeRun (c *l2wg.Context)()  {

}

func Run (c *l2wg.Context)()  {

}

func AfterRun (c *l2wg.Context)()  {

}

func main()  {
	// 返回一个路由
	router := l2wg.AppRun()
	router.GET("/hello/:name", BeforeRun, Run, AfterRun)

}
