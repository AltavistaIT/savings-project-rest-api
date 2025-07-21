package infra_db

import (
	"github.com/ssssshel/sp-api/src/domain/aggregates"
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

// TODO: no en uso
func (r *transactionRepository) GetTransactionsByTableID(tableID uint64) ([]*aggregates.TransactionWithPosition, error) {
	var results []struct {
		entities.Transaction
		Position int
	}

	err := r.db.
		Table("transactions").
		Select("transactions.*, table_transactions.position").
		Joins("JOIN table_transactions ON table_transactions.transaction_id = transactions.id").
		Where("table_transactions.table_id = ?", tableID).
		Order("table_transactions.position ASC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	transactionsWithPosition := make([]*aggregates.TransactionWithPosition, 0, len(results))
	for _, r := range results {
		tx := r.Transaction
		transactionsWithPosition = append(transactionsWithPosition, &aggregates.TransactionWithPosition{
			Transaction: &tx,
			Position:    r.Position,
		})
	}

	return transactionsWithPosition, nil
}

func (r *transactionRepository) CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	return transaction, r.db.Create(transaction).Error
}

func (r *transactionRepository) UpdateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	return transaction, r.db.Model(transaction).Updates(transaction).Error

}

func (r *transactionRepository) DeleteTransaction(id uint64) error {
	return r.db.Where("id = ?", id).Delete(&entities.Transaction{}).Error
}
