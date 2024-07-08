package dataaccess

import (
	"context"
	"errors"
	"statsforagerapi/domain"
)

type ImpressionsRepo struct {
	dataStore statsDataStore
}

func NewImpressionsRepo(dataStore statsDataStore) ImpressionsRepo {
	return ImpressionsRepo{dataStore}
}

func (impRepo ImpressionsRepo) SaveImpression(
	context context.Context,
	impression domain.Impression) error {

	const sql string = `
			INSERT INTO impressions 
			(
				impression_id,
				ip_address,
				user_agent,
				language,
				location,
				referrer,
				date_time,
				event_date_time_utc,
				site_key,
				is_leaving) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	cmdTag, err := impRepo.dataStore.Exec(
		context,
		sql,
		impression.ImpressionId,
		impression.IpAddress,
		impression.UserAgent,
		impression.Language,
		impression.Location,
		impression.Referrer,
		impression.DateTime,
		impression.EventDateTimeUtc,
		impression.SiteKey,
		impression.IsLeaving)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() < 1 {
		return errors.New("Insert impression affected no rows")
	}
	return nil
}
