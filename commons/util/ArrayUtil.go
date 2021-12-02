package util

import "errors"

// CopyOfRange Copy the data of the specified range in the array
func CopyOfRange(src []byte, srcOffset int, size int) ([]byte, error) {
	srcLen := len(src)

	if srcOffset > srcLen || size > srcLen {
		return nil, errors.New("source buffer Index out of range")
	}

	dst := make([]byte, size-srcOffset)

	index := 0
	for i := srcOffset; i < size; i++ {
		dst[index] = src[i]
		index++
	}
	return dst, nil
}
