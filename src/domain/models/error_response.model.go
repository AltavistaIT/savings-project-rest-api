package models

type ErrorResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
