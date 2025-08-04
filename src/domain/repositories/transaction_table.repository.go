package repositories

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type TransactionTableRepository interface {
	CreateTransactionTable(transactionTable *entities.TableTransaction) (*entities.TableTransaction, error)
	DeleteTransactionTableByTxID(txID uint64) error
}
