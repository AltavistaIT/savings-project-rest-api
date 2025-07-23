package usecases_transaction

import (
	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type UpdateTransactionUsecase interface {
	Execute(transaction *dtos.UpdateTransactionDto) (*entities.Transaction, error)
}

type updateTransactionUsecase struct {
	transactionRepository repositories.TransactionRepository
}

func NewUpdateTransactionUsease(txRepository repositories.TransactionRepository) UpdateTransactionUsecase {
	return &updateTransactionUsecase{
		transactionRepository: txRepository,
	}
}

func (u *updateTransactionUsecase) Execute(transaction *dtos.UpdateTransactionDto) (*entities.Transaction, error) {

	updatedTx, err := u.transactionRepository.UpdateTransaction(&entities.Transaction{
		ID:          transaction.ID,
		Amount:      transaction.Amount,
		TypeID:      transaction.TypeID,
		Description: transaction.Description,
		Date:        transaction.Date,
	})

	if err != nil {
		return nil, err
	}

	return updatedTx, nil
}
