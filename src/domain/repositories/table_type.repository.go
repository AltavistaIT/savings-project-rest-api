package repositories

import "github.com/ssssshel/sp-api/src/domain/entities"

type TableTypeRepository interface {
	GetAll() ([]*entities.TableType, error)
}
