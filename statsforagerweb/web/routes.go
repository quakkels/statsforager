package web

import (
	"context"
	"net/http"
	"statsforagerweb/domain"

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
	sitesManager domain.SitesManager) {
	// routes
	mux.Handle("GET /static/", http.FileServerFS(staticFs))
	mux.HandleFunc("PUT /api/sites/{siteKey}/impressions/{impressionId}", putImpressionHandler(impressionsManager))
	mux.HandleFunc("OPTIONS /api/sites/{siteKey}/impressions/{impressionId}", optionsCorsHandler())
	mux.HandleFunc("GET /dashboard", getDashboardHandler(sitesManager, impressionsManager))
	mux.HandleFunc("GET /health", healthHandler(appInfo, statsdatastore))
	mux.HandleFunc("GET /register", getRegisterHandler())
	mux.HandleFunc("POST /register", postRegisterHandler())
	mux.HandleFunc("GET /", getHomeHandler())
}
