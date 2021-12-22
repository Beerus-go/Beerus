package route

import (
	"github.com/yuyenews/Beerus/commons/util"
	"github.com/yuyenews/Beerus/network/http/commons"
	"strings"
)

// interceptorMap
//Store the map of the interceptor
var interceptorMap = make(map[string]func(req *commons.BeeRequest, res *commons.BeeResponse) bool)

// afterReloadingInterceptorMap
//When the service is started, the interceptor pattern and wroute are matched and then stored here to improve the efficiency of getting the interceptor based on the wroute.
var afterReloadingInterceptorMap = make(map[string]map[int]func(req *commons.BeeRequest, res *commons.BeeResponse) bool)

// AddInterceptor
// Add an interceptor
func AddInterceptor(pattern string, before func(req *commons.BeeRequest, res *commons.BeeResponse) bool) {
	interceptorMap[pattern] = before
}

// GetInterceptor
// Get interceptors based on routes
func GetInterceptor(path string) map[int]func(req *commons.BeeRequest, res *commons.BeeResponse) bool {
	return afterReloadingInterceptorMap[path]
}

// ReloadMatchToUrl
// When the service is started, the interceptor and the wroute are matched and then stored according to the wroute, so that it is easy to get the interceptor according to the wroute
func ReloadMatchToUrl() {
	if len(interceptorMap) <= 0 || len(afterReloadingInterceptorMap) > 0 {
		return
	}

	index := 0

	for key, value := range interceptorMap {
		for routePath, _ := range routeMap {
			last := strings.LastIndex(routePath, "/")
			routePath = routePath[:last]

			if util.Match(routePath, key) == false {
				continue
			}

			var interceptorArray = afterReloadingInterceptorMap[routePath]
			if interceptorArray == nil || len(interceptorArray) <= 0 {
				interceptorArray = make(map[int]func(req *commons.BeeRequest, res *commons.BeeResponse) bool)
			}

			interceptorArray[index] = value
			afterReloadingInterceptorMap[routePath] = interceptorArray

			index++
		}
	}
}
