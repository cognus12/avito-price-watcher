package avito

import (
	"apricescrapper/internal/handlers"
	"apricescrapper/internal/scrapper"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
}

func New() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(r *httprouter.Router) {
	r.GET("/avito/parse", h.ParseHandler)
}

func (h *handler) ParseHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	city := r.URL.Query().Get("city")
	category := r.URL.Query().Get("category")
	slug := r.URL.Query().Get("slug")

	const baseUrl = "https://www.avito.ru/"

	if city == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No city provided"))
	}

	if category == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No category provided"))
	}

	if slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No slug provided"))
	}

	url := baseUrl + city + "/" + category + "/" + slug

	price := scrapper.GetPrice(url)

	w.Write([]byte("city: " + city + ", category: " + category + ", slug: " + slug + ", price: " + strconv.FormatUint(price, 10)))
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
