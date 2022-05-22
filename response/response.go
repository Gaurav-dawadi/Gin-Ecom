package response
import ("net/http")


type ResponseMessage struct{
	message  string  `json:"message"`
	status   int64 	 `json:"status"`
	code     string	 `json:"code"`
}

func ResponseOK(message string) *ResponseMessage{
	return &ResponseMessage{
		message: message,
		status: http.StatusOK,
		code: "OK",
	}
}

func ResponseCreated(message string) *ResponseMessage{
	return &ResponseMessage{
		message: message,
		status: http.StatusCreated,
		code: "created",
	}
}

func ResponseBadRequest(message string) *ResponseMessage{
	return &ResponseMessage{
		message: message,
		status: http.StatusBadRequest,
		code: "bad_request",
	}
}