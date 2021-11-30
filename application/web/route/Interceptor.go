package route

import (
	"github.com/yuyenews/Beerus/commons/util"
	"github.com/yuyenews/Beerus/network/http/commons"
	"strings"
)

// interceptorMap
//Store the map of the interceptor
var interceptorMap = make(map[string]func(req *commons.BeeRequest, res *commons.BeeResponse) string)

// afterReloadingInterceptorMap
//When the service is started, the interceptor pattern and route are matched and then stored here to improve the efficiency of getting the interceptor based on the route.
var afterReloadingInterceptorMap = make(map[string]map[int]func(req *commons.BeeRequest, res *commons.BeeResponse) string)

// AddInterceptor
// Add an interceptor
func AddInterceptor(pattern string, before func(req *commons.BeeRequest, res *commons.BeeResponse) string) {
	interceptorMap[pattern] = before
}

// GetInterceptor
// Get interceptors based on routes
func GetInterceptor(path string) map[int]func(req *commons.BeeRequest, res *commons.BeeResponse) string {
	return afterReloadingInterceptorMap[path]
}

// ReloadMatchToUrl
// When the service is started, the interceptor and the route are matched and then stored according to the route, so that it is easy to get the interceptor according to the route
func ReloadMatchToUrl() {
	if len(interceptorMap) <= 0 || len(afterReloadingInterceptorMap) > 0 {
		return
	}

	index := 0

	for key, value := range interceptorMap {
		for routePath, _ := range routeMap {
			last := strings.LastIndex(routePath, "/")
			routePath = routePath[:last]

			if string_util.Match(routePath, key) == false {
				continue
			}

			var interceptorArray = afterReloadingInterceptorMap[routePath]
			if interceptorArray == nil || len(interceptorArray) <= 0 {
				interceptorArray = make(map[int]func(req *commons.BeeRequest, res *commons.BeeResponse) string)
			}

			interceptorArray[index] = value
			afterReloadingInterceptorMap[routePath] = interceptorArray

			index++
		}
	}
}
