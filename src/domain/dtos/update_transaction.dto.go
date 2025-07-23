package dtos

import "time"

type UpdateTransactionDto struct {
	ID          uint64    `json:"id" validate:"required"`
	TypeID      uint64    `json:"type_id" validate:""`
	Amount      float64   `json:"amount" validate:""`
	Date        time.Time `json:"date" validate:"required"`
	Description string    `json:"description" validate:""`
}
