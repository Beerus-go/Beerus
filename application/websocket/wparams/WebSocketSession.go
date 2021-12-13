package wparams

import (
	"errors"
	"github.com/yuyenews/Beerus/commons/util"
	"net"
)

// WebSocketSession WebSocket session, mainly used to send messages to the client
type WebSocketSession struct {
	Id         uint64
	Connection net.Conn
}

// SendString String message
func (ws WebSocketSession) SendString(msg string) {
	ws.Send(util.StrToBytes(msg))
}

// Send byte[] message
func (ws WebSocketSession) Send(msg []byte) error {
	startIndex := 2
	var boardCastData []byte

	if len(msg) < 126 {
		boardCastData = make([]byte, 2+len(msg))
		boardCastData[0] = 0x81
		boardCastData[1] = byte(len(msg))
	} else if len(msg) >= 126 && len(msg) < 65535 {
		boardCastData = make([]byte, 4+len(msg))
		bytes := util.IntToBytes(len(msg), 2)
		boardCastData[0] = 0x81
		boardCastData[1] = 126
		boardCastData[2] = bytes[0]
		boardCastData[3] = bytes[1]
		startIndex = 4
	} else {
		return errors.New("maximum supported message length is 65534 bytes")
	}

	index := 0
	for i := startIndex; i < len(boardCastData); i++ {
		boardCastData[i] = msg[index]
		index++
	}

	ws.Connection.Write(boardCastData)

	return nil
}
