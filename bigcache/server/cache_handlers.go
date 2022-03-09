package main

import (
	"github.com/allegro/bigcache/v3"
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/middleware"
	"github.com/pkg/errors"
	"net/http"
)

// ---------handles GET requests----------

type GetCacheRequest struct {
	Key string `path:"key" minLength:"1" maxLength:"36" example:"my-unique-key" doc:"The unique key for value to be cached"`
}

type GetCacheResponse struct {
	Value string `json:"value" example:"a9b3f69f7fedb34d20c80087c969d4689c3d94aaa69a8a2c4eb714d3b8e1daab" doc:"The value stored in cache"`
}

func getCacheHandler(ctx huma.Context, input GetCacheRequest) {
	logger := middleware.GetLogger(ctx)

	if ctx.HasError() {
		logger.Errorw("Got a bad request", "key", input.Key)
		ctx.WriteError(http.StatusBadRequest, "Bad input")
	}

	entry, err := cache.Get(input.Key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			logger.Infow("Entry not found", "key", input.Key)
			ctx.WriteHeader(http.StatusNotFound)
			return
		}

		logger.Errorw("Fatal: failed to get entry", "error", err)
		ctx.WriteError(http.StatusInternalServerError, "Fatal: failed to get entry", err)
		return
	}

	ctx.WriteModel(http.StatusFound, &GetCacheResponse{Value: string(entry)})

	logger.Infof("Responsed object for key %q", input.Key)
}

// ---------handles PUT request----------

type PutCacheRequest struct {
	Value string `json:"value" example:"a9b3f69f7fedb34d20c80087c969d4689c3d94aaa69a8a2c4eb714d3b8e1daab" doc:"The value to be stored in cache"`
}

type PutCacheRequest_ struct {
	Key  string          `path:"key" minLength:"1" maxLength:"36" example:"my-unique-key" doc:"The unique key for value to be cached"`
	Body PutCacheRequest `format:"binary" doc:"The binary content to be cached"`
}

func putCacheHandler(ctx huma.Context, input PutCacheRequest_) {
	logger := middleware.GetLogger(ctx)

	if ctx.HasError() {
		logger.Errorw("Got a bad request", "key", input.Key)
		ctx.WriteError(http.StatusBadRequest, "Bad input")
	}

	if err := cache.Set(input.Key, []byte(input.Body.Value)); err != nil {
		logger.Errorw("Fatal: failed to cache entry", "error", err)
		ctx.WriteError(http.StatusInternalServerError, "Fatal: failed to cache entry", err)
		return
	}

	logger.Infof("Stored %q in cache", input.Key)

	ctx.WriteHeader(http.StatusCreated)
}

// ---------handles DELETE request----------

type DeleteCacheRequest struct {
	Key string `path:"key" minLength:"1" maxLength:"36" example:"my-unique-key" doc:"The unique key for value to be cached"`
}

func deleteCacheHandler(ctx huma.Context, input DeleteCacheRequest) {
	logger := middleware.GetLogger(ctx)

	if ctx.HasError() {
		logger.Errorw("Got a bad request", "key", input.Key)
		ctx.WriteError(http.StatusBadRequest, "Bad input")
	}

	if err := cache.Delete(input.Key); err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			logger.Infow("Entry not found", "key", input.Key)
			ctx.WriteHeader(http.StatusNotFound)
			return
		}

		logger.Errorw("Fatal: failed to delete entry", "error", err)
		ctx.WriteError(http.StatusInternalServerError, "Fatal: failed to delete entry", err)
		return
	}

	logger.Infof("Delete %q from cache", input.Key)

	// this is what the RFC says to use when calling DELETE.
	ctx.WriteHeader(http.StatusOK)
}
