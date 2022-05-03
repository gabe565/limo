package server

import (
	"database/sql"
	"net/http"
)

type Server struct {
	Db *sql.DB
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.Handler())
}
