package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ImpressionsManager struct {
	ImpressionsRepo ImpressionsRepository
	SitesRepo       SitesRepository
}

type Impression struct {
	SiteKey      string
	ImpressionId string
	UserAgent    string
	Language     string
	Location     string
	Referrer     string
	IpAddress    string
	StartedUtc   time.Time
	CompletedUtc time.Time
	CreatedUtc   time.Time
}

type ImpressionsRepository interface {
	SaveImpression(context.Context, Impression) error
	GetAllImpressions(context.Context) ([]Impression, error)
	GetLocationCount(context.Context, string) (map[string]int, error)
}

func NewImpressionsManager(
	repo ImpressionsRepository,
	sitesrepo SitesRepository) ImpressionsManager {
	manager := ImpressionsManager{
		ImpressionsRepo: repo,
		SitesRepo:       sitesrepo,
	}
	return manager
}

func (manager *ImpressionsManager) SaveImpression(
	context context.Context,
	siteKey string,
	impressionId string,
	userAgent string,
	language string,
	location string,
	referrer string,
	ipAddress string,
	startedUtc time.Time,
	completedUtc time.Time) (*validationResult, error) {

	dto := Impression{
		SiteKey:      siteKey,
		ImpressionId: impressionId,
		UserAgent:    userAgent,
		Language:     language,
		Location:     location,
		Referrer:     referrer,
		IpAddress:    ipAddress,
		StartedUtc:   startedUtc,
		CompletedUtc: completedUtc,
	}

	vr, err := manager.ValidateImpression(context, dto)
	if err != nil {
		return vr, err
	}

	if !vr.IsSuccess {
		return vr, nil
	}

	err = manager.ImpressionsRepo.SaveImpression(context, dto)
	return vr, err
}

func (manager *ImpressionsManager) ValidateImpression(
	context context.Context,
	impression Impression) (*validationResult, error) {
	const emptyOrMissing = "Empty or missing value"
	const notUuid = "Not a UUID"
	messages := make(map[string]string)

	isSiteKeyUuid := true
	if err := uuid.Validate(impression.SiteKey); err != nil {
		messages["SiteKey"] = notUuid
		isSiteKeyUuid = false
	}

	if err := uuid.Validate(impression.ImpressionId); err != nil {
		messages["ImpressionId"] = notUuid
	}

	if impression.UserAgent == "" {
		messages["UserAgent"] = emptyOrMissing
	}

	if impression.Location == "" {
		messages["Location"] = emptyOrMissing
	} else if isSiteKeyUuid {
		site, err := manager.SitesRepo.GetSite(context, impression.SiteKey)
		if err != nil {
			return nil, err
		}
		if !site.HasLocation(impression.Location) {
			messages["Location"] = fmt.Sprint(
				impression.Location,
				"is not valid for this SiteKey",
				impression.SiteKey)
		}
	}

	if impression.IpAddress == "" {
		messages["IpAddress"] = emptyOrMissing
	}

	if impression.StartedUtc.IsZero() {
		messages["StartedUtc"] = emptyOrMissing
	}

	if impression.CompletedUtc.IsZero() {
		messages["CompletedUtc"] = emptyOrMissing
	}

	return NewValidationResult(messages), nil
}

func (manager *ImpressionsManager) GetAllImpressions(ctx context.Context) ([]Impression, error) {
	return manager.ImpressionsRepo.GetAllImpressions(ctx)
}

func (manager *ImpressionsManager) GetLocationCounts(ctx context.Context, siteKey string) (map[string]int, error) {
	return manager.ImpressionsRepo.GetLocationCount(ctx, siteKey)
}
