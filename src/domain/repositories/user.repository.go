package repositories

import "github.com/ssssshel/sp-api/src/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
}
