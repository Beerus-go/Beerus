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
func StartUDP(function func(data []byte), port int) {
	udp.StartUdpServer(function, port)
}
