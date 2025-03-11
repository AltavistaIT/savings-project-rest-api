package tables

import (
	"errors"

	"github.com/ssssshel/sp-api/src/config"
	"github.com/ssssshel/sp-api/src/models"
	dto "github.com/ssssshel/sp-api/src/validators/tables"
)

func CreateTable(table *dto.CreateTableRequest) (*models.Table, error) {
	db := config.DBConn

	var existingTable models.Table
	err := db.Where("user_id = ? AND type_id = ? AND month_year = ?", table.UserID, table.TypeID, table.MonthYear).First(&existingTable).Error

	if err == nil {
		return nil, err
	}

	if existingTable.ID != 0 {
		return nil, errors.New("table already exists")
	}

	tableModel := &models.Table{
		UserID:    table.UserID,
		TypeID:    table.TypeID,
		MonthYear: table.MonthYear,
	}

	err = db.Create(&tableModel).Error

	if err != nil {
		return nil, err
	}

	return tableModel, nil
}
