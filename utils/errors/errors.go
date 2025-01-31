package errors

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NotFoundError(message string) *AppError {
	return &AppError{
		Code: http.StatusNotFound,
        Message: message,
	}
}

func InternalServerError(message string) *AppError {
	return &AppError{
		Code: http.StatusNotFound,
        Message: message,
	}
}

func UnauthorizedError(message string) *AppError {
	return &AppError{
		Code: http.StatusUnauthorized,
        Message: message,
	}
}