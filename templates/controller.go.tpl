package controller
{{#with Config}}
import (
	"{{Module}}/infrastruture/web"
	"{{Module}}/model"
	"{{Module}}/service/contract"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)
{{/with}}
{{#with Meta}}
type {{ModelName}}Controller struct {
	svc contract.{{ModelName}}Service
}


// @Summary 获取{{TableComment}}分页列表
// @Description 
// @Tags　{{TableComment}}管理
// @Produce application/json
// @Param object query model.{{ModelName}}QueryReq false "查询参数"
// @Param page.size query integer false "分页大小" default(20)
// @Param page.no query integer false "页码" default(1)
// @Success 200 {object} web.PageResp{records=[]model.{{ModelName}}DetailResp}
// @Failure 400 {object} web.ErrorResp
// @Router {{BasePath}} [get]
func (c *{{ModelName}}Controller) Get(ctx iris.Context,param model.{{ModelName}}QueryReq) mvc.Result {
	rst := web.ValidateRequest(param)
	if rst != nil {
        return rst
	}

	list,err := c.svc.Query(&param); if err != nil {
		return web.WrapError(err)
	}

	return web.WrapResp(list)
}

// @Summary 创建{{TableComment}}记录
// @Description 
// @Tags　{{TableComment}}管理
// @Accept application/json
// @Produce application/json
// @Param body body model.{{ModelName}}InfoReq true "创建参数"
// @Success 200 {object} model.{{ModelName}}DetailResp
// @Failure 400 {object} web.ErrorResp
// @Router {{BasePath}} [post]
func (c *{{ModelName}}Controller) Post(ctx iris.Context,param model.{{ModelName}}InfoReq) mvc.Result {
	invalid := web.ValidateRequest(param); if invalid != nil {
        return invalid
	}

	vo, err := c.svc.Create(&param); if err != nil {
		return web.WrapError(err)
		
	}
	return web.WrapResp(vo)
}


// @Summary 获取单个{{TableComment}}记录
// @Description 
// @Tags　{{TableComment}}管理
// @Produce application/json
// @Param id path int64 true "ID"
// @Success 200 {object} model.{{ModelName}}DetailResp
// @Failure 404 
// @Router {{BasePath}}/{id} [get]
func (c *{{ModelName}}Controller) GetBy(ctx iris.Context,id int64) mvc.Result {
	vo, err := c.svc.Item(id); if err != nil {
		return web.WrapError(err)
	}
	return web.WrapResp(vo)
}



// @Summary 更新{{TableComment}}记录
// @Description 
// @Tags　{{TableComment}}管理
// @Accept application/json
// @Produce application/json
// @Param id path int64 true "ID"
// @Param body body model.{{ModelName}}InfoReq true "修改参数"
// @Success 200 {object} model.{{ModelName}}DetailResp
// @Failure 400 {object} web.ErrorResp
// @Router {{BasePath}}/{id} [put]
func (c *{{ModelName}}Controller) PutBy(ctx iris.Context,id int64,param model.{{ModelName}}InfoReq) mvc.Result {
	invalid := web.ValidateRequest(param); if invalid != nil {
        return invalid
	}

	vo, err := c.svc.Update(id, &param); if err != nil {
		return web.WrapError(err)
		
	}
	return web.WrapResp(vo)
}

// @Summary 修改{{TableComment}}记录
// @Description 
// @Tags　{{TableComment}}管理
// @Accept application/json
// @Produce application/json
// @Param id path int64 true "ID"
// @Param body body model.{{ModelName}}InfoReq true "修改参数"
// @Success 200 {object} model.{{ModelName}}DetailResp
// @Failure 400 {object} web.ErrorResp
// @Router {{BasePath}}/{id} [patch]
func (c *{{ModelName}}Controller) PatchBy(ctx iris.Context,id int64,param model.{{ModelName}}InfoReq) mvc.Result {
	invalid := web.ValidateRequest(param); if invalid != nil {
        return invalid
	}

	vo, err := c.svc.Patch(id, &param); if err != nil {
		return web.WrapError(err)
		
	}
	return web.WrapResp(vo)
}

// @Summary 删除{{TableComment}}记录
// @Description 
// @Tags　{{TableComment}}管理
// @Produce application/json
// @Param id path int64 true "ID"
// @Success 200 {object} model.{{ModelName}}DetailResp
// @Failure 400 {object} web.ErrorResp
// @Router {{BasePath}}/{id} [delete]
func (c *{{ModelName}}Controller) DeleteBy(ctx iris.Context,id int64) mvc.Result {
	err := c.svc.Delete(id); if err != nil {
		return web.WrapError(err)
	}
	return web.WrapSuccess()
}


/*
// 注册自定义路由
//
func (m *{{ModelName}}Controller) BeforeActivation(b mvc.BeforeActivation) {

    // 1-> 方法
    // 2-> 路径
     // 3-> 控制器函数的名称将被解析未一个处理程序 [ handler ]
     // 4-> 任何应该在 MyCustomHandler 之前运行的处理程序[ handlers ]
     //b.Handle("GET", "/something/{id:long}", "MyCustomHandler", anyMiddleware...)
}
*/

// 返回路由根路径
func (c {{ModelName}}Controller) Route() string {
	return "{{BasePath}}"
}

// 构造器
func New{{ModelName}}Controller(svc contract.{{ModelName}}Service) *{{ModelName}}Controller {
	return &{{ModelName}}Controller {
		svc: svc,
	}
}
{{/with}}