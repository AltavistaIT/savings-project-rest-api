package responses

import (
	"github.com/ssssshel/sp-api/src/domain/aggregates"
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type GetTableByIdResponse struct {
	Table        *entities.Table                       `json:"table"`
	Transactions []*aggregates.TransactionWithPosition `json:"transactions"`
}
