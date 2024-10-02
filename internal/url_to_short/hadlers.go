package urltoshort

import (
	"go_api/internal/storage"
	"go_api/pkg/handler_utils"
	"go_api/pkg/logger"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type NewAliasRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type NewAliasResponse struct {
	handler_utils.Response
	Alias string `json:"alias,omitempty"`
}


func NewAlias(log *slog.Logger, urlSaveQuery storage.QueryFunctions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const place = "handlers.newAlias"
		log = log.With(
			slog.String("handler", place),
			slog.String("reques_id", middleware.GetReqID(r.Context())),
		)
	
		var req NewAliasRequest
	
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", logger.Err(err))
			render.JSON(w, r, handler_utils.Error("failed to decode request"))
	
			return
		}
	
		log.Info("Request decoded", slog.Any("request", req))
		
	}
}

