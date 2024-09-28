package web

import (
	"fmt"
	"net/http"
	"statsforagerweb/domain"
)

func getDashboardHandler(
	sitesManager domain.SitesManager,
	impressionsManager domain.ImpressionsManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		sites, err := sitesManager.GetAllSites(r.Context())
		if err != nil {
			fmt.Println(err)
		}

		impressions, err := impressionsManager.GetAllImpressions(r.Context())
		if err != nil {
			fmt.Println(err)
		}

		locationCount, err := impressionsManager.GetLocationCounts(r.Context(), sites[0].SiteKey)
		if err != nil {
			fmt.Println(err)
		}

		model := struct {
			Parameters struct {
				SiteKey       string
				TimeUnitCount int
				TimeUnit      string
			}
			Sites         []domain.Site
			Impressions   []domain.Impression
			LocationCount map[string]int
		}{
			Parameters: struct {
				SiteKey       string
				TimeUnitCount int
				TimeUnit      string
			}{
				SiteKey:       "test",
				TimeUnitCount: 10,
				TimeUnit:      "day",
			},
			Sites:         sites,
			Impressions:   impressions,
			LocationCount: locationCount,
		}
		render(w, r, "dashboard.html", model)
	}
}
