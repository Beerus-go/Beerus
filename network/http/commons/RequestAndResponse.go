package commons

import (
	"fmt"
	"github/yuyenews/Beerus/commons/util"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

// BeeRequest ----------- 二次封装过的请求对象，主要是为了方便获取json传参 -----------
// 二次封装了一部分 高频使用的函数，其它函数可以从 Request 里面取
type BeeRequest struct {
	Request *http.Request
	Json    string
}

// FormValue 获取请求参数
func (req BeeRequest) FormValue(key string) string {
	return req.Request.FormValue(key)
}

// HeaderValue 获取请求头
func (req BeeRequest) HeaderValue(key string) string {
	return req.Request.Header.Get(key)
}

// HeaderValues 获取请求头
func (req BeeRequest) HeaderValues(key string) []string {
	return req.Request.Header.Values(key)
}

// getFile 获取请求文件
func (req BeeRequest) getFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return req.Request.FormFile(key)
}

// BeeResponse ----------- 二次封装过的请求对象，主要是为了将参数的获取方式进行统一 -----------
// 对响应的部分增强了一下，提供了一些高频使用的函数
type BeeResponse struct {
	Response http.ResponseWriter
}

// SetHeader 设置响应头
func (res BeeResponse) SetHeader(key string, value string) BeeResponse {
	res.Response.Header().Set(key, value)
	return res
}

// SendJson 响应json字符串
func (res BeeResponse) SendJson(value string) {
	res.SetHeader(ContentType, "application/json;charset=UTF-8")
	res.Response.Write(string_util.StrToBytes(value))
}

// SendText 响应文本
func (res BeeResponse) SendText(value string) {
	res.SetHeader(ContentType, "text/plain;charset=UTF-8")
	res.Response.Write(string_util.StrToBytes(value))
}

// SendHtml 响应html
func (res BeeResponse) SendHtml(value string) {
	res.SetHeader(ContentType, "text/html;charset=UTF-8")
	res.Response.Write(string_util.StrToBytes(value))
}

// SendStream 响应文件流
func (res BeeResponse) SendStream(fileName string, file []byte) {
	res.SetHeader(ContentType, "application/octet-stream")
	res.SetHeader(ContentDisposition, "attachment; filename="+url.PathEscape(fileName))
	res.Response.Write(file)
}

// SendData 响应其他自定义ContentType的数据
func (res BeeResponse) SendData(value string) {
	res.Response.Write(string_util.StrToBytes(value))
}

// SendErrorMsg 以json格式返回错误信息
func (res BeeResponse) SendErrorMsg(code int, msg string) {
	res.SendJson(fmt.Sprintf(ErrorMsg, strconv.Itoa(code), msg))
}
