package handler_auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	usecases_auth "github.com/ssssshel/sp-api/src/usecases/Auth"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	loginUsecase    usecases_auth.LoginUsecase
	registerUsecase usecases_auth.RegisterUsecase
}

func NewAuthHandler(loginUsecase usecases_auth.LoginUsecase, registerUsecase usecases_auth.RegisterUsecase) AuthHandler {
	return &authHandler{
		loginUsecase:    loginUsecase,
		registerUsecase: registerUsecase,
	}
}

func (ah *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var payload dtos.LoginDto

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.New().Struct(payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	token, err := ah.loginUsecase.Execute(&payload)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusOK, handler.HttpMessage[http.StatusOK], token)
}

func (ah *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var payload dtos.RegisterDto

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.New().Struct(payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	err := ah.registerUsecase.Execute(&payload)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusCreated, handler.HttpMessage[http.StatusCreated], nil)
}
