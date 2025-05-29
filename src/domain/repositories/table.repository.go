package repositories

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/models"
)

type TableRepository interface {
	CreateTable(table *models.CreateTableModel) (*entities.Table, error)
}
