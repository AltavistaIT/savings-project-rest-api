package infra_db

import (
	"errors"

	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/models"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"gorm.io/gorm"
)

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) repositories.TableRepository {
	return &tableRepository{
		db: db,
	}
}

func (r *tableRepository) GetTableById(id uint64) (*entities.Table, error) {
	var table entities.Table
	err := r.db.Where("id = ?", id).First(&table).Error
	return &table, err
}

func (r *tableRepository) CreateTable(table *models.CreateTableModel) (*entities.Table, error) {
	tableModel := &entities.Table{
		UserID:    table.UserID,
		TypeID:    table.TypeID,
		MonthYear: table.MonthYear,
	}

	result := r.db.Where("user_id = ? AND type_id = ? AND month_year = ?", table.UserID, table.TypeID, table.MonthYear).Attrs(tableModel).FirstOrCreate(tableModel)

	if result.RowsAffected == 0 {
		return nil, errors.New("table not created")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return tableModel, nil
}
