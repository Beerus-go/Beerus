package croute

import (
	"github.com/yuyenews/Beerus/application/cloud"
	"github.com/yuyenews/Beerus/application/cloud/cmodel"
	"github.com/yuyenews/Beerus/application/web/route"
	"time"
)

// CreateCommunicationRoute
// Create communication routes between nodes
func CreateCommunicationRoute() {
	route.POST(cloud.CommunicationRoute, func(paramModel cmodel.ParamModel) (*cmodel.ParamModel, error) {

		// Save the received interface
		for key, val := range paramModel.ParamRoutes {
			for _, va := range val {

				routeMap := cloud.LocalRouteCacheMap[key]
				if routeMap == nil {
					routeMap = make(map[string]*LocalRouteCacheManager)
					routeMap[va.Path] = new(LocalRouteCacheManager)
				}

				va.CreateTime = time.Now().Unix()
				routeMap[va.Path].AddRoute(va)

				cloud.LocalRouteCacheMap[key] = routeMap
			}
		}

		// Returns a list of interfaces cached by this node
		paramModels := new(cmodel.ParamModel)
		paramModels.ParamRoutes = make(map[string][]*cmodel.LocalRouteCacheModel)

		for key, val := range cloud.LocalRouteCacheMap {
			if val == nil {
				continue
			}

			for _, va := range val {
				paramModels.ParamRoutes[key] = va.LocalRouteCacheModelArray
			}
		}
		return paramModels, nil
	})
}