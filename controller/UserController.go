package controller

import (
	"log"

	"github.com/kaixinhupo/quick/infrastruture/web"
	"github.com/kaixinhupo/quick/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct{}


// @Summary 获取用户分页列表
// @Description 
// @Tags　用户管理
// @Produce application/json
// @Param object query model.UserQueryReq false "查询参数"
// @Param page.size query integer false "分页大小" default(20)
// @Param page.no query integer false "页码" default(1)
// @Success 200 {object} web.PageResp{records=[]model.UserDetailResp}
// @Failure 400 {object} web.ErrorResp
// @Router /user [get]
func (m *UserController) Get(ctx iris.Context,param model.UserQueryReq) mvc.Result {
	log.Println("GET /user param:",param)

	rst := web.ValidateRequest(param)
	if rst != nil {
        return rst
	}

	list := make([]int64,16)  

	return web.WrapPage(list,10,1)
}


// @Summary 获取单个用户记录
// @Description 
// @Tags　用户管理
// @Produce application/json
// @Param id path int64 true "ID"
// @Success 200 {object} model.UserDetailResp
// @Failure 404 
// @Router /user/{id} [get]
func (m *UserController) GetBy(id int64) model.UserDetailResp {
	return model.UserDetailResp{
		Id: 1,
	}
}

/*
// 注册自定义路由
//
func (m *UserController) BeforeActivation(b mvc.BeforeActivation) {

    // 1-> 方法
    // 2-> 路径
     // 3-> 控制器函数的名称将被解析未一个处理程序 [ handler ]
     // 4-> 任何应该在 MyCustomHandler 之前运行的处理程序[ handlers ]
     //b.Handle("GET", "/something/{id:long}", "MyCustomHandler", anyMiddleware...)
}
*/

// 返回路由根路径
func (c UserController) Route() string {
	return "/user"
}

func init() {
	RegisterRoute(new(UserController))
}