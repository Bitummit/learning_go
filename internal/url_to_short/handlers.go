package urltoshort

import (
	"context"
	"errors"
	"go_api/internal/storage"
	"go_api/internal/utils"
	"go_api/pkg/handler_utils"
	"go_api/pkg/logger"
	"log/slog"
	"net/http"
	"slices"

	"github.com/go-chi/chi"
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


func NewAlias(log *slog.Logger, queryTool storage.QueryFunctionsWithContext) http.HandlerFunc {
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
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, handler_utils.Error("failed to decode request"))
	
			return
		}
		log.Info("Request decoded", slog.Any("request", req))
		
		if err := validator.New().Struct(req); err != nil {
			validErr := err.(validator.ValidationErrors)

			log.Error("invalid request", logger.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, validErr)

			return
		}

		// TODO: move to config
		const aliasLength = 6

		generatedAlias := utils.NewRandomString(aliasLength)

		all_aliases, err := queryTool.GetAllAliases(context.TODO())
		if err != nil {
			log.Error("Can not fetch all aliases", logger.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, handler_utils.Error("Db query error"))
			
			return
		}

		for slices.Contains(all_aliases, generatedAlias) {
			generatedAlias = utils.NewRandomString(aliasLength)
		}

		id, err := queryTool.SaveURL(context.TODO(), req.URL, generatedAlias)
		if err != nil {
			log.Error("Failed to create create new url in DB", logger.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, handler_utils.Error("failed to save url"))
			
			return
		}
		
		log.Info("Created new url", slog.Int64("id", id))
		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, NewAliasResponse{
			Response: handler_utils.OK(),
			Alias: generatedAlias,
		})

	}
}


type RedirectAliasResponse struct {
	Response handler_utils.Response
}


func RedirectAlias(log *slog.Logger, queryTool storage.QueryFunctionsWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const place = "handlers.redirectAlias"
		// log = log.With(
		// 	slog.String("handler", place),
		// 	slog.String("request_id", middleware.GetReqID(r.Context())),
		// )
	
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, handler_utils.Error("alias not found"))
			return
		}

		URL, err := queryTool.GetURL(context.TODO(), alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				log.Error("no such url", logger.Err(err))
				w.WriteHeader(http.StatusBadRequest)
				render.JSON(w, r, handler_utils.Error("no such url"))
				return
			}
			log.Error("Failed to get url from db", logger.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, handler_utils.Error("failed to get url from db"))
		}

		log.Info("Got url", slog.String("url", URL))

		http.Redirect(w, r, URL, http.StatusFound)

	}
}


