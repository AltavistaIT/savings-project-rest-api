package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ssssshel/sp-api/src/services"
	"github.com/ssssshel/sp-api/src/utils"
	"github.com/ssssshel/sp-api/src/validators"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a user"))
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request validators.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid request", err)
		return
	}

	if err := validators.ValidateCreateUserRequest(&request); err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid request", err)
		return
	}

	user, err := services.CreateUser(&request)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	utils.HandleSuccess(w, http.StatusCreated, "User created successfully", user)
}
