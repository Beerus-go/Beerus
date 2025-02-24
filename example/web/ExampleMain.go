package main

import (
	"github.com/Beerus-go/Beerus"
	"github.com/Beerus-go/Beerus/example/web/interceptor"
	"github.com/Beerus-go/Beerus/example/web/routes"
)

func main() {

	// Interceptors, routes, etc. Loading of data requires its own calls

	//routes.CreateRoute()
	routes.CreateJsonRoute()

	routes.CreateWebSocketRoute()
	interceptor.CreateInterceptor()

	// Listen the service and listen to port 8080
	beerus.ListenHTTP(8080)
}
