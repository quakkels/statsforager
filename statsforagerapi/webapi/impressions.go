package webapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"statsforagerapi/domain"
	"time"
)

func PutImpressionHandler(impressionsManager domain.ImpressionsManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		siteKey := r.PathValue("siteKey")
		impressionId := r.PathValue("impressionId")
		ipAddress := r.Header.Get(http.CanonicalHeaderKey("x-forwarded-for"))
		if ipAddress == "" {
			var err error
			ipAddress, _, err = net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				log.Fatal("Error finding remote IP: ", err.Error())
				WriteJSON(w, http.StatusInternalServerError, "Error finding remote IP")
			}
		}

		type impressionModel struct {
			UserAgent    string    `json:"userAgent"`
			Language     string    `json:"language"`
			Location     string    `json:"location"`
			Referrer     string    `json:"referrer"`
			StartedUtc   time.Time `json:"startedUtc"`
			CompletedUtc time.Time `json:"completedUtc"`
		}

		var model impressionModel
		enc := json.NewDecoder(r.Body)
    enc.Decode(&model)

		result, err := impressionsManager.SaveImpression(
			r.Context(),
			siteKey,
			impressionId,
			model.UserAgent,
			model.Language,
			model.Location,
			model.Referrer,
			ipAddress,
			model.StartedUtc,
			model.CompletedUtc)

		if err != nil {
			fmt.Println("Error: could not save impression. ", err.Error())
			WriteJSON(w, http.StatusInternalServerError, errorResponse{"Error saving impression"})
			return
		}
		
		if !result.IsSuccess {
			WriteJSON(w, http.StatusBadRequest, result.Messages)
		} else {
			WriteJSON(w, http.StatusOK, model)
		}
	}
}
