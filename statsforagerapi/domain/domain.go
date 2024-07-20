package domain

import (
	"context"
	"strings"
)

type validationResult struct {
	IsSuccess bool
	Messages  map[string]string
}

func NewValidationResult(messages map[string]string) *validationResult {
	vr := &validationResult{}
	if messages == nil || len(messages) == 0 {
		vr.IsSuccess = true
		return vr
	}
	vr.IsSuccess = false
	vr.Messages = messages
	return vr
}

type Site struct {
	SiteKey string
	Domain  string
}

type SitesRepository interface {
	GetSite(context.Context, string) (Site, error)
}

func (site Site) HasLocation(location string) bool {
	hasLocation := strings.HasPrefix(location, "http://"+site.Domain) ||
		strings.HasPrefix(location, "https://"+site.Domain)
	return hasLocation
}
