package string_util

import "unsafe"

func StrToBytes(val string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&val))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func BytesToString(val []byte) string {
	return *(*string)(unsafe.Pointer(&val))
}
