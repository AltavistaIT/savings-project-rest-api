package infra_db

import (
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"github.com/ssssshel/sp-api/src/domain/responses"
	"gorm.io/gorm"
)

type monthYearRepository struct {
	db *gorm.DB
}

func NewMonthYearRepository(db *gorm.DB) repositories.MonthYearRepository {
	return &monthYearRepository{
		db: db,
	}
}

func (r *monthYearRepository) GetConfigDates() (*responses.GetConfigDatesResponse, error) {
	var months []int
	var years []int
	if err := r.db.Table("month_years").Select("DISTINCT month").Order("month").Pluck("month", &months).Error; err != nil {
		return nil, err
	}

	if err := r.db.Table("month_years").Select("DISTINCT year").Order("year").Pluck("year", &years).Error; err != nil {
		return nil, err
	}

	return &responses.GetConfigDatesResponse{
		Months: months,
		Years:  years,
	}, nil
}
