package urltoshort

import (
	"go_api/internal/storage"
	"go_api/internal/utils"
	"go_api/pkg/handler_utils"
	"go_api/pkg/logger"
	"log/slog"
	"net/http"
	"slices"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

)

type NewAliasRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type NewAliasResponse struct {
	Response handler_utils.Response
	Alias string `json:"alias,omitempty"`
}


func NewAlias(log *slog.Logger, queryTool storage.QueryFunctions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const place = "handlers.newAlias"
		log = log.With(
			slog.String("handler", place),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
	
		var req NewAliasRequest
	
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", logger.Err(err))
			w.WriteHeader(500)
			render.JSON(w, r, handler_utils.Error("failed to decode request"))
	
			return
		}
		log.Info("Request decoded", slog.Any("request", req))
		
		if err := validator.New().Struct(req); err != nil {
			validErr := err.(validator.ValidationErrors)

			log.Error("invalid request", logger.Err(err))
			w.WriteHeader(400)
			render.JSON(w, r, validErr)

			return
		}

		// TODO: move to config
		const aliasLength = 6

		generatedAlias := utils.NewRandomString(aliasLength)

		all_aliases, err := queryTool.GetAllAliases()
		if err != nil {
			log.Error("Can not fetch all aliases", logger.Err(err))
			w.WriteHeader(500)
			render.JSON(w, r, handler_utils.Error("Db query error"))
			
			return
		}

		for slices.Contains(all_aliases, generatedAlias) {
			generatedAlias = utils.NewRandomString(aliasLength)
		}

		id, err := queryTool.SaveURL(req.URL, generatedAlias)
		if err != nil {
			log.Error("Failed to create create new url in DB", logger.Err(err))
			w.WriteHeader(500)
			render.JSON(w, r, handler_utils.Error("failed to save url"))
			
			return
		}
		
		log.Info("Created new url", slog.Int64("id", id))

		render.JSON(w, r, NewAliasResponse{
			Response: handler_utils.OK(),
			Alias: generatedAlias,
		})

	}
}

