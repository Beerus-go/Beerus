package cmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yuyenews/Beerus/application/cloud/cparams"
	"github.com/yuyenews/Beerus/commons/util"
	"io/ioutil"
	"log"
	"strings"
)

// DoCommunication
// Send locally cached interfaces to other nodes and get the list of interfaces from other nodes
func DoCommunication() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("DoCommunication Error:", err)
		}
	}()

	paramModel := LocalRouteCacheMapToParamModel(LocalRouteCacheMap)
	bytesData, _ := json.Marshal(paramModel)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	if strings.LastIndex(cparams.ConnectionUrl, "/") > -1 {
		cparams.ConnectionUrl = cparams.ConnectionUrl[:len(cparams.ConnectionUrl)-1]
	}
	resp, err := util.RequestBody(cparams.ConnectionUrl+cparams.CommunicationRoute,
		"POST",
		header, bytes.NewReader(bytesData))

	if err != nil {
		log.Println("DoCommunication Error:", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("DoCommunication Error:", err)
	}

	if body == nil && len(body) < 1 {
		return
	}

	resultParamModel := new(cparams.ParamModel)
	json.Unmarshal(body, resultParamModel)

	err = ParamModelInsertToLocalRouteCacheMap(resultParamModel)
	if err != nil {
		log.Println("DoCommunication Error:", err)
	}
}

// InitLocalCacheRouteMap
// Initialize this node's interfaces to the local interface cache,
// executed once when this node is started
func InitLocalCacheRouteMap() {
	// TODO
}
