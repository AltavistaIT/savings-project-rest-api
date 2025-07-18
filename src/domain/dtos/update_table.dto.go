package dtos

type UpdateTableAmountDto struct {
	ID     uint64  `json:"id" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}
