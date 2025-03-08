package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/ssssshel/sp-api/src/services"
	"github.com/ssssshel/sp-api/src/utils"
	"github.com/ssssshel/sp-api/src/validators"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	request := validators.GetUserRequest{
		ID: userId,
	}

	if err != nil || validators.ValidateGetUserRequest(&request) != nil {
		utils.HandleHttpError(w, http.StatusBadRequest, "Invalid request", errors.New("invalid request"))
		return
	}

	user, err := services.GetUser(&request)
	if err != nil {
		utils.HandleHttpError(w, http.StatusInternalServerError, "Error getting user", err)
		return
	}

	utils.HandleHttpSuccess(w, http.StatusOK, "User retrieved successfully", user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request validators.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.HandleHttpError(w, http.StatusBadRequest, "Invalid request", err)
		return
	}

	if err := validators.ValidateCreateUserRequest(&request); err != nil {
		utils.HandleHttpError(w, http.StatusBadRequest, "Invalid request", err)
		return
	}

	user, err := services.CreateUser(&request)
	if err != nil {
		utils.HandleHttpError(w, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	utils.HandleHttpSuccess(w, http.StatusCreated, "User created successfully", user)
}
