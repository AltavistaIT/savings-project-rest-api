package transactions

import (
	"log"

	"github.com/ssssshel/sp-api/src/config"
	"github.com/ssssshel/sp-api/src/models"
	dto "github.com/ssssshel/sp-api/src/validators/transactions"
)

func CreateTransaction(tx *dto.CreateTransactionRequest) (*models.Transaction, error) {
	db := config.DBConn

	var lastTransaction models.TableTransaction

	db.Order("position DESC").First(&lastTransaction)

	log.Printf("Last transaction position: %d", lastTransaction.Position)

	expectedPosition := lastTransaction.Position + 1

	transaction := &models.Transaction{
		Description: tx.Description,
		TypeID:      tx.TypeID,
		Budget:      tx.Budget,
		CurrencyID:  tx.CurrencyID,
	}

	if err := db.Create(&transaction).Error; err != nil {
		return nil, err
	}

	tableTransaction := &models.TableTransaction{
		TableID:       tx.TableID,
		TransactionID: transaction.ID,
		Position:      expectedPosition,
	}

	if err := db.Create(&tableTransaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}
