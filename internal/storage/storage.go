package storage

import (
	"context"
	"errors"
)

type URL struct {
	Id int64
	Long_url string
	Alias string
}


type QueryFunctions interface {
	SaveURL(string, string) (int64, error)
	GetURL(string) (string, error)
	GetAllAliases() ([]string, error)
	DeleteURL(string) error
}

type QueryFunctionsWithContext interface {
	SaveURL(context.Context, string, string) (int64, error)
	GetURL(context.Context, string) (string, error)
	GetAllAliases(context.Context) ([]string, error)
	DeleteURL(context.Context, string) error
	GetURLList(context.Context) ([]URL, error),
}


var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists = errors.New("url exists")
)