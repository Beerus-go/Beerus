package web

import (
	"github/yuyenews/Beerus/network/http/commons"
	"github/yuyenews/Beerus/web/params"
	"github/yuyenews/Beerus/web/route"
	"strings"
)

// ExecuteApi Execute the interceptor and the corresponding interface of the route
func ExecuteApi(request *commons.BeeRequest, response *commons.BeeResponse) {

	method := request.Request.Method
	routePath := getRoutePath(request)
	function := route.GetRoute(routePath + "/" + strings.ToUpper(method))

	if function == nil {
		response.SendErrorMsg(400, "This route does not exist, please check if the route path and request method are correct")
		return
	}

	// exec interceptors
	var interceptors = route.GetInterceptor(routePath)
	for _, value := range interceptors {
		var result = value(request, response)
		if result != params.SUCCESS {
			response.SendErrorMsg(500, result)
			return
		}
	}

	// Execute the function on the route
	function(request, response)
}

// getRoutePath Get the route path to request
func getRoutePath(request *commons.BeeRequest) string {
	url := request.Request.RequestURI
	var lastIndex = strings.LastIndex(url, "?")
	if lastIndex > -1 {
		url = url[:lastIndex]
	}

	return url
}
