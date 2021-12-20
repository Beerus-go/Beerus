package routes

import (
	"github.com/yuyenews/Beerus/application/web"
	"github.com/yuyenews/Beerus/application/web/params"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/network/http/commons"
	"io/ioutil"
)

func CreateRoute() {

	route.JsonMode = true

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
	route.POST("/example/post", func(param DemoParam, req commons.BeeRequest, res commons.BeeResponse) map[string]string {

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
		return msg
	})

	// Example of parameter conversion to struct + checksum in one step
	route.PUT("/example/put", func(req commons.BeeRequest, res commons.BeeResponse) map[string]string {

		param := DemoParam{}

		// Extraction parameters, Generally used in scenarios where verification is not required or you want to verify manually
		params.ToStruct(req, &param, param)

		// Separate validation of data in struct, this feature can be used independently in any case and is not limited to the routing layer.
		// json mode does not require manual validation, this code can be omitted, here is used to demonstrate the non-json mode, how to validate the parameters
		var result = params.Validation(req, &param, param)
		if result != params.SUCCESS {

			// Non-json mode also can not be returned in this way, you need to call the Send function in the res object to return the result to the front end
			msg := make(map[string]string)
			msg["code"] = "1128"
			msg["msg"] = result
			return msg
		}

		// Extraction of parameters + validation
		// json mode does not require manual validation, this code can be omitted, here is used to demonstrate the non-json mode, how to validate the parameters
		result = params.ToStructAndValidation(req, &param, param)
		if result != params.SUCCESS {

			// Non-json mode also can not be returned in this way, you need to call the Send function in the res object to return the result to the front end
			msg := make(map[string]string)
			msg["code"] = "1128"
			msg["msg"] = result
			return msg
		}

		println(param.TestStringReception)
		println(param.TestIntReception)
		println(param.TestInt64Reception)
		println(param.TestFloatReception)
		println(param.TestUintReception)
		println(param.TestUint64Reception)
		println(param.TestBoolReception)

		msg := make(map[string]string)
		msg["msg"] = "success"
		return msg
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
	TestUint32Reception    uint32  `max:"123" min:"32" msg:"TestUint32Reception The value range must be between 32 - 123"`
	TestUint64Reception    uint64  `max:"123" min:"32" msg:"TestUint64Reception The value range must be between 32 - 123"`
	TestFloatReception     float32 `max:"123" min:"32" msg:"TestFloatReception The value range must be between 32 - 123"`
	TestStringRegReception string  `reg:"^[a-z]+$" msg:"TestStringRegReception Does not meet the regular"`
	TestBoolReception      bool
	TestBeeFileReception   commons.BeeFile

	TestArrayReception []string `notnull:"true" msg:"TestArrayReception Cannot be empty"`
}
