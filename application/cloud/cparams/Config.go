package cparams

var (

	// BeerusCloud
	// Whether to enable microservice mode
	BeerusCloud = false

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
	CommunicationRoute = "/beeruscc/04bdcbf5-2b06-4ac2-a283-9343b66ef16c"

	// LocalCacheTimeout
	// Local route cache expiry time
	// Unit: second
	LocalCacheTimeout int64 = 10

	// SecretKey
	// Used to access the built-in interface to view data such as the node's local cache list
	SecretKey = "04jsdskd5-2b06-4ac2-a283-9343b66ef89v"

	// DataSynchronisationInterval
	// Data synchronisation interval, default is cache expiry time -1
	DataSynchronisationInterval = LocalCacheTimeout - 2
)
