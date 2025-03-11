package users

import (
	"github.com/ssssshel/sp-api/src/config"
	"github.com/ssssshel/sp-api/src/models"
	dto "github.com/ssssshel/sp-api/src/validators/users"
)

func GetUser(user *dto.GetUserRequest) (*models.User, error) {
	db := config.DBConn

	userModel := &models.User{}

	err := db.Where("id = ?", user.ID).First(&userModel).Error

	if err != nil {
		return nil, err
	}

	return userModel, nil
}
