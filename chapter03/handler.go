package chapter03

import "net/http"

type Handler struct {
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

}

func register(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/value", h.Get)

	return mux
}
