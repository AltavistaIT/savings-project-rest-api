package seeders

import (
	"github.com/ssssshel/sp-api/src/config"
	"github.com/ssssshel/sp-api/src/models"
)

func TableTypes() {
	db := config.DBConn

	tableTypes := []models.TableType{
		{Key: "invoices", Description: "Invoices"},
		{Key: "expenses", Description: "Expenses"},
		{Key: "savings", Description: "Savings"},
	}

	SeedData(db, tableTypes, "key", func(tableType models.TableType) interface{} {
		return tableType.Key
	})
}
