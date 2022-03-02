package avito

import (
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
	r.GET("/avito/parse", h.ParseHandler)
}

func (h *handler) ParseHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	city := helpers.GetQueryParam(r, "city")
	category := helpers.GetQueryParam(r, "category")
	slug := helpers.GetQueryParam(r, "slug")

	params := urlParams{city, category, slug}

	adInfo, err := h.AvitoService.GetAdInfo(params)

	if err != nil {
		msg := err.Error()

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))

		h.logger.Error(msg)
		return
	}

	h.responseJson(w, adInfo)
}

func (h *handler) responseJson(w http.ResponseWriter, s interface{}) {
	jsonBytes, err := json.Marshal(s)

	if err != nil {

		msg := err.Error()

		h.logger.Error(msg)

		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
