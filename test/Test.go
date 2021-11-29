package main

import (
	"github/yuyenews/Beerus"
	"github/yuyenews/Beerus/test/controller"
)

func main() {
	controller.CreateRoute()

	beerus.StartHttp(8080)
}
