package web

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"statsforagerweb/web/tplhelpers"

	"github.com/justinas/nosurf"
)

//go:embed templates/*.html
var tplFs embed.FS

//go:embed static/*
var staticFs embed.FS

type errorResponse struct {
	Message string `json:"message"`
}

var funcMap = template.FuncMap{
	"select":  tplhelpers.Select,
	"makeMap": tplhelpers.MakeMap,
}

var tplGlob = template.Must(
	template.
		New("base").
		Funcs(funcMap).
		ParseFS(tplFs, "templates/*.html"))

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
	Token       string
	Model       any
}

func (self *modelWrapper) CsrfInputElement() template.HTML {
	element := "<input type='hidden' name='csrf_token' value='" + self.Token + "'>"
	return template.HTML(element)
}

func render(w http.ResponseWriter, request *http.Request, templateName string, model any) {
	modelWrapper := &modelWrapper{
		Model: model,
		Token: nosurf.Token(request),
	}

	context := request.Context()

	accountCode, ok := context.Value("accountCode").(string)
	if ok {
		modelWrapper.AccountCode = accountCode
	}

	if err := tplGlob.ExecuteTemplate(w, templateName, modelWrapper); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}
}
