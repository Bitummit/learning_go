package urltoshort

import (
	"go_api/internal/storage"
	"log/slog"
	"net/http"
)

type Request struct {
	URL string `json:"url" validate:"required,url"`
}

type Response struct {
	Status string `json:"status"`
	Error string `json:"error,omitempty"`
	Alias string `json:"alias,omitempty"`
}


func NewAlias(log *slog.Logger, urlSaver storage.QueryFunctions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}