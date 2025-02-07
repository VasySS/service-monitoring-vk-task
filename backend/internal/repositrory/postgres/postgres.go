package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type Postgres struct {
	db *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Postgres {
	return &Postgres{
		db: pool,
	}
}
