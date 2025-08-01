package repositories

import (
	"github.com/ssssshel/sp-api/src/domain/aggregates"
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type TransactionRepository interface {
	GetTransactionsByTableID(tableID uint64) ([]*aggregates.TransactionWithPosition, error)
	CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
	UpdateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
	DeleteTransaction(id uint64) error
}
