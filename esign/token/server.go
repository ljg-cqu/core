package main

import (
	"context"
	"fmt"
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/middleware"
	"github.com/danielgtaylor/huma/responses"
	"github.com/go-chi/chi"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

const (
	// API version.
	version = "1.0.0"

	esignSandBoxHost = "https://smlopenapi.esign.cn"
	getTokenPath     = "/v1/oauth2/access_token"
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

	client := resty.New()

	app = cli.NewRouter("Esign token API\nThis is esign token API.", version)

	app.Contact("zealy", "ljg_cqu@126.com", "https://github.com/ljg-cqu/core")

	app.DocsHandler(huma.SwaggerUIHandler(huma.New("Test API", version)))

	token := app.Resource("/esign/token")
	token.Get("get-token", "Get a token from esign",
		responses.OK().Model(GetResponse{}),
		responses.BadRequest(),
		responses.InternalServerError(),
	).Run(GetToken(client))

	fmt.Printf(app.OpenAPI().StringIndent("", " "))
	app.Run()
}
