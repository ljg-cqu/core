package main

import (
	"github.com/danielgtaylor/huma"
	"net/http"
)

// ---------return the cache's statistics---------

type Stats struct {
	Hits       int64 `json:"hits" example:"10" doc:"Hits is a number of successfully found keys"`
	Misses     int64 `json:"misses" example:"2" doc:"Misses is a number of not found keys"`
	DelHits    int64 `json:"delete_hits" example:"1" doc:"DelHits is a number of successfully deleted keys"`
	DelMisses  int64 `json:"delete_misses" example:"1" doc:"DelMisses is a number of not deleted keys"`
	Collisions int64 `json:"collisions" example:"1" doc:"Collisions is a number of happened key-collisions"`
}

func getCacheStatsHandler(ctx huma.Context, input GetCacheRequest) {
	_stats := cache.Stats()
	var stats Stats
	stats.Hits = _stats.Hits
	stats.Misses = _stats.Misses
	stats.DelHits = _stats.DelHits
	stats.DelMisses = _stats.DelMisses
	stats.Collisions = _stats.Collisions

	ctx.WriteModel(http.StatusOK, &stats)
	return
}
