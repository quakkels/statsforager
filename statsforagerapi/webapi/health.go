package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HealthHandler(appVersion string, statsdatastore StatsDataStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var dbversion string
		statsdatastore.QueryRow(r.Context(), "SELECT version FROM db_version").Scan(&dbversion)

		type health struct {
			DatabaseVersion string `json:"database_version"`
			ApiVersion      string `json:"api_version"`
		}

		model := health{
			DatabaseVersion: dbversion,
			ApiVersion:      appVersion,
		}
		fmt.Println(model)
		result, err := json.Marshal(model)
		if err != nil {
			panic(err)
		}

		w.Write(result)
	}
}
