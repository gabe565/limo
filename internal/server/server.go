package server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct{}

func (s *Server) ListenAndServe(addr string) error {
	log.WithField("address", addr).Info("listening for requests")
	return http.ListenAndServe(addr, s.Handler())
}
