package web

import (
	"fmt"
	"net/http"
)

func getRegisterHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		model := struct {
			Email string
		}{
			Email: "here",
		}

		if err := tpl["register"].Execute(w, model); err != nil {
			fmt.Println(err)
		}
	}
}

