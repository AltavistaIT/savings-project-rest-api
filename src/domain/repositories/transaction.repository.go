package repositories

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type TransactionRepository interface {
	CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
}
