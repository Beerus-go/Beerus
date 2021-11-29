package route

import "github/yuyenews/Beerus/network/http/commons"

// Store the map of the route
var routeMap = make(map[string]func(req *commons.BeeRequest, res *commons.BeeResponse))

// GetRouteMap Get all routes
func GetRouteMap() map[string]func(req *commons.BeeRequest, res *commons.BeeResponse) {
	return routeMap
}

// GetRoute Get routes
func GetRoute(path string) func(req *commons.BeeRequest, res *commons.BeeResponse) {
	return routeMap[path]
}

// GET Add a route for GET request method
func GET(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "GET", function)
}

// POST Add a route for POST request method
func POST(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "GET", function)
}

// PUT Add a route for PUT request method
func PUT(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "PUT", function)
}

// DELETE Add a route for DELETE request method
func DELETE(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "DELETE", function)
}

// PATCH Add a route for PATCH request method
func PATCH(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "PATCH", function)
}

// HEAD Add a route for HEAD request method
func HEAD(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "HEAD", function)
}

// OPTIONS Add a route for OPTIONS request method
func OPTIONS(path string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	AddRoute(path, "OPTIONS", function)
}

// AddRoute Add a route, and if you need to use another request method, you can use this
func AddRoute(path string, method string, function func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	routeMap[path+"/"+method] = function
}