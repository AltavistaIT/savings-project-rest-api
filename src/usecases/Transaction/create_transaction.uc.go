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

// Execute creates a new transaction and adds it to the specified table.
//
// It first finds the position for the new transaction by getting the last transaction table
// for the specified table id. If it doesn't exist, the position is set to 1. Otherwise,
// the position is set to the position of the last transaction plus one.
//
// The transaction is then created using the transaction repository.
//
// If the transaction is created successfully, a new transaction table is created using the
// transaction table repository, with the table id and transaction id set to the specified
// table id and the id of the newly created transaction, respectively. The position is set
// to the position determined above.
//
// If the transaction table is created successfully, the amount of the table is updated
// using the table repository, with the id set to the specified table id and the amount
// set to the specified budget.
//
// If any error occurs during the above steps, it is returned. Otherwise, the newly created
// transaction is returned.
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
