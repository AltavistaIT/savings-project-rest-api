package infra_db

import (
	"fmt"

	"github.com/ssssshel/sp-api/src/shared/config"
	"github.com/ssssshel/sp-api/src/shared/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnections struct {
	DBConn *gorm.DB
}

func InitConnections() (*DBConnections, error) {
	config := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Lima search_path=%s", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSchema)
	DBConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Fatal("Error connecting to database => ", err)
		return nil, err
	}
	logger.Info("Database connection successful")

	return &DBConnections{DBConn: DBConn}, nil
}
