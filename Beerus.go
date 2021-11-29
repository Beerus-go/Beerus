package beerus

import (
	"github/yuyenews/beerus/network/http"
	"strconv"
)

func StartHttp(port int) {
	http.StartHttpServer(strconv.Itoa(port))
}
