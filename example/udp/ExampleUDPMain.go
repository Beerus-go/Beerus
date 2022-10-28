package main

import (
	"github.com/Beerus-go/Beerus"
	"github.com/Beerus-go/Beerus/commons/util"
)

func main() {

	// Listening to a UDP service
	// The first parameter is the handler
	// The second parameter is the data separator
	// The third parameter is the port
	beerus.ListenUDP(updHandler, []byte("|"), 8080)

}

func updHandler(data []byte) {
	// data is the data you received
	println(util.BytesToString(data))
}
