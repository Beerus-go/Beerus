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

	//jsonStr := "{\"Name\":\"张三\",\"Age\":19,\"Girlfriend\":[\"王五\",\"赵六\"]}"
	//
	//paramStruct := &JsonStruct{}
	//
	//json.Unmarshal(string_util.StrToBytes(jsonStr), paramStruct)
	//
	//println(paramStruct.Name)
}

type JsonStruct struct {
	Name       string
	Age        int
	Girlfriend []string
}
