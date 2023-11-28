package domain

type CustomError struct {
	Field      string `json:"field"`
	ErrMessage string `json:"error_message"`
}
