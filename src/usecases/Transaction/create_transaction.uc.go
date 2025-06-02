package usecases_transaction

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/models"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"gorm.io/gorm"
)

type CreateTransactionUsecase interface {
	Execute(transaction *models.CreateTransactionModel) (*entities.Transaction, error)
}

type createTransactionUseCase struct {
	transactionRepository      repositories.TransactionRepository
	transactionTableRepository repositories.TransactionTableRepository
}

func NewCreateTransactionUsecase(txRepository repositories.TransactionRepository, txTableRepository repositories.TransactionTableRepository) CreateTransactionUsecase {
	return &createTransactionUseCase{
		transactionRepository:      txRepository,
		transactionTableRepository: txTableRepository,
	}
}

func (uc *createTransactionUseCase) Execute(transaction *models.CreateTransactionModel) (*entities.Transaction, error) {
	var expectedPosition int
	lastTxTable, err := uc.transactionTableRepository.GetLastTransactionTableByTableID(transaction.TableID)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if lastTxTable != nil {
		expectedPosition = lastTxTable.Position + 1
	} else {
		expectedPosition = 1
	}

	createdTx, err := uc.transactionRepository.CreateTransaction(&entities.Transaction{
		Description: transaction.Description,
		TypeID:      transaction.TypeID,
		Budget:      transaction.Budget,
		CurrencyID:  transaction.CurrencyID,
	})

	if err != nil {
		return nil, err
	}

	if _, err := uc.transactionTableRepository.CreateTransactionTable(&entities.TableTransaction{
		TableID:       transaction.TableID,
		TransactionID: createdTx.ID,
		Position:      expectedPosition,
	}); err != nil {
		return nil, err
	}

	return createdTx, nil
}
