package config

import (
	"fmt"
	"log"
	"reflect"

	"github.com/ssssshel/sp-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func DBConnection() {
	config := DatabaseConfig()

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Lima search_path=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.Schema)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database => ", err)
	}
	log.Println("Database connection successful")
}

func MigrateDB() {
	modelInstances := []interface{}{
		models.TransactionType{},
		models.ReportType{},
		models.TableType{},
		models.Currency{},
		models.Period{},
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

	err := DBConn.AutoMigrate(validModels...)
	if err != nil {
		log.Fatal("Error migrating models => ", err)
	}

	log.Println("Models migrated successfully")
}
