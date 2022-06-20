package subscription

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
	r.HandlerFunc(http.MethodPost, "/api/subscribe", apperror.Middleware(h.Subscribe))
	r.HandlerFunc(http.MethodPost, "/api/unsubscribe", apperror.Middleware(h.Unsubscribe))
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

	err = h.Service.Subscribe(dto.Url, dto.Email)

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

	err = h.Service.Unsubscribe(dto.Url, dto.Email)

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
