package usecases_config

import (
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"github.com/ssssshel/sp-api/src/domain/responses"
)

type GetConfigUsecase interface {
	Execute() (*responses.GetConfigResponse, error)
}

type getConfigUsecase struct {
	currencyRepository        repositories.CurrencyRepository
	transactionTypeRepository repositories.TransactionTypeRepository
	monthYearRepository       repositories.MonthYearRepository
	tableTypeRepository       repositories.TableTypeRepository
}

func NewGetConfigUsecase(currencyRepository repositories.CurrencyRepository, transactionTypeRepository repositories.TransactionTypeRepository, monthYearRepository repositories.MonthYearRepository, tableTypeRepository repositories.TableTypeRepository) GetConfigUsecase {
	return &getConfigUsecase{
		currencyRepository:        currencyRepository,
		transactionTypeRepository: transactionTypeRepository,
		monthYearRepository:       monthYearRepository,
		tableTypeRepository:       tableTypeRepository,
	}
}

func (uc *getConfigUsecase) Execute() (*responses.GetConfigResponse, error) {
	currencies, err := uc.currencyRepository.GetAll()
	if err != nil {
		return nil, err
	}

	transactionTypes, err := uc.transactionTypeRepository.GetAll()
	if err != nil {
		return nil, err
	}

	monthYears, err := uc.monthYearRepository.GetConfigDates()
	if err != nil {
		return nil, err
	}

	tableTypes, err := uc.tableTypeRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return &responses.GetConfigResponse{
		Currencies:       currencies,
		TransactionTypes: transactionTypes,
		MonthYears:       monthYears,
		TableTypes:       tableTypes,
	}, nil
}
