package beerus

import (
	"github.com/yuyenews/Beerus/network/http"
	"github.com/yuyenews/Beerus/network/udp"
	"strconv"
)

// StartHttp Start an udp service
func StartHttp(port int) {
	http.StartHttpServer(strconv.Itoa(port))
}

// StartUDP Start an udp service
func StartUDP(handler func(data []byte), separator []byte, port int) {
	udp.StartUdpServer(handler, separator, port)
}
