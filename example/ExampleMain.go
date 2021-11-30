package main

import (
	"github.com/yuyenews/Beerus"
	"github.com/yuyenews/Beerus/example/interceptor"
	"github.com/yuyenews/Beerus/example/routes"
)

func main() {
	routes.CreateRoute()
	interceptor.CreateInterceptor()
	//
	beerus.StartHttp(8080)

	//println(string_util.Match("/api/posts", "/*/post"))
	//string_util.Match("/api/post", "/api/post")
	//string_util.Match("/api/post", "/api/*ost")
}
