package routes

import (
	"github.com/Beerus-go/Beerus/application/websocket/wparams"
	"github.com/Beerus-go/Beerus/application/websocket/wroute"
)

// CreateWebSocketRoute Creating websocket routes
func CreateWebSocketRoute() {
	wroute.AddWebSocketRoute("/ws/test", onConnection, onMessage, onClose)
	wroute.AddWebSocketRoute("/ws/test2", onConnection, onMessage, onClose)
}

// In order to save time, only three functions are used below. In practice, you can configure a set of functions for each wroute

func onConnection(session *wparams.WebSocketSession, msg string) {
	println(msg + "-------------------------------")
	session.SendString("connection success")
}

func onMessage(session *wparams.WebSocketSession, msg string) {
	println(msg + "-------------------------------")
	session.SendString("I got the message.")
}

func onClose(session *wparams.WebSocketSession, msg string) {
	println(msg + "-------------------------------")
}
