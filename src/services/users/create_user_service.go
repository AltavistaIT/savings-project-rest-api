package users

import (
	"github.com/ssssshel/sp-api/src/config"
	"github.com/ssssshel/sp-api/src/models"
	dto "github.com/ssssshel/sp-api/src/validators/users"
)

func CreateUser(user *dto.CreateUserRequest) (*models.User, error) {
	db := config.DBConn

	userModel := &models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Status:   true,
	}

	err := db.Create(&userModel).Error

	if err != nil {
		return nil, err
	}

	return userModel, nil
}
