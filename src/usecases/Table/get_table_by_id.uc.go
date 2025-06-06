package usecases_table

import (
	"github.com/ssssshel/sp-api/src/domain/aggregates"
	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type GetTableByIdUsecase interface {
	Execute(id uint64) (*aggregates.TableWithTransactions, error)
}

type getTableByIdUsecase struct {
	tableRepository       repositories.TableRepository
	transactionRepository repositories.TransactionRepository
}

func NewGetTableByIdUsecase(tableRepository repositories.TableRepository, transactionRepository repositories.TransactionRepository) GetTableByIdUsecase {
	return &getTableByIdUsecase{
		tableRepository:       tableRepository,
		transactionRepository: transactionRepository,
	}
}

func (u *getTableByIdUsecase) Execute(id uint64) (*aggregates.TableWithTransactions, error) {
	table, err := u.tableRepository.GetTableById(id)
	if err != nil {
		return nil, err
	}

	transactions, err := u.transactionRepository.GetTransactionsByTableID(table.ID)

	if err != nil {
		return nil, err
	}

	return &aggregates.TableWithTransactions{
		Table:        table,
		Transactions: transactions,
	}, nil
}
