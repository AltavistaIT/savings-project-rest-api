package usecases_transaction

import (
	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type CreateTransactionUsecase interface {
	Execute(transaction *dtos.CreateTransactionDto) (*entities.Transaction, error)
}

type createTransactionUsecase struct {
	transactionRepository      repositories.TransactionRepository
	transactionTableRepository repositories.TransactionTableRepository
	tableRepository            repositories.TableRepository
}

func NewCreateTransactionUsecase(txRepository repositories.TransactionRepository, txTableRepository repositories.TransactionTableRepository, tableRepository repositories.TableRepository) CreateTransactionUsecase {
	return &createTransactionUsecase{
		transactionRepository:      txRepository,
		transactionTableRepository: txTableRepository,
		tableRepository:            tableRepository,
	}
}

func (uc *createTransactionUsecase) Execute(transaction *dtos.CreateTransactionDto) (*entities.Transaction, error) {
	createdTx, err := uc.transactionRepository.CreateTransaction(&entities.Transaction{
		Description: transaction.Description,
		TypeID:      transaction.TypeID,
		Amount:      transaction.Amount,
		CurrencyID:  transaction.CurrencyID,
		Date:        transaction.Date,
	})

	if err != nil {
		return nil, err
	}

	if _, err := uc.transactionTableRepository.CreateTransactionTable(&entities.TableTransaction{
		TableID:       transaction.TableID,
		TransactionID: createdTx.ID,
	}); err != nil {
		return nil, err
	}

	return createdTx, nil
}
