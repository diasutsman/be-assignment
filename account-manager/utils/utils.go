package utils

import (
	"be-test-concrete-ai/account-manager/models/responses"

	"github.com/go-playground/validator/v10"
)

func MsgForTag(tag string) string {
    switch tag {
    case "required":
        return "This field is required"
    case "email":
        return "Invalid email"
	case "oneof":
		return "Invalid type"
    }
	
    return ""
}

func GetErrorArray(validationErrors validator.ValidationErrors) []responses.ApiError {
	errorArray := []responses.ApiError{}
	for _, validationError := range validationErrors {
		errorArray = append(errorArray, responses.ApiError{
			Field: validationError.Field(),
			Msg:   MsgForTag(validationError.Tag()),
		})
	}
	return errorArray
}