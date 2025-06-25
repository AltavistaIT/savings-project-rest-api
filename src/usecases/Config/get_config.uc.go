package usecases_config

import (
	"github.com/ssssshel/sp-api/src/domain/aggregates"
	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type GetConfigUsecase interface {
	Execute() (*aggregates.Config, error)
}

type getConfigUsecase struct {
	currencyRepository        repositories.CurrencyRepository
	transactionTypeRepository repositories.TransactionTypeRepository
}

func NewGetConfigUsecase(currencyRepository repositories.CurrencyRepository, transactionTypeRepository repositories.TransactionTypeRepository) GetConfigUsecase {
	return &getConfigUsecase{
		currencyRepository:        currencyRepository,
		transactionTypeRepository: transactionTypeRepository,
	}
}

func (uc *getConfigUsecase) Execute() (*aggregates.Config, error) {
	currencies, err := uc.currencyRepository.GetAll()

	if err != nil {
		return nil, err
	}

	transactionTypes, err := uc.transactionTypeRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return &aggregates.Config{
		Currencies:       currencies,
		TransactionTypes: transactionTypes,
	}, nil
}
