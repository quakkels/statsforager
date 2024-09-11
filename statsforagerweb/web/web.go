package web

import (
	"embed"
	"encoding/json"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var tplFs embed.FS

//go:embed static/*
var staticFs embed.FS

type errorResponse struct {
	Message string `json:"message"`
}

var tplGlob = template.Must(template.ParseFS(tplFs, "templates/*.html"))

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func setupCors(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set(
		"Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func optionsCorsHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupCors(w)
	}
}


