package web

import (
	"github/yuyenews/beerus/network/http/commons"
	"github/yuyenews/beerus/web/route"
	"strings"
)

// ExecuteApi 执行api
func ExecuteApi(request *commons.BeeRequest, response *commons.BeeResponse) {

	var function = route.GetRoute(getRoutePath(request))

	if function == nil {
		response.SendErrorMsg(400, "This route does not exist, please check if the route path and request method are correct")
		return
	}

	function(request, response)
}

// getRoutePath 拼接要请求的路由路径
func getRoutePath(request *commons.BeeRequest) string {
	url := request.Request.RequestURI
	method := request.Request.Method

	var lastIndex = strings.LastIndex(url, "?")
	if lastIndex > -1 {
		url = url[:lastIndex]
	}

	return url + "/" + strings.ToUpper(method)
}
