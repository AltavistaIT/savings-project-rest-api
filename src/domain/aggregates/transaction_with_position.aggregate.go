package aggregates

import "github.com/ssssshel/sp-api/src/domain/entities"

type TransactionWithPosition struct {
	*entities.Transaction
	Position int `json:"position"`
}
