package responses

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type GetConfigResponse struct {
	Currencies       []*entities.Currency        `json:"currencies"`
	TransactionTypes []*entities.TransactionType `json:"transaction_types"`
	MonthYears       []*GetConfigDatesResponse   `json:"month_years"`
}
