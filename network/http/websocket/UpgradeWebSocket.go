package websocket

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"github.com/yuyenews/Beerus/application/websocket"
	"github.com/yuyenews/Beerus/commons/util"
	"github.com/yuyenews/Beerus/network/http/commons"
	"log"
	"net"
	"net/http"
	"strings"
)

// UpgradeToWebSocket Upgrade to websocket
func UpgradeToWebSocket(write http.ResponseWriter, request *commons.BeeRequest) {

	h, ok := write.(http.Hijacker)

	if ok == false {
		return
	}

	netConn, brw, err := h.Hijack()
	if err != nil {
		log.Println(err)
	}

	if brw.Reader.Buffered() > 0 {
		log.Println("WebSocket client sent data before handshake is complete")
		netConn.Close()
		return
	}

	stringBuilder := strings.Builder{}

	stringBuilder.WriteString(commons.ResponseOnline)
	stringBuilder.WriteString(commons.CarriageReturn)
	stringBuilder.WriteString("Upgrade:websocket")
	stringBuilder.WriteString(commons.CarriageReturn)
	stringBuilder.WriteString("Connection:Upgrade")
	stringBuilder.WriteString(commons.CarriageReturn)
	stringBuilder.WriteString("Sec-WebSocket-Accept:" + getAccept(request))
	stringBuilder.WriteString(commons.CarriageReturn)
	stringBuilder.WriteString(commons.CarriageReturn)

	_, err = netConn.Write(string_util.StrToBytes(stringBuilder.String()))

	if err != nil {
		log.Println("WebSocket Link establishment failure with client, " + err.Error())
		netConn.Close()
	}

	websocket.ExecuteConnection(request.RoutePath, netConn)

	// Open a goroutine that listens to this link
	go connProcessing(netConn, request.RoutePath)
}

// connProcessing listens to this link
func connProcessing(conn net.Conn, routePath string) {
	defer conn.Close()

	buf := new(bytes.Buffer)
	for {
		// Read messages from the client
		readByte := make([]byte, 1024)
		ln, err := conn.Read(readByte)
		if err != nil {
			websocket.ExecuteClose(routePath)
			log.Println("WebSocket client abnormally disconnected, " + err.Error())
			break
		}

		if ln <= 0 {
			continue
		}
		buf.Write(readByte)

		// Parse the message and call the corresponding handler for business processing
		size, errMsg := Processing(buf, routePath)
		if errMsg != "ok" {
			websocket.ExecuteClose(routePath)
			log.Println("WebSocket Exceptions in parsing messages, " + errMsg)
			break
		}

		// Remove used data from the cache
		if size == buf.Len() {
			buf.Reset()
		} else {
			remaining, errorMsg := string_util.SubBytes(buf.Bytes(), 0, size)
			if errorMsg != "ok" {
				websocket.ExecuteClose(routePath)
				log.Println("WebSocket Removing exceptions from already processed data, " + errMsg)
				break
			}
			buf = bytes.NewBuffer(remaining)
		}
	}
}

// Get the logo for connecting to the client
func getAccept(request *commons.BeeRequest) string {
	secKey := request.HeaderValue(commons.SecWebsocketKey)

	swKey := secKey + commons.SocketSecretKey

	hash := sha1.New()

	hash.Write(string_util.StrToBytes(swKey))

	hashResult := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(hashResult)
}
