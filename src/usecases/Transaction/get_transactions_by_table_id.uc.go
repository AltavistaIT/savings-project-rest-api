package usecases_transaction

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type GetTransactionsByTableIdUsecase interface {
	Execute(tableID uint64) ([]*entities.Transaction, error)
}

type getTransactionsByTableIdUsecase struct {
	transactionRepository repositories.TransactionRepository
}

func NewGetTransactionsByTableIdUsecase(txRepository repositories.TransactionRepository) GetTransactionsByTableIdUsecase {
	return &getTransactionsByTableIdUsecase{
		transactionRepository: txRepository,
	}
}

func (u *getTransactionsByTableIdUsecase) Execute(tableID uint64) ([]*entities.Transaction, error) {
	return u.transactionRepository.GetTransactionsByTableID(tableID)
}
