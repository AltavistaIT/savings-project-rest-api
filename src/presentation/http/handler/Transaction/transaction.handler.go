package handler_transaction

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ssssshel/sp-api/src/domain/models"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	usecases_transaction "github.com/ssssshel/sp-api/src/usecases/Transaction"
)

type TransactionHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type transactionHandler struct {
	createTransactionUsecase usecases_transaction.CreateTransactionUsecase
}

func NewTransactionHandler(createTransactionUsecase usecases_transaction.CreateTransactionUsecase) TransactionHandler {
	return &transactionHandler{
		createTransactionUsecase: createTransactionUsecase,
	}
}

func (h *transactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var payload models.CreateTransactionModel

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, handler.HttpMessage[http.StatusBadRequest], err)
		return
	}

	if err := validator.New().Struct(payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, handler.HttpMessage[http.StatusBadRequest], err)
		return
	}

	createdTransaction, err := h.createTransactionUsecase.Execute(&payload)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, handler.HttpMessage[http.StatusInternalServerError], err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusCreated, handler.HttpMessage[http.StatusCreated], createdTransaction)
}
