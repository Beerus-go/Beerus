package controller

import (
	"github/yuyenews/Beerus/network/http/commons"
	"github/yuyenews/Beerus/web/params"
	"github/yuyenews/Beerus/web/route"
	"io/ioutil"
)

func CreateRoute() {

	// Example of file download
	route.GET("/downLoad/file", func(req *commons.BeeRequest, res *commons.BeeResponse) {
		file, err := ioutil.ReadFile("/Users/yeyu/Downloads/goland-2021.2.4.dmg")
		if err == nil {

		}
		res.SendStream("goland.dmg", file)
	})

	// Example of parameter conversion to struct and parameter checksum
	route.POST("/example/post", func(req *commons.BeeRequest, res *commons.BeeResponse) {

		var paramStruct = params.ToStruct(req, DemoParam{})
		println(paramStruct)

		var result = params.Verification(paramStruct)
		if result != params.SUCCESS {
			res.SendErrorMsg(1128, result)
			return
		}

		res.SendJson("{\"msg\":\"SUCCESS\"}")
	})

	// Example of parameter conversion to struct + checksum in one step
	route.PUT("/example/put", func(req *commons.BeeRequest, res *commons.BeeResponse) {
		var paramStruct, result = params.ToStructAndVerification(req, DemoParam{})

		if result != params.SUCCESS {
			res.SendErrorMsg(1128, result)
			return
		}

		println(paramStruct)

		res.SendJson("{\"msg\":\"SUCCESS\"}")
	})

}

// DemoParam If you have a struct like this, and you want to put all the parameters from the request into this struct
type DemoParam struct {
	// You can customize any field
	// the name of the field must be exactly the same as the name of the requested parameter, and is case-sensitive
}
