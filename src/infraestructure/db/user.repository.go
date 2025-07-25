package infra_db

import (
	"errors"

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

func (r *userRepository) GetUserById(id uint64) (*entities.User, error) {
	var user entities.User
	return &user, r.db.Where("id = ?", id).First(&user).Error
}

func (r *userRepository) CreateUser(user *entities.User) (*entities.User, error) {
	return user, r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
