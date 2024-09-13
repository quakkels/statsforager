package web

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
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

type modelWrapper struct {
	AccountCode string
	Model any
}
func render(w http.ResponseWriter, context context.Context, templateName string, model any) {
	modelWrapper := modelWrapper{Model: model}

	accountCode, ok := context.Value("accountCode").(string)
	if ok {
		modelWrapper.AccountCode = accountCode
	}

	if err := tplGlob.ExecuteTemplate(w, templateName, modelWrapper); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}
}
