package beerus

import (
	"github.com/yuyenews/Beerus/network/http"
	"github.com/yuyenews/Beerus/network/udp"
	"strconv"
)

// Port
// Record the port number,
// if there are other places that need to get the port of this service, you can use this variable directly
var Port = 8080

// ListenHTTP Start an udp service
func ListenHTTP(port int) {
	Port = port
	http.StartHttpServer(strconv.Itoa(port))
}

// ListenUDP Start an udp service
func ListenUDP(handler func(data []byte), separator []byte, port int) {
	udp.StartUdpServer(handler, separator, port)
}
