package shared

import "github.com/go-playground/validator/v10"

func ValidateRequest(request interface{}) error {
	return validator.New().Struct(request)
}
