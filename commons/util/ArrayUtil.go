package string_util

func CopyOfRange(src []byte, srcOffset int, size int) ([]byte, string) {
	srcLen := len(src)

	if srcOffset > srcLen || size > srcLen {
		return nil, "Source buffer Index out of range"
	}

	dst := make([]byte, size-srcOffset)

	index := 0
	for i := srcOffset; i < size; i++ {
		dst[index] = src[srcOffset]
		index++
	}
	return dst, ""
}

func ArrayCopy(src []byte, srcOffset int, dst []byte, start int, size int) []byte {
	return nil
}
