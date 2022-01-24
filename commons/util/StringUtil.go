package util

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

// StrToBytes string to byte[]
func StrToBytes(val string) []byte {
	if val == "" {
		return nil
	}
	x := (*[2]uintptr)(unsafe.Pointer(&val))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesToString byte[] to string
func BytesToString(val []byte) string {
	if val == nil || len(val) < 1 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&val))
}

// Match Determines if the first parameter matches the second parameter with a wildcard
func Match(source string, reg string) bool {

	var index = strings.Index(reg, "*")
	if index < 0 {
		return source == reg
	}

	if reg == "*" {
		return true
	}

	reg = strings.ReplaceAll(reg, "*", "([a-zA-Z1-9/]+)")
	regular := regexp.MustCompile("^" + reg + "$")

	return regular.MatchString(source)
}

// ToString
// Unbox interface{} into string
func ToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		return strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		return strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		return strconv.Itoa(it)
	case uint:
		it := value.(uint)
		return strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		return strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		return strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		return strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		return strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		return strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		return strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		return strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		return strconv.FormatUint(it, 10)
	case string:
		return value.(string)
	case []byte:
		return BytesToString(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		return BytesToString(newValue)
	}
}
