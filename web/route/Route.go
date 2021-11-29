package route

import "github/yuyenews/beerus/network/http/commons"

// 存路由的map
var routeMap = make(map[string]func(req *commons.BeeRequest, res *commons.BeeResponse))

// GetRouteMap 获取所有路由
func GetRouteMap() map[string]func(req *commons.BeeRequest, res *commons.BeeResponse) {
	return routeMap
}

// GetRoute 获取路由
func GetRoute(path string) func(req *commons.BeeRequest, res *commons.BeeResponse) {
	return routeMap[path]
}

// GET 添加GET请求方式的路由
func GET(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "GET", function)
}

// POST 添加GET请求方式的路由
func POST(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "GET", function)
}

// PUT 添加GET请求方式的路由
func PUT(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "PUT", function)
}

// DELETE 添加GET请求方式的路由
func DELETE(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "DELETE", function)
}

// PATCH 添加GET请求方式的路由
func PATCH(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "PATCH", function)
}

// HEAD 添加GET请求方式的路由
func HEAD(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "HEAD", function)
}

// OPTIONS 添加GET请求方式的路由
func OPTIONS(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "OPTIONS", function)
}

// AddRoute 添加路由，如果需要采用其他请求方式，可以用这个
func AddRoute(path string, method string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	routeMap[path+"/"+method] = function
}
