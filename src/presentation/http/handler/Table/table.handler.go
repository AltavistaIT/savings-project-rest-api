package handler_table

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ssssshel/sp-api/src/domain/models"
	"github.com/ssssshel/sp-api/src/presentation/http/handler"
	usecases_table "github.com/ssssshel/sp-api/src/usecases/Table"
)

type TableHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type tableHandler struct {
	usecases_table.CreateTableUsecase
}

func NewTableHandler(createTableUsecase usecases_table.CreateTableUsecase) TableHandler {
	return &tableHandler{
		CreateTableUsecase: createTableUsecase,
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

	createdTable, err := h.CreateTableUsecase.Execute(&payload)

	if err != nil {
		handler.HandleHttpError(w, http.StatusInternalServerError, handler.HttpMessage[http.StatusInternalServerError], err)
		return
	}

	handler.HandleHttpSuccess(w, http.StatusCreated, handler.HttpMessage[http.StatusCreated], createdTable)
}
