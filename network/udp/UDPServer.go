package udp

import (
	"bytes"
	"github.com/Beerus-go/Beerus/commons/util"
	"log"
	"net"
)

// StartUdpServer Start an udp service
func StartUdpServer(handler func(data []byte), separator []byte, port int) {
	if separator == nil || len(separator) <= 0 {
		log.Println("The separator must not be empty")
		return
	}

	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})

	if err != nil {
		log.Println("Listen failed,", err)
		return
	}

	udpConnection(udpConn, separator, handler)
}

// udpConnection Handling UDP connections
func udpConnection(conn *net.UDPConn, separator []byte, handler func(data []byte)) {
	defer conn.Close()

	buf := new(bytes.Buffer)
	readSizeCache := 0
	for {
		var data = make([]byte, 1024)
		ln, addr, err := conn.ReadFromUDP(data)

		if err != nil {
			log.Printf("Read from udp server:%s failed,err:%s", addr, err)
			break
		}

		if ln <= 0 {
			continue
		}

		readSizeCache += ln
		buf.Write(data)

		// It may be possible to read more than one message at a time, so a loop needs to be made here
		for {
			separatorIndex := util.ByteIndexOf(buf.Bytes(), separator)

			if separatorIndex <= 0 {
				break
			}

			message, errMsg := util.SubBytes(buf.Bytes(), 0, separatorIndex)

			if errMsg != nil {
				log.Printf("Read from udp server:%s failed,err:%s", addr, errMsg)
				break
			}

			handler(message)

			processedLength := len(message) + len(separator)

			if processedLength == readSizeCache {
				buf.Reset()
				readSizeCache = 0
				break
			} else {
				remaining, errorMsg := util.CopyOfRange(buf.Bytes(), processedLength, readSizeCache)
				if errorMsg != nil {
					log.Printf("Read from udp server:%s failed,err:%s", addr, errMsg)
					return
				}
				buf = bytes.NewBuffer(remaining)
				readSizeCache = readSizeCache - processedLength
				if readSizeCache < 0 {
					readSizeCache = 0
				}
			}
		}
	}
}
