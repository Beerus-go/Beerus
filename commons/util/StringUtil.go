package string_util

import (
	"regexp"
	"strings"
	"unsafe"
)

// StrToBytes string to byte[]
func StrToBytes(val string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&val))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesToString byte[] to string
func BytesToString(val []byte) string {
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
	regular := regexp.MustCompile(reg)

	return regular.MatchString(source)
}
