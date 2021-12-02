package main

import (
	"github.com/yuyenews/Beerus"
	"github.com/yuyenews/Beerus/example/web/interceptor"
	"github.com/yuyenews/Beerus/example/web/routes"
)

func main() {

	// Interceptors, routes, etc. Loading of data requires its own calls
	routes.CreateRoute()
	routes.CreateWebSocketRoute()
	interceptor.CreateInterceptor()

	// Start the service and listen to port 8080
	beerus.ListenHTTP(8080)
}
