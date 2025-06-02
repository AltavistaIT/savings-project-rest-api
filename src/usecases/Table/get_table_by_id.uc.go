package usecases_table

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type GetTableByIdUsecase interface {
	Execute(id uint64) (*entities.Table, error)
}

type getTableByIdUsecase struct {
	tableRepository repositories.TableRepository
}

func NewGetTableByIdUsecase(tableRepository repositories.TableRepository) GetTableByIdUsecase {
	return &getTableByIdUsecase{
		tableRepository: tableRepository,
	}
}

func (u *getTableByIdUsecase) Execute(id uint64) (*entities.Table, error) {
	return u.tableRepository.GetTableById(id)
}
