package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ssssshel/sp-api/src/domain/models"
	"github.com/ssssshel/sp-api/src/shared/logger"
)

func HandleHttpError(w http.ResponseWriter, statusCode int, message error) {
	response := models.ErrorResponse{
		Success: false,
	}

	if message != nil {
		response.Message = message.Error()
	} else {
		response.Message = HttpMessage[statusCode]
	}

	logger.Info("%s", message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}

func HandleHttpSuccess(w http.ResponseWriter, statusCode int, message string, data ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := models.SuccessResponse{
		Message: message,
		Success: true,
	}

	if len(data) == 1 && data[0] != nil {
		response.Data = data[0]
	}
	if len(data) > 1 {
		response.Data = data
	}

	json.NewEncoder(w).Encode(response)
}
