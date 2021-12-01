package main

import (
	"github.com/yuyenews/Beerus"
	"github.com/yuyenews/Beerus/example/web/interceptor"
	"github.com/yuyenews/Beerus/example/web/routes"
)

func main() {
	routes.CreateRoute()
	routes.CreateWebSocketRoute()
	interceptor.CreateInterceptor()

	beerus.StartHttp(8080)

	//println(string_util.Match("/api/posts", "/*/post"))
	//string_util.Match("/api/post", "/api/post")
	//string_util.Match("/api/post", "/api/*ost")
}
