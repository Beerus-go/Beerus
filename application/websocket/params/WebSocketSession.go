package params

import (
	"github.com/yuyenews/Beerus/commons/util"
	"net"
)

// WebSocketSession WebSocket session, mainly used to send messages to the client
type WebSocketSession struct {
	Id         string
	Connection net.Conn
}

// SendString String message
func (ws WebSocketSession) SendString(msg string) {
	ws.Send(string_util.StrToBytes(msg))
}

// Send TODO byte[] message
func (ws WebSocketSession) Send(msg []byte) {

}
