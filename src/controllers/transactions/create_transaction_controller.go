package transactions

import (
	"encoding/json"
	"net/http"

	service "github.com/ssssshel/sp-api/src/services/transactions"
	"github.com/ssssshel/sp-api/src/utils"
	"github.com/ssssshel/sp-api/src/validators"
	dto "github.com/ssssshel/sp-api/src/validators/transactions"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.HandleHttpError(w, http.StatusBadRequest, utils.HttpMessage[http.StatusBadRequest], err)
		return
	}

	if err := validators.ValidateRequest(&request); err != nil {
		utils.HandleHttpError(w, http.StatusBadRequest, utils.HttpMessage[http.StatusBadRequest], err)
		return
	}

	tx, err := service.CreateTransaction(&request)
	if err != nil {
		utils.HandleHttpError(w, http.StatusInternalServerError, utils.HttpMessage[http.StatusInternalServerError], err)
		return
	}

	utils.HandleHttpSuccess(w, http.StatusCreated, "Transaction created successfully", tx)
}
