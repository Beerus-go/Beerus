package interceptor

import (
	"fmt"
	"github/yuyenews/Beerus/application/web/params"
	"github/yuyenews/Beerus/application/web/route"
	"github/yuyenews/Beerus/network/http/commons"
)

func CreateInterceptor() {
	route.AddInterceptor("*", loginInterceptorBefore)
}

func loginInterceptorBefore(req *commons.BeeRequest, res *commons.BeeResponse) string {
	res.SetHeader("abc", "hahahaha").SetHeader("hello", "word")

	fmt.Println("exec interceptor")
	return params.SUCCESS
}
