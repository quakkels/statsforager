package web

import (
	"fmt"
	"net/http"
)

func getHomeHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" { // only routes that end with '/' need this
			http.NotFound(w, r)
			return
		}

		model := struct {
			Content string `json:"database_version"`
		}{
			Content: "Welcome to StatsForager",
		}

		if err := tplGlob.ExecuteTemplate(w, "home.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}
