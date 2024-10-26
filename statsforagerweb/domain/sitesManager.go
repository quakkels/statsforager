package domain

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

type SitesRepository interface {
	GetSite(context.Context, string) (Site, error)
	GetAllSites(context.Context) ([]Site, error)
	GetSites(context.Context, string) ([]Site, error)
	SaveSite(context.Context, Site) error
}

type Site struct {
	SiteKey      string
	Domain       string
	OwnerAccount string
	SiteName     string
}

func (site Site) HasLocation(location string) bool {
	hasLocation := strings.HasPrefix(strings.ToLower(location), "http://"+site.Domain) ||
		strings.HasPrefix(strings.ToLower(location), "https://"+site.Domain)
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

func (manager *SitesManager) GetSites(ctx context.Context, accountCode string) ([]Site, error) {
	return manager.SitesRepo.GetSites(ctx, accountCode)
}

func (manager *SitesManager) SaveSite(ctx context.Context, site Site) (validationResult, error) {
	return validationResult{}, nil
}

func (manager *SitesManager) validateSite(site Site) validationResult {
	messages := make(map[string]string)
	if err := uuid.Validate(site.SiteKey); err != nil {
		messages["SiteKey"] = "SiteKey is incorrect"
	}

	if strings.HasPrefix(strings.ToLower(site.Domain), "http") {
		messages["Domain"] = "Domain should not start with http"
	}
	return *NewValidationResult(messages)
}
