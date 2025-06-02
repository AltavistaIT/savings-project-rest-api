package infra_db

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *entities.User) (*entities.User, error) {
	return user, r.db.Create(user).Error
}

func (r *userRepository) GetUserById(id uint64) (*entities.User, error) {
	var user entities.User
	return &user, r.db.Where("id = ?", id).First(&user).Error
}
