package webapi

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type StatsDataStore interface {
	QueryRow(context.Context, string) pgx.Row
	Close()
}

func RegisterRoutes(
	mux *http.ServeMux,
	version string,
	builddate string,
	hash string,
	statsdatastore StatsDataStore) {
	mux.HandleFunc("PUT /api/sites/{siteKey}/impression/{impressionId}", PutImpressionHandler(statsdatastore))
	mux.HandleFunc("PUT /api/sites/{siteKey}/impression/{impressionId}/end", PutImpressionLeavingHandler(statsdatastore))
	mux.HandleFunc("GET /health", HealthHandler(version, builddate, hash, statsdatastore))
}
