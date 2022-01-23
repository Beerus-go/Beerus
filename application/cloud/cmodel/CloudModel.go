package cmodel

// LocalRouteCacheModel
// Save information about each route
type LocalRouteCacheModel struct {
	Method     string
	Path       string
	Url        string
	CreateTime int64
}

type ParamModel struct {
	ParamRoutes map[string][]*LocalRouteCacheModel
}
