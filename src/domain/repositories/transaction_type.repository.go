package repositories

import "github.com/ssssshel/sp-api/src/domain/entities"

type TransactionTypeRepository interface {
	GetAll() ([]*entities.TransactionType, error)
}
