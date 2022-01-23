package constant

import "github.com/yuyenews/Beerus/application/cloud/croute"

var (
	// ServerName
	// Service name, if not set it means it is the first node to start
	// Which must be the same for multiple nodes of the same service, must be unique between services.
	ServerName = "firstNode"

	// ServerUrl
	// The url of the current node, the default is the intranet IP:PORT of the server where the node is located
	ServerUrl = ""

	// ConnectionUrl
	// The url of the node to connect to, if it is the first node to be started, can be set without
	ConnectionUrl = ""

	// CommunicationRoute
	// Routing of the inter-node communication interface,
	// used to synchronize the routing list of each service to each node, enabling services to call each other
	CommunicationRoute = "beeruscc/04bdcbf5-2b06-4ac2-a283-9343b66ef16c"

	// LocalCacheTimeout
	// Local route cache expiry time
	// Unit: second
	LocalCacheTimeout int64 = 10

	// DataSynchronisationInterval
	// Data synchronisation interval, default is cache expiry time -1
	DataSynchronisationInterval = LocalCacheTimeout - 2

	// LocalRouteCacheMap
	// A local route cache, including routes for this node and all nodes in the entire microservice, is the map from which nodes fetch URLs, methods and other relevant information when they call another node.
	LocalRouteCacheMap = make(map[string]map[string]*croute.LocalRouteCacheManager)
)
