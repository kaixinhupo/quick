package controller

type RoleController struct{}

func (m *RoleController) Get() string {
	return "role"
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

func (c RoleController) Route() string {
	return "/role"
}

func NewRoleController() *RoleController {
	return &RoleController{}
}