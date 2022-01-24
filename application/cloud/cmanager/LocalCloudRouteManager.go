package cmanager

import (
	"github.com/yuyenews/Beerus/application/cloud/cparams"
	"sync"
	"time"
)

// LocalRouteCacheMap
// A local route cache, including routes for this node and all nodes in the entire microservice, is the map from which nodes fetch URLs, methods and other relevant information when they call another node.
var LocalRouteCacheMap = make(map[string]map[string]*LocalRouteCacheManager)

// LocalRouteCacheManager
// Local route cache management
type LocalRouteCacheManager struct {
	LocalRouteCacheModelArray []*cparams.LocalRouteCacheModel
	index                     int
	mutex                     sync.Mutex
}

// ClearExpiredRoutes
// Clear out-of-date cache
func (localRouteCacheManager *LocalRouteCacheManager) clearExpiredRoutes() {
	if localRouteCacheManager.LocalRouteCacheModelArray == nil || len(localRouteCacheManager.LocalRouteCacheModelArray) < 1 {
		return
	}

	newLocalRouteCacheModelArray := make([]*cparams.LocalRouteCacheModel, 0)

	for _, val := range localRouteCacheManager.LocalRouteCacheModelArray {
		if (time.Now().Unix() - val.CreateTime) < cparams.LocalCacheTimeout {
			newLocalRouteCacheModelArray = append(newLocalRouteCacheModelArray, val)
		}
	}

	localRouteCacheManager.LocalRouteCacheModelArray = newLocalRouteCacheModelArray
}

// AddRoute
// Adding a route cache
func (localRouteCacheManager *LocalRouteCacheManager) AddRoute(localRouteCacheModel *cparams.LocalRouteCacheModel) {
	if localRouteCacheManager.LocalRouteCacheModelArray == nil {
		localRouteCacheManager.LocalRouteCacheModelArray = make([]*cparams.LocalRouteCacheModel, 0)
	}

	for _, val := range localRouteCacheManager.LocalRouteCacheModelArray {
		if val.Url == localRouteCacheModel.Url && val.Method == localRouteCacheModel.Method {
			val.CreateTime = time.Now().Unix()
			return
		}
	}

	localRouteCacheManager.LocalRouteCacheModelArray = append(localRouteCacheManager.LocalRouteCacheModelArray, localRouteCacheModel)
}

// GetRouteInfo
// Get routing information in a polled fashion
func (localRouteCacheManager *LocalRouteCacheManager) GetRouteInfo() *cparams.LocalRouteCacheModel {
	localRouteCacheManager.mutex.Lock()
	defer localRouteCacheManager.mutex.Unlock()

	// Clean up expired routes first
	localRouteCacheManager.clearExpiredRoutes()

	if localRouteCacheManager.LocalRouteCacheModelArray == nil || len(localRouteCacheManager.LocalRouteCacheModelArray) < 1 {
		return nil
	}

	// Then get the required route
	if (len(localRouteCacheManager.LocalRouteCacheModelArray) - 1) > localRouteCacheManager.index {
		localRouteCacheManager.index++
	} else {
		localRouteCacheManager.index = 0
	}

	return localRouteCacheManager.LocalRouteCacheModelArray[localRouteCacheManager.index]
}
