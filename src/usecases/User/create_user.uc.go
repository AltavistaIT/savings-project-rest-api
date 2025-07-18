package usecases_user

import (
	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
)

type CreateUserUsecase interface {
	Execute(createUserRequest *dtos.CreateUserDto) (*entities.User, error)
}

type createUserUsecase struct {
	userRepository repositories.UserRepository
}

func NewCreateUserUsecase(userRepository repositories.UserRepository) CreateUserUsecase {
	return &createUserUsecase{
		userRepository: userRepository,
	}
}

func (u *createUserUsecase) Execute(createUserRequest *dtos.CreateUserDto) (*entities.User, error) {

	userEntity := &entities.User{
		Username: createUserRequest.Username,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
		Status:   true,
	}

	return u.userRepository.CreateUser(userEntity)
}
