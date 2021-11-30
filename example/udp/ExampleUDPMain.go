package main

import beerus "github.com/yuyenews/Beerus"

func main() {
	beerus.StartUDP(func(data []byte) {

		// data is the data you received

	}, 8080)
}
