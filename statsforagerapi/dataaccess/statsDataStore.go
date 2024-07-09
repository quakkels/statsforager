package dataaccess

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
		}

		err = db.Ping(ctx)
		if err != nil {
			poolError = fmt.Errorf("Unable to establish connection with database: %w", err)
		}

		if err == nil {
			fmt.Println("Created db connection pool and established connection")
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

func (pg *statsDataStore) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return pg.db.Exec(ctx, sql, arguments)
}
