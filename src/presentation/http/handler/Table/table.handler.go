package handler_table

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/ssssshel/sp-api/src/domain/models"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	usecases_table "github.com/ssssshel/sp-api/src/usecases/Table"
)

type TableHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type tableHandler struct {
	createTableUsecase  usecases_table.CreateTableUsecase
	getTableByIdUsecase usecases_table.GetTableByIdUsecase
}

func NewTableHandler(createTableUsecase usecases_table.CreateTableUsecase, getTableByIdUsecase usecases_table.GetTableByIdUsecase) TableHandler {
	return &tableHandler{
		createTableUsecase:  createTableUsecase,
		getTableByIdUsecase: getTableByIdUsecase,
	}
}

func (h *tableHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var payload models.CreateTableModel

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, handler.HttpMessage[http.StatusBadRequest], err)
		return
	}

	if err := validator.New().Struct(payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, handler.HttpMessage[http.StatusBadRequest], err)
		return
	}

	createdTable, err := h.createTableUsecase.Execute(&payload)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, handler.HttpMessage[http.StatusInternalServerError], err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusCreated, handler.HttpMessage[http.StatusCreated], createdTable)
}

func (h *tableHandler) GetById(w http.ResponseWriter, r *http.Request) {
	tableIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	tableId, err := strconv.ParseUint(tableIdStr, 10, 64)

	var payload models.GetTableModel
	payload.ID = tableId

	if err != nil || validator.New().Struct(payload) != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, handler.HttpMessage[http.StatusBadRequest], errors.New("invalid request"))
		return
	}

	table, err := h.getTableByIdUsecase.Execute(payload.ID)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, handler.HttpMessage[http.StatusInternalServerError], err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusOK, handler.HttpMessage[http.StatusOK], table)
}
