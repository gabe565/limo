package server

import (
	"github.com/gabe565/limo/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (s *Server) PutFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := filepath.Join("/", chi.URLParam(r, "name"))
		if models.Files(Where("name=?", name)).ExistsP(r.Context(), s.Db) {
			http.Error(w, "Already exists", http.StatusUnprocessableEntity)
			return
		}

		var file models.File
		file.Name = name
		file.InsertP(r.Context(), s.Db, boil.Infer())

		out, err := os.Create(filepath.Join("data/files", file.Name))
		if err != nil {
			panic(err)
		}
		defer func(out *os.File) {
			_ = out.Close()
		}(out)

		if _, err := io.Copy(out, r.Body); err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusCreated)
	}
}
