package main

import (
	"context"
	"fmt"
	"net/http"

	"statsforagerapi/dataaccess"
	"statsforagerapi/domain"
	"statsforagerapi/webapi"
	"statsforagerapi/webapi/middleware"
)

var (
	Version   = "0.0.1"
	BuildDate = "No Date"
	Hash      = "No Hash"
)

func main() {
	appInfo := webapi.AppInfo{Version: Version, BuildDate: BuildDate, Hash: Hash}

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

	impressionsRepo := dataaccess.NewImpressionsRepo(*statsDataStore)
	impressionsManager := domain.NewImpressionsManager(&impressionsRepo)

	mux := http.NewServeMux()
	webapi.RegisterRoutes(
		mux,
		appInfo,
		statsDataStore,
		impressionsManager)

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8000",
		Handler: middlewareStack(mux),
	}
	fmt.Println("---about to listen and serve")
	err = server.ListenAndServe()
	fmt.Println("---after listening and serving")
	if err != nil {
		fmt.Println(err)
	}
}
