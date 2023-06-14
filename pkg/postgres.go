package pkg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToPostgres(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.TODO(), dsn)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.TODO())
	if err != nil {
		return nil, err
	}

	return pool, err
}
