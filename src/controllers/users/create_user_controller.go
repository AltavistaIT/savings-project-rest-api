package users

import (
	"encoding/json"
	"net/http"

	service "github.com/ssssshel/sp-api/src/services/users"
	"github.com/ssssshel/sp-api/src/utils"
	"github.com/ssssshel/sp-api/src/validators"
	dto "github.com/ssssshel/sp-api/src/validators/users"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.HandleHttpError(w, http.StatusBadRequest, utils.HttpMessage[http.StatusBadRequest], err)
		return
	}

	if err := validators.ValidateRequest(&request); err != nil {
		utils.HandleHttpError(w, http.StatusBadRequest, utils.HttpMessage[http.StatusBadRequest], err)
		return
	}

	user, err := service.CreateUser(&request)
	if err != nil {
		utils.HandleHttpError(w, http.StatusInternalServerError, utils.HttpMessage[http.StatusInternalServerError], err)
		return
	}

	utils.HandleHttpSuccess(w, http.StatusCreated, "User created successfully", map[string]interface{}{"user_id": user.ID})
}
