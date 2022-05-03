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
)

func (s *Server) GetFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := filepath.Join("/", chi.URLParam(r, "name"))
		file, err := models.Files(Where("name=?", name)).One(r.Context(), s.Db)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			panic(err)
		}

		in, err := os.Open(filepath.Join("data/files", file.Name))
		if err != nil {
			panic(err)
		}

		if _, err = io.Copy(w, in); err != nil {
			panic(err)
		}

	}
}
