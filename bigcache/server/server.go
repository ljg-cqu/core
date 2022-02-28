package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/middleware"
	"github.com/danielgtaylor/huma/responses"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"log"
)

const (
	// server version.
	version = "1.0.0"
)

var (
	logfile string

	// cache-specific settings.
	cache  *bigcache.BigCache
	config = bigcache.Config{}
)

func init() {
	flag.BoolVar(&config.Verbose, "v", false, "Verbose logging.")
	flag.IntVar(&config.Shards, "shards", 1024, "Number of shards for the cache.")
	flag.IntVar(&config.MaxEntriesInWindow, "maxInWindow", 1000*10*60, "Used only in initial memory allocation.")
	flag.DurationVar(&config.LifeWindow, "lifetime", 100000*100000*60, "Lifetime of each cache object.")
	flag.IntVar(&config.HardMaxCacheSize, "max", 8192, "Maximum amount of data in the cache in MB.")
	flag.IntVar(&config.MaxEntrySize, "maxShardEntrySize", 500, "The maximum size of each object stored in a shard. Used only in initial memory allocation.")
	flag.StringVar(&logfile, "logfile", "", "Location of the logfile.")
}

func main() {
	flag.Parse()

	fmt.Printf("BigCache HTTP Server v%s", version)

	var err error
	cache, err = bigcache.NewBigCache(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("cache initialised.")
	router := huma.New("Cache API\nThis cache api is implemented base on allegro/bigcache.", version)
	app := cli.New(router)

	l, err := middleware.NewDefaultLogger()
	if err != nil {
		panic(err)
	}

	l = l.With()
	middleware.NewLogger = func() (*zap.Logger, error) {
		return l, nil
	}

	app.Middleware(
		middleware.OpenTracing,
		middleware.Logger,
		middleware.Recovery(func(ctx context.Context, err error, request string) {
			log := middleware.GetLogger(ctx)
			log = log.With(zap.Error(err))
			log.With(
				zap.String("http.request", request),
				zap.String("http.template", chi.RouteContext(ctx).RoutePattern()),
			).Error("Caught panic")
		}),
		middleware.ContentEncoding,
		middleware.PreferMinimal,
		requestMetrics(l))

	app = cli.NewRouter("Cache API\nThis cache api is implemented base on allegro/bigcache.", version)

	app.Contact("zealy", "ljg_cqu@126.com", "https://github.com/ljg-cqu/core/tree/main/bigcache/server")

	app.DocsHandler(huma.SwaggerUIHandler(huma.New("Test API", version)))
	//app.DocsHandler(huma.ReDocHandler(huma.New("Test API", version)))
	//app.DocsHandler(huma.RapiDocHandler(huma.New("Test API", version)))

	cache := app.Resource("/v1/cache/{key}")
	cache.Get("getCache", "Return value for the given key", responses.OK().Model(&GetCacheResponse{})).Run(getCacheHandler)
	cache.Put("putCache", "Store value in cache").Run(putCacheHandler)
	cache.Delete("deleteCache", "Delete value in cache").Run(deleteCacheHandler)

	stats := app.Resource("/v1/stats/{key}")
	stats.Get("getStats", "Return the cache's statistics", responses.OK().Model(&Stats{})).Run(getCacheStatsHandler)

	app.Run()
}
