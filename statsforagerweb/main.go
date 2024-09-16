package main

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/didip/tollbooth"
	"github.com/joho/godotenv"
	"github.com/justinas/nosurf"

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

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	var (
		postgresHost     = os.Getenv("postgres_host")
		postgresPort     = os.Getenv("postgres_port")
		postgresUser     = os.Getenv("postgres_user")
		postgresPassword = os.Getenv("postgres_password")
		postgresDbname   = os.Getenv("postgres_dbname")
	)

	// data
	connString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDbname)
	statsDataStore, err := dataaccess.NewStatsDataStore(context.Background(), connString)
	if err != nil {
		panic(err)
	}
	defer statsDataStore.Close()

	impressionsRepo := dataaccess.NewImpressionsRepo(*statsDataStore)
	sitesRepo := dataaccess.NewSitesRepo(*statsDataStore)
	accountsRepo := dataaccess.NewAccountsRepo(*statsDataStore)

	// business domain
	smtpIsLive, err := strconv.ParseBool(os.Getenv("smtp_is_live"))
	if err != nil {
		log.Panic("smtp_is_live could not be parsed.", err)
	}
	mail, err := domain.NewMail(
		domain.SmtpConfig{
			User:     os.Getenv("smtp_user"),
			From:     os.Getenv("smtp_from"),
			Password: os.Getenv("smtp_password"),
			Host:     os.Getenv("smtp_host"),
			Port:     os.Getenv("smtp_port"),
			IsLive:   smtpIsLive,
		},
	)
	if err != nil {
		panic(err)
	}

	impressionsManager := domain.NewImpressionsManager(
		&impressionsRepo,
		&sitesRepo)
	sitesManager := domain.NewSitesManager(&sitesRepo)
	accountsManager := domain.NewAccountsManager(
		domain.AccountsConfig{
			AppRoot: os.Getenv("app_root"),
		},
		&accountsRepo,
		mail,
	)

	// web
	gob.Register(domain.OtpToken{}) // scs requires custom types to be registered in gob
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode
	sessionManager.Cookie.Secure = strings.HasPrefix(os.Getenv("app_root"), "https://")

	ham := middleware.NewHydrateAccountMiddleware(sessionManager)

	
	lmt := tollbooth.NewLimiter(1, nil)
	lmt.SetMethods([]string{"GET", "POST"})

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
		nosurf.NewPure,
		sessionManager.LoadAndSave,
		ham.Apply,
	)

	mux := http.NewServeMux()
	web.RegisterRoutes(
		mux,
		appInfo,
		statsDataStore,
		impressionsManager,
		sitesManager,
		accountsManager,
		sessionManager,
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
