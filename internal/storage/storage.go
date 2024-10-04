package storage

import "errors"

type URL struct {
	id int64
	long_url string
	aslias string
}


type QueryFunctions interface {
	SaveURL(string, string) (int64, error)
	GetURL(string) (string, error)
	GetAllAliases() ([]string, error)
	DeleteURL(string) error
}


var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists = errors.New("url exists")
)