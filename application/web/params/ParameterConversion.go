package params

import (
	"encoding/json"
	"github.com/Beerus-go/Beerus/commons"
	"github.com/Beerus-go/Beerus/commons/util"
	"github.com/Beerus-go/Beerus/network/http/commons"
	"log"
	"reflect"
	"strconv"
)

const (
	Field = "field"
)

// ToStruct Take out the parameters and wrap them in struct
func ToStruct(request commons.BeeRequest, pointParamStruct interface{}) {

	contentType := request.ContentType()

	if commons.IsJSON(contentType) {
		if request.Json != "" {
			json.Unmarshal(util.StrToBytes(request.Json), pointParamStruct)
		}
		return
	}

	var paramElem = reflect.ValueOf(pointParamStruct).Elem()
	var paramType = paramElem.Type()

	fieldNum := paramType.NumField()
	for i := 0; i < fieldNum; i++ {
		SetValue(paramType, paramElem, request, i)
	}
}

// ToStructAndValidation Take the parameters out, wrap them in a struct and check the parameters
func ToStructAndValidation(request commons.BeeRequest, pointParamStruct interface{}) string {
	ToStruct(request, pointParamStruct)
	var result = Validation(request, pointParamStruct)
	return result
}

// SetValue Assigning values to fields
func SetValue(paramType reflect.Type, paramElem reflect.Value, request commons.BeeRequest, i int) {
	var structField = paramType.Field(i)
	fieldName := structField.Name
	fieldTag := structField.Tag
	fieldType := GetFieldType(structField)

	field := paramElem.FieldByName(fieldName)

	var paramValues []string

	if fieldTag != "" {
		fieldTagName := fieldTag.Get(Field)
		if fieldTagName != "" {
			paramValues = request.FormValues(fieldTagName)
		}
	}

	if paramValues == nil || len(paramValues) < 1 {
		paramValues = request.FormValues(fieldName)
	}

	if (paramValues == nil || len(paramValues) < 1) && fieldType != data_type.BeeFile {
		return
	}

	// In the actual scenario, most of the fields are definitely not slices, so here we determine if the first value is empty, and if it is, we don't need to go further
	oneParam := paramValues[0]
	if fieldType != data_type.BeeFile && oneParam == "" {
		return
	}

	switch fieldType {
	case data_type.Int:
		val, err := strconv.ParseInt(oneParam, 10, 64)
		if err != nil {
			errorPrint(fieldName, err)
			return
		}
		field.SetInt(val)
	case data_type.Uint:
		val, err := strconv.ParseUint(oneParam, 10, 64)
		if err != nil {
			errorPrint(fieldName, err)
			return
		}
		field.SetUint(val)
		break
	case data_type.Float:
		val, err := strconv.ParseFloat(oneParam, 64)
		if err != nil {
			errorPrint(fieldName, err)
			return
		}
		field.SetFloat(val)
		break
	case data_type.Bool:
		val, err := strconv.ParseBool(oneParam)
		if err != nil {
			errorPrint(fieldName, err)
			return
		}
		field.SetBool(val)
		break
	case data_type.String:
		field.SetString(oneParam)
		break
	case data_type.Slice:
		field.Set(reflect.ValueOf(paramValues))
		break
	case data_type.BeeFile:
		contentType := request.ContentType()
		if commons.IsFormData(contentType) {
			beeFile, err := request.GetFile(fieldName)
			if err != nil {
				errorPrint(fieldName, err.(error))
				return
			}
			field.Set(reflect.ValueOf(*beeFile))
		}
		break
	}
}

// errorPrint
func errorPrint(fieldName string, err error) {
	if err != nil {
		log.Println("field:" + fieldName + "Setting value Exception occurs, " + err.Error())
	}
}
