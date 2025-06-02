package usecases_user

import (
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type GetUserByIdUsecase interface {
	Execute(id uint64) (*entities.User, error)
}

type getUserByIdUsecase struct {
	repository repositories.UserRepository
}

func NewGetUserByIdUsecase(repository repositories.UserRepository) GetUserByIdUsecase {
	return &getUserByIdUsecase{
		repository: repository,
	}
}

func (u *getUserByIdUsecase) Execute(id uint64) (*entities.User, error) {
	return u.repository.GetUserById(id)
}
