package dataaccess

import (
	"context"
	"statsforagerweb/domain"
)

type SitesRepo struct {
	dataStore statsDataStore
}

func NewSitesRepo(dataStore statsDataStore) SitesRepo {
	return SitesRepo{dataStore}
}

func (repo *SitesRepo) GetSite(context context.Context, siteKey string) (domain.Site, error) {
	var site domain.Site
	err := repo.dataStore.
		QueryRow(context, "SELECT \"site_key\", domain FROM sites").
		Scan(&site.SiteKey, &site.Domain)
	return site, err
}

func (repo *SitesRepo) GetAllSites(context context.Context) ([]domain.Site, error) {
	var sites []domain.Site

	rows, err := repo.dataStore.Query(context, "SELECT domain, site_key FROM sites")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var site domain.Site
		err := rows.Scan(&site.Domain, &site.SiteKey)
		if err != nil {
			return nil, err
		}

		sites = append(sites, site)
	}

	err = rows.Err() // get error from rows.Next() or rows.Scan()

	return sites, err
}
