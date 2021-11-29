package main

import string_util "github/yuyenews/Beerus/commons/util"

func main() {
	//controller.CreateRoute()
	//interceptor.CreateInterceptor()
	//
	//beerus.StartHttp(8080)

	println(string_util.Match("/api/posts", "/*/post"))
	//string_util.Match("/api/post", "/api/post")
	//string_util.Match("/api/post", "/api/*ost")
}
