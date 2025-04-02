package transactions

type CreateTransactionRequest struct {
	Description string  `json:"description" validate:"required"`
	TableID     uint64  `json:"table_id" validate:"required"`
	TypeID      uint64  `json:"type_id" validate:"required"`
	Budget      float64 `json:"budget" validate:"required"`
	CurrencyID  uint64  `json:"currency_id" validate:"required"`
}
