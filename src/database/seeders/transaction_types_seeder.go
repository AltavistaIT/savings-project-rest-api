package seeders

import (
	"log"

	"github.com/ssssshel/sp-api/src/config"
	"github.com/ssssshel/sp-api/src/models"
)

func TransactionTypes() {
	db := config.DBConn

	tableTypes := []models.TableType{}

	if err := db.Find(&tableTypes).Error; err != nil {
		log.Fatal("Error finding table types => ", err)
	}
	log.Println(tableTypes)

	tableTypeMap := make(map[string]*models.TableType)
	for i := range tableTypes {
		tableTypeMap[tableTypes[i].Key] = &tableTypes[i]
	}

	transactionTypes := []models.TransactionType{
		{Key: "transport", Description: "Transport", TableType: tableTypeMap["expenses"]},
		{Key: "food", Description: "Food", TableType: tableTypeMap["expenses"]},
		{Key: "entertainment", Description: "Entertainment", TableType: tableTypeMap["expenses"]},
		{Key: "groceries", Description: "Groceries", TableType: tableTypeMap["expenses"]},
		{Key: "rent", Description: "Rent", TableType: tableTypeMap["expenses"]},
		{Key: "utilities", Description: "Utilities", TableType: tableTypeMap["invoices"]},
	}

	SeedData(db, transactionTypes, "key", func(transactionType models.TransactionType) interface{} {
		return transactionType.Key
	})
}
