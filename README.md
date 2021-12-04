# [Beerus](https://www.ww.com) | <img src="https://img.shields.io/badge/licenes-MIT-brightgreen.svg"/> <img src="https://img.shields.io/badge/golang-1.17.3-brightgreen.svg"/> <img src="https://img.shields.io/badge/release-master-brightgreen.svg"/>

Beerus is a web framework developed entirely in go, 
whether it is json passing, form submission, formdata, etc., 
it is easy to extract parameters from the request into the structure and automatically complete parameter validation.

- ***NetWork:*** Based on go's own package: net/http, with many extensions added

- ***WebSocket:*** Extensions to net/http to automatically analyze WebSocket requests and upgrade to the websocket protocol for message listening and sending and receiving

- ***Http:*** Most of the functions are developed based on net/http, the Json pass-through is handled, and each request is automatically routed to the corresponding function

- ***Enhanced routing management:*** Enhanced the way handlers are configured and unified through routing, so that each request will be automatically routed to the corresponding function

- ***Automatic processing parameters:*** can automatically extract any type of parameter from request to struct, and support automatic parameter validation, regular, range, non-empty validation, and custom hint messages


## Installation

## Documentation

## Examples

### HTTP example

Create a function to manage the routing configuration

```go
func CreateRoute() {
	// post route example
    route.POST("/example/post", func (req *commons.BeeRequest, res *commons.BeeResponse) {
        
        res.SendJson(`{"msg":"SUCCESS"}`)
    })

    // get route example
    route.GET("/example/get", func (req *commons.BeeRequest, res *commons.BeeResponse) {
    
        res.SendJson(`{"msg":"SUCCESS"}`)
    })
}
```

Start Service

```go
func main() {
    // Interceptors, routes, etc. Loading of data requires its own calls
    routes.CreateRoute()
    
    // Listen the service and listen to port 8080
    beerus.ListenHTTP(8080)
}
```

If you want to put the parameters inside struct and complete the parameter checks

```go
func CreateRoute() {
    // Example of parameter conversion to struct and parameter checksum
    route.POST("/example/post", func (req *commons.BeeRequest, res *commons.BeeResponse) {
        param := DemoParam{}
        
        // Extraction parameters, Generally used in scenarios where verification is not required or you want to verify manually
        params.ToStruct(req, &param, param)
        
        // Separate validation of data in struct, this feature can be used independently in any case and is not limited to the routing layer.
        var result = params.Validation(req, &param, param)
        if result != params.SUCCESS {
            res.SendErrorMsg(1128, result)
            return
        }
        
        // You can also convert + validate the parameters in one step
        // Extraction of parameters + validation
        var result = params.ToStructAndValidation(req, &param, param)
        if result != params.SUCCESS {
            res.SendErrorMsg(1128, result)
            return
        }
        
        
        res.SendJson(`{"msg":"SUCCESS"}`)
    })
}

// DemoParam If you have a struct like this, and you want to put all the parameters from the request into this struct
type DemoParam struct {
    // You can customize any field
    // the name of the field must be exactly the same as the name of the requested parameter, and is case-sensitive
    TestStringReception  string  `notnull:"true" msg:"TestStringReception Cannot be empty" routes:"/example/put"`
    TestIntReception     int     `max:"123" min:"32" msg:"TestIntReception The value range must be between 32 - 123" routes:"/example/post"`
    TestUintReception    uint    `max:"123" min:"32" msg:"TestUintReception The value range must be between 32 - 123"`
    TestFloatReception   float32 `max:"123" min:"32" msg:"TestFloatReception The value range must be between 32 - 123"`
    TestBoolReception    bool
    TestStringRegReception string `reg:"^[a-z]+$" msg:"TestStringRegReception Does not meet the regular"`
    TestBeeFileReception commons.BeeFile
    
    TestJsonReception []string
}
```

### WebSocket example

CreateWebSocketRoute Creating websocket routes

```go
func CreateWebSocketRoute() {
	route.AddWebSocketRoute("/ws/test", onConnection, onMessage, onClose)
	route.AddWebSocketRoute("/ws/test2", onConnection, onMessage, onClose)
}

// In order to save time, only three functions are used below. In practice, you can configure a set of functions for each route

func onConnection(session *params.WebSocketSession, msg string) {
	session.SendString("connection success")
}

func onMessage(session *params.WebSocketSession, msg string) {
	session.SendString("I got the message.")
}

func onClose(session *params.WebSocketSession, msg string) {
    println(msg + "-------------------------------")
}
```

Start Service

```go
func main() {
    // Interceptors, routes, etc. Loading of data requires its own calls
    routes.CreateRoute()
    routes.CreateWebSocketRoute()
    
    // Listen the service and listen to port 8080
    beerus.ListenHTTP(8080)
}
```

[Complete sample code](https://github.com/yuyenews/Beerus/tree/master/example)

## License

Beerus is [MIT licensed](https://github.com/yuyenews/Beerus/blob/master/LICENSE)
