package web

import (
	"context"
	"html/template"
	"net/http"
	"statsforagerapi/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type StatsDataStore interface {
	QueryRow(context.Context, string, ...any) pgx.Row
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
}

type AppInfo struct {
	Version   string
	BuildDate string
	Hash      string
}

// cache the template to avoid reparsing it on every request

var t = template.Must(template.ParseFS(tplFs, "templates/layout.html"))

func RegisterRoutes(
	mux *http.ServeMux,
	appInfo AppInfo,
	statsdatastore StatsDataStore,
	impressionsManager domain.ImpressionsManager) {
	// routes
	mux.HandleFunc("PUT /api/sites/{siteKey}/impressions/{impressionId}", putImpressionHandler(impressionsManager))
	mux.HandleFunc("OPTIONS /api/sites/{siteKey}/impressions/{impressionId}", optionsCorsHandler())
	mux.HandleFunc("GET /health", healthHandler(appInfo, statsdatastore))
	mux.Handle("GET /static/", http.FileServer(http.FS(staticFs)))
	mux.HandleFunc("GET /", getHomeHandler())
}
