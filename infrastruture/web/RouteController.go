package web

// 定义用于注册的路由分组地址
type RouteController interface {
	// 返回路由分组地址
	Route() string
}