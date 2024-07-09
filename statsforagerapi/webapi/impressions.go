package webapi

import (
	"net/http"
	"statsforagerapi/domain"
)

func PostImpressionHandler(impressionsManager domain.ImpressionsManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo: how to make a manager or service with a lifecycle for just this request?

		// var dbversion string
		// statsdatastore.QueryRow(r.Context(), "SELECT version FROM db_version").Scan(&dbversion)

		impressionId := r.PathValue("impressionId")
		siteKey := r.PathValue("siteKey")

		result := struct {
			SiteKey      string
			ImpressionId string
		}{
			SiteKey:      siteKey,
			ImpressionId: impressionId,
		}

		WriteJSON(w, http.StatusOK, result)
	}
}
