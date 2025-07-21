package usecases_transaction

import "github.com/ssssshel/sp-api/src/domain/repositories"

type DeleteTransactionUsecase interface {
	Execute(id uint64) error
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

func (u *deleteTransactionUsecase) Execute(id uint64) error {
	err := u.transactionTableRepository.DeleteTransactionTableByTxID(id)

	if err != nil {
		return err
	}

	return u.transactionRepository.DeleteTransaction(id)
}
