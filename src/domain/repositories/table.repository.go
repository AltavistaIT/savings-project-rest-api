package repositories

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/models"
)

type TableRepository interface {
	CreateTable(table *models.CreateTableModel) (*entities.Table, error)
	GetTableById(id uint64) (*entities.Table, error)
	UpdateTableAmount(table *models.UpdateTableAmountModel) (*entities.Table, error)
}
