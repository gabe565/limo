package server

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) ListenAndServe(addr string) error {
	log.WithField("address", addr).Info("listening for requests")
	return http.ListenAndServe(addr, s.Handler())
}
