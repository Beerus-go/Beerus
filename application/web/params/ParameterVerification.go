package params

import (
	"fmt"
	"github/yuyenews/Beerus/commons"
	"github/yuyenews/Beerus/commons/util"
	"github/yuyenews/Beerus/network/http/commons"
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

// Verification Checking the parameters of the struct
func Verification(request *commons.BeeRequest, pointParamStruct interface{}, paramStruct interface{}) string {
	var paramType = reflect.TypeOf(paramStruct)
	var paramElem = reflect.ValueOf(pointParamStruct).Elem()
	var requestPath = request.RoutePath

	i := 0
	fieldNum := paramType.NumField()
	for i < fieldNum {
		var field = paramType.Field(i)
		var fieldName = field.Name
		var fieldType = field.Type.Name()
		var fieldTag = field.Tag

		if fieldTag == "" {
			i++
			continue
		}

		// Get the tag information of the field and validate the field data based on this information
		var notNull = fieldTag.Get(ValidKeyNotNull)
		var max = fieldTag.Get(ValidKeyMax)
		var min = fieldTag.Get(ValidKeyMin)
		var reg = fieldTag.Get(ValidKeyReg)
		var msg = fieldTag.Get(ValidKeyMsg)
		var apis = fieldTag.Get(ValidKeyRoutes)

		if notNull == "" && reg == "" && max == "" && min == "" {
			i++
			continue
		}

		// Whether the route of this request is within the scope of this field check
		if apis != "" {
			var isContain = false

			var apisArray = strings.Split(apis, ",")
			for _, apiPath := range apisArray {
				if string_util.Match(requestPath, apiPath) {
					isContain = true
				}
			}

			if isContain == false {
				i++
				continue
			}
		}

		// If the user does not set a prompt message, give a default value
		if msg == "" {
			msg = "Parameters [" + fieldName + "] do not meet the calibration rules"
		}

		// Unify the handling of numeric variable types to remove the bit identifiers and facilitate the following judgments
		var fType = GetFieldType(fieldType)
		if fType != "" {
			fieldType = fType
		}

		// Start checking the fields
		var fieldObj = paramElem.FieldByName(fieldName)
		var result = isSuccess(fieldType, fieldObj, fieldName, notNull, reg, max, min, msg)

		// If the verification does not pass, a message is returned directly
		if result != SUCCESS {
			return result
		}

		i++
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
	}

	return SUCCESS
}
