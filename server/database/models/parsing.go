package models

import (
	"clueless-cartel-server/api/apiModels"
	"reflect"
)

func ValidateData(params interface{}, res *apiModels.Response) {
	values := reflect.ValueOf(params)
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface() == "" {
			res.Errors = append(res.Errors, "Missing "+values.Type().Field(i).Name)
		}
	}
}
