package users

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	service "github.com/ssssshel/sp-api/src/services/users"
	"github.com/ssssshel/sp-api/src/utils"
	"github.com/ssssshel/sp-api/src/validators"
	dto "github.com/ssssshel/sp-api/src/validators/users"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	request := dto.GetUserRequest{
		ID: userId,
	}

	if err != nil || validators.ValidateRequest(&request) != nil {
		utils.HandleHttpError(w, http.StatusBadRequest, utils.HttpMessage[http.StatusBadRequest], errors.New("invalid request"))
		return
	}

	user, err := service.GetUser(&request)
	if err != nil {
		utils.HandleHttpError(w, http.StatusInternalServerError, utils.HttpMessage[http.StatusInternalServerError], err)
		return
	}

	utils.HandleHttpSuccess(w, http.StatusOK, "User retrieved successfully", user)
}
