package response

import "net/http"

type CustomError struct {
	Code           string      `json:"code"`
	StatusCode     int         `json:"status_code"`
	Message        string      `json:"message"`
	AdditionalInfo interface{} `json:"additional_info,omitempty"`
}

var (
	generalError = CustomError{
		Code:       "ERR0001",
		StatusCode: http.StatusInternalServerError,
		Message:    "INTERNAL SERVER ERROR",
	}
	repositoryError = CustomError{
		Code:       "ERR0002",
		StatusCode: http.StatusInternalServerError,
		Message:    "REPOSITORY ERROR",
	}
	notFoundError = CustomError{
		Code:       "ERR0003",
		StatusCode: http.StatusNotFound,
		Message:    "NOT FOUND ERROR",
	}
)

func GeneralError() *CustomError {
	err := generalError
	return &err
}
func GeneralErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := generalError
	err.AdditionalInfo = info
	return &err
}

func RepositoryError() *CustomError {
	err := repositoryError
	return &err
}

func RepositoryErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := repositoryError
	err.AdditionalInfo = info
	return &err
}
func NotFoundError() *CustomError {
	err := notFoundError
	return &err
}

func NotFoundErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := repositoryError
	err.AdditionalInfo = info
	return &err
}
