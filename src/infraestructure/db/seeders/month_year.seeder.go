package seeders

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
)

func MonthYear(dbConnection *infra_db.DBConnections) {
	db := dbConnection.DBConn

	var monthYears []entities.MonthYear

	for year := 2010; year <= 2090; year++ {
		for month := 1; month <= 12; month++ {
			monthYears = append(monthYears, entities.MonthYear{
				Month: month,
				Year:  year,
			})
		}
	}

	SeedDataConcurrent(db, monthYears, "month,year", func(m entities.MonthYear) interface{} {
		return []interface{}{m.Month, m.Year}
	})
}
