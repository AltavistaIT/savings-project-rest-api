package usecases_auth

import (
	"errors"

	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/domain/entities"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUsecase interface {
	Execute(payload *dtos.RegisterDto) error
}

type registerUsecase struct {
	userRepository repositories.UserRepository
}

func NewRegisterUsecase(userRepository repositories.UserRepository) RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
	}
}

func (u *registerUsecase) Execute(payload *dtos.RegisterDto) error {
	existentUser, err := u.userRepository.GetUserByEmail(payload.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existentUser != nil {
		return errors.New("user already exists")
	}

	hashedPsw, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
	if err != nil {
		return err
	}
	payload.Password = string(hashedPsw)

	_, err = u.userRepository.CreateUser(&entities.User{
		Name:     payload.Name,
		Surname:  payload.Surname,
		Email:    payload.Email,
		Password: payload.Password,
	})

	return err
}
