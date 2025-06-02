package handler_transaction

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/ssssshel/sp-api/src/domain/models"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	"github.com/ssssshel/sp-api/src/shared/logger"
	usecases_transaction "github.com/ssssshel/sp-api/src/usecases/Transaction"
)

type TransactionHandler interface {
	GetByTableId(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type transactionHandler struct {
	getTransactionsByTableId usecases_transaction.GetTransactionsByTableIdUsecase
	createTransactionUsecase usecases_transaction.CreateTransactionUsecase
}

func NewTransactionHandler(createTransactionUsecase usecases_transaction.CreateTransactionUsecase, getTransactionsByTableId usecases_transaction.GetTransactionsByTableIdUsecase) TransactionHandler {
	return &transactionHandler{
		getTransactionsByTableId: getTransactionsByTableId,
		createTransactionUsecase: createTransactionUsecase,
	}
}

func (h *transactionHandler) GetByTableId(w http.ResponseWriter, r *http.Request) {
	tableIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	tableId, err := strconv.ParseUint(tableIdStr, 10, 64)

	var payload models.GetTransactionsByTableIDModel
	payload.TableID = tableId

	if err != nil || validator.New().Struct(payload) != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, handler.HttpMessage[http.StatusBadRequest], errors.New("invalid request"))
		return
	}

	transactions, err := h.getTransactionsByTableId.Execute(payload.TableID)
	logger.Info("%+v", transactions)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, handler.HttpMessage[http.StatusInternalServerError], err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusOK, handler.HttpMessage[http.StatusOK], transactions)
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
