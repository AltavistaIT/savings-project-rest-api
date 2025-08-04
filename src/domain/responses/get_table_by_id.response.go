package responses

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type GetTableResponse struct {
	Table        *entities.Table         `json:"table"`
	Transactions []*entities.Transaction `json:"transactions"`
}
