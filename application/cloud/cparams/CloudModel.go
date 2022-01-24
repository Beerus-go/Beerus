package cparams

// LocalRouteCacheModel
// Save information about each route
type LocalRouteCacheModel struct {
	Method     string
	Path       string
	Url        string
	CreateTime int64
}

// ParamModel
// For data exchange between nodes
type ParamModel struct {
	ParamRoutes map[string][]*LocalRouteCacheModel
}

// ParamSearchModel
// Parameters for client queries to the local interface cache
type ParamSearchModel struct {
	ParamSecretKey string `field:"paramSecretKey" notnull:true msg:"Please add the correct secret key"`
}

// ResultDataModel
// Data structure returned by the client when querying the local interface cache
type ResultDataModel struct {
	ApiCount  int
	NodeCount int
	ApiList   map[string][]*LocalRouteCacheModel
}
