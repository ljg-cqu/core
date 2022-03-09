package main

import (
	"context"
	"fmt"
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/middleware"
	"github.com/danielgtaylor/huma/responses"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

const (
	// API version.
	version = "1.0.0"
)

func main() {
	router := huma.New("Contract review API\nThis is just a demonstration of contract review API.", version)
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
		middleware.PreferMinimal)

	app = cli.NewRouter("Contract review API\nThis is just a demonstration of contract review API.", version)

	app.Contact("zealy", "ljg_cqu@126.com", "https://github.com/ljg-cqu/core")

	app.DocsHandler(huma.SwaggerUIHandler(huma.New("Test API", version)))

	abfpaas := app.Resource("/v1/abfpaas")
	review := abfpaas.SubResource("/review")

	review.SubResource("/cancel/{session-id}").
		Post("review-cancel", "Cancel contract review", responses.OK().Model(ResponseOk{})).Run(Cancel)
	//review.SubResource("/at").
	//	Post("review-at", "At someone something", responses.OK().Model(ResponseOk{})).Run(At)
	//review.SubResource("/comment").
	//	Post("review-comment", "Comment in a contract session", responses.OK().Model(ResponseOk{})).Run(Comment)

	//
	//review := app.Resource("/v1/contract-review/cancel/{review-api}")
	//review.Post("contract-review-cancel", "Cancel contract review", responses.OK().Model(ResponseOk{})).Run(Cancel)
	//review.Post("contract-review-at", "At someone something", responses.OK().Model(ResponseOk{})).Run(At)

	fmt.Printf(app.OpenAPI().StringIndent("", " "))
	app.Run()
}
