package main

import (
	"log"

	"github.com/kaixinhupo/quick/infrastructure/web"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// 控制器列表
var routers = [...]web.RouteController{
	InitUserController(),
	InitGenController(),
	InitRoleController(),
}

// ConfigureRouter 注册路由
func ConfigureRouter(app *iris.Application) {
	for _, c := range routers {
		p := c.Route()
		mvc.New(app.Party(p)).Handle(c)
	}
	r := app.GetRoutes()
	for _, v := range r {
		log.Println(v.Method, v.Path)
	}
}
