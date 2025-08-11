package usecases_transaction

import (
	"errors"

	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type UpdateTransactionUsecase interface {
	Execute(transaction *dtos.UpdateTransactionDto) (*entities.Transaction, error)
}

type updateTransactionUsecase struct {
	transactionRepository repositories.TransactionRepository
	tableRepository       repositories.TableRepository
}

func NewUpdateTransactionUsease(txRepository repositories.TransactionRepository, tableRepository repositories.TableRepository) UpdateTransactionUsecase {
	return &updateTransactionUsecase{
		transactionRepository: txRepository,
		tableRepository:       tableRepository,
	}
}

func (u *updateTransactionUsecase) Execute(transaction *dtos.UpdateTransactionDto) (*entities.Transaction, error) {
	// Validate tx belongs to user
	table, err := u.tableRepository.GetTableById(transaction.TableID)
	if err != nil {
		return nil, err
	}

	if table.UserID != transaction.UserID {
		return nil, errors.New("forbidden")
	}

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
