package params

import (
	"fmt"
	"github.com/yuyenews/Beerus/commons"
	"github.com/yuyenews/Beerus/commons/util"
	"github.com/yuyenews/Beerus/network/http/commons"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	SUCCESS         = "SUCCESS"
	ErrorMsg        = "Field:[%s]->Tag:[%s] setting is incorrect,[%s]"
	ValidKeyNotNull = "notnull"
	ValidKeyReg     = "reg"
	ValidKeyMax     = "max"
	ValidKeyMin     = "min"
	ValidKeyRoutes  = "routes"
	ValidKeyMsg     = "msg"
)

// Validation Checking the parameters of the struct
func Validation(request commons.BeeRequest, pointParamStruct interface{}, paramStruct interface{}) string {
	var paramType = reflect.TypeOf(paramStruct)
	var paramElem = reflect.ValueOf(pointParamStruct).Elem()
	return ValidationReflect(request, paramElem, paramType)
}

// ValidationReflect Checking the parameters of the struct
func ValidationReflect(request commons.BeeRequest, paramElem reflect.Value, paramType reflect.Type) string {
	var requestPath = request.RoutePath

	fieldNum := paramType.NumField()
	for i := 0; i < fieldNum; i++ {
		var field = paramType.Field(i)
		var fieldName = field.Name
		var fieldType = GetFieldType(field)
		var fieldTag = field.Tag

		if fieldTag == "" {
			continue
		}

		// Get the tag information of the field and validate the field data based on this information
		var notNull = fieldTag.Get(ValidKeyNotNull)
		var max = fieldTag.Get(ValidKeyMax)
		var min = fieldTag.Get(ValidKeyMin)
		var reg = fieldTag.Get(ValidKeyReg)
		var msg = fieldTag.Get(ValidKeyMsg)
		var routes = fieldTag.Get(ValidKeyRoutes)

		if notNull == "" && reg == "" && max == "" && min == "" {
			continue
		}

		// Whether the wroute of this request is within the scope of this field check
		if routes != "" {
			var isContain = false

			var apisArray = strings.Split(routes, ",")
			for _, apiPath := range apisArray {
				if util.Match(requestPath, apiPath) {
					isContain = true
				}
			}

			if isContain == false {
				continue
			}
		}

		// If the user does not set a prompt message, give a default value
		if msg == "" {
			msg = "Parameters [" + fieldName + "] do not meet the calibration rules"
		}

		// Start checking the fields
		var fieldObj = paramElem.FieldByName(fieldName)
		var result = isSuccess(fieldType, fieldObj, fieldName, notNull, reg, max, min, msg)

		// If the verification does not pass, a message is returned directly
		if result != SUCCESS {
			return result
		}
	}

	return SUCCESS
}

// isSuccess Verify that the value of this field meets the requirements
func isSuccess(fieldType string, field reflect.Value, fieldName string, notNull string, reg string, max string, min string, msg string) string {
	switch fieldType {
	case data_type.Int:
		val := field.Int()
		if min != "" {
			intMin, err := strconv.ParseInt(min, 10, 64)
			if err != nil {
				return fmt.Sprintf(ErrorMsg, fieldName, "min", "Should be set to int type")
			}

			if val < intMin {
				return msg
			}
		}
		if max != "" {
			intMax, err := strconv.ParseInt(max, 10, 64)
			if err != nil {
				return fmt.Sprintf(ErrorMsg, fieldName, "max", "Should be set to int type")
			}

			if val > intMax {
				return msg
			}
		}
		break
	case data_type.Uint:
		var val = field.Uint()
		if min != "" {
			intMin, err := strconv.ParseUint(min, 10, 64)
			if err != nil {
				return fmt.Sprintf(ErrorMsg, fieldName, "min", "Should be set to int type")
			}

			if val < intMin {
				return msg
			}
		}
		if max != "" {
			intMax, err := strconv.ParseUint(max, 10, 64)
			if err != nil {
				return fmt.Sprintf(ErrorMsg, fieldName, "max", "Should be set to int type")
			}

			if val > intMax {
				return msg
			}
		}
		break
	case data_type.Float:
		var val = field.Float()
		if min != "" {
			intMin, err := strconv.ParseFloat(min, 64)
			if err != nil {
				return fmt.Sprintf(ErrorMsg, fieldName, "min", "Should be set to float type")
			}

			if val < intMin {
				return msg
			}
		}
		if max != "" {
			intMax, err := strconv.ParseFloat(max, 64)
			if err != nil {
				return fmt.Sprintf(ErrorMsg, fieldName, "max", "Should be set to float type")
			}

			if val > intMax {
				return msg
			}
		}
		break
	case data_type.String:
		var val = field.String()
		if notNull != "" {
			isNotNull, err := strconv.ParseBool(notNull)
			if err != nil {
				return fmt.Sprintf(ErrorMsg, fieldName, "notnull", "Must be true or false")
			}

			if isNotNull && val == "" {
				return msg
			}
		}
		if reg != "" {
			regular := regexp.MustCompile(reg)
			var isMatch = regular.MatchString(val)
			if !isMatch {
				return msg
			}
		}
		break
	case data_type.Slice:
		if notNull != "" && field.IsNil() {
			return msg
		}
		break
	}

	return SUCCESS
}
