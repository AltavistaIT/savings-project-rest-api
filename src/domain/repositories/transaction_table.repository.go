package repositories

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type TransactionTableRepository interface {
	GetLastTransactionTableByTableID(tableID uint64) (*entities.TableTransaction, error)
	CreateTransactionTable(transactionTable *entities.TableTransaction) (*entities.TableTransaction, error)
	DeleteTransactionTableByTxID(txID uint64) error
}
