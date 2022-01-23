package croute

import (
	"github.com/yuyenews/Beerus/application/cloud"
	"github.com/yuyenews/Beerus/application/cloud/cmodel"
	"sync"
	"time"
)

// LocalRouteCacheManager
// Local route cache management
type LocalRouteCacheManager struct {
	LocalRouteCacheModelArray []*cmodel.LocalRouteCacheModel
	index                     int
	mutex                     sync.Mutex
}

// ClearExpiredRoutes
// Clear out-of-date cache
func (localRouteCacheManager *LocalRouteCacheManager) clearExpiredRoutes() {
	if localRouteCacheManager.LocalRouteCacheModelArray == nil || len(localRouteCacheManager.LocalRouteCacheModelArray) < 1 {
		return
	}

	newLocalRouteCacheModelArray := make([]*cmodel.LocalRouteCacheModel, 0)

	for _, val := range localRouteCacheManager.LocalRouteCacheModelArray {
		if (time.Now().Unix() - val.CreateTime) < cloud.LocalCacheTimeout {
			newLocalRouteCacheModelArray = append(newLocalRouteCacheModelArray, val)
		}
	}

	localRouteCacheManager.LocalRouteCacheModelArray = newLocalRouteCacheModelArray
}

// AddRoute
// Adding a route cache
func (localRouteCacheManager *LocalRouteCacheManager) AddRoute(localRouteCacheModel *cmodel.LocalRouteCacheModel) {
	if localRouteCacheManager.LocalRouteCacheModelArray == nil {
		localRouteCacheManager.LocalRouteCacheModelArray = make([]*cmodel.LocalRouteCacheModel, 0)
	}
	localRouteCacheManager.LocalRouteCacheModelArray = append(localRouteCacheManager.LocalRouteCacheModelArray, localRouteCacheModel)
}

// GetIndex
// Get routing information in a polled fashion
func (localRouteCacheManager *LocalRouteCacheManager) GetIndex() *cmodel.LocalRouteCacheModel {
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
