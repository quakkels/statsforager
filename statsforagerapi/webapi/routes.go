package webapi

import (
	"context"
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

func RegisterRoutes(
	mux *http.ServeMux,
	appInfo AppInfo,
	statsdatastore StatsDataStore,
	impressionsManager domain.ImpressionsManager) {
	// routes
	mux.HandleFunc("PUT /api/sites/{siteKey}/impression/{impressionId}", PutImpressionHandler(impressionsManager))
	mux.HandleFunc("GET /health", HealthHandler(appInfo, statsdatastore))
}
