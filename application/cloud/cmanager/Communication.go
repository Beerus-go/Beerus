package cmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yuyenews/Beerus/application/cloud/cparams"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/commons/util"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// DoCommunication
// Send locally cached interfaces to other nodes and get the list of interfaces from other nodes
func DoCommunication() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("DoCommunication Error:", err)
		}
	}()

	if cparams.ConnectionUrl == "" {
		log.Println("ConnectionUrl Cannot be empty")
		return
	}

	// Get a list of local cache interfaces for propagation to other nodes
	paramModel := LocalRouteCacheMapToParamModel(LocalRouteCacheMap)
	bytesData, _ := json.Marshal(paramModel)

	// Propagation in json format
	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	connectionUrlArray := strings.Split(cparams.ConnectionUrl, ",")
	for _, connectionUrl := range connectionUrlArray {

		connectionUrl = strings.Trim(connectionUrl, " ")

		// Remove the last / avoid conflict with routePath
		if strings.LastIndex(connectionUrl, "/") > -1 {
			connectionUrl = connectionUrl[:len(connectionUrl)-1]
		}

		// Initiate requests and interact with other nodes
		resp, err := util.RequestBody(connectionUrl+cparams.CommunicationRoute,
			"POST",
			header, bytes.NewReader(bytesData))

		if err != nil {
			log.Println("DoCommunication Error:", err)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("DoCommunication Error:", err)
			continue
		}

		if body == nil && len(body) < 1 {
			continue
		}

		// Parsing the received data and caching it to LocalRouteCacheMap
		resultParamModel := new(cparams.ParamModel)
		json.Unmarshal(body, resultParamModel)

		err = ParamModelInsertToLocalRouteCacheMap(resultParamModel)
		if err != nil {
			log.Println("DoCommunication Error:", err)
			continue
		}
	}
}

// InitLocalCacheRouteMap
// Initialize this node's interfaces to the local interface cache,
// executed once when this node is started
func InitLocalCacheRouteMap() {

	if cparams.ServerName == "" {
		log.Println("ServerName Cannot be empty")
		return
	}

	if cparams.ServerUrl == "" {
		log.Println("ServerUrl Cannot be empty")
		return
	}

	routeMap := route.GetRouteMap()

	for key := range routeMap {
		lastIndex := strings.LastIndex(key, "/")
		apiPath := key[:lastIndex]
		if apiPath == cparams.CommunicationRoute {
			continue
		}

		// Remove the last / avoid conflict with routePath
		if strings.LastIndex(cparams.ServerUrl, "/") > -1 {
			cparams.ServerUrl = cparams.ServerUrl[:len(cparams.ServerUrl)-1]
		}

		cacheModel := new(cparams.LocalRouteCacheModel)
		cacheModel.Url = cparams.ServerUrl + apiPath
		cacheModel.Path = apiPath
		cacheModel.CreateTime = time.Now().Unix()
		cacheModel.Method = key[(lastIndex + 1):]

		cacheManagerMap := LocalRouteCacheMap[cparams.ServerName]
		if cacheManagerMap == nil {
			cacheManagerMap = make(map[string]*LocalRouteCacheManager)
			LocalRouteCacheMap[cparams.ServerName] = cacheManagerMap
		}

		cacheManager := LocalRouteCacheMap[cparams.ServerName][apiPath]
		if cacheManager == nil {
			cacheManager = new(LocalRouteCacheManager)
			LocalRouteCacheMap[cparams.ServerName][apiPath] = cacheManager
		}
		cacheManager.AddRoute(cacheModel)
	}
}
