package error_handler

import "fmt"

type CustomError struct {
	Code    int32
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func NewCustomError(code int32, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}