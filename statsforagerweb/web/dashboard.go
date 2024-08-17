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

		model := struct {
			Sites       []domain.Site
			Impressions []domain.Impression
		}{
			Sites:       sites,
			Impressions: impressions,
		}

		if err := tpl["dashboard"].Execute(w, model); err != nil {
			fmt.Println(err)
		}
	}
}
