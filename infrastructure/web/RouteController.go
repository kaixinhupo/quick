package web

// RouteController 定义用于注册的路由分组地址
type RouteController interface {
	// Route 返回路由分组地址
	Route() string
}
