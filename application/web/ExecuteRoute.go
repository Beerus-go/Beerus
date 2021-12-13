package web

import (
	"github.com/yuyenews/Beerus/application/web/params"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/network/http/commons"
	"strings"
)

// ExecuteRoute Execute the interceptor and the corresponding interface of the wroute
func ExecuteRoute(request *commons.BeeRequest, response *commons.BeeResponse) {

	method := request.Request.Method
	routePath := request.RoutePath
	routeFunction := route.GetRoute(routePath + "/" + strings.ToUpper(method))

	if routeFunction == nil {
		response.SendErrorMsg(400, "This wroute does not exist, please check if the wroute path and request method are correct")
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

	// Execute the routeFunction on the wroute
	routeFunction(request, response)
}
