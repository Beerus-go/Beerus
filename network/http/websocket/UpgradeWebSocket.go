package websocket

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"github.com/yuyenews/Beerus/application/websocket"
	"github.com/yuyenews/Beerus/application/websocket/route"
	"github.com/yuyenews/Beerus/commons/util"
	"github.com/yuyenews/Beerus/network/http/commons"
	"log"
	"net"
	"net/http"
	"strings"
)

// UpgradeToWebSocket Upgrade to websocket
func UpgradeToWebSocket(write http.ResponseWriter, request *commons.BeeRequest) {

	// Does routing exist
	if route.Exist(request.RoutePath) == false {
		log.Println("WebSocket route does not exist, connection failed")
		return
	}

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

	_, err = netConn.Write(util.StrToBytes(stringBuilder.String()))

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

	// Message data that has been read but not yet processed
	buf := new(bytes.Buffer)

	// Length of messages already read
	readSizeCache := 0

	for {
		// Read messages from the client
		readByte := make([]byte, 500)
		ln, err := conn.Read(readByte)
		if err != nil {
			websocket.ExecuteClose(routePath)
			log.Println("WebSocket client abnormally disconnected, " + err.Error())
			break
		}

		if ln <= 0 {
			continue
		}

		readSizeCache += ln
		buf.Write(readByte)

		// Parse the message and call the corresponding handler for business processing
		size, errMsg := Processing(buf, readSizeCache, routePath)
		if errMsg != nil {
			websocket.ExecuteClose(routePath)
			log.Println("WebSocket Exceptions in parsing messages, " + errMsg.Error())
			break
		}

		// Parsed data If = 0, it means that a complete message has not been read yet, so continue reading
		if size <= 0 {
			continue
		}

		// Remove used data from the cache
		if size == readSizeCache {
			buf.Reset()
			readSizeCache = 0
		} else {
			remaining, errorMsg := util.CopyOfRange(buf.Bytes(), size, readSizeCache)
			if errorMsg != nil {
				websocket.ExecuteClose(routePath)
				log.Println("WebSocket Removing exceptions from already processed data, " + errMsg.Error())
				break
			}
			buf = bytes.NewBuffer(remaining)
			readSizeCache = readSizeCache - size
		}
	}
}

// Get the logo for connecting to the client
func getAccept(request *commons.BeeRequest) string {
	secKey := request.HeaderValue(commons.SecWebsocketKey)

	swKey := secKey + commons.SocketSecretKey

	hash := sha1.New()

	hash.Write(util.StrToBytes(swKey))

	hashResult := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(hashResult)
}
