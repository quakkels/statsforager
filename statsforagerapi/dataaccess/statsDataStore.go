package dataaccess

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type statsDataStore struct {
	db *pgxpool.Pool
}

var pgInstance *statsDataStore
var pgOnce sync.Once

func NewStatsDataStore(ctx context.Context, connString string) (*statsDataStore, error) {
	var poolError error = nil

	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			poolError = fmt.Errorf("Unable to create connection pool: %w", err)
			return
		} else {
			fmt.Println("no error when making pool")
		}

		pgInstance = &statsDataStore{db}
	})

	return pgInstance, poolError
}

func (pg *statsDataStore) Close() {
	pg.db.Close()
}

func (pg *statsDataStore) QueryRow(ctx context.Context, sql string) pgx.Row {
	return pg.db.QueryRow(ctx, sql)
}