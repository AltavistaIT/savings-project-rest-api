package migrations

import (
	"log"
	"reflect"

	"github.com/ssssshel/sp-api/src/domain/entities"
	infra_db "github.com/ssssshel/sp-api/src/infraestructure/db"
	"github.com/ssssshel/sp-api/src/infraestructure/db/seeders"
)

func MigrateDB(dbConnection *infra_db.DBConnections) {
	modelInstances := []interface{}{
		entities.TableType{},
		entities.TransactionType{},
		entities.ReportType{},
		entities.Currency{},
		entities.User{},
		entities.Table{},
		entities.Transaction{},
		entities.Report{},
		entities.TableTransaction{},
		entities.MonthYear{},
	}

	validModels := []interface{}{}

	for _, model := range modelInstances {
		if reflect.TypeOf(model).Kind() == reflect.Struct {
			validModels = append(validModels, model)
		}
	}

	err := dbConnection.DBConn.AutoMigrate(validModels...)
	if err != nil {
		log.Fatal("Error migrating entities => ", err)
	}

	log.Println("Models migrated successfully")

	// Seeders
	seeders.TableTypes(dbConnection)
	seeders.TransactionTypes(dbConnection)
	seeders.Currencies(dbConnection)
	seeders.MonthYear(dbConnection)
}
