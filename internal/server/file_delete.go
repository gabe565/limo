package server

import (
	"database/sql"
	"errors"
	"github.com/gabe565/limo/internal/models"
	"github.com/go-chi/chi/v5"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"path/filepath"
)

func (s *Server) DeleteFile() http.HandlerFunc {
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
		file.DeleteGP(r.Context(), false)

		w.WriteHeader(http.StatusNoContent)
	}
}
