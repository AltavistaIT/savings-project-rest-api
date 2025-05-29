package models

type CreateUserModel struct {
	Username string `validate:"required,min=5,max=100"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
