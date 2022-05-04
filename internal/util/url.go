package util

import (
	"net/http"
	"net/url"
)

func NewUrl(r *http.Request, path string) url.URL {
	return url.URL{
		Scheme: SchemeFromRequest(r),
		Host:   r.Host,
		Path:   path,
	}
}

func SchemeFromRequest(r *http.Request) string {
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		return "https"
	}
	return "http"
}
