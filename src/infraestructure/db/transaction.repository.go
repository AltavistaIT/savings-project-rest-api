package infra_db

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
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
		Table("transactions").
		Select("transactions.*").
		Joins("JOIN table_transactions ON table_transactions.transaction_id = transactions.id").
		Where("table_transactions.table_id = ?", tableID).
		Scan(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepository) CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	return transaction, r.db.Create(transaction).Error
}

func (r *transactionRepository) UpdateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	return transaction, r.db.Model(transaction).Updates(transaction).Error

}

func (r *transactionRepository) DeleteTransaction(id uint64) error {
	result := r.db.Where("id = ?", id).Delete(&entities.Transaction{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}
