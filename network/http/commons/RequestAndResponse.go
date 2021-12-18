package commons

import (
	"fmt"
	"github.com/yuyenews/Beerus/commons/util"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// BeeRequest ----------- Secondary wrapping over the request object, mainly to facilitate the acquisition of json passing parameters -----------
// Secondary encapsulation of a part of the high-frequency use of the function, other functions can be taken from the Request inside
type BeeRequest struct {
	Request   *http.Request
	Json      string
	RoutePath string
	Params    map[string][]string
}

// FormValue Get request parameters
func (req *BeeRequest) FormValue(key string) string {
	values := req.FormValues(key)
	if values == nil || len(values) < 1 {
		return ""
	}
	return values[0]
}

// FormValues Get request parameters
func (req *BeeRequest) FormValues(key string) []string {
	values := req.Params[key]
	if values == nil || len(values) < 1 {
		values = make([]string, 1)
		values[0] = req.Request.FormValue(key)
		return values
	}
	return values
}

// HeaderValue Get request header
func (req *BeeRequest) HeaderValue(key string) string {
	return req.Request.Header.Get(key)
}

// HeaderValues Get request headers
func (req *BeeRequest) HeaderValues(key string) []string {
	return req.Request.Header.Values(key)
}

// GetFile Get request file
func (req *BeeRequest) GetFile(key string) (*BeeFile, error) {
	file, fileHeader, error := req.Request.FormFile(key)

	var beeFile = new(BeeFile)
	beeFile.File = file
	beeFile.FileHeader = fileHeader

	return beeFile, error
}

// ContentType get Content-Type
func (req *BeeRequest) ContentType() string {
	contentType := req.HeaderValue(ContentType)
	if contentType == "" {
		contentType = req.HeaderValue(ContentType2)
	}
	return contentType
}

// AddParam add param
func (req *BeeRequest) AddParam(name string, val string) {
	name = strings.TrimSpace(name)
	val = strings.TrimSpace(val)

	paramArray := req.Params[name]

	if paramArray == nil {
		paramArray = make([]string, 0)
	}

	paramArray = append(paramArray, val)

	req.Params[name] = paramArray
}

// BeeResponse ----------- Secondary wrapping of the response object, the response part is enhanced a bit, providing some high-frequency use of the function -----------
type BeeResponse struct {
	Response http.ResponseWriter
}

// SetHeader Set the response header
func (res *BeeResponse) SetHeader(key string, value string) *BeeResponse {
	res.Response.Header().Set(key, value)
	return res
}

// SendJson Send json string to client
func (res *BeeResponse) SendJson(value string) {
	res.SetHeader(ContentType, "application/json;charset=UTF-8")
	res.Response.Write(util.StrToBytes(value))
}

// SendText Send text to client
func (res *BeeResponse) SendText(value string) {
	res.SetHeader(ContentType, "text/plain;charset=UTF-8")
	res.Response.Write(util.StrToBytes(value))
}

// SendHtml Send html string to client
func (res *BeeResponse) SendHtml(value string) {
	res.SetHeader(ContentType, "text/html;charset=UTF-8")
	res.Response.Write(util.StrToBytes(value))
}

// SendStream Send file to client
func (res *BeeResponse) SendStream(fileName string, file []byte) {
	res.SetHeader(ContentType, "application/octet-stream")
	res.SetHeader(ContentDisposition, "attachment; filename="+url.PathEscape(fileName))
	res.Response.Write(file)
}

// SendData Sending other custom ContentType data to the client
func (res *BeeResponse) SendData(value string) {
	res.Response.Write(util.StrToBytes(value))
}

// SendErrorMsg Return error messages in json format
func (res *BeeResponse) SendErrorMsg(code int, msg string) {
	res.SendJson(fmt.Sprintf(ErrorMsg, strconv.Itoa(code), msg))
}
