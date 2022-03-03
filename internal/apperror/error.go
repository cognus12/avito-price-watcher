package apperror

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Err     error  `json:"-"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

var (
	ErrNotFound         = NewAppError(nil, "not found", http.StatusNotFound)
	InternalServerError = NewAppError(nil, "internal server error", http.StatusInternalServerError)
)

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppError(err error, message string, code int) *AppError {
	return &AppError{
		Err:     err,
		Message: message,
		Code:    code,
	}
}
