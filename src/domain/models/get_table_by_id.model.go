package models

type GetTableByIdModel struct {
	ID uint64 `json:"id" validate:"required"`
}
