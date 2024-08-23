package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"statsforagerweb/dataaccess"
	"statsforagerweb/domain"
	"statsforagerweb/web"
	"statsforagerweb/web/middleware"
)

var (
	Version   = "0.0.1"
	BuildDate = "No Date"
	Hash      = "No Hash"
)

func main() {
	appInfo := web.AppInfo{Version: Version, BuildDate: BuildDate, Hash: Hash}

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
	sitesRepo := dataaccess.NewSitesRepo(*statsDataStore)
	impressionsManager := domain.NewImpressionsManager(
		&impressionsRepo,
		&sitesRepo)
	sitesManager := domain.NewSitesManager(&sitesRepo)

	mux := http.NewServeMux()
	web.RegisterRoutes(
		mux,
		appInfo,
		statsDataStore,
		impressionsManager,
		sitesManager)

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8000",
		Handler: middlewareStack(mux),
	}

	// reference: https://dev.to/antonkuklin/golang-graceful-shutdown-3n6d
	// may need to revisit and add WaitGroups as services are added.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen and serve returned err: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	if err := server.Shutdown(context.TODO()); err != nil {
		log.Printf("server shutdown returned an err: %v\n", err)
	}

	log.Println("final")
}
