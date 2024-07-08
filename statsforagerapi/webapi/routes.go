package webapi

import (
	"context"
	"net/http"
	"statsforagerapi/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type StatsDataStore interface {
	QueryRow(context.Context, string) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

func RegisterRoutes(
	mux *http.ServeMux,
	version string,
	builddate string,
	hash string,
	statsdatastore StatsDataStore,
	impressionsManager domain.ImpressionsManager) {
	// deps

	// routes
	mux.HandleFunc("PUT /api/sites/{siteKey}/impression/{impressionId}", PutImpressionHandler(impressionsManager))
	mux.HandleFunc("PUT /api/sites/{siteKey}/impression/{impressionId}/end", PutImpressionLeavingHandler(impressionsManager))
	mux.HandleFunc("GET /health", HealthHandler(version, builddate, hash, statsdatastore))
}
