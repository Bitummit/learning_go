package postgresql

import (
	"context"
	"go_api/internal/storage"
	"os"
	"time"

	// "github.com/jackc/pgx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


type Storage struct {
	DB *pgxpool.Pool
}

func NewConnectionPool(ctx context.Context) (*Storage, error) {
	ctx, cancel := context.WithTimeout(ctx, 90 * time.Second)
	defer cancel()

	database_url := os.Getenv("DATABASE_URL")
	db, err := pgxpool.New(ctx, database_url)
	if err != nil {
		return nil, err
	}

	return &Storage{DB:db}, nil

}

func (s *Storage) SaveURL(ctx context.Context, URL string, alias string) (int64, error) {
	query := `
		INSERT INTO url (url, alias) VALUES(@URL, @alias) RETURNING id;
	`
	args := pgx.NamedArgs{
		"URL": URL,
		"alias": alias,
	}

	var id int64

	err := s.DB.QueryRow(ctx, query, args).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}


func(s *Storage) GetURL(ctx context.Context, alias string) (string, error) {
	query := `
		SELECT url FROM url where alias=@alias
	`
	args := pgx.NamedArgs{
		"alias": alias,
	}
	var URL string
	err := s.DB.QueryRow(ctx, query, args).Scan(&URL)
	if err != nil {
		return "", nil
	}

	return URL, nil
}


func (s *Storage) GetAllAliases(ctx context.Context) ([]string, error) {
	query := `
		SELECT alias FROM url
	`
	rows, err := s.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aliases []string

	for rows.Next() {
		var alias string
		err = rows.Scan(&alias)
		if err != nil {
			return nil, err
		}
		aliases = append(aliases, alias)
	}

	return aliases, nil
}


func (s *Storage) DeleteURL(ctx context.Context, alias string) error {
	query := `
		DELETE FROM url WHERE alias=@alias
	`
	args := pgx.NamedArgs{
		"alias": alias,
	}
	_, err := s.DB.Exec(ctx, query, args)
	
	return err

}


func (s *Storage) GetURLList(ctx context.Context) ([]storage.URL, error) {
	var URLs []storage.URL
	query := `
		SELECT * FROM url;
	`
	rows, err := s.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var URL storage.URL
		err := rows.Scan(&URL.Id, &URL.Long_url, &URL.Alias)
		if err != nil {
			return nil, err
		}
		URLs = append(URLs, URL)
	}

	return URLs, nil
}