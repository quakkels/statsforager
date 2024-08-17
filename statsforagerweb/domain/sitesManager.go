package domain

import (
	"context"
	"strings"
)

type SitesRepository interface {
	GetSite(context.Context, string) (Site, error)
	GetAllSites(context.Context) ([]Site, error)
}

type Site struct {
	SiteKey string
	Domain  string
}

func (site Site) HasLocation(location string) bool {
	hasLocation := strings.HasPrefix(location, "http://"+site.Domain) ||
		strings.HasPrefix(location, "https://"+site.Domain)
	return hasLocation
}

type SitesManager struct {
	SitesRepo SitesRepository
}

func NewSitesManager(sitesrepo SitesRepository) SitesManager {
	manager := SitesManager{
		SitesRepo: sitesrepo,
	}
	return manager
}

func (manager *SitesManager) GetAllSites(ctx context.Context) ([]Site, error) {
	return manager.SitesRepo.GetAllSites(ctx)
}
