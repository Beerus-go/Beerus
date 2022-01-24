package cmanager

import (
	"errors"
	"github.com/yuyenews/Beerus/application/cloud/cparams"
	"time"
)

// LocalRouteCacheMapToParamModel
// Converting the local interface cache into a ParamModel
func LocalRouteCacheMapToParamModel(from map[string]map[string]*LocalRouteCacheManager) *cparams.ParamModel {
	paramModels := new(cparams.ParamModel)
	paramModels.ParamRoutes = make(map[string][]*cparams.LocalRouteCacheModel)

	for key, val := range from {
		if val == nil {
			continue
		}

		for _, va := range val {
			paramModels.ParamRoutes[key] = va.LocalRouteCacheModelArray
		}
	}

	return paramModels
}

// ParamModelInsertToLocalRouteCacheMap
// Insert the ParamModel into the local interface cache
func ParamModelInsertToLocalRouteCacheMap(paramModel *cparams.ParamModel) error {
	paramRoutes := paramModel.ParamRoutes
	if paramRoutes == nil {
		return errors.New("")
	}

	for key, val := range paramRoutes {
		if LocalRouteCacheMap[key] == nil {
			routeManagerMap := make(map[string]*LocalRouteCacheManager)
			add(routeManagerMap, val)
			LocalRouteCacheMap[key] = routeManagerMap
		} else {
			routeManagerMap := LocalRouteCacheMap[key]
			add(routeManagerMap, val)
		}
	}

	return nil
}

// Perform an insert operation
func add(routeManagerMap map[string]*LocalRouteCacheManager, val []*cparams.LocalRouteCacheModel) {
	for _, v := range val {
		v.CreateTime = time.Now().Unix()

		if routeManagerMap[v.Path] == nil {
			routeCacheManager := new(LocalRouteCacheManager)
			routeCacheManager.AddRoute(v)
			routeManagerMap[v.Path] = routeCacheManager
		} else {
			routeManagerMap[v.Path].AddRoute(v)
		}
	}
}
