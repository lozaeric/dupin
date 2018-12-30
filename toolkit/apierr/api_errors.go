package apierr

import (
	"fmt"
)

type ApiError struct {
	StatusCode int
	Message    string
}

func New(statusCode int, message string) *ApiError {
	// check if it's a valid status code
	return &ApiError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("status code = %v, %s", e.StatusCode, e.Message)
}
