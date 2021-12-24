package web

import (
	"encoding/json"
	"github.com/yuyenews/Beerus/application/web/params"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/commons"
	"github.com/yuyenews/Beerus/commons/util"
	"github.com/yuyenews/Beerus/network/http/commons"
	"log"
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
			log.Println("[This is a friendly reminder, not an error or a warning]: your interceptor returns false, make sure you have called the res.SendXXX function to respond to the front end before returning false")
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

		// Prompt the user if there is no return value, in JSON mode there must be
		if result == nil || len(result) < 1 {
			log.Println("[ERROR]: If you turn on json mode, then all routes must have a return value to give the front-end response through the return value")
			response.SendErrorMsg(500, "If you turn on json mode, then all routes must have a return value to give the front-end response through the return value")
			return
		}

		// If it returns a download, this is a file download route and the front-end response has already been given inside the route, so just return it here
		if result[0].String() == Download {
			return
		}

		// If there is more than one return value, then determine if the second value is of type error
		if len(result) > 1 {
			// If it is not of type error, then the user is prompted that it must be of type error
			if result[1].Type().Name() != "error" {
				log.Println("[ERROR]: In JSON mode, the second return value of the route must be of type error, or only one return value must be set")
				response.SendErrorMsg(500, "In JSON mode, the second return value of the route must be of type error, or only one return value must be set")
				return
			}

			// If it is an error type, then determine whether the value is empty,
			//if not, it means that there is an exception inside the route, and respond directly to the front-end with the error message
			errInterface := result[1].Interface()
			if errInterface != nil {
				err := errInterface.(error)
				if err != nil {
					response.SendErrorMsg(500, err.Error())
					return
				}
			}
		}

		// If there is only one return value, or if the second return value is empty,
		// the request is normal and the first return value is converted into a json response to the front-end.
		resultJson, _ := util.ToJSONString(result[0].Interface())
		response.SendJson(resultJson)
	}
}
