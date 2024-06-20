package dataaccess

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

var pgInstance *postgres
var pgOnce sync.Once

func NewConnPool(ctx context.Context, connString string) (*postgres, error) {

	var poolError error = nil

	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			poolError = fmt.Errorf("Unable to create connection pool: %w", err)
			return
		} else {
			fmt.Println("no error when making pool")
		}

		pgInstance = &postgres{db}
	})
	
	return pgInstance, poolError
}

func (pg *postgres) Close() {
	pg.db.Close()
}

func (pg *postgres) QueryRow(ctx context.Context, sql string) pgx.Row {
	return pg.db.QueryRow(ctx, sql)
}
