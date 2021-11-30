package params

import (
	"encoding/json"
	"fmt"
	"github/yuyenews/Beerus/commons"
	"github/yuyenews/Beerus/commons/util"
	"github/yuyenews/Beerus/network/http/commons"
	"reflect"
	"strconv"
)

// ToStruct Take out the parameters and wrap them in struct
func ToStruct(request *commons.BeeRequest, pointParamStruct interface{}, paramStruct interface{}) {

	contentType := request.ContentType()

	if commons.IsJSON(contentType) {
		json.Unmarshal(string_util.StrToBytes(request.Json), pointParamStruct)
		return
	}

	var paramType = reflect.TypeOf(paramStruct)
	var paramElem = reflect.ValueOf(pointParamStruct).Elem()

	i := 0
	fieldNum := paramType.NumField()
	for i < fieldNum {
		setValue(paramType, paramElem, request, i)
		i++
	}
}

// ToStructAndVerification Take the parameters out, wrap them in a struct and check the parameters
func ToStructAndVerification(request *commons.BeeRequest, pointParamStruct interface{}, paramStruct interface{}) string {
	ToStruct(request, pointParamStruct, paramStruct)
	var result = Verification(request, pointParamStruct, paramStruct)
	return result
}

// setValue Assigning values to fields
func setValue(paramType reflect.Type, paramElem reflect.Value, request *commons.BeeRequest, i int) {
	var structField = paramType.Field(i)
	fieldName := structField.Name
	fieldType := structField.Type.Name()

	field := paramElem.FieldByName(fieldName)
	paramValue := request.FormValue(fieldName)

	if paramValue == "" && fieldType != data_type.BeeFile {
		return
	}

	// Unify the handling of numeric variable types to remove the bit identifiers and facilitate the following judgments
	var fType = GetFieldType(fieldType)
	if fType != "" {
		fieldType = fType
	}

	switch fieldType {
	case data_type.Int:
		val, err := strconv.ParseInt(paramValue, 10, 64)
		if err != nil {
			errorPrint(fieldName, err)
			return
		}
		field.SetInt(val)
	case data_type.Uint:
		val, err := strconv.ParseUint(paramValue, 10, 64)
		if err != nil {
			errorPrint(fieldName, err)
			return
		}
		field.SetUint(val)
		break
	case data_type.Float:
		val, err := strconv.ParseFloat(paramValue, 64)
		if err != nil {
			errorPrint(fieldName, err)
			return
		}
		field.SetFloat(val)
		break
	case data_type.Bool:
		val, err := strconv.ParseBool(paramValue)
		if err != nil {
			errorPrint(fieldName, err)
			return
		}
		field.SetBool(val)
		break
	case data_type.String:
		field.SetString(paramValue)
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
		fmt.Println("field:" + fieldName + "Setting value Exception occurs, " + err.Error())
	}
}
