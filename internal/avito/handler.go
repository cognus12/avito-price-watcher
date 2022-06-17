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
	r.HandlerFunc(http.MethodGet, "/api/info", apperror.Middleware(h.ParseHandler))
	r.HandlerFunc(http.MethodPost, "/api/subscribe", apperror.Middleware(h.Subscribe))
	r.HandlerFunc(http.MethodPost, "/api/unsubscribe", apperror.Middleware(h.Unsubscribe))

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

func (h *handler) Subscribe(w http.ResponseWriter, r *http.Request) error {
	h.logger.Info("Call subscribe method")

	defer r.Body.Close()

	var dto SubscribtionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		return err
	}

	err = dto.Validate()

	if err != nil {
		return err
	}

	err = h.AvitoService.Subscribe(dto.Url, dto.Email)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Call subscribe"))

	return nil
}

func (h *handler) Unsubscribe(w http.ResponseWriter, r *http.Request) error {
	h.logger.Info("Call unsubscribe method")

	defer r.Body.Close()

	var dto SubscribtionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		return err
	}

	err = dto.Validate()

	if err != nil {
		return err
	}

	err = h.AvitoService.Unsubscribe(dto.Url, dto.Email)

	if err != nil {
		return err
	}

	h.logger.Info("Call unsubscribe method with %+v", dto)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Call unsubscribe"))

	return nil
}

func (h *handler) responseJson(w http.ResponseWriter, s interface{}) error {
	jsonBytes, err := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

	return err
}
