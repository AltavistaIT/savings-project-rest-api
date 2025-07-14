package usecases_table

import (
	"github.com/ssssshel/sp-api/src/domain/models"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"github.com/ssssshel/sp-api/src/domain/responses"
)

type GetTableByParamsUsecase interface {
	Execute(table *models.GetTableByParamsModel) (*responses.GetTableResponse, error)
}

type getTableByParamsUsecase struct {
	tableRepository       repositories.TableRepository
	transactionRepository repositories.TransactionRepository
}

func NewGetTableByParamsUsecase(tableRepository repositories.TableRepository, transactionRepository repositories.TransactionRepository) GetTableByParamsUsecase {
	return &getTableByParamsUsecase{
		tableRepository:       tableRepository,
		transactionRepository: transactionRepository,
	}
}

func (u *getTableByParamsUsecase) Execute(getTableModel *models.GetTableByParamsModel) (*responses.GetTableResponse, error) {
	table, err := u.tableRepository.GetTableByParams(getTableModel)
	if err != nil {
		return nil, err
	}

	transactions, err := u.transactionRepository.GetTransactionsByTableID(table.ID)
	if err != nil {
		return nil, err
	}

	return &responses.GetTableResponse{
		Table:        table,
		Transactions: transactions,
	}, nil
}
