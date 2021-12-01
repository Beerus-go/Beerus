package websocket

import (
	"bytes"
	"github.com/yuyenews/Beerus/application/websocket"
)

// Processing TODO Parsing Package Text and Calling Business Processes
func Processing(buffer *bytes.Buffer, routePath string) (int, string) {

	// When a complete message is parsed, the route is invoked for service processing
	websocket.ExecuteMessage(routePath, "")

	// When the connection is properly disconnected, the route is called for service processing
	websocket.ExecuteClose(routePath)
	return 0, "ok"
}
