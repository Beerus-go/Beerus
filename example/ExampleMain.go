package main

import (
	"github/yuyenews/Beerus"
	"github/yuyenews/Beerus/example/interceptor"
	"github/yuyenews/Beerus/example/routes"
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
