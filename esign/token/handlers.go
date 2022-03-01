package main

import (
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/middleware"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func GetToken(client *resty.Client) func(ctx huma.Context, req GetRequest) {
	return func(ctx huma.Context, req GetRequest) {
		logger := middleware.GetLogger(ctx)

		if ctx.HasError() {
			logger.Errorw("Got a bad request", "key", spew.Sdump(&req))
			ctx.WriteError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}

		res := GetResponse{}

		_, err := client.R().SetQueryParams(map[string]string{
			"appId":     req.AppId,
			"secret":    req.Secret,
			"grantType": req.GrantType,
		}).SetResult(&res).Get(esignSandBoxHost + getTokenPath)

		if err != nil {
			logger.Errorw("Fatal: failed to get token from esign", "key", "error", err.Error())
			ctx.AddError(err)
			ctx.WriteError(http.StatusInternalServerError, "failed to get token from esign", err)
			return
		}

		logger.Debugf("Get a token from esign:%v", res.Data)
		ctx.WriteModel(http.StatusOK, res)
	}
}
