package web

import (
	"fmt"
	"net/http"
	"statsforagerweb/domain"
	"strconv"
)

func getDashboardHandler(
	sitesManager domain.SitesManager,
	impressionsManager domain.ImpressionsManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		currentSiteKey := params.Get("SiteKey")
		timeUnitCount, err := strconv.Atoi(params.Get("TimeUnitCount"))
		if err != nil {
			timeUnitCount = 10
		}
		timeUnit := params.Get("TimeUnit")
		if timeUnit != "day" && timeUnit != "week" && timeUnit != "month" {
			timeUnit = "week"
		}

		sites, err := sitesManager.GetAllSites(r.Context())
		if err != nil {
			fmt.Println(err)
		}

		if len(currentSiteKey) == 0 && len(sites) != 0 {
			currentSiteKey = sites[0].SiteKey
		}

		impressions, err := impressionsManager.GetAllImpressions(r.Context())
		if err != nil {
			fmt.Println(err)
		}

		locationCount, err := impressionsManager.GetLocationCounts(r.Context(), currentSiteKey)
		if err != nil {
			fmt.Println(err)
		}

		model := struct {
			Parameters struct {
				SiteKey       string
				TimeUnitCount int
				TimeUnit      string
			}
			SiteSelect    map[string]string
			Impressions   []domain.Impression
			LocationCount map[string]int
		}{
			Parameters: struct {
				SiteKey       string
				TimeUnitCount int
				TimeUnit      string
			}{
				SiteKey:       currentSiteKey,
				TimeUnitCount: timeUnitCount,
				TimeUnit:      timeUnit,
			},
			SiteSelect: func(s []domain.Site) map[string]string {
				m := make(map[string]string)
				for _, site := range s {
					m[site.SiteKey] = site.Domain
				}
				return m
			}(sites),
			Impressions:   impressions,
			LocationCount: locationCount,
		}
		render(w, r, "dashboard.html", model)
	}
}
