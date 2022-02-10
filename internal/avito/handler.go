package avito

import (
	"apricescrapper/internal/handlers"
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
	city := r.URL.Query().Get("city")
	category := r.URL.Query().Get("category")
	slug := r.URL.Query().Get("slug")

	params := urlParams{city, category, slug}

	adInfo, err := h.AvitoService.GetAdInfo(params)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	jsonBytes, err := json.Marshal(adInfo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.Write(jsonBytes)
}

/*

func sendJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}s

*/
