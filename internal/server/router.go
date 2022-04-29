package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello world"))
	})

	return r
}
