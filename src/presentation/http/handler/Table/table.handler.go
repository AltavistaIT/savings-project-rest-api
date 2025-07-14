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
	GetByParams(w http.ResponseWriter, r *http.Request)
}

type tableHandler struct {
	createTableUsecase      usecases_table.CreateTableUsecase
	getTableByIdUsecase     usecases_table.GetTableByIdUsecase
	getTableByParamsUsecase usecases_table.GetTableByParamsUsecase
}

func NewTableHandler(createTableUsecase usecases_table.CreateTableUsecase, getTableByIdUsecase usecases_table.GetTableByIdUsecase, getTableByParamsUsecase usecases_table.GetTableByParamsUsecase) TableHandler {
	return &tableHandler{
		createTableUsecase:      createTableUsecase,
		getTableByIdUsecase:     getTableByIdUsecase,
		getTableByParamsUsecase: getTableByParamsUsecase,
	}
}

func (h *tableHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var payload models.CreateTableModel

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.New().Struct(payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	createdTable, err := h.createTableUsecase.Execute(&payload)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusCreated, handler.HttpMessage[http.StatusCreated], createdTable)
}

func (h *tableHandler) GetById(w http.ResponseWriter, r *http.Request) {
	tableIdStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	tableId, err := strconv.ParseUint(tableIdStr, 10, 64)

	var payload models.GetTableByIdModel
	payload.ID = tableId

	if err != nil || validator.New().Struct(payload) != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, errors.New("invalid request"))
		return
	}

	table, err := h.getTableByIdUsecase.Execute(payload.ID)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusOK, handler.HttpMessage[http.StatusOK], table)
}

func (h *tableHandler) GetByParams(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	userIDStr := query.Get("user_id")
	typeIDStr := query.Get("type_id")
	monthYear := query.Get("month_year")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, errors.New("invalid user_id"))
		return
	}

	typeID, err := strconv.ParseUint(typeIDStr, 10, 64)
	if err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, errors.New("invalid type_id"))
		return
	}

	model := &models.GetTableByParamsModel{
		UserID:    userID,
		TypeID:    typeID,
		MonthYear: monthYear,
	}

	if err := validator.New().Struct(model); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	model.MonthYear = model.MonthYear + "-01"
	table, err := h.getTableByParamsUsecase.Execute(model)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusOK, handler.HttpMessage[http.StatusOK], table)
}
