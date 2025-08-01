package dtos

import "time"

type CreateTransactionDto struct {
	Description string    `json:"description" validate:"required"`
	TableID     uint64    `json:"table_id" validate:"required"`
	TypeID      uint64    `json:"type_id" validate:"required"`
	Amount      float64   `json:"amount" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	CurrencyID  uint64    `json:"currency_id" validate:"required"`
}
