package cloud

import (
	"bytes"
	"encoding/json"
	"github.com/yuyenews/Beerus/application/cloud/cmanager"
	"github.com/yuyenews/Beerus/commons/util"
	"io/ioutil"
	"net/http"
)

func RequestJson(serverName string, routePath string, header map[string]string, param interface{}) (string, error) {

	localRouteCacheModel := cmanager.LocalRouteCacheMap[serverName][routePath].GetRouteInfo()

	var resp *http.Response
	var err error

	if localRouteCacheModel.Method == "GET" {

		resp, err = util.Get(localRouteCacheModel.Url, param)
	} else {
		resp, err = request(serverName, localRouteCacheModel.Url, header, param)
	}

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return util.BytesToString(body), nil
}

func RequestStream(serverName string, url string, header map[string]string, param interface{}) ([]byte, error) {
	resp, err := request(serverName, url, header, param)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func RequestUpload(serverName string, url string, files [][]byte) (string, error) {
	return "", nil
}

func request(serverName string, url string, header map[string]string, param interface{}) (*http.Response, error) {
	bytesData, _ := json.Marshal(param)

	header["Content-Type"] = "application/json"

	return util.RequestBody(url, "POST", header, bytes.NewReader(bytesData))
}
