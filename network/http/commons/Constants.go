package commons

import "strings"

const (
	UrlEncode          = "application/x-www-form-urlencoded"
	FormData           = "multipart/form-data"
	JSON               = "application/json"
	ContentType        = "Content-Type"
	ContentType2       = "content-type"
	ContentDisposition = "Content-Disposition"
	ErrorMsg           = "{\"code\":%s, \"msg\":\"%s\"}"
	CarriageReturn     = "\r\n"
)

const (
	Connection           = "Connection"
	Upgrade              = "Upgrade"
	SecWebsocketKey      = "Sec-WebSocket-Key"
	SecWebsocketProtocol = "Sec-WebSocket-Protocol"
	ResponseOnline       = "HTTP/1.1 101 Switching Protocols"

	// SocketSecretKey SecretKey encryption secret key
	SocketSecretKey = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
)

// ----------- Determine whether the Content-Type matches the requirements -----------

func IsJSON(contentType string) bool {
	contentType = strings.ToLower(contentType)

	if JSON == contentType || strings.HasPrefix(contentType, JSON) {
		return true
	}
	return false
}

func IsFormData(contentType string) bool {
	contentType = strings.ToLower(contentType)

	if FormData == contentType || strings.HasPrefix(contentType, FormData) {
		return true
	}
	return false
}

func IsUrlEncode(contentType string) bool {
	contentType = strings.ToLower(contentType)

	if UrlEncode == contentType || strings.HasPrefix(contentType, UrlEncode) {
		return true
	}
	return false
}
