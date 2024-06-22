package webapi

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type StatsDataStore interface {
	QueryRow(context.Context, string) pgx.Row
	Close()
}

func RegisterRoutes(
	mux *http.ServeMux,
	statsdatastore StatsDataStore) {
	mux.HandleFunc("GET /route/{siteKey}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("siteKey")
		var dbversion string
		statsdatastore.QueryRow(r.Context(), "SELECT version FROM db_version").Scan(&dbversion)
		w.Write([]byte("you found me: " + id + "\n\n"))
		w.Write([]byte("<p>db version: " + dbversion + "</p>\n\n"))
	})
}
