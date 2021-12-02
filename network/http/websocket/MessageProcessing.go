package websocket

import (
	"bytes"
	"errors"
	"github.com/yuyenews/Beerus/application/websocket"
	"github.com/yuyenews/Beerus/commons/util"
)

const (
	CLOSE   = "connection close"
	READING = "reading"
	BLANK   = "blank"
)

// Processing Parsing Package Text and Calling Business Processes
func Processing(buffer *bytes.Buffer, readSizeCache int, routePath string) (int, error) {

	message, isClose, size := readMessage(buffer, readSizeCache)

	// When the connection is properly disconnected, the route is called for service processing
	if isClose != nil {
		websocket.ExecuteClose(routePath)
		return 0, errors.New(CLOSE)
	}

	// When a complete message is parsed, the route is invoked for service processing
	if message != BLANK && message != READING {
		websocket.ExecuteMessage(routePath, message)
		return size, nil
	}

	return 0, nil
}

// readMessage Parsing messages
func readMessage(buffer *bytes.Buffer, readSizeCache int) (string, error, int) {

	bytesData := buffer.Bytes()

	opcode := bytesData[0] & 15
	if opcode == 8 {
		return BLANK, errors.New(CLOSE), 0
	}
	if readSizeCache < 2 {
		return READING, nil, 0
	}

	payloadLength := int(bytesData[1] & 0x7f)
	if payloadLength < 1 {
		return READING, nil, 0
	}
	mask := bytesData[1] >> 7

	maskStartIndex := 2

	if payloadLength == 126 {
		length := getLength(bytesData, 2, 2)
		payloadLength = util.BytesToInt(length, 0, len(length))
		maskStartIndex = 4
	} else if payloadLength == 127 {
		length := getLength(bytesData, 2, 8)
		payloadLength = util.BytesToInt(length, 0, len(length))
		maskStartIndex = 10
	}

	maskEndIndex := maskStartIndex + 4

	if readSizeCache < (payloadLength + maskEndIndex) {
		return READING, nil, 0
	}

	maskByte, errMsg := util.CopyOfRange(bytesData, maskStartIndex, maskEndIndex)
	payloadData, errMsg2 := util.CopyOfRange(bytesData, maskEndIndex, payloadLength+maskEndIndex)

	if errMsg != nil || errMsg2 != nil {
		return BLANK, errors.New(CLOSE), 0
	}

	if len(payloadData) < payloadLength {
		return READING, nil, 0
	}

	if mask == 1 {
		for i := 0; i < len(payloadData); i++ {
			payloadData[i] = payloadData[i] ^ maskByte[i%4]
		}
	}

	return util.BytesToString(payloadData), nil, maskEndIndex + len(payloadData)
}

// getLength Parsing out the length of the message body from the message
func getLength(bytesData []byte, start int, size int) []byte {
	index := 0
	length := make([]byte, size)

	for i := start; i < (start + size); i++ {
		length[index] = bytesData[i]
		index++
	}
	return length
}
