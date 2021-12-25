package routes

import (
	"github.com/yuyenews/Beerus/application/web"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/network/http/commons"
	"io/ioutil"
)

func CreateJsonRoute() {

	// Example of file download
	route.GET("/downLoad/file", func(req commons.BeeRequest, res commons.BeeResponse) string {
		file, err := ioutil.ReadFile("/Users/yeyu/Downloads/goland-2021.2.4.dmg")
		if err == nil {

		}
		//req.GetFile()
		res.SendStream("goland.dmg", file)

		return web.Download
	})

	// Example of parameter conversion to struct and parameter checksum
	route.POST("/example/post", func(param DemoParam, req commons.BeeRequest, res commons.BeeResponse) DemoParam {

		println(param.TestStringReception)
		println(param.TestIntReception)
		println(param.TestInt64Reception)
		println(param.TestFloatReception)
		println(param.TestUintReception)
		println(param.TestUint64Reception)
		println(param.TestBoolReception)

		//print(param.TestBeeFileReception.FileHeader.Filename)
		//print(": ")
		//println(param.TestBeeFileReception.FileHeader.Size)

		msg := make(map[string]string)
		msg["msg"] = "success"
		return param
	})
}
