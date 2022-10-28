package route

import (
	"github.com/Beerus-go/Beerus/commons/util"
	"github.com/Beerus-go/Beerus/network/http/commons"
	"strings"
)

// interceptorMap
//Store the map of the interceptor
var interceptorMap = make(map[string]func(req *commons.BeeRequest, res *commons.BeeResponse) bool)

// afterReloadingInterceptorMap
//When the service is started, the interceptor pattern and wroute are matched and then stored here to improve the efficiency of getting the interceptor based on the wroute.
var afterReloadingInterceptorMap = make(map[string][]func(req *commons.BeeRequest, res *commons.BeeResponse) bool)

// AddInterceptor
// Add an interceptor
func AddInterceptor(pattern string, before func(req *commons.BeeRequest, res *commons.BeeResponse) bool) {
	interceptorMap[pattern] = before
}

// GetInterceptor
// Get interceptors based on routes
func GetInterceptor(path string) []func(req *commons.BeeRequest, res *commons.BeeResponse) bool {
	return afterReloadingInterceptorMap[path]
}

// ReloadMatchToUrl
// When the service is started, the interceptor and the wroute are matched and then stored according to the wroute, so that it is easy to get the interceptor according to the wroute
func ReloadMatchToUrl() {
	if len(interceptorMap) <= 0 || len(afterReloadingInterceptorMap) > 0 {
		return
	}

	for key, value := range interceptorMap {
		for routePath, _ := range routeMap {
			last := strings.LastIndex(routePath, "/")
			routePath = routePath[:last]

			if util.Match(routePath, key) == false {
				continue
			}

			var interceptorArray = afterReloadingInterceptorMap[routePath]
			if interceptorArray == nil || len(interceptorArray) <= 0 {
				interceptorArray = make([]func(req *commons.BeeRequest, res *commons.BeeResponse) bool, 0)
			}

			interceptorArray = append(interceptorArray, value)
			afterReloadingInterceptorMap[routePath] = interceptorArray
		}
	}
}
