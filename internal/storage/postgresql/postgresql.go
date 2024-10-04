package postgresql

import (
	"context"
	"os"

	// "github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)


type storage struct {
	db *pgxpool.Pool
}

func NewConnectionPool(ctx context.Context) (*storage, error) {
	database_url := os.Getenv("DATABASE_URL")
	db, err := pgxpool.New(ctx, database_url)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS url (
		id SERIAL PRIMARY KEY,
		url VARCHAR(256) NOT NULL,
		alias VARCHAR(256) NOT NULL UNIQUE
		);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		return nil, err
	}

	return &storage{db:db}, nil

}