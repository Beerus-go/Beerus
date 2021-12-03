package route

import "github.com/yuyenews/Beerus/application/websocket/params"

const (
	OnConnection = "onConnection"
	OnMessage    = "onMessage"
	OnClose      = "onClose"
)

// Save map of WebSocket routes
var webSocketRouteMap = make(map[string]map[string]func(session *params.WebSocketSession, msg string))

// AddWebSocketRoute Add a route, the first function is triggered when the connection is successful, the second function is triggered when a message is received, and the third function is triggered when the link is broken
func AddWebSocketRoute(routePath string, onConnection func(session *params.WebSocketSession, msg string), onMessage func(session *params.WebSocketSession, msg string), onClose func(session *params.WebSocketSession, msg string)) {
	funcMap := make(map[string]func(session *params.WebSocketSession, msg string))
	funcMap[OnConnection] = onConnection
	funcMap[OnMessage] = onMessage
	funcMap[OnClose] = onClose

	webSocketRouteMap[routePath] = funcMap
}

// GetWebSocketRoute Get the required function based on the route
func GetWebSocketRoute(routePath string, funcName string) func(session *params.WebSocketSession, msg string) {
	funcMap := webSocketRouteMap[routePath]

	if funcMap == nil || len(funcMap) <= 0 {
		return nil
	}
	return funcMap[funcName]
}

// WebSocketRouteExist Does routing exist
func WebSocketRouteExist(routePath string) bool {
	funcMap := webSocketRouteMap[routePath]

	if funcMap == nil || len(funcMap) <= 0 {
		return false
	}

	return true
}
