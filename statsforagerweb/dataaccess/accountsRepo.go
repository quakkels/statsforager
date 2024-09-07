package dataaccess

import (
	"context"
	"errors"
	"fmt"
	"statsforagerweb/domain"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type AccountsRepo struct {
	dataStore statsDataStore
}

func NewAccountsRepo(dataStore statsDataStore) AccountsRepo {
	return AccountsRepo{dataStore}
}

func (repo *AccountsRepo) GetAccountByEmail(context context.Context, email string) (domain.Account, error) {
	var account domain.Account
	err := repo.dataStore.
		QueryRow(context, "SELECT email, is_active FROM accounts WHERE email = $1;", email).
		Scan(&account.Email, &account.IsActive)
	return account, err
}

func (repo *AccountsRepo) SaveAccount(context context.Context, account domain.Account) error {
	const sql = `
	MERGE INTO accounts target
	USING (SELECT 
		@email AS email,
		@isActive::boolean AS is_active) AS source
	ON target.email = source.email
	WHEN NOT MATCHED THEN
		INSERT (email, is_active)
		VALUES (source.email, source.is_active)
	WHEN MATCHED THEN
		UPDATE SET
			email = source.email,
			is_active = source.is_active;
	`

	cmdTag, err := repo.dataStore.Exec(
		context,
		sql,
		pgx.NamedArgs{
			"email":    strings.ToLower(account.Email),
			"isActive": account.IsActive})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println("Postgres error message: ", pgErr.Message)
			fmt.Println("Postgres error code: ", pgErr.Code)
		}
		return err
	}
	if cmdTag.RowsAffected() < 1 {
		return errors.New("Save account affected no rows")
	}
	return nil
}

func (repo *AccountsRepo) RegisterAccount(context context.Context, email string) error {
	const sql = `
	MERGE INTO accounts target
	USING (SELECT 
		@email AS email) AS source
	ON target.email = source.email
	WHEN NOT MATCHED THEN
		INSERT (email)
		VALUES (source.email)
	WHEN MATCHED THEN
		UPDATE SET
			email = source.email;
	`

	cmdTag, err := repo.dataStore.Exec(
		context,
		sql,
		pgx.NamedArgs{"email": strings.ToLower(email)})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println("Postgres error message: ", pgErr.Message)
			fmt.Println("Postgres error code: ", pgErr.Code)
		}
		return err
	}
	if cmdTag.RowsAffected() < 1 {
		return errors.New("Register account affected no rows")
	}
	return nil
}
