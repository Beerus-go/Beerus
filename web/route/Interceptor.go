package route

import "github/yuyenews/Beerus/network/http/commons"

var interceptorMap = make(map[string][]func(req *commons.BeeRequest, res *commons.BeeResponse))

var AfterReloadingInterceptorMap = make(map[string][]func(req *commons.BeeRequest, res *commons.BeeResponse))

func AddInterceptor(pattern string, before func(req *commons.BeeRequest, res *commons.BeeResponse), after func(req *commons.BeeRequest, res *commons.BeeResponse)) {
	var interceptor = make([]func(req *commons.BeeRequest, res *commons.BeeResponse), 2)
	interceptor[0] = before
	interceptor[1] = after

	interceptorMap[pattern] = interceptor
}

func GetInterceptor(path string) []func(req *commons.BeeRequest, res *commons.BeeResponse) {
	return AfterReloadingInterceptorMap[path]
}

func ReloadMatchToUrl() {

}
