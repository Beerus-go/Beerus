package routes

import (
	"github.com/yuyenews/Beerus/application/web/params"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/network/http/commons"
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

		// Extraction parameters, Generally used in scenarios where verification is not required or you want to verify manually
		params.ToStruct(req, &param, param)

		// Separate validation of data in struct, this feature can be used independently in any case and is not limited to the routing layer.
		var result = params.Validation(req, &param, param)
		if result != params.SUCCESS {
			res.SendErrorMsg(1128, result)
			return
		}

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

		res.SendJson(`{"msg":"SUCCESS"}`)
	})

	// Example of parameter conversion to struct + checksum in one step
	route.PUT("/example/put", func(req *commons.BeeRequest, res *commons.BeeResponse) {
		param := DemoParam{}

		// Extraction of parameters + validation
		var result = params.ToStructAndValidation(req, &param, param)

		if result != params.SUCCESS {
			res.SendErrorMsg(1128, result)
			return
		}

		println(param.TestStringReception)
		println(param.TestIntReception)
		println(param.TestInt64Reception)
		println(param.TestFloatReception)
		println(param.TestUintReception)
		println(param.TestUint64Reception)
		println(param.TestBoolReception)

		res.SendJson(`{"msg":"SUCCESS"}`)
	})

}

// DemoParam If you have a struct like this, and you want to put all the parameters from the request into this struct
type DemoParam struct {
	// You can customize any field
	// the name of the field must be exactly the same as the name of the requested parameter, and is case-sensitive
	TestStringReception    string  `notnull:"true" msg:"TestStringReception Cannot be empty" routes:"/example/put"`
	TestIntReception       int     `max:"123" min:"32" msg:"TestIntReception The value range must be between 32 - 123" routes:"/example/post"`
	TestInt64Reception     int64   `max:"123" min:"32" msg:"TestInt64Reception The value range must be between 32 - 123"`
	TestUintReception      uint    `max:"123" min:"32" msg:"TestUintReception The value range must be between 32 - 123"`
	TestUint32Reception    uint32  `max:"123" min:"32" msg:"TTestUint32Reception The value range must be between 32 - 123"`
	TestUint64Reception    uint64  `max:"123" min:"32" msg:"TestUint64Reception The value range must be between 32 - 123"`
	TestFloatReception     float32 `max:"123" min:"32" msg:"TestFloatReception The value range must be between 32 - 123"`
	TestStringRegReception string  `reg:"^[a-z]+$" msg:"TestStringRegReception Does not meet the regular"`
	TestBoolReception      bool
	TestBeeFileReception   commons.BeeFile

	TestJsonReception []string `notnull:"true" msg:"TestJsonReception Cannot be empty"`
}
