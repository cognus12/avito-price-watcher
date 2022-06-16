package apperror

import (
	"apricescrapper/pkg/logger"
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := h(w, r)

		if err != nil {
			proccessError(err, w)
		}
	}
}

func proccessError(e error, w http.ResponseWriter) {
	logger := logger.GetInstance()

	logger.Error(e.Error())

	var appErr *AppError

	if errors.As(e, &appErr) {
		if errors.Is(e, ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(ErrNotFound.Marshal())
			return
		}

		if errors.Is(e, ErrAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
			w.Write(ErrAlreadyExists.Marshal())
			return
		}

		e = e.(*AppError)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(appErr.Marshal())
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(InternalServerError.Marshal())

	return
}
