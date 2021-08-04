package controller

import (
	"github.com/kaixinhupo/quick/infrastructure/web"
	"github.com/kaixinhupo/quick/model"
	"github.com/kaixinhupo/quick/service/contract"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type RoleController struct {
	svc contract.RoleService
}

// Get 获取角色分页列表
// @Summary 获取角色分页列表
// @Description
// @Tags　角色管理
// @Produce application/json
// @Param object query model.RoleQueryReq false "查询参数"
// @Param page.size query integer false "分页大小" default(20)
// @Param page.no query integer false "页码" default(1)
// @Success 200 {object} web.PageResp{records=[]model.RoleDetailResp}
// @Failure 400 {object} web.ErrorResp
// @Router  [get]
func (c *RoleController) Get(ctx iris.Context, param model.RoleQueryReq) mvc.Result {
	rst := web.ValidateRequest(param)
	if rst != nil {
		return rst
	}

	list, err := c.svc.Query(&param)
	if err != nil {
		return web.WrapError(err)
	}

	return web.WrapResp(list)
}

// Post 创建角色记录
// @Summary 创建角色记录
// @Description
// @Tags　角色管理
// @Accept application/json
// @Produce application/json
// @Param body body model.RoleInfoReq true "创建参数"
// @Success 200 {object} model.RoleDetailResp
// @Failure 400 {object} web.ErrorResp
// @Router  [post]
func (c *RoleController) Post(ctx iris.Context, param model.RoleInfoReq) mvc.Result {
	invalid := web.ValidateRequest(param)
	if invalid != nil {
		return invalid
	}

	vo, err := c.svc.Create(&param)
	if err != nil {
		return web.WrapError(err)

	}
	return web.WrapResp(vo)
}

// GetBy 获取单个角色记录
// @Summary 获取单个角色记录
// @Description
// @Tags　角色管理
// @Produce application/json
// @Param id path int64 true "ID"
// @Success 200 {object} model.RoleDetailResp
// @Failure 404
// @Router /{id} [get]
func (c *RoleController) GetBy(ctx iris.Context, id int64) mvc.Result {
	vo, err := c.svc.Item(id)
	if err != nil {
		return web.WrapError(err)
	}
	return web.WrapResp(vo)
}

// PutBy 更新角色记录
// @Summary 更新角色记录
// @Description
// @Tags　角色管理
// @Accept application/json
// @Produce application/json
// @Param id path int64 true "ID"
// @Param body body model.RoleInfoReq true "修改参数"
// @Success 200 {object} model.RoleDetailResp
// @Failure 400 {object} web.ErrorResp
// @Router /{id} [put]
func (c *RoleController) PutBy(ctx iris.Context, id int64, param model.RoleInfoReq) mvc.Result {
	invalid := web.ValidateRequest(param)
	if invalid != nil {
		return invalid
	}

	vo, err := c.svc.Update(id, &param)
	if err != nil {
		return web.WrapError(err)

	}
	return web.WrapResp(vo)
}

// PatchBy 修改角色记录
// @Summary 修改角色记录
// @Description
// @Tags　角色管理
// @Accept application/json
// @Produce application/json
// @Param id path int64 true "ID"
// @Param body body model.RoleInfoReq true "修改参数"
// @Success 200 {object} model.RoleDetailResp
// @Failure 400 {object} web.ErrorResp
// @Router /{id} [patch]
func (c *RoleController) PatchBy(ctx iris.Context, id int64, param model.RoleInfoReq) mvc.Result {
	invalid := web.ValidateRequest(param)
	if invalid != nil {
		return invalid
	}

	vo, err := c.svc.Patch(id, &param)
	if err != nil {
		return web.WrapError(err)

	}
	return web.WrapResp(vo)
}

// DeleteBy 删除角色记录
// @Summary 删除角色记录
// @Description
// @Tags　角色管理
// @Produce application/json
// @Param id path int64 true "ID"
// @Success 200 {object} model.RoleDetailResp
// @Failure 400 {object} web.ErrorResp
// @Router /{id} [delete]
func (c *RoleController) DeleteBy(ctx iris.Context, id int64) mvc.Result {
	err := c.svc.Delete(id)
	if err != nil {
		return web.WrapError(err)
	}
	return web.WrapSuccess()
}

/*
// 注册自定义路由
//
func (m *RoleController) BeforeActivation(b mvc.BeforeActivation) {

    // 1-> 方法
    // 2-> 路径
     // 3-> 控制器函数的名称将被解析未一个处理程序 [ handler ]
     // 4-> 任何应该在 MyCustomHandler 之前运行的处理程序[ handlers ]
     //b.Handle("GET", "/something/{id:long}", "MyCustomHandler", anyMiddleware...)
}
*/

// Route 返回路由根路径
func (c RoleController) Route() string {
	return "/role"
}

// NewRoleController 构造器
func NewRoleController(svc contract.RoleService) *RoleController {
	return &RoleController{
		svc: svc,
	}
}
