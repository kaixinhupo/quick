package controller

import (
	"github.com/kaixinhupo/quick/infrastruture/web"
	"github.com/kaixinhupo/quick/model"
	"github.com/kaixinhupo/quick/service/contract"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	svc contract.UserService
}


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
func (c *UserController) Get(ctx iris.Context,param model.UserQueryReq) mvc.Result {
	rst := web.ValidateRequest(param)
	if rst != nil {
        return rst
	}

	list,err := c.svc.Query(&param); if err != nil {
		return web.WrapError(err)
	}

	return web.WrapResp(list)
}

// @Summary 创建用户记录
// @Description 
// @Tags　用户管理
// @Accept application/json
// @Produce application/json
// @Param body body model.UserInfoReq true "创建参数"
// @Success 200 {object} model.UserDetailResp
// @Failure 400 {object} web.ErrorResp
// @Router /user [post]
func (c *UserController) Post(ctx iris.Context,param model.UserInfoReq) mvc.Result {
	invalid := web.ValidateRequest(param); if invalid != nil {
        return invalid
	}

	vo, err := c.svc.Create(&param); if err != nil {
		return web.WrapError(err)
		
	}
	return web.WrapResp(vo)
}


// @Summary 获取单个用户记录
// @Description 
// @Tags　用户管理
// @Produce application/json
// @Param id path int64 true "ID"
// @Success 200 {object} model.UserDetailResp
// @Failure 404 
// @Router /user/{id} [get]
func (c *UserController) GetBy(ctx iris.Context,id int64) mvc.Result {
	vo, err := c.svc.Item(id); if err != nil {
		return web.WrapError(err)
	}
	return web.WrapResp(vo)
}



// @Summary 更新用户记录
// @Description 
// @Tags　用户管理
// @Accept application/json
// @Produce application/json
// @Param id path int64 true "ID"
// @Param body body model.UserInfoReq true "修改参数"
// @Success 200 {object} model.UserDetailResp
// @Failure 400 {object} web.ErrorResp
// @Router /user/{id} [put]
func (c *UserController) PutBy(ctx iris.Context,id int64,param model.UserInfoReq) mvc.Result {
	invalid := web.ValidateRequest(param); if invalid != nil {
        return invalid
	}

	vo, err := c.svc.Update(id, &param); if err != nil {
		return web.WrapError(err)
		
	}
	return web.WrapResp(vo)
}

// @Summary 修改用户记录
// @Description 
// @Tags　用户管理
// @Accept application/json
// @Produce application/json
// @Param id path int64 true "ID"
// @Param body body model.UserInfoReq true "修改参数"
// @Success 200 {object} model.UserDetailResp
// @Failure 400 {object} web.ErrorResp
// @Router /user/{id} [patch]
func (c *UserController) PatchBy(ctx iris.Context,id int64,param model.UserInfoReq) mvc.Result {
	invalid := web.ValidateRequest(param); if invalid != nil {
        return invalid
	}

	vo, err := c.svc.Patch(id, &param); if err != nil {
		return web.WrapError(err)
		
	}
	return web.WrapResp(vo)
}

// @Summary 删除用户记录
// @Description 
// @Tags　用户管理
// @Produce application/json
// @Param id path int64 true "ID"
// @Success 200 {object} model.UserDetailResp
// @Failure 400 {object} web.ErrorResp
// @Router /user/{id} [delete]
func (c *UserController) DeleteBy(ctx iris.Context,id int64) mvc.Result {
	err := c.svc.Delete(id); if err != nil {
		return web.WrapError(err)
	}
	return web.WrapSuccess()
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

// 构造器
func NewUserController(userService contract.UserService) *UserController {
	return &UserController {
		svc: userService,
	}
}