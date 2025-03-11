package migrations

import (
	"log"
	"reflect"

	"github.com/ssssshel/sp-api/src/config"
	"github.com/ssssshel/sp-api/src/database/seeders"
	"github.com/ssssshel/sp-api/src/models"
)

func MigrateDB() {
	modelInstances := []interface{}{
		models.TransactionType{},
		models.ReportType{},
		models.TableType{},
		models.Currency{},
		models.User{},
		models.Table{},
		models.Transaction{},
		models.Report{},
		models.TableTransaction{},
	}

	validModels := []interface{}{}

	for _, model := range modelInstances {
		if reflect.TypeOf(model).Kind() == reflect.Struct {
			validModels = append(validModels, model)
		}
	}

	err := config.DBConn.AutoMigrate(validModels...)
	if err != nil {
		log.Fatal("Error migrating models => ", err)
	}

	log.Println("Models migrated successfully")

	// Seeders
	seeders.TableTypes()
}
