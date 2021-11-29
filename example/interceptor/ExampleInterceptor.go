package interceptor

import (
	"fmt"
	"github/yuyenews/Beerus/network/http/commons"
	"github/yuyenews/Beerus/web/params"
	"github/yuyenews/Beerus/web/route"
)

func CreateInterceptor() {
	route.AddInterceptor("/downLoad/*", loginInterceptorBefore)
}

func loginInterceptorBefore(req *commons.BeeRequest, res *commons.BeeResponse) string {
	fmt.Println("exec interceptor")
	return params.SUCCESS
}
