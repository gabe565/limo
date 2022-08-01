package server

import (
	"errors"
	"github.com/gabe565/limo/internal/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
)

func (s *Server) DeleteFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := filepath.Base(chi.URLParam(r, "name"))
		var file models.File
		if err := s.DB.Where("name=?", name).First(&file).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			panic(err)
		}
		if err := s.DB.Delete(&file).Error; err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
