package core

import "strings"

type BaseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HandleError(err error) BaseResponse {
	var message string

	switch {
	case strings.Contains(err.Error(), "duplicate key value violates unique constraint"):
		message = "Duplicate key value"
	default:
		message = "Internal Server Error"
	}

	response := BaseResponse{
		Message: message,
		Status:  "Error",
	}

	return response
}
