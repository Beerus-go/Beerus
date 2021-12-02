package routes

import (
	"github.com/yuyenews/Beerus/application/websocket/params"
	"github.com/yuyenews/Beerus/application/websocket/route"
)

// CreateWebSocketRoute Creating websocket routes
func CreateWebSocketRoute() {
	route.AddWebSocketRoute("/ws/test", onConnection, onMessage, onClose)
	route.AddWebSocketRoute("/ws/test2", onConnection, onMessage, onClose)
}

// In order to save time, only three functions are used below. In practice, you can configure a set of functions for each route

func onConnection(session *params.WebSocketSession, msg string) {
	println(msg + "-------------------------------")
	session.SendString("connection success")
}

func onMessage(session *params.WebSocketSession, msg string) {
	println(msg + "-------------------------------")

	session.SendString("我收到消息了")
}

func onClose(session *params.WebSocketSession, msg string) {
	println(msg + "-------------------------------")
}