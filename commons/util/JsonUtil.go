package util

import (
	"encoding/json"
	"log"
)

// ToJSONString Structs to json strings
func ToJSONString(stc interface{}) (string, error) {
	jsonString, err := json.Marshal(stc)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return BytesToString(jsonString), nil
}

// ParseStruct Json string to struct
func ParseStruct(jsonString string, result interface{}) error {
	err := json.Unmarshal(StrToBytes(jsonString), result)
	if err != nil {
		return err
	}
	return nil
}
