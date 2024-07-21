package dataaccess

import (
	"context"
	"statsforagerapi/domain"
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
