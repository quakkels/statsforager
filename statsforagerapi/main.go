package main

import (
	"context"
	"fmt"
	"net/http"

	"statsforagerapi/dataaccess"
	"statsforagerapi/webapi"
	"statsforagerapi/webapi/middleware"
)

func MakeHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("works\n"))
	}
}

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "stats"
	)

	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	statsDataStore, err := dataaccess.NewStatsDataStore(context.Background(), connString)
	if err != nil {
		panic(err)
	}
	defer statsDataStore.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /thing/{siteKey}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("siteKey")
		var dbversion string
		statsDataStore.QueryRow(r.Context(), "SELECT version FROM db_version").Scan(&dbversion)
		w.Write([]byte("you found me: " + id + "\n\n"))
		w.Write([]byte("<p>db version: " + dbversion + "</p>\n\n"))
	})
	mux.HandleFunc("GET /thing/makething", MakeHandler())

	webapi.RegisterRoutes(
		mux,
		statsDataStore)

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8000",
		Handler: middlewareStack(mux),
	}

	server.ListenAndServe()
}
