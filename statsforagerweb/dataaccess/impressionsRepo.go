package dataaccess

import (
	"context"
	"errors"
	"fmt"
	"statsforagerweb/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type ImpressionsRepo struct {
	dataStore statsDataStore
}

func NewImpressionsRepo(dataStore statsDataStore) ImpressionsRepo {
	return ImpressionsRepo{dataStore}
}

func (impRepo *ImpressionsRepo) SaveImpression(
	context context.Context,
	impression domain.Impression) error {

	const sql = `
	
	MERGE INTO impressions target
	USING (SELECT 
		@impressionId::uuid AS impression_id,
		@ipAddress::inet AS ip_address,
		@userAgent AS user_agent,
		@language AS language,
		@location AS location,
		@referrer AS referrer,
		@siteKey::uuid AS site_key,
		@startedUtc::timestamp AS started_utc,
		@completedUtc::timestamp AS completed_utc) AS source
	ON target.impression_id = source.impression_id
	WHEN NOT MATCHED THEN
		INSERT (impression_id,  ip_address, user_agent, language,  location,  referrer,  site_key, started_utc,	completed_utc)
		VALUES (source.impression_id, source.ip_address, source.user_agent, source.language,
		source.location, source.referrer, source.site_key, source.started_utc, source.completed_utc)
	WHEN MATCHED THEN
		UPDATE SET
			impression_id = source.impression_id,
			ip_address = source.ip_address,
			user_agent = source.user_agent,
			language = source.language,
			location = source.location,
			referrer = source.referrer,
			site_key = source.site_key,
			started_utc = source.started_utc,
			completed_utc = source.completed_utc;
	`

	cmdTag, err := impRepo.dataStore.Exec(
		context,
		sql,
		pgx.NamedArgs{
			"impressionId": impression.ImpressionId,
			"ipAddress":    impression.IpAddress,
			"userAgent":    impression.UserAgent,
			"language":     impression.Language,
			"location":     impression.Location,
			"referrer":     impression.Referrer,
			"siteKey":      impression.SiteKey,
			"startedUtc":   impression.StartedUtc,
			"completedUtc": impression.CompletedUtc})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println("Postgres error message: ", pgErr.Message)
			fmt.Println("Postgres error code: ", pgErr.Code)
		}
		return err
	}
	if cmdTag.RowsAffected() < 1 {
		return errors.New("Insert impression affected no rows")
	}
	return nil
}

func (repo *ImpressionsRepo) GetAllImpressions(ctx context.Context) ([]domain.Impression, error) {
	fmt.Println("In repo.GetAllImpressions()")
	var impressions []domain.Impression
	const sql = "SELECT impression_id, ip_address, user_agent, language, location, referrer, site_key, started_utc,	completed_utc FROM impressions;"
	rows, err := repo.dataStore.Query(ctx, sql)
	defer rows.Close()
	if err != nil {
		fmt.Println("line 90")
		return impressions, nil
	}

		fmt.Println("line 94")
	for rows.Next() {
		fmt.Println("line 96 in loop")
		var impression domain.Impression
		err := rows.Scan(
			&impression.ImpressionId,
			&impression.IpAddress,
			&impression.UserAgent,
			&impression.Language,
			&impression.Location,
			&impression.Referrer,
			&impression.SiteKey,
			&impression.StartedUtc,
			&impression.CompletedUtc)

		if err != nil {
		fmt.Println("line 110 in error")//can't scan into dest[1]: cannot scan inet (OID 869) in binary format into *s

			return nil, err
		}

		impressions = append(impressions, impression)
	}

	err = rows.Err() // get error from rows.Next() or rows.Scan()

	fmt.Println("Impression count", len(impressions))
	return impressions, err
}
