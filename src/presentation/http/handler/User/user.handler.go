package handler_user

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	usecases_user "github.com/ssssshel/sp-api/src/usecases/User"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	createUserUsecase  usecases_user.CreateUserUsecase
	getUserByIdUsecase usecases_user.GetUserByIdUsecase
}

func NewUserHandler(createUserUsecase usecases_user.CreateUserUsecase, getUserByIdUsecase usecases_user.GetUserByIdUsecase) UserHandler {
	return &userHandler{
		createUserUsecase:  createUserUsecase,
		getUserByIdUsecase: getUserByIdUsecase,
	}
}

func (h *userHandler) GetById(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	userId, err := strconv.ParseUint(userIdStr, 10, 64)

	if err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	user, err := h.getUserByIdUsecase.Execute(userId)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusOK, handler.HttpMessage[http.StatusOK], user)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var payload dtos.CreateUserDto

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.New().Struct(payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	createdUser, err := h.createUserUsecase.Execute(&payload)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusCreated, handler.HttpMessage[http.StatusCreated], createdUser)
}
