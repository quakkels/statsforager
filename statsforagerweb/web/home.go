package web

import (
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
		render(w, r, "home.html", model)
	}
}
