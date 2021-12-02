package util

import "errors"

// SubBytes Extract the byte[] between the specified coordinates
func SubBytes(source []byte, startIndex int, endIndex int) ([]byte, error) {
	if startIndex > endIndex {
		return nil, errors.New("the start coordinate cannot be greater than the end coordinate")
	}

	if (len(source) - 1) < startIndex {
		return nil, errors.New("start coordinates are out of length")
	}

	if len(source) < (endIndex - startIndex) {
		return nil, errors.New("the length to be Extract is already greater than the data being Extract")
	}

	length := endIndex - startIndex
	bytes := make([]byte, length)

	i := startIndex
	for index := 0; index < length; index++ {
		if i > (len(source) - 1) {
			break
		}

		bytes[index] = source[i]
		i++
	}

	return bytes, nil
}

// ByteIndexOf Find the coordinates of the corresponding data from byte[]
func ByteIndexOf(source []byte, targetByte []byte) int {

	startIndex := 0
	endIndex := len(targetByte)

	for {
		index := 0
		exist := true

		if (len(source) - startIndex) < len(targetByte) {
			return -1
		}

		for j := startIndex; j < endIndex; j++ {
			if index > len(targetByte)-1 {
				return -1
			}

			if source[j] != targetByte[index] {
				startIndex++
				endIndex++

				if startIndex > (len(source)-1) || endIndex > len(source) {
					return -1
				}

				exist = false
				break
			}
			index++
		}

		if exist {
			return startIndex
		}
	}
	return -1
}

// BytesToInt byte[] to int
func BytesToInt(b []byte, start int, length int) int {
	sum := 0
	end := start + length

	for i := start; i < end; i++ {
		n := int(b[i]) & 0xff
		length--
		n <<= length * 8
		sum += n
	}
	return sum
}

// IntToBytes int to byte[]
func IntToBytes(n int, length int) []byte {
	b := make([]byte, length)

	for i := length; i > 0; i-- {
		b[(i - 1)] = (byte)(n >> 8 * (length - i) & 0xFF)
	}
	return b
}
