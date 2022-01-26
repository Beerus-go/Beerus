package main

import (
	"github.com/yuyenews/Beerus"
	"github.com/yuyenews/Beerus/application/cloud/cparams"
	"github.com/yuyenews/Beerus/example/web/interceptor"
	"github.com/yuyenews/Beerus/example/web/routes"
)

func main() {

	cparams.BeerusCloud = true
	cparams.ServerUrl = "http://qqq.com"

	// Interceptors, routes, etc. Loading of data requires its own calls

	//routes.CreateRoute()
	routes.CreateJsonRoute()

	routes.CreateWebSocketRoute()
	interceptor.CreateInterceptor()

	// Listen the service and listen to port 8080
	beerus.ListenHTTP(8080)
}
