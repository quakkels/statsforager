package web

import (
	"fmt"
	"net/http"
)

var count = 1

func getRegisterHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		model := struct {
			Email string
			Count int
		}{
			Email: "here",
			Count: count,
		}

		if err := tplGlob.ExecuteTemplate(w, "register.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}

func postRegisterHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		count++

		model := struct {
			Email string
			Count int
		}{
			Email: "here",
			Count: count,
		}

		if err := tplGlob.ExecuteTemplate(w, "register.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}
