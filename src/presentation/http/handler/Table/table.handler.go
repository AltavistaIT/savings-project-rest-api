package handler_table

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/ssssshel/sp-api/src/domain/dtos"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	usecases_table "github.com/ssssshel/sp-api/src/usecases/Table"
)

type TableHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByParams(w http.ResponseWriter, r *http.Request)
}

type tableHandler struct {
	createTableUsecase      usecases_table.CreateTableUsecase
	getTableByParamsUsecase usecases_table.GetTableByParamsUsecase
}

func NewTableHandler(createTableUsecase usecases_table.CreateTableUsecase, getTableByParamsUsecase usecases_table.GetTableByParamsUsecase) TableHandler {
	return &tableHandler{
		createTableUsecase:      createTableUsecase,
		getTableByParamsUsecase: getTableByParamsUsecase,
	}
}

func (h *tableHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var payload dtos.CreateTableDto

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.New().Struct(payload); err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, err)
		return
	}

	// Validate table belongs to user
	userID := r.Context().Value("user_id")
	if userID != payload.UserID {
		handler.HandleHttpError(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}

	createdTable, err := h.createTableUsecase.Execute(&payload)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusCreated, handler.HttpMessage[http.StatusCreated], createdTable)
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

	// Validate table belongs to user
	userIDCtx := r.Context().Value("user_id")
	if userIDCtx != userID {
		handler.HandleHttpError(w, http.StatusForbidden, errors.New("forbidden"))
		return
	}

	typeID, err := strconv.ParseUint(typeIDStr, 10, 64)
	if err != nil {
		handler.HandleHttpError(w, http.StatusBadRequest, errors.New("invalid type_id"))
		return
	}

	model := &dtos.GetTableByParamsDto{
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
