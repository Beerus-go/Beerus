package interceptor

import (
	"fmt"
	"github.com/yuyenews/Beerus/application/web/params"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/network/http/commons"
)

func CreateInterceptor() {
	route.AddInterceptor("*", loginInterceptorBefore)
}

func loginInterceptorBefore(req *commons.BeeRequest, res *commons.BeeResponse) string {
	res.SetHeader("abc", "hahahaha").SetHeader("hello", "word")

	fmt.Println("exec interceptor")
	return params.SUCCESS
}
