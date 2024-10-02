package sqlite

import (
	"database/sql"
	"errors"
	"go_api/pkg/storage"

	"github.com/mattn/go-sqlite3"
)


type Storage struct {
	db *sql.DB
}

func New(storagePAth string) (*Storage, error){

	db, err := sql.Open("sqlite3", "./storage/storage.db")
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url (
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(); if err != nil {
		return nil, err
	}

	return &Storage{db:db}, nil
}

func (s *Storage) SaveURL(URL string, alias string) (int64, error) {
	
	stmt, err := s.db.Prepare("INSERT INTO url(url, alias) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(URL, alias)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintCheck{
			return 0, storage.ErrURLExists
		}
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	stmt, err := s.db.Prepare("SELECT alias FROM url WHERE alias=?")
	if err != nil {
		return "", err
	}

	var URL string
	err = stmt.QueryRow(alias).Scan(&URL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrURLNotFound
		}
		return "", err
	}

	return URL, nil
}