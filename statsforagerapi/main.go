package main

import (
	"context"
	"fmt"
	"net/http"

	"statsforagerapi/dataaccess"
	"statsforagerapi/webapi"
	"statsforagerapi/webapi/middleware"
)

var (
	Version   = "0.0.1"
	BuildDate = "No Date"
	Hash      = "No Hash"
)

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

	webapi.RegisterRoutes(
		mux,
		Version,
		BuildDate,
		Hash,
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
