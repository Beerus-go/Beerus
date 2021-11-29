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
)

// ----------- 判断Content-Type 是否符合要求 -----------

// IsJSON 判断是否是json
func IsJSON(contentType string) bool {
	contentType = strings.ToLower(contentType)

	if JSON == contentType || strings.HasPrefix(contentType, JSON) {
		return true
	}
	return false
}

// IsFormData 判断是否是formData
func IsFormData(contentType string) bool {
	contentType = strings.ToLower(contentType)

	if FormData == contentType || strings.HasPrefix(contentType, FormData) {
		return true
	}
	return false
}

// IsUrlEncode 判断是否是普通表单
func IsUrlEncode(contentType string) bool {
	contentType = strings.ToLower(contentType)

	if UrlEncode == contentType || strings.HasPrefix(contentType, UrlEncode) {
		return true
	}
	return false
}
