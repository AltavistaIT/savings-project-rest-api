package seeders

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
)

func Currencies(dbConnection *infra_db.DBConnections) {
	db := dbConnection.DBConn

	currencies := []entities.Currency{
		{Description: "US Dollar", Abbreviation: "USD", Symbol: "$", Status: true},
		{Description: "Euro", Abbreviation: "EUR", Symbol: "€", Status: true},
		{Description: "Peruvian Sol", Abbreviation: "PEN", Symbol: "S/", Status: true},
	}

	SeedData(db, currencies, "abbreviation", func(currency entities.Currency) interface{} {
		return currency.Abbreviation
	})
}
