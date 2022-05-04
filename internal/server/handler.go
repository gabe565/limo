package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) Handler() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/api/health"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/files", func(r chi.Router) {
		r.Put("/{name}", s.PutFile())
		r.Put("/", s.PutFile())
		r.Get("/{name}", s.GetFile())
		r.Delete("/{name}", s.DeleteFile())
	})

	r.Get("/raw/{name}", s.GetFile())
	r.Get("/f/{name}", s.GetFile())

	return r
}
