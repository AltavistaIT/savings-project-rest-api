package repositories

import "github.com/ssssshel/sp-api/src/domain/entities"

type CurrencyRepository interface {
	GetAll() ([]*entities.Currency, error)
}
