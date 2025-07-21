package handler_transaction

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	usecases_transaction "github.com/ssssshel/sp-api/src/usecases/Transaction"
)

type TransactionHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type transactionHandler struct {
	createTransactionUsecase usecases_transaction.CreateTransactionUsecase
	updateTransactionUsecase usecases_transaction.UpdateTransactionUsecase
	deleteTransactionUsecase usecases_transaction.DeleteTransactionUsecase
}

func NewTransactionHandler(createTransactionUsecase usecases_transaction.CreateTransactionUsecase, updateTransactionUsecase usecases_transaction.UpdateTransactionUsecase, deleteTransactionUsecase usecases_transaction.DeleteTransactionUsecase) TransactionHandler {
	return &transactionHandler{
		createTransactionUsecase: createTransactionUsecase,
		updateTransactionUsecase: updateTransactionUsecase,
		deleteTransactionUsecase: deleteTransactionUsecase,
	}
}

func (h *transactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var payload dtos.CreateTransactionDto

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.New().Struct(payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	createdTransaction, err := h.createTransactionUsecase.Execute(&payload)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusCreated, handler.HttpMessage[http.StatusCreated], createdTransaction)
}

func (h *transactionHandler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	transactionIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	transactionId, err := strconv.ParseUint(transactionIdStr, 10, 64)

	if err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var payload dtos.UpdateTransactionDto

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	payload.ID = transactionId

	if err := validator.New().Struct(payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}
	updatedTransaction, err := h.updateTransactionUsecase.Execute(&payload)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusOK, handler.HttpMessage[http.StatusOK], updatedTransaction)
}

func (h *transactionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	transactionIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	transactionId, err := strconv.ParseUint(transactionIdStr, 10, 64)

	if err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	err = h.deleteTransactionUsecase.Execute(transactionId)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusOK, handler.HttpMessage[http.StatusOK], nil)
}
