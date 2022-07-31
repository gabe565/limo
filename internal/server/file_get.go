package server

import (
	"github.com/gabe565/limo/internal/models"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (s *Server) GetFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := filepath.Join("/", chi.URLParam(r, "name"))
		var file models.File
		if err := s.DB.Where("name=?", name).First(&file).Error; err != nil {
			panic(err)
		}

		fpath := filepath.Join("data/files", file.Name)
		info, err := os.Stat(fpath)
		if err != nil {
			panic(err)
		}
		in, err := os.Open(fpath)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Length", strconv.Itoa(int(info.Size())))
		if _, err = io.Copy(w, in); err != nil {
			panic(err)
		}

	}
}
