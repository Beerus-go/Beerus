package wroute

import "github.com/Beerus-go/Beerus/application/websocket/wparams"

const (
	OnConnection = "onConnection"
	OnMessage    = "onMessage"
	OnClose      = "onClose"
)

// Save map of WebSocket routes
var webSocketRouteMap = make(map[string]map[string]func(session *wparams.WebSocketSession, msg string))

// AddWebSocketRoute Add a wroute, the first function is triggered when the connection is successful, the second function is triggered when a message is received, and the third function is triggered when the link is broken
func AddWebSocketRoute(routePath string, onConnection func(session *wparams.WebSocketSession, msg string), onMessage func(session *wparams.WebSocketSession, msg string), onClose func(session *wparams.WebSocketSession, msg string)) {
	funcMap := make(map[string]func(session *wparams.WebSocketSession, msg string))
	funcMap[OnConnection] = onConnection
	funcMap[OnMessage] = onMessage
	funcMap[OnClose] = onClose

	webSocketRouteMap[routePath] = funcMap
}

// GetWebSocketRoute Get the required function based on the wroute
func GetWebSocketRoute(routePath string, funcName string) func(session *wparams.WebSocketSession, msg string) {
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
