package server

import (
	"database/sql"
	"errors"
	"github.com/gabe565/limo/internal/models"
	"github.com/go-chi/chi/v5"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (s *Server) GetFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := filepath.Join("/", chi.URLParam(r, "name"))
		file, err := models.Files(Where("name=?", name)).OneG(r.Context())
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
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
