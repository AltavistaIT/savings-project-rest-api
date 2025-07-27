package responses

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data omitempty"`
	Success bool        `json:"success"`
}
