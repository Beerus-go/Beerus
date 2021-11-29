package params

import "github/yuyenews/Beerus/network/http/commons"

// ToStruct Take out the parameters and wrap them in struct
func ToStruct(request *commons.BeeRequest, paramStruct interface{}) interface{} {

	return nil
}

// ToStructAndVerification Take the parameters out, wrap them in a struct and check the parameters
func ToStructAndVerification(request *commons.BeeRequest, paramStruct interface{}) (interface{}, string) {
	var params = ToStruct(request, paramStruct)
	var result = Verification(params)
	return params, result
}
