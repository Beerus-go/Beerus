package main

import "github/yuyenews/Beerus/commons/util"

func main() {
	//routes.CreateRoute()
	//interceptor.CreateInterceptor()
	//
	//beerus.StartHttp(8080)

	println(string_util.Match("/api/posts", "/*/post"))
	//string_util.Match("/api/post", "/api/post")
	//string_util.Match("/api/post", "/api/*ost")
}
