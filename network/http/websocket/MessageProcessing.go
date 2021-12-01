package websocket

import (
	"bytes"
	"github.com/yuyenews/Beerus/application/websocket"
	"github.com/yuyenews/Beerus/commons/util"
)

// Processing TODO Parsing Package Text and Calling Business Processes
func Processing(buffer *bytes.Buffer, routePath string) (int, string) {

	message, isClose, size := readMessage(buffer)

	if isClose != "" {
		// When the connection is properly disconnected, the route is called for service processing
		websocket.ExecuteClose(routePath)
		return 0, "Connection close"
	}

	if message != "" {
		// When a complete message is parsed, the route is invoked for service processing
		websocket.ExecuteMessage(routePath, message)

		return size, "ok"
	}

	return 0, "ok"
}

func readMessage(buffer *bytes.Buffer) (string, string, int) {

	bytesData := buffer.Bytes()

	opcode := bytesData[0] & 0x0f
	if opcode == 8 {
		return "", "close", 0
	}
	if len(bytesData) < 2 {
		return "reading", "", 0
	}

	payloadLength := int(bytesData[1] & 0x7f)
	if payloadLength < 1 {
		return "reading", "", 0
	}

	maskStartIndex := 2

	if payloadLength == 126 {
		length := getLength(bytesData, 2, 2)
		payloadLength = string_util.BytesToInt(length, 0, len(length))
		maskStartIndex = 4
	} else if payloadLength == 127 {
		length := getLength(bytesData, 2, 8)
		payloadLength = string_util.BytesToInt(length, 0, len(length))
		maskStartIndex = 10
	}

	maskEndIndex := maskStartIndex + 4

	if len(bytesData) < (payloadLength + maskEndIndex) {
		return "reading", "", 0
	}

	mask, _ := string_util.CopyOfRange(bytesData, maskStartIndex, maskEndIndex)
	payloadData, _ := string_util.CopyOfRange(bytesData, maskEndIndex, payloadLength+maskEndIndex)

	if len(payloadData) < payloadLength {
		return "reading", "", 0
	}

	i := 0
	for i < len(payloadData) {
		payloadData[i] = payloadData[i] ^ mask[i%4]
		i++
	}

	return string_util.BytesToString(payloadData), "", maskEndIndex + len(payloadData)
}

func getLength(bytesData []byte, start int, size int) []byte {
	index := 0
	length := make([]byte, size)
	i := start
	for i < (start + size) {
		length[index] = bytesData[i]

		index++
		i++
	}
	return length
}
