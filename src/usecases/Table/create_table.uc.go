package usecases_table

import (
	"time"

	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/models"
	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type CreateTableUsecase interface {
	Execute(table *models.CreateTableModel) (*entities.Table, error)
}

type createTableUsecase struct {
	repository repositories.TableRepository
}

func NewCreateTableUsecase(repository repositories.TableRepository) CreateTableUsecase {
	return &createTableUsecase{
		repository: repository,
	}
}

func (uc *createTableUsecase) Execute(table *models.CreateTableModel) (*entities.Table, error) {
	monthYear, err := time.Parse("2006-01", table.MonthYear)
	if err != nil {
		return nil, err
	}

	createdTable, err := uc.repository.CreateTable(&entities.Table{
		UserID:    table.UserID,
		TypeID:    table.TypeID,
		MonthYear: monthYear,
	})

	if err != nil {
		return nil, err
	}

	return createdTable, nil
}
