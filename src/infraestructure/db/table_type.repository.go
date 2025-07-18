package infra_db

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"gorm.io/gorm"
)

type tableTypeRepository struct {
	db *gorm.DB
}

func NewTableTypeRepository(db *gorm.DB) repositories.TableTypeRepository {
	return &tableTypeRepository{
		db: db,
	}
}

func (r *tableTypeRepository) GetAll() ([]*entities.TableType, error) {
	var tableTypes []*entities.TableType
	if err := r.db.Select("id", "key", "description").Find(&tableTypes).Error; err != nil {
		return nil, err
	}
	return tableTypes, nil
}
