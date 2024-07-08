package webapi

import (
	"net/http"
	"statsforagerapi/domain"
)

func PutImpressionHandler(impressionsManager domain.ImpressionsManager) func(http.ResponseWriter, *http.Request) {
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

func PutImpressionLeavingHandler(impressionsManager domain.ImpressionsManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo: how to make a manager or service with a lifecycle for just this request?

		impressionId := r.PathValue("impressionId")
		siteKey := r.PathValue("siteKey")

		result := struct {
			SiteKey      string `json:"site_key"`
			ImpressionId string `json:"impression_id"`
			End          bool   `json:"end"`
		}{
			SiteKey:      siteKey,
			ImpressionId: impressionId,
			End:          true,
		}

		WriteJSON(w, http.StatusOK, result)
	}
}
