package token

import (
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
	"github.com/go-resty/resty/v2"
	"github.com/ljg-cqu/core/esignbox/common"
	"net/http"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyiiorw&namespace=opendoc%2Fsaas_api

type GetTokenResponse struct {
	Code int                  `json:"code" required:"true" doc:"业务码，0表示成功"`
	Msg  string               `json:"message,omitempty" doc:"信息"`
	Data GetTokenResponseData `json:"data,omitempty" doc:"业务信息"`
}

type GetTokenResponseData struct {
	Token        string `json:"token" doc:"授权码注意：120分钟失效，请在expiresIn参数的有效截止时间失效前重新获取token，建议提前5分钟重新获取。"`
	ExpiresIn    string `json:"expiresIn" doc:" 有效截止时间（毫秒）"`
	RefreshToken string `json:"refreshToken" doc:"刷新授权码"`
}

func GetToken(client *resty.Client) func(ctx huma.Context) {
	return func(ctx huma.Context) {
		parsedResp := GetTokenResponse{}

		restyResp, err := client.R().SetQueryParams(map[string]string{
			"appId":     AppId,
			"secret":    Secret,
			"grantType": GrantType,
		}).SetResult(&parsedResp).Get(EsignSandBoxGetTokenPath)

		if common.WriteError(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
			return
		}

		ctx.WriteModel(http.StatusOK, parsedResp.Data)
	}
}

func RunGetToken(r *huma.Resource, client *resty.Client) {
	r.Get("GetToken", "获取鉴权Token",
		responses.BadRequest(),
		responses.OK().Model(GetTokenResponseData{}),
	).Run(GetToken(client))
}
