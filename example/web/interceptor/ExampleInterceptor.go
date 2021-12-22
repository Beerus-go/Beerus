package interceptor

import (
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/network/http/commons"
	"log"
)

func CreateInterceptor() {
	route.AddInterceptor("/example/*", loginInterceptorBefore)
}

func loginInterceptorBefore(req *commons.BeeRequest, res *commons.BeeResponse) bool {
	res.SetHeader("hello", "hello word").SetHeader("hello2", "word2")

	log.Println("exec interceptor")
	return true
}
