package main

import (
	"github.com/yuyenews/Beerus"
	"github.com/yuyenews/Beerus/commons/util"
)

func main() {
	beerus.StartUDP(func(data []byte) {

		// data is the data you received
		println(util.BytesToString(data))

	}, []byte("|"), 8080)
}
