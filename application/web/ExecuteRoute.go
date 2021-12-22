package web

import (
	"encoding/json"
	"github.com/yuyenews/Beerus/application/web/params"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/commons"
	"github.com/yuyenews/Beerus/commons/util"
	"github.com/yuyenews/Beerus/network/http/commons"
	"reflect"
	"strings"
)

const Download = "download-b7c39bbf-d22f-42f6-ad84-bd472dbc9e9a"

// ExecuteRoute Execute the interceptor and the corresponding interface of the route
func ExecuteRoute(request *commons.BeeRequest, response *commons.BeeResponse) {

	method := request.Request.Method
	routePath := request.RoutePath
	routeFunction := route.GetRoute(routePath + "/" + strings.ToUpper(method))

	if routeFunction == nil {
		response.SendErrorMsg(400, "This route does not exist, please check if the route path and request method are correct")
		return
	}

	// exec interceptors
	var interceptors = route.GetInterceptor(routePath)
	for _, inter := range interceptors {
		var result = inter(request, response)
		if result == false {
			// TODO 这个版本发行以后，一定要在文档里标注一下，拦截器如果返回false，那么必须 在返回之前，在拦截器里调用res.SendXXX 函数给前端一个响应，不然本次请求就卡死了
			return
		}
	}

	// Execute the routeFunction on the route
	executeFunction(request, response, routeFunction)
}

// Execute the routing function
func executeFunction(request *commons.BeeRequest, response *commons.BeeResponse, routeFunction interface{}) {
	method := reflect.ValueOf(routeFunction)
	paramNum := method.Type().NumIn()

	paramArray := make([]reflect.Value, 0)

	for i := 0; i < paramNum; i++ {
		param := method.Type().In(i)
		paramObj := reflect.New(param)
		paramElem := paramObj.Elem()

		if strings.ToLower(param.Kind().String()) != data_type.Struct {
			paramArray = append(paramArray, paramElem)
			continue
		}

		if strings.ToLower(param.Name()) == data_type.BeeRequest {
			paramArray = append(paramArray, reflect.ValueOf(*request))
			continue
		}

		if strings.ToLower(param.Name()) == data_type.BeeResponse {
			paramArray = append(paramArray, reflect.ValueOf(*response))
			continue
		}

		// Assigning values to the fields inside the parameters
		if commons.IsJSON(request.ContentType()) {
			if request.Json != "" {
				json.Unmarshal(util.StrToBytes(request.Json), paramObj.Interface())
				paramElem = paramObj.Elem()
			}

			paramArray = append(paramArray, paramElem)

		} else {
			for j := 0; j < param.NumField(); j++ {
				params.SetValue(param, paramElem, *request, j)
			}

			paramArray = append(paramArray, paramElem)
		}

		// If json mode is turned on, then automated parameter validation will be performed and the response message in json format will be given to the front-end based on the validation result.
		if route.JsonMode {
			result := params.ValidationReflect(*request, paramElem, param)
			if result != params.SUCCESS {
				response.SendErrorMsg(1128, result)
				return
			}
		}
	}

	// If json mode is turned off
	if route.JsonMode == false {
		method.Call(paramArray)
		return
	}

	// If json mode is turned on, then the return value of the function is converted to json and used as the response data
	if route.JsonMode {
		result := method.Call(paramArray)

		if result == nil || len(result) < 1 {
			response.SendErrorMsg(500, "If you turn on json mode, then all routes must have a return value to give the front-end response through the return value")
			return
		}
		if result[0].String() == Download {
			return
		}

		resultJson, _ := util.ToJSONString(result[0].Interface())
		response.SendJson(resultJson)
	}
}
