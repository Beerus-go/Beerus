package websocket

import (
	"github.com/yuyenews/Beerus/application/websocket/params"
	"github.com/yuyenews/Beerus/application/websocket/route"
	"github.com/yuyenews/Beerus/commons/util"
	"net"
)

// SessionMap Save each linked session
var sessionMap = make(map[string]*params.WebSocketSession)

// ExecuteConnection Triggers the onConnection function inside the route
func ExecuteConnection(routePath string, conn net.Conn) {
	session := new(params.WebSocketSession)
	session.Connection = conn
	session.Id = util.UniqueId()
	sessionMap[routePath] = session

	route.GetWebSocketRoute(routePath, route.OnConnection)(session, "The client is already connected")
}

// ExecuteMessage Triggers the onMessage function inside the route
func ExecuteMessage(routePath string, message string) {
	route.GetWebSocketRoute(routePath, route.OnMessage)(sessionMap[routePath], message)
}

// ExecuteClose Triggers the onClose function inside the route
func ExecuteClose(routePath string) {
	if sessionMap[routePath] == nil {
		return
	}

	delete(sessionMap, routePath)
	route.GetWebSocketRoute(routePath, route.OnClose)(sessionMap[routePath], "The client is disconnected")
}
