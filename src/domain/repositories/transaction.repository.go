package repositories

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type TransactionRepository interface {
	GetTransactionsByTableID(tableID uint64) ([]*entities.Transaction, error)
	CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
}
