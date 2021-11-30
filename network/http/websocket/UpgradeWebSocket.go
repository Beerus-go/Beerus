package websocket

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/yuyenews/Beerus/commons/util"
	"github.com/yuyenews/Beerus/network/http/commons"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Upgrade(write http.ResponseWriter, request *http.Request) {

	h, ok := write.(http.Hijacker)

	if ok {
		netConn, brw, err := h.Hijack()
		if err != nil {
			log.Println(err)
		}

		if brw.Reader.Buffered() > 0 {
			netConn.Close()
			log.Println("websocket: client sent data before handshake is complete")
		}

		stringBuilder := strings.Builder{}

		stringBuilder.WriteString(commons.ResponseOnline)
		stringBuilder.WriteString(commons.CarriageReturn)
		stringBuilder.WriteString("Upgrade:websocket")
		stringBuilder.WriteString(commons.CarriageReturn)
		stringBuilder.WriteString("upgrade:Upgrade")
		stringBuilder.WriteString(commons.CarriageReturn)
		stringBuilder.WriteString("Sec-WebSocket-Accept:" + getAccept(request))
		stringBuilder.WriteString(commons.CarriageReturn)
		stringBuilder.WriteString(commons.CarriageReturn)

		var wr = string_util.StrToBytes(stringBuilder.String())

		netConn.SetDeadline(time.Time{})

		len, err := netConn.Write(wr)
		log.Printf("websocket write %s size\r\n", strconv.Itoa(len))
		if err != nil {
			println("-----------------------")
			netConn.Close()
		}
	}

	//stringBuilder := strings.Builder{}
	//
	//stringBuilder.WriteString(commons.ResponseOnline)
	//stringBuilder.WriteString(commons.CarriageReturn)
	//stringBuilder.WriteString("Upgrade:websocket")
	//stringBuilder.WriteString(commons.CarriageReturn)
	//stringBuilder.WriteString("upgrade:Upgrade")
	//stringBuilder.WriteString(commons.CarriageReturn)
	//stringBuilder.WriteString("Sec-WebSocket-Accept:" + getAccept(request))
	//stringBuilder.WriteString(commons.CarriageReturn)
	//stringBuilder.WriteString(commons.CarriageReturn)
	//
	//s := stringBuilder.String()

	//write.Write(string_util.StrToBytes("OK"))

	//for {
	//
	//}
}

func getAccept(request *http.Request) string {
	secKey := request.Header.Get(commons.SecWebsocketKey)

	swKey := secKey + commons.SocketSecretKey

	hash := sha1.New()

	hash.Write(string_util.StrToBytes(swKey))

	hashResult := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(hashResult)
}
