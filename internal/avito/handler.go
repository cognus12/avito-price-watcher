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

type SubscribtionDTO struct {
	Url   string
	Email string
}

func NewHandler(avitoService Service, logger logger.Logger) handlers.Handler {
	return &handler{
		AvitoService: avitoService,
		logger:       logger,
	}
}

func (h *handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, "/api/info", apperror.Middleware(h.ParseHandler))
	r.HandlerFunc(http.MethodPost, "/api/subscribe", h.Subscribe)
	r.HandlerFunc(http.MethodPost, "/api/unsubscribe", h.Unsubscribe)
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

func (h *handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Call subscribe method")

	defer r.Body.Close()

	var dto SubscribtionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Info("Call subscribe method with %+v", dto)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Call subscribe"))
}

func (h *handler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Call unsubscribe method")

	defer r.Body.Close()

	var dto SubscribtionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Info("Call unsubscribe method with %+v", dto)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Call unsubscribe"))
}

func (h *handler) responseJson(w http.ResponseWriter, s interface{}) error {
	jsonBytes, err := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

	return err
}
