package services

import (
	"github.com/ssssshel/sp-api/src/config"
	"github.com/ssssshel/sp-api/src/models"
	"github.com/ssssshel/sp-api/src/validators"
)

func GetUser(user *validators.GetUserRequest) (*models.User, error) {
	db := config.DBConn

	userModel := &models.User{}

	err := db.Where("id = ?", user.ID).First(&userModel).Error

	if err != nil {
		return nil, err
	}

	return userModel, nil
}

func CreateUser(user *validators.CreateUserRequest) (*models.User, error) {
	db := config.DBConn

	userModel := &models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Status:   true,
	}

	err := db.Raw(
		"INSERT INTO public.users (username, email, password, status) VALUES (?, ?, ?, ?)",
		userModel.Username,
		userModel.Email,
		userModel.Password,
		userModel.Status,
	).Scan(&userModel.ID).Error

	if err != nil {
		return nil, err
	}

	return userModel, nil
}
