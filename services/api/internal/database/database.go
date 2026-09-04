package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error){
	dsn := "postgres://strata:strata_dev@localhost:5432/strata_db"

	pool,err := pgxpool.New(context.Background(),dsn) 
	if err != nil{
		return nil, fmt.Errorf("Creating Database Connection pool: %w", err)
	}

	err = pool.Ping(context.Background())
	if err != nil{
		pool.Close()
		return nil, fmt.Errorf("creating Database Connection pool: %w", err)
	}

	return pool, nil
}