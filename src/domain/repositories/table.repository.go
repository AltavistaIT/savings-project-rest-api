package repositories

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/models"
)

type TableRepository interface {
	CreateTable(table *entities.Table) (*entities.Table, error)
	GetTableById(id uint64) (*entities.Table, error)
	GetTableByParams(table *models.GetTableByParamsModel) (*entities.Table, error)
	UpdateTableAmount(table *models.UpdateTableAmountModel) (*entities.Table, error)
}
