package usecases_transaction

import (
	"errors"

	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type DeleteTransactionUsecase interface {
	Execute(txId, tableId uint64) error
}

type deleteTransactionUsecase struct {
	transactionRepository      repositories.TransactionRepository
	transactionTableRepository repositories.TransactionTableRepository
}

func NewDeleteTransactionUsecase(transactionRepository repositories.TransactionRepository, transactionTableRepository repositories.TransactionTableRepository) DeleteTransactionUsecase {
	return &deleteTransactionUsecase{
		transactionRepository:      transactionRepository,
		transactionTableRepository: transactionTableRepository,
	}
}

func (u *deleteTransactionUsecase) Execute(transactionID, userID uint64) error {
	// Validate tx belongs to user
	txTable, err := u.transactionTableRepository.GetTransactionTableByTxID(transactionID)
	if err != nil {
		return err
	}

	if txTable.Table.UserID != userID {
		return errors.New("forbidden")
	}

	err = u.transactionTableRepository.DeleteTransactionTableByTxID(transactionID)

	if err != nil {
		return err
	}

	return u.transactionRepository.DeleteTransaction(transactionID)
}
