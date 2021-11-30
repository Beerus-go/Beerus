package http

import (
	"github.com/yuyenews/Beerus/application/web"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/commons/util"
	"github.com/yuyenews/Beerus/network/http/commons"
	"github.com/yuyenews/Beerus/network/http/websocket"
	"io/ioutil"
	"log"
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

	if isWebSocket(request) {
		websocket.Upgrade(write, request)
		return
	}

	var req = new(commons.BeeRequest)
	var res = new(commons.BeeResponse)

	req.Request = request
	res.Response = write

	setRoutePath(req)

	var error = parsingJson(req)

	if error != nil {
		res.SendErrorMsg(500, error.Error())
		return
	}

	web.ExecuteApi(req, res)
}

// parsingJson Parsing json parameters
func parsingJson(request *commons.BeeRequest) error {

	contentType := request.ContentType()

	if strings.ToUpper(request.Request.Method) != "GET" && commons.IsJSON(contentType) {
		var result, error = ioutil.ReadAll(request.Request.Body)
		if error != nil {
			log.Print("Exception for parsing json parameters", error)
			return error
		}

		request.Json = string_util.BytesToString(result)
	}

	return nil
}

// setRoutePath Set the route path to request
func setRoutePath(request *commons.BeeRequest) {
	url := request.Request.RequestURI
	var lastIndex = strings.LastIndex(url, "?")
	if lastIndex > -1 {
		url = url[:lastIndex]
	}

	request.RoutePath = url
}

// isWebSocket Is it a websocket
func isWebSocket(request *http.Request) bool {
	upgrade := request.Header.Get(commons.Upgrade)
	connection := request.Header.Get(commons.Connection)
	secKey := request.Header.Get(commons.SecWebsocketKey)

	if upgrade == "" || connection == "" || secKey == "" {
		return false
	}

	if strings.ToUpper(connection) != strings.ToUpper(commons.Upgrade) {
		return false
	}

	return true
}
