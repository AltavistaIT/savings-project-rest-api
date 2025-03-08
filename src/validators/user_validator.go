package validators

import "github.com/go-playground/validator/v10"

var validate = validator.New()

// type CreateUserRequest struct {
// 	Username string `json:"username" validate:"required"`
// 	Email    string `json:"email" validate:"required"`
// 	Password string `json:"password" validate:"required"`
// }

type CreateUserRequest struct {
	Username string `validate:"required,min=5,max=100"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func ValidateCreateUserRequest(request *CreateUserRequest) error {
	return validate.Struct(request)
}
