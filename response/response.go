package response

import (
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
}

func ResponseOK(message string) *ResponseMessage {
	return &ResponseMessage{
		Message: message,
		Status:  http.StatusOK,
		Code:    "OK",
	}
}

func ResponseCreated(message string) *ResponseMessage {
	return &ResponseMessage{
		Message: message,
		Status:  http.StatusCreated,
		Code:    "created",
	}
}

func ResponseBadRequest(message string) ResponseMessage {
	return ResponseMessage{
		Message: message,
		Status:  http.StatusBadRequest,
		Code:    "bad_request",
	}
}
