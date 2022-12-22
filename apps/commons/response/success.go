package response

import "net/http"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload,omitempty"`
}

var (
	generalSuccess = Response{
		StatusCode: http.StatusOK,
		Message:    "SUCCESS",
	}
	createSuccess = Response{
		StatusCode: http.StatusCreated,
		Message:    "CREATED SUCCESS",
	}
)

func GeneralSuccess() *Response {
	succ := generalSuccess
	return &succ
}

func GeneralSuccessCustomMessageAndPayload(message string, payload interface{}) *Response {
	succ := generalSuccess
	succ.Message = message
	succ.Payload = payload
	return &succ
}

func CreatedSuccess() *Response {
	succ := createSuccess
	return &succ
}

func CreatedSuccessWithPayload(payload interface{}) *Response {
	succ := createSuccess
	succ.Payload = payload
	return &succ
}
