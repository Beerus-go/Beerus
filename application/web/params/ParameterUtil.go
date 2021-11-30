package params

import (
	"github/yuyenews/Beerus/commons"
	"strings"
)

// GetFieldType 获取字段类型
func GetFieldType(fieldType string) string {
	if strings.HasPrefix(fieldType, data_type.Int) {
		return data_type.Int
	}

	if strings.HasPrefix(fieldType, data_type.Float) {
		return data_type.Float
	}

	if strings.HasPrefix(fieldType, data_type.Uint) {
		return data_type.Uint
	}

	return ""
}
