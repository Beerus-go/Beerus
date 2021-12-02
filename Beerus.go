package beerus

import (
	"github.com/yuyenews/Beerus/network/http"
	"github.com/yuyenews/Beerus/network/udp"
	"strconv"
)

// ListenHTTP Start an udp service
func ListenHTTP(port int) {
	http.StartHttpServer(strconv.Itoa(port))
}

// ListenUDP Start an udp service
func ListenUDP(handler func(data []byte), separator []byte, port int) {
	udp.StartUdpServer(handler, separator, port)
}
