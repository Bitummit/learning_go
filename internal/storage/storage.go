package storage

import "errors"

type QueryFunctions interface {
	SaveURL(string, string) (int64, error)
	GetURL(string) (string, error)
	DeleteURL(string) error
}

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists = errors.New("url exists")
)