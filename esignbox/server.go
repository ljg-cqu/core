package main

import (
	"context"
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/middleware"
	"github.com/go-chi/chi"
	"github.com/go-resty/resty/v2"
	"github.com/ljg-cqu/core/esignbox/template"
	"github.com/ljg-cqu/core/esignbox/token"
	"go.uber.org/zap"
)

const (
	EsignSandBoxHost = "https://smlopenapi.esign.cn"
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

	client := resty.New().SetDebug(true).SetBaseURL(EsignSandBoxHost)

	app = cli.NewRouter("e签宝模板管理API\ne签宝模板管理官方API：https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fawgyis&namespace=opendoc%2Fsaas_api", version)

	app.Contact("罗继高", "ljg_cqu@126.com", "https://github.com/ljg-cqu/core")

	app.DocsHandler(huma.SwaggerUIHandler(huma.New("Test API", version)))

	tokenResource := app.Resource("/esignbox/token")
	token.RunGetToken(tokenResource, client)

	templResource := app.Resource("/esignbox/template/info")
	template.RunGetTemplInfo(templResource, client)

	componentResource := app.Resource("/esignbox/template/{templateId}/components")
	template.RunAddFillControl(componentResource, client)

	templUploadResource := app.Resource("/esignbox/template/upload_url")
	template.RunGetTemplUploadUrl(templUploadResource, client)
	template.RunUploadDocTemplFile(templUploadResource, client)

	templUploadStatusResource := app.Resource("/esignbox/template/{templateId}/upload_status")
	template.RunQueryTemplUploadStatus(templUploadStatusResource, client)

	templDetailsResource := app.Resource("/esignbox/template/{templateId}/details")
	template.RunQueryTemplDetails(templDetailsResource, client)

	fillTemplResource := app.Resource("/esignbox/template/fill")
	template.RunFillTemplateContent(fillTemplResource, client)
	pdfFileDetailsResource := app.Resource("/esignbox/pdf_file/{fileId}/details")
	template.RunQueryPdfFileDetails(pdfFileDetailsResource, client)

	deleteFillControlResource := app.Resource("/esignbox/docTemplate/{templateId}/components/{ids}")
	template.RunDeleteFillControl(deleteFillControlResource, client)

	app.Run()
}
