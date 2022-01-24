package http

import (
	"github.com/yuyenews/Beerus/application/cloud/cmanager"
	"github.com/yuyenews/Beerus/application/cloud/cparams"
	"github.com/yuyenews/Beerus/application/cloud/croute"
	"github.com/yuyenews/Beerus/application/web"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/commons/util"
	"github.com/yuyenews/Beerus/network/http/commons"
	"github.com/yuyenews/Beerus/network/http/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// StartHttpServer Start an http service
func StartHttpServer(port string) {
	if cparams.ServerName != "firstNode" && cparams.ConnectionUrl != "" {
		croute.CreateCommunicationRoute()
		cmanager.InitLocalCacheRouteMap()
		cmanager.DoCommunication()
	}

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
	req.Params = make(map[string][]string)

	setRoutePath(req)

	// If WebSocket, upgrade the protocol
	if isWebSocket(req) {
		websocket.UpgradeToWebSocket(write, req)
		return
	}

	// Not WebSocket will handle http normally
	error := parsingParam(req)

	if error != nil {
		log.Println("[ERROR]: Parsing parameter exception, " + error.Error())
		res.SendErrorMsg(500, error.Error())
		return
	}

	web.ExecuteRoute(req, res)
}

// parsingParam Parsing json parameters
func parsingParam(request *commons.BeeRequest) error {

	contentType := request.ContentType()

	if strings.ToUpper(request.Request.Method) == "GET" {
		url := request.Request.RequestURI
		paramIndex := strings.LastIndex(url, "?")

		if paramIndex > -1 {
			paramStr := url[(paramIndex + 1):]
			extractionParameters(paramStr, request)
		}

	} else {
		if commons.IsUrlEncode(contentType) {
			body := request.Request.Body
			if body == nil {
				return nil
			}

			resultFrom, err := ioutil.ReadAll(body)
			if err != nil {
				log.Print("Exception for parsing urlEncode parameters", err.Error())
				return err
			}

			extractionParameters(util.BytesToString(resultFrom), request)
		} else if commons.IsJSON(contentType) {
			body := request.Request.Body
			if body == nil {
				return nil
			}

			resultJson, err := ioutil.ReadAll(body)
			if err != nil {
				log.Print("Exception for parsing json parameters", err.Error())
				return err
			}

			request.Json = util.BytesToString(resultJson)
		}
	}

	return nil
}

// extractionParameters Extraction parameters
func extractionParameters(paramStr string, request *commons.BeeRequest) {
	if paramStr == "" {
		return
	}

	paramStr, err := url.QueryUnescape(paramStr)
	if err != nil {
		// Here it just fails to decode, so there is no need to stop, as the user receives the parameters and can decode them themselves.
		log.Println(err.Error())
	}

	paramArray := strings.Split(paramStr, "&")
	for _, param := range paramArray {
		if param == "" {
			continue
		}

		paramKeyAndVal := strings.Split(param, "=")
		if len(paramKeyAndVal) <= 0 {
			continue
		}

		request.AddParam(paramKeyAndVal[0], paramKeyAndVal[1])
	}
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
func isWebSocket(request *commons.BeeRequest) bool {
	upgrade := request.HeaderValue(commons.Upgrade)
	connection := request.HeaderValue(commons.Connection)
	secKey := request.HeaderValue(commons.SecWebsocketKey)

	if upgrade == "" || connection == "" || secKey == "" {
		return false
	}

	if strings.Index(strings.ToUpper(connection), strings.ToUpper(commons.Upgrade)) <= -1 {
		return false
	}

	return true
}
