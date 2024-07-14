package webapi

import (
	"net/http"
)

func healthHandler(
	appInfo AppInfo,
	statsdatastore StatsDataStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var dbversion string
		statsdatastore.QueryRow(r.Context(), "SELECT version FROM db_version").Scan(&dbversion)

		type health struct {
			DatabaseVersion string `json:"database_version"`
			ApiVersion      string `json:"api_version"`
			ApiBuildDate    string `json:"api_build_date"`
			ApiHash         string `json:"api_hash"`
		}

		model := health{
			DatabaseVersion: dbversion,
			ApiVersion:      appInfo.Version,
			ApiBuildDate:    appInfo.BuildDate,
			ApiHash:         appInfo.Hash,
		}

		WriteJson(w, http.StatusOK, model)
	}
}
