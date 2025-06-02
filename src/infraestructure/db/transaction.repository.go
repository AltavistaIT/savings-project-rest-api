package infra_db

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"github.com/ssssshel/sp-api/src/shared/logger"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) repositories.TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) GetTransactionsByTableID(tableID uint64) ([]*entities.Transaction, error) {
	var transactions []*entities.Transaction

	err := r.db.
		Joins("JOIN table_transactions ON table_transactions.transaction_id = transactions.id").
		Where("table_transactions.table_id = ?", tableID).
		Order("table_transactions.position ASC").
		Find(&transactions).Error

	logger.Info("%+v", transactions)

	return transactions, err
}

func (r *transactionRepository) CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	return transaction, r.db.Create(transaction).Error
}
