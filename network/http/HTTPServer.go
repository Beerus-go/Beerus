package http

import (
	"fmt"
	"github/yuyenews/Beerus/commons/util"
	"github/yuyenews/Beerus/network/http/commons"
	"github/yuyenews/Beerus/web"
	"github/yuyenews/Beerus/web/route"
	"io/ioutil"
	"net/http"
	"strings"
)

// StartHttpServer Start an http service
func StartHttpServer(port string) {

	route.ReloadMatchToUrl()

	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:"+port, nil)

}

// handler
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

// parsingJson Parsing json parameters
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
