package usecases_auth

import (
	"errors"

	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/domain/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginUsecase interface {
	Execute(payload *dtos.LoginDto) (string, error)
}

type loginUsecase struct {
	userRepository repositories.UserRepository
	jwtUsecase     JWTUsecase
}

func NewLoginUsecase(userRepository repositories.UserRepository, jwtUsecase JWTUsecase) LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		jwtUsecase:     jwtUsecase,
	}
}

func (u *loginUsecase) Execute(payload *dtos.LoginDto) (string, error) {
	user, err := u.userRepository.GetUserByEmail(payload.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid credentials")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := u.jwtUsecase.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
