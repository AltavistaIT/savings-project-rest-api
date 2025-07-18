package usecases_table

import (
	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"github.com/ssssshel/sp-api/src/domain/responses"
)

type GetTableByParamsUsecase interface {
	Execute(table *dtos.GetTableByParamsDto) (*responses.GetTableResponse, error)
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

func (u *getTableByParamsUsecase) Execute(getTableModel *dtos.GetTableByParamsDto) (*responses.GetTableResponse, error) {
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
