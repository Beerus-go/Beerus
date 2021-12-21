package route

// JsonMode true means on, false means off, if on, then the entire framework will enter json mode
var JsonMode = true

// Store the map of the wroute
var routeMap = make(map[string]interface{})

// GetRouteMap Get all routes
func GetRouteMap() map[string]interface{} {
	return routeMap
}

// GetRoute Get routes
func GetRoute(path string) interface{} {
	return routeMap[path]
}

// GET Add a wroute for GET request method
func GET(path string, function interface{}) {
	AddRoute(path, "GET", function)
}

// POST Add a wroute for POST request method
func POST(path string, function interface{}) {
	AddRoute(path, "POST", function)
}

// PUT Add a wroute for PUT request method
func PUT(path string, function interface{}) {
	AddRoute(path, "PUT", function)
}

// DELETE Add a wroute for DELETE request method
func DELETE(path string, function interface{}) {
	AddRoute(path, "DELETE", function)
}

// PATCH Add a wroute for PATCH request method
func PATCH(path string, function interface{}) {
	AddRoute(path, "PATCH", function)
}

// HEAD Add a wroute for HEAD request method
func HEAD(path string, function interface{}) {
	AddRoute(path, "HEAD", function)
}

// OPTIONS Add a wroute for OPTIONS request method
func OPTIONS(path string, function interface{}) {
	AddRoute(path, "OPTIONS", function)
}

// Any add a router for all request method
func Any(path string, function interface{}) {
	GET(path, function)
	POST(path, function)
	PUT(path, function)
	DELETE(path, function)
	PATCH(path, function)
	HEAD(path, function)
	OPTIONS(path, function)
}

// AddRoute Add a wroute, and if you need to use another request method, you can use this
func AddRoute(path string, method string, function interface{}) {
	routeMap[path+"/"+method] = function
}
