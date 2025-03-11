package tables

type CreateTableRequest struct {
	UserID    uint64 `json:"user_id" validate:"required"`
	TypeID    uint64 `json:"type_id" validate:"required"`
	MonthYear string `json:"month_year" validate:"required"`
}
