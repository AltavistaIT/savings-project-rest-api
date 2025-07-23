package infra_db

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"gorm.io/gorm"
)

type transactionTableRepository struct {
	db *gorm.DB
}

func NewTransactionTableRepository(db *gorm.DB) repositories.TransactionTableRepository {
	return &transactionTableRepository{
		db: db,
	}
}

func (r *transactionTableRepository) GetLastTransactionTableByTableID(tableID uint64) (*entities.TableTransaction, error) {
	var transactionTable entities.TableTransaction
	return &transactionTable, r.db.Where("table_id = ?", tableID).Order("position DESC").First(&transactionTable).Error
}

func (r *transactionTableRepository) CreateTransactionTable(transactionTable *entities.TableTransaction) (*entities.TableTransaction, error) {
	return transactionTable, r.db.Create(transactionTable).Error
}

func (r *transactionTableRepository) DeleteTransactionTableByTxID(txID uint64) error {
	result := r.db.Where("transaction_id = ?", txID).Delete(&entities.TableTransaction{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
