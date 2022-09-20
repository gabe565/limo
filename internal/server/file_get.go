package server

import (
	"errors"
	"github.com/gabe565/limo/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (s *Server) GetFile() http.HandlerFunc {
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

		fpath := filepath.Join(filepath.Join(viper.GetString("data-dir"), "files"), file.Name)
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
