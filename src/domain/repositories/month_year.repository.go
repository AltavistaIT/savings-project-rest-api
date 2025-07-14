package repositories

import (
	"github.com/ssssshel/sp-api/src/domain/responses"
)

type MonthYearRepository interface {
	GetConfigDates() (*responses.GetConfigDatesResponse, error)
}
