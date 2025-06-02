package repositories

import "github.com/ssssshel/sp-api/src/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUserById(id uint64) (*entities.User, error)
}
