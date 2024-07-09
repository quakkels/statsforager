package domain

import (
	"context"
	"errors"
	"time"
)

type ImpressionsManager struct {
	ImpressionRepo ImpressionRepository
}

type Impression struct {
	ImpressionId string
	IpAddress    string
	UserAgent    string
	Language     string
	Location     string
	Referrer     string
	SiteKey      string
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

func (manager *ImpressionsManager) SaveImpression(impression Impression) []error {
	return make([]error, errors.New("Not yet implemented"))
}

