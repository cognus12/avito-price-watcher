package advt

import (
	"apricescrapper/internal/apperror"
	"apricescrapper/internal/handlers"
	"apricescrapper/pkg/helpers"
	"apricescrapper/pkg/logger"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	Service Service
	logger  logger.Logger
}

func NewHandler(service Service, logger logger.Logger) handlers.Handler {
	return &handler{
		Service: service,
		logger:  logger,
	}
}

func (h *handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, "/api/ad-info", apperror.Middleware(h.AdInfo))
}

func (h *handler) AdInfo(w http.ResponseWriter, r *http.Request) error {
	url := helpers.GetQueryParam(r, "url")

	adInfo, err := h.Service.GetAdInfo(url)

	if err != nil {
		return err
	}

	h.responseJson(w, adInfo)
	return nil
}

func (h *handler) responseJson(w http.ResponseWriter, s interface{}) error {
	jsonBytes, err := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

	return err
}
