package domain

import (
	"context"
	"time"
)

type ImpressionsManager struct {
	ImpressionRepo ImpressionRepository
}

type Impression struct {
	ImpressionId     string
	IpAddress        string
	UserAgent        string
	Language         string
	Location         string
	Referrer         string
	DateTime         time.Time
	IsLeaving        bool
	EventDateTimeUtc time.Time
	SiteKey          string
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
