package aggregates

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type TableWithTransactions struct {
	Table        *entities.Table            `json:"table"`
	Transactions []*TransactionWithPosition `json:"transactions"`
}
