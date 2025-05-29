package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleHttpError(w http.ResponseWriter, statusCode int, message string, err error) {
	response := map[string]interface{}{
		"message": message,
		"error":   true,
	}

	if err != nil {
		log.Printf("%s: %v", message, err)
		response["error"] = err.Error()
	} else {
		log.Printf("%s", message)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response)
}

func HandleHttpSuccess(w http.ResponseWriter, statusCode int, message string, data ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"message": message,
		"error":   nil,
	}

	if len(data) == 1 && data[0] != nil {
		response["data"] = data[0]
	}
	if len(data) > 1 {
		response["data"] = data
	}

	json.NewEncoder(w).Encode(response)
}
