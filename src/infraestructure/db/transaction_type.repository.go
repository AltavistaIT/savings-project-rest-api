package infra_db

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"gorm.io/gorm"
)

type transactionTypeRepository struct {
	db *gorm.DB
}

func NewTransactionTypeRepository(db *gorm.DB) repositories.TransactionTypeRepository {
	return &transactionTypeRepository{
		db: db,
	}
}

func (r *transactionTypeRepository) GetAll() ([]*entities.TransactionType, error) {
	var transactionTypes []*entities.TransactionType
	err := r.db.Find(&transactionTypes).Error
	return transactionTypes, err
}
