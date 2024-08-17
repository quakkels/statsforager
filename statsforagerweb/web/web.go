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

var tpl = make(map[string]*template.Template)

func init() {
	tpl["home"] = template.Must(template.ParseFS(tplFs, "templates/base.html", "templates/home.html"))
	tpl["register"] = template.Must(template.ParseFS(tplFs, "templates/base.html", "templates/register.html"))
	tpl["dashboard"] = template.Must(template.ParseFS(tplFs, "templates/base.html", "templates/dashboard.html"))
}

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