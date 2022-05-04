package server

import (
	"encoding/json"
	"github.com/gabe565/limo/internal/models"
	"github.com/gabe565/limo/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (s *Server) PutFile() http.HandlerFunc {
	type Response struct {
		URL       string    `json:"url"`
		RawURL    string    `json:"raw_url"`
		ExpiresAt null.Time `json:"expiresAt"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			rand, err := util.RandSlug(4)
			if err != nil {
				panic(err)
			}
			name = rand
		}
		name = filepath.Join("/", name)

		if models.Files(Where("name=?", name)).ExistsGP(r.Context()) {
			http.Error(w, "File already exists", http.StatusUnprocessableEntity)
			return
		}

		var file models.File
		file.Name = name
		file.InsertGP(r.Context(), boil.Infer())

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

		if err := out.Close(); err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusCreated)

		publicUrl := util.NewUrl(r, "/f"+file.Name)
		switch r.Header.Get("Accept") {
		case "application/json":
			rawUrl := util.NewUrl(r, "/raw"+file.Name)
			resp := Response{
				RawURL:    rawUrl.String(),
				URL:       publicUrl.String(),
				ExpiresAt: file.ExpiresAt,
			}
			if err = json.NewEncoder(w).Encode(resp); err != nil {
				panic(err)
			}
		default:
			if _, err = w.Write([]byte(publicUrl.String() + "\n")); err != nil {
				panic(err)
			}
		}
	}
}
