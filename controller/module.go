package controller

import (
	"github.com/kaixinhupo/quick/infrastruture/web"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

var routerMap = make( map[string]interface{})

func Configure(app *iris.Application) {
	for p,c :=range routerMap {
		mvc.New(app.Party(p)).Handle(c)
	}
}

func RegisterRoute(controller web.RouteController)  {
	routerMap[controller.Route()]= controller
}