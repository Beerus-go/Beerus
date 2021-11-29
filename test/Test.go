package main

import (
	"github/yuyenews/beerus"
	"github/yuyenews/beerus/test/controller"
)

func main() {
	controller.CreateRoute()

	beerus.StartHttp(8080)
}
