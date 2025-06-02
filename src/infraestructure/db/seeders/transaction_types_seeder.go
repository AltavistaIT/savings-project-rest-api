package seeders

import (
	"log"

	"github.com/ssssshel/sp-api/src/domain/entities"
	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
)

func TransactionTypes(dbConnection *infra_db.DBConnections) {
	db := dbConnection.DBConn

	tableTypes := []entities.TableType{}

	if err := db.Find(&tableTypes).Error; err != nil {
		log.Fatal("Error finding table types => ", err)
	}

	tableTypeMap := make(map[string]*entities.TableType)
	for i := range tableTypes {
		tableTypeMap[tableTypes[i].Key] = &tableTypes[i]
	}

	transactionTypes := []entities.TransactionType{
		{Key: "transport", Description: "Transport", TableType: tableTypeMap["expenses"]},
		{Key: "food", Description: "Food", TableType: tableTypeMap["expenses"]},
		{Key: "entertainment", Description: "Entertainment", TableType: tableTypeMap["expenses"]},
		{Key: "groceries", Description: "Groceries", TableType: tableTypeMap["expenses"]},
		{Key: "rent", Description: "Rent", TableType: tableTypeMap["expenses"]},
		{Key: "utilities", Description: "Utilities", TableType: tableTypeMap["invoices"]},
	}

	SeedData(db, transactionTypes, "key", func(transactionType entities.TransactionType) interface{} {
		return transactionType.Key
	})
}
