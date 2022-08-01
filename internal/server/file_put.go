package server

import (
	"encoding/json"
	"github.com/gabe565/limo/internal/models"
	"github.com/gabe565/limo/internal/util"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type PutFileResponse struct {
	URL       string     `json:"url"`
	RawURL    string     `json:"raw_url"`
	ExpiresAt *time.Time `json:"expiresAt"`
}

func (s *Server) PutFile() http.HandlerFunc {
	shouldRandomize := func(r *http.Request) bool {
		v := strings.ToLower(r.Header.Get("Random"))
		random, err := strconv.ParseBool(v)
		return (err == nil && random) || v == "yes"
	}

	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if shouldRandomize(r) || name == "" {
			rand, err := util.RandSlug(4)
			if err != nil {
				panic(err)
			}
			name = rand + util.SmarterExt(name)
		}
		name = filepath.Base(name)

		var file models.File
		if err := s.DB.Where("name=?", name).Find(&file).Error; err != nil {
			panic(err)
		}

		if file.ID != 0 {
			http.Error(w, "File already exists", http.StatusUnprocessableEntity)
			return
		}

		file.Name = name

		if expIn := r.Header.Get("ExpiresIn"); expIn != "" {
			d, err := time.ParseDuration(expIn)
			if err != nil {
				panic(err)
			}
			_ = file.ExpiresAt.Scan(time.Now().Add(d))
		}

		if err := s.DB.Create(&file).Error; err != nil {
			panic(err)
		}

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

		publicUrl := util.NewUrl(r, path.Join("/f", file.Name))
		switch r.Header.Get("Accept") {
		case "application/json":
			rawUrl := util.NewUrl(r, path.Join("/raw", file.Name))
			resp := PutFileResponse{
				RawURL: rawUrl.String(),
				URL:    publicUrl.String(),
			}
			if file.ExpiresAt.Valid {
				resp.ExpiresAt = &file.ExpiresAt.Time
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
