package routes

import (
	"github.com/yuyenews/Beerus/application/websocket/params"
	"github.com/yuyenews/Beerus/application/websocket/wroute"
)

// CreateWebSocketRoute Creating websocket routes
func CreateWebSocketRoute() {
	wroute.AddWebSocketRoute("/ws/test", onConnection, onMessage, onClose)
	wroute.AddWebSocketRoute("/ws/test2", onConnection, onMessage, onClose)
}

// In order to save time, only three functions are used below. In practice, you can configure a set of functions for each wroute

func onConnection(session *params.WebSocketSession, msg string) {
	println(msg + "-------------------------------")
	session.SendString("connection success")
}

func onMessage(session *params.WebSocketSession, msg string) {
	println(msg + "-------------------------------")
	session.SendString("I got the message.")
}

func onClose(session *params.WebSocketSession, msg string) {
	println(msg + "-------------------------------")
}
