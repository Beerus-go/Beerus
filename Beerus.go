package beerus

import (
	"github.com/yuyenews/Beerus/network/http"
	"strconv"
)

func StartHttp(port int) {
	http.StartHttpServer(strconv.Itoa(port))
}
