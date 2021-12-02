# [Beerus](https://www.ww.com) | <img src="https://img.shields.io/badge/licenes-MIT-brightgreen.svg"/> <img src="https://img.shields.io/badge/golang-1.17.3-brightgreen.svg"/> <img src="https://img.shields.io/badge/release-master-brightgreen.svg"/>

Beerus is a web framework developed entirely in go, 
whether it is json passing, form submission, formdata, etc., 
it is easy to extract parameters from the request into the structure and automatically complete parameter validation.

***NetWork:*** Based on go's own package: net/http, with many extensions added

***WebSocket:*** Extensions to net/http to automatically analyze WebSocket requests and upgrade to the websocket protocol for message listening and sending and receiving

***Http:*** Most of the functions are developed based on net/http, the Json pass-through is handled, and each request is automatically routed to the corresponding function

***Enhanced routing management:*** Enhanced the way handlers are configured and unified through routing, so that each request will be automatically routed to the corresponding function

***Automatic processing parameters:*** can automatically extract any type of parameter from request to struct, and support automatic parameter validation, regular, range, non-empty validation, and custom hint messages


## Installation

## Documentation

## Examples

## License

Beerus is [MIT licensed.](https://github.com/yuyenews/Beerus/blob/master/LICENSE)
