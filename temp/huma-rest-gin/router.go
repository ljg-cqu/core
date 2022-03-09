package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest"
	"github.com/swaggest/rest/chirouter"
	"github.com/swaggest/rest/jsonschema"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/openapi"
	"github.com/swaggest/rest/request"
	"github.com/swaggest/rest/response"
	"github.com/swaggest/rest/response/gzip"
	v4 "github.com/swaggest/swgui/v4"
	"net/http"
)

func NewRouter() http.Handler {
	apiSchema := &openapi.Collector{}
	validatorFactory := jsonschema.NewFactory(apiSchema, apiSchema)
	decoderFactory := request.NewDecoderFactory()
	decoderFactory.SetDecoderFunc(rest.ParamInPath, chirouter.PathToURLValues)

	apiSchema.Reflector().SpecEns().Info.Title = "Advanced Example"
	apiSchema.Reflector().SpecEns().Info.WithDescription("This app showcases a variety of features.")
	apiSchema.Reflector().SpecEns().Info.Version = "v1.2.3"

	r := chirouter.NewWrapper(chi.NewRouter())

	r.Use(
		middleware.Recoverer,                          // Panic recovery.
		nethttp.OpenAPIMiddleware(apiSchema),          // Documentation collector.
		request.DecoderMiddleware(decoderFactory),     // Request decoder setup.
		request.ValidatorMiddleware(validatorFactory), // Request validator setup.
		response.EncoderMiddleware,                    // Response encoder setup.

		// Response validator setup.
		//
		// It might be a good idea to disable this middleware in production to save performance,
		// but keep it enabled in dev/test/staging environments to catch logical issues.
		response.ValidatorMiddleware(validatorFactory),
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)

	// Annotations can be used to alter documentation of operation identified by method and path.
	apiSchema.Annotate(http.MethodPost, "/validation", func(op *openapi3.Operation) error {
		if op.Description != nil {
			*op.Description = *op.Description + " Custom annotation."
		}

		return nil
	})

	r.Method(http.MethodPost, "/file-multi-upload", nethttp.NewHandler(fileMultiUploader()))

	// Swagger UI endpoint at /docs.
	r.Method(http.MethodGet, "/docs/openapi.json", apiSchema)
	r.Mount("/docs", v4.NewHandler(apiSchema.Reflector().Spec.Info.Title,
		"/docs/openapi.json", "/docs"))

	return r
}
