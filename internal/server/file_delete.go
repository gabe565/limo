package server

import (
	"github.com/gabe565/limo/internal/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"path/filepath"
)

func (s *Server) DeleteFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := filepath.Join("/", chi.URLParam(r, "name"))
		var file models.File
		if err := s.DB.Where("name=?", name).First(&file).Error; err != nil {
			panic(err)
		}
		if err := s.DB.Delete(&file).Error; err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
