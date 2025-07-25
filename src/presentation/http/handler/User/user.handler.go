package handler_user

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	usecases_user "github.com/ssssshel/sp-api/src/usecases/User"
)

type UserHandler interface {
	GetById(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	getUserByIdUsecase usecases_user.GetUserByIdUsecase
}

func NewUserHandler(getUserByIdUsecase usecases_user.GetUserByIdUsecase) UserHandler {
	return &userHandler{
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
