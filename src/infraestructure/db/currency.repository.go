package infra_db

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"gorm.io/gorm"
)

type currencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) repositories.CurrencyRepository {
	return &currencyRepository{
		db: db,
	}
}

func (r *currencyRepository) GetAll() ([]*entities.Currency, error) {
	var currencies []*entities.Currency
	err := r.db.Find(&currencies).Error
	return currencies, err
}
