package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
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

type ImpressionsManager struct {
	ImpressionRepo ImpressionRepository
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

type ImpressionRepository interface {
	SaveImpression(context.Context, Impression) error
}

func NewImpressionsManager(repo ImpressionRepository) ImpressionsManager {
	manager := ImpressionsManager{
		ImpressionRepo: repo,
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

	err = manager.ImpressionRepo.SaveImpression(context, dto)
	return vr, err
}

func (manager *ImpressionsManager) ValidateImpression(
	context context.Context,
	impression Impression) (*validationResult, error) {
	const emptyOrMissing = "Empty or missing value"
	const notUuid = "Not a UUID"
	messages := make(map[string]string)

	if err := uuid.Validate(impression.SiteKey); err != nil {
		fmt.Println("impression.SiteKey", impression.SiteKey)
		messages["SiteKey"] = notUuid
	}

	if err := uuid.Validate(impression.ImpressionId); err != nil {
		messages["ImpressionId"] = notUuid
	}

	if impression.UserAgent == "" {
		messages["UserAgent"] = emptyOrMissing
	}

	if impression.Location == "" {
		messages["Location"] = emptyOrMissing
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
