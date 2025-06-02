package models

type GetTransactionsByTableIDModel struct {
	TableID uint64 `validate:"required"`
}
