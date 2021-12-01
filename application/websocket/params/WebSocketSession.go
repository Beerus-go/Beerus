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
func (ws WebSocketSession) Send(msg []byte) string {
	startIndex := 2
	var boardCastData []byte

	if len(msg) < 126 {
		boardCastData = make([]byte, 2+len(msg))
		boardCastData[0] = 0x81
		boardCastData[1] = byte(len(msg))

	} else if len(msg) >= 126 && len(msg) < 65535 {
		boardCastData = make([]byte, 4+len(msg))
		bytes := string_util.IntToBytes(len(msg), 2)
		boardCastData[0] = 0x81
		boardCastData[1] = 126
		boardCastData[2] = bytes[0]
		boardCastData[3] = bytes[1]
		startIndex = 4
	} else {
		return "最大支持的消息长度为65534个字节"
	}

	//System.arraycopy(message, 0, boardCastData, startIndex, message.length)

	string_util.ArrayCopy(msg, 0, boardCastData, startIndex, len(msg))

	ws.Connection.Write(boardCastData)

	return ""
}
