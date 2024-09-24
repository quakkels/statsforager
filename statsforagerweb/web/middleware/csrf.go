package middleware

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func Csrf(h http.Handler) http.Handler {
	surfing := nosurf.New(h)
	surfing.ExemptRegexps("/api/(.*)")
	surfing.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Failed to validate CSRF token:", nosurf.Reason(r))
		w.WriteHeader(http.StatusBadRequest)
	}))
	return surfing
}
