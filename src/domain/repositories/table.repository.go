package repositories

import (
	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/domain/entities"
)

type TableRepository interface {
	CreateTable(table *entities.Table) (*entities.Table, error)
	GetTableByParams(table *dtos.GetTableByParamsDto) (*entities.Table, error)
	GetTableById(id uint64) (*entities.Table, error)
}
