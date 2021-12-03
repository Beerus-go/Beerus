package util

import (
	"bytes"
	"encoding/gob"
)

// Serialization Serialize data to []byte
func Serialization(data interface{}) ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}

// DeSerialization Deserialize source to dst
func DeSerialization(source []byte, dst interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(source))
	err := decoder.Decode(dst)
	if err != nil {
		return err
	}
	return nil
}
