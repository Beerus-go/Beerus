package routes

import (
	"github/yuyenews/Beerus/application/web/params"
	"github/yuyenews/Beerus/application/web/route"
	"github/yuyenews/Beerus/network/http/commons"
	"io/ioutil"
)

func CreateRoute() {

	// Example of file download
	route.GET("/downLoad/file", func(req *commons.BeeRequest, res *commons.BeeResponse) {
		file, err := ioutil.ReadFile("/Users/yeyu/Downloads/goland-2021.2.4.dmg")
		if err == nil {

		}
		//req.GetFile()
		res.SendStream("goland.dmg", file)
	})

	// Example of parameter conversion to struct and parameter checksum
	route.POST("/example/post", func(req *commons.BeeRequest, res *commons.BeeResponse) {
		param := DemoParam{}

		params.ToStruct(req, &param, param)

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

		var result = params.Verification(req, &param, param)
		if result != params.SUCCESS {
			res.SendErrorMsg(1128, result)
			return
		}

		res.SendJson(`{"msg":"SUCCESS"}`)
	})

	// Example of parameter conversion to struct + checksum in one step
	route.PUT("/example/put", func(req *commons.BeeRequest, res *commons.BeeResponse) {
		param := DemoParam{}
		var result = params.ToStructAndVerification(req, &param, param)

		println(param.TestStringReception)
		println(param.TestIntReception)
		println(param.TestInt64Reception)
		println(param.TestFloatReception)
		println(param.TestUintReception)
		println(param.TestUint64Reception)
		println(param.TestBoolReception)

		if result != params.SUCCESS {
			res.SendErrorMsg(1128, result)
			return
		}

		//println(paramStruct)

		res.SendJson(`{"msg":"SUCCESS"}`)
	})

}

// DemoParam If you have a struct like this, and you want to put all the parameters from the request into this struct
type DemoParam struct {
	// You can customize any field
	// the name of the field must be exactly the same as the name of the requested parameter, and is case-sensitive
	TestStringReception  string  `notnull:"true" msg:"TestStringReception不可以为空" routes:"/example/put"`
	TestIntReception     int     `max:"123" min:"32" msg:"TestIntReception取值范围必须在32 - 123之间" routes:"/example/post"`
	TestInt64Reception   int64   `max:"123" min:"32" msg:"TestInt64Reception取值范围必须在32 - 123之间"`
	TestUintReception    uint    `max:"123" min:"32" msg:"TestUintReception取值范围必须在32 - 123之间"`
	TestUint32Reception  uint32  `max:"123" min:"32" msg:"TTestUint32Reception取值范围必须在32 - 123之间"`
	TestUint64Reception  uint64  `max:"123" min:"32" msg:"TestUint64Reception取值范围必须在32 - 123之间"`
	TestFloatReception   float32 `max:"123" min:"32" msg:"TestFloatReception取值范围必须在32 - 123之间"`
	TestBoolReception    bool
	TestBeeFileReception commons.BeeFile

	TestJsonReception []string
}
