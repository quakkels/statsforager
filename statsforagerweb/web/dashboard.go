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
			Sites         []domain.Site
			Impressions   []domain.Impression
			LocationCount map[string]int
		}{
			Sites:         sites,
			Impressions:   impressions,
			LocationCount: locationCount,
		}
		render(w, r, "dashboard.html", model)
	}
}
