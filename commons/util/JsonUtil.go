package util

import (
	"encoding/json"
	"log"
)

// ToJSONString Structs to json strings
func ToJSONString(stc interface{}) string {
	jsonString, err := json.Marshal(stc)
	if err != nil {
		log.Println(err)
		return ""
	}

	return BytesToString(jsonString)
}

// ParseStruct Json string to struct
func ParseStruct(jsonString string, result interface{}) {
	err := json.Unmarshal(StrToBytes(jsonString), result)
	if err != nil {
		log.Println(err)
	}
}
