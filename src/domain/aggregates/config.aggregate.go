package aggregates

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type Config struct {
	Currencies       []*entities.Currency        `json:"currencies"`
	TransactionTypes []*entities.TransactionType `json:"transactionTypes"`
}
