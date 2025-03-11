package users

type GetUserRequest struct {
	ID uint64 `validate:"required"`
}
