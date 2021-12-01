package string_util

// SubBytes Extract the byte[] between the specified coordinates
func SubBytes(source []byte, startIndex int, endIndex int) ([]byte, string) {
	if startIndex > endIndex {
		return nil, "The start coordinate cannot be greater than the end coordinate"
	}

	if (len(source) - 1) < startIndex {
		return nil, "Start coordinates are out of length"
	}

	if len(source) < (endIndex - startIndex) {
		return nil, "The length to be Extract is already greater than the data being Extract"
	}

	length := endIndex - startIndex

	bytes := make([]byte, length)

	index := 0
	i := startIndex

	for index < length {
		if i > (len(source) - 1) {
			break
		}

		bytes[index] = source[i]
		i++
		index++
	}

	return bytes, "ok"
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

		j := startIndex
		for j < endIndex {
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
			j++
		}

		if exist {
			return startIndex
		}
	}
	return -1
}

func BytesToInt(b []byte, start int, length int) int {
	sum := 0
	end := start + length

	i := start

	for i < end {
		n := int(b[i]) & 0xff
		length--
		n <<= length * 8
		sum += n
		i++
	}
	return sum
}

func IntToBytes(n int, length int) []byte {
	b := make([]byte, length)

	i := length

	for i > 0 {
		b[(i - 1)] = (byte)(n >> 8 * (length - i) & 0xFF)
		i--
	}
	return b
}
