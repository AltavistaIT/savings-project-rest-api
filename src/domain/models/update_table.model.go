package models

type UpdateTableAmountModel struct {
	ID     uint64  `json:"id" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}
