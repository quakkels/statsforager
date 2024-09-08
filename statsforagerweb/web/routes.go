package web

import (
	"context"
	"net/http"
	"statsforagerweb/domain"

	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type StatsDataStore interface {
	QueryRow(context.Context, string, ...any) pgx.Row
	Query(context.Context, string, ...any) (pgx.Rows, error)
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
}

type AppInfo struct {
	Version   string
	BuildDate string
	Hash      string
}

func RegisterRoutes(
	mux *http.ServeMux,
	appInfo AppInfo,
	statsdatastore StatsDataStore,
	impressionsManager domain.ImpressionsManager,
	sitesManager domain.SitesManager,
	accountsManager domain.AccountsManager,
	sessionManager *scs.SessionManager,
) {
	// routes
	mux.Handle("GET /static/", http.FileServerFS(staticFs))
	mux.HandleFunc("PUT /api/sites/{siteKey}/impressions/{impressionId}", putImpressionHandler(impressionsManager))
	mux.HandleFunc("OPTIONS /api/sites/{siteKey}/impressions/{impressionId}", optionsCorsHandler())
	mux.HandleFunc("GET /dashboard", getDashboardHandler(sitesManager, impressionsManager))
	mux.HandleFunc("GET /health", healthHandler(appInfo, statsdatastore))
	mux.HandleFunc("GET /login/confirm/{otp}", getLoginConfirmHandler(accountsManager, sessionManager))
	mux.HandleFunc("GET /login", getLoginHandler())
	mux.HandleFunc("POST /login", postLoginHandler(accountsManager, sessionManager))
	mux.HandleFunc("GET /register", getRegisterHandler())
	mux.HandleFunc("POST /register", postRegisterHandler(accountsManager))
	mux.HandleFunc("GET /", getHomeHandler())
}
