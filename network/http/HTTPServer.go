package http

import (
	"fmt"
	"github/yuyenews/beerus/commons/util"
	"github/yuyenews/beerus/network/http/commons"
	"github/yuyenews/beerus/web"
	"io/ioutil"
	"net/http"
	"strings"
)

// StartHttpServer 启动一个http服务
func StartHttpServer(port string) {

	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:"+port, nil)

}

// handler 数据解析器
func handler(write http.ResponseWriter, request *http.Request) {

	var req = new(commons.BeeRequest)
	var res = new(commons.BeeResponse)

	req.Request = request
	res.Response = write

	var error = parsingJson(req)

	if error != nil {
		res.SendErrorMsg(500, error.Error())
		return
	}

	web.ExecuteApi(req, res)
}

// parsingJson 解析json参数
func parsingJson(request *commons.BeeRequest) error {

	contentType := request.HeaderValue(commons.ContentType)
	if contentType == "" {
		contentType = request.HeaderValue(commons.ContentType2)
	}

	if strings.ToUpper(request.Request.Method) != "GET" && commons.IsJSON(contentType) {
		var result, error = ioutil.ReadAll(request.Request.Body)
		if error != nil {
			fmt.Print("Exception for parsing json parameters", error)
			return error
		}

		request.Json = string_util.BytesToString(result)
	}

	return nil
}
