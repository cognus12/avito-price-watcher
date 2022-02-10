package avito

import (
	"apricescrapper/internal/handlers"
	"apricescrapper/pkg/helpers"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	AvitoService Service
}

func NewHandler(avitoService Service) handlers.Handler {
	return &handler{
		AvitoService: avitoService,
	}
}

func (h *handler) Register(r *httprouter.Router) {
	r.GET("/avito/parse", h.ParseHandler)
}

func (h *handler) ParseHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	city := helpers.GetQueryParam(r, "city")
	category := helpers.GetQueryParam(r, "city")
	slug := helpers.GetQueryParam(r, "slug")

	params := urlParams{city, category, slug}

	adInfo, err := h.AvitoService.GetAdInfo(params)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	h.responseJson(w, adInfo)
}

func (h *handler) responseJson(w http.ResponseWriter, s interface{}) {
	jsonBytes, err := json.Marshal(s)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
