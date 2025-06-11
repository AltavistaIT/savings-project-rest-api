package models

type GetTableModel struct {
	ID uint64 `json:"id" validate:"required"`
}
