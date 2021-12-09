package websocket

import (
	"github.com/yuyenews/Beerus/application/websocket/params"
	"github.com/yuyenews/Beerus/application/websocket/route"
	"github.com/yuyenews/Beerus/commons/util"
	"log"
	"net"
)

// SessionMap Save each linked session
var sessionMap = make(map[string]*params.WebSocketSession)

var snowflake *util.SnowFlake

// ExecuteConnection Triggers the onConnection function inside the route
func ExecuteConnection(routePath string, conn net.Conn) {

	var snowflakeId uint64
	var err error
	if snowflake == nil {
		snowflake, err = util.New(2)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	snowflakeId, err = snowflake.Generate()
	if err != nil {
		log.Println(err.Error())
		return
	}

	session := new(params.WebSocketSession)
	session.Connection = conn
	session.Id = snowflakeId
	sessionMap[routePath] = session

	route.GetWebSocketRoute(routePath, route.OnConnection)(session, "The client is already connected")
}

// ExecuteMessage Triggers the onMessage function inside the route
func ExecuteMessage(routePath string, message string) {
	route.GetWebSocketRoute(routePath, route.OnMessage)(sessionMap[routePath], message)
}

// ExecuteClose Triggers the onClose function inside the route
func ExecuteClose(routePath string) {
	if sessionMap[routePath] == nil {
		return
	}

	delete(sessionMap, routePath)
	route.GetWebSocketRoute(routePath, route.OnClose)(sessionMap[routePath], "The client is disconnected")
}
