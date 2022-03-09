package avito

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
	AvitoService Service
	logger       logger.Logger
}

func NewHandler(avitoService Service, logger logger.Logger) handlers.Handler {
	return &handler{
		AvitoService: avitoService,
		logger:       logger,
	}
}

func (h *handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, "/avito/parse", apperror.Middleware(h.ParseHandler))
}

func (h *handler) ParseHandler(w http.ResponseWriter, r *http.Request) error {
	city := helpers.GetQueryParam(r, "city")
	category := helpers.GetQueryParam(r, "category")
	slug := helpers.GetQueryParam(r, "slug")

	params := urlParams{city, category, slug}

	adInfo, err := h.AvitoService.GetAdInfo(params)

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
