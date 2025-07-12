package seeders

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
)

func TableTypes(dbConnection *infra_db.DBConnections) {
	db := dbConnection.DBConn

	tableTypes := []entities.TableType{
		{Key: "invoices", Description: "Invoices"},
		{Key: "expenses", Description: "Expenses"},
		{Key: "savings", Description: "Savings"},
	}

	SeedData(db, tableTypes, "key", func(tableType entities.TableType) interface{} {
		return tableType.Key
	})
}
