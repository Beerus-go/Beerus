package beerus

import (
	"github/yuyenews/Beerus/network/http"
	"strconv"
)

func StartHttp(port int) {
	http.StartHttpServer(strconv.Itoa(port))
}
