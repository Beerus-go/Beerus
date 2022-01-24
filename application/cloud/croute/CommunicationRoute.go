package croute

import (
	"errors"
	"github.com/yuyenews/Beerus/application/cloud/cmanager"
	"github.com/yuyenews/Beerus/application/cloud/cparams"
	"github.com/yuyenews/Beerus/application/web/route"
	"github.com/yuyenews/Beerus/network/http/commons"
)

// CreateCommunicationRoute
// Create communication routes between nodes
func CreateCommunicationRoute() {

	// Communication with other nodes, synchronisation/exchange List of interfaces
	route.POST(cparams.CommunicationRoute, func(paramModel cparams.ParamModel) (*cparams.ParamModel, error) {

		// Save the received interface
		cmanager.ParamModelInsertToLocalRouteCacheMap(&paramModel)

		// Returns a list of interfaces cached by this node
		paramModels := cmanager.LocalRouteCacheMapToParamModel(cmanager.LocalRouteCacheMap)

		return paramModels, nil
	})

	// View a cached list of all interfaces on that node,
	// theoretically all interfaces for the entire microservice (after all services have been started and data synchronised).
	route.GET("beeruscc/searchApiList", func(search cparams.ParamSearchModel, response commons.BeeResponse) (*cparams.ResultDataModel, error) {
		if search.ParamSecretKey != cparams.SecretKey {
			return nil, errors.New("please add the correct secret key")
		}

		paramModels := cmanager.LocalRouteCacheMapToParamModel(cmanager.LocalRouteCacheMap)

		result := new(cparams.ResultDataModel)
		result.ApiList = paramModels.ParamRoutes
		result.NodeCount = len(paramModels.ParamRoutes)
		for _, val := range paramModels.ParamRoutes {
			result.ApiCount = result.ApiCount + len(val)
		}

		return result, nil
	})
}
