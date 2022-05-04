package server

import (
	"net/http"
)

type Server struct{}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.Handler())
}
