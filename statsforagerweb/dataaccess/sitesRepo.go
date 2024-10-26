package dataaccess

import (
	"context"
	"errors"
	"fmt"
	"statsforagerweb/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
		QueryRow(context, "SELECT site_key, domain, owner_account, site_name FROM sites WHERE site_key=$1;", siteKey).
		Scan(&site.SiteKey, &site.Domain, &site.OwnerAccount, &site.SiteName)
	return site, err
}

func (repo *SitesRepo) GetAllSites(context context.Context) ([]domain.Site, error) {
	rows, err := repo.dataStore.Query(context, "SELECT domain, site_key, owner_account, site_name FROM sites;")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	return mapRowsToSites(rows)
}

func (repo *SitesRepo) GetSites(context context.Context, accountCode string) ([]domain.Site, error) {
	rows, err := repo.dataStore.Query(context, "SELECT domain, site_key, owner_account, site_name FROM sites WHERE owner_account=$1;", accountCode)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	return mapRowsToSites(rows)
}

func (repo *SitesRepo) SaveSite(context context.Context, site domain.Site) error {
	const sql = `
		MERGE INTO sites target
		USING (SELECT 
				@siteKey::uuid AS site_key,
				@domain AS domain,
				@ownerAccount AS owner_account
				@siteName AS site_name) AS source
			ON target.site_key = source.site_key
		WHEN NOT MATCHED THEN
			INSERT (site_key, domain, owner_account, site_name)
			VALUES (source.site_key, source.domain, source.owner_account, source.site_name)
		WHEN MATCHED THEN
			UPDATE SET
			site_key = source.site_key,
			domain = source.domain,
			owner_account = source.owner_account,
			site_name = source.site_name;
	`

	cmdTag, err := repo.dataStore.Exec(
		context,
		sql,
		pgx.NamedArgs{
			"siteKey":      site.SiteKey,
			"domain":       site.Domain,
			"ownerAccount": site.OwnerAccount,
			"siteName": site.SiteName})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println("sitesRepo.go: SaveSite> Postgres error message: ", pgErr.Message)
			fmt.Println("sitesRepo.go: SaveSite> Postgres error code: ", pgErr.Code)
		}
		return err
	}
	if cmdTag.RowsAffected() < 1 {
		return errors.New("Save Site affected no rows")
	}
	return nil
}

func mapRowsToSites(rows pgx.Rows) ([]domain.Site, error) {
	var sites []domain.Site
	for rows.Next() {
		var site domain.Site
		err := rows.Scan(&site.Domain, &site.SiteKey, &site.OwnerAccount, &site.SiteName)
		if err != nil {
			return nil, err
		}

		sites = append(sites, site)
	}

	err := rows.Err() // get error from rows.Next() or rows.Scan()
	return sites, err
}
