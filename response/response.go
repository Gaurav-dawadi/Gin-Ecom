package response

import (
	"go-practice/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
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

func JWTErrorResponse(err error) ResponseMessage {
	var code string
	var status int

	code = utils.BAD_REQUEST
	status = http.StatusBadRequest

	if err == jwt.ErrSignatureInvalid {
		status = http.StatusBadRequest
		code = utils.BAD_REQUEST
	}

	if err == jwt.ErrTokenExpired {
		status = http.StatusUnauthorized
		code = utils.TOKEN_EXPIRED
	}

	return ResponseMessage{
		Message: err.Error(),
		Status:  status,
		Code:    code,
	}
}
