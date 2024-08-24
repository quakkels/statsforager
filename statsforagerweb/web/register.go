package web

import (
	"fmt"
	"net/http"
)

type registerModel struct {
	Email string
	IsPostSuccess bool
	Errors []string
}

func getRegisterHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		model := registerModel{}

		if err := tplGlob.ExecuteTemplate(w, "register.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}

func postRegisterHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		model := registerModel{
			Email: r.Form.Get("Email"),
			Errors: []string{"This is the first error.", "Second error.", "Third error right here."},
			IsPostSuccess: false,
		}

		if err := tplGlob.ExecuteTemplate(w, "register.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}
