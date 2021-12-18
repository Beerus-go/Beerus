package params

import (
	"github.com/yuyenews/Beerus/commons"
	"reflect"
	"strings"
)

// GetFieldType 获取字段类型
func GetFieldType(structField reflect.StructField) string {
	fieldType := structField.Type.Kind().String()
	if fieldType == "" || fieldType == data_type.Struct {
		fieldType = structField.Type.Name()
	}

	if fieldType == "" {
		return ""
	}

	fieldType = strings.ToLower(fieldType)

	if strings.HasPrefix(fieldType, data_type.Int) {
		return data_type.Int
	}

	if strings.HasPrefix(fieldType, data_type.Float) {
		return data_type.Float
	}

	if strings.HasPrefix(fieldType, data_type.Uint) {
		return data_type.Uint
	}

	return fieldType
}
