package exceptions

import "net/http"

type APIError struct {
	Code    int
	Message string
}

func Error(e APIError) string {
	return e.Message
}

func NotFoundError(message string) APIError {
	return APIError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func Unprocessable(message string) APIError {
	return APIError{
		Code: http.StatusUnprocessableEntity, 
		Message: message,
	}
}
