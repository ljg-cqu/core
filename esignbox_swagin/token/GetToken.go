package token

import (
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/long2ice/swagin/router"
	"github.com/wI2L/fizz/markdown"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyiiorw&namespace=opendoc%2Fsaas_api

type GetTokenRequest struct{}

type GetTokenResponse struct {
	Code int                  `json:"code" description:"业务码，0表示成功"`
	Msg  string               `json:"message,omitempty" description:"信息"`
	Data GetTokenResponseData `json:"data,omitempty" description:"业务信息"`
}

type GetTokenResponseData struct {
	Token        string `json:"token" binding:"required" description:"授权码注意：120分钟失效，请在expiresIn参数的有效截止时间失效前重新获取token，建议提前5分钟重新获取。" example:"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJnSWQiOiIyNGM4YTFhOWU3YWM0MTAwOTcxZTJhNTcyMmZkZGE0NCIsImFwcElkIjoiNzQzODkwMjUwMyIsIm9JZCI6IjY1MmNlOWIxNzNhMDQ5M2FhNzE3MzM4OGVhYjg2NzNiIiwidGltZXN0YW1wIjoxNjQ2NDgwODM1NTY1fQ.OshaEH5qv8WXWClStydMO7FrmwiTAm1ZdJg2hDEtDZw"`
	ExpiresIn    string `json:"expiresIn" binding:"required" description:"有效截止时间（毫秒）" example:"1646488035568"`
	RefreshToken string `json:"refreshToken" binding:"required" description:"刷新授权码" example:"a5696858941e06312a5a681c45dc70aa"`
}

func (req *GetTokenRequest) Handler(ctx *gin.Context) {
	parsedResp := GetTokenResponse{}

	restyResp, err := common.Client.R().SetQueryParams(map[string]string{
		"appId":     AppId,
		"secret":    Secret,
		"grantType": GrantType,
	}).SetResult(&parsedResp).Get(EsignSandBoxGetTokenPath)

	if common.WriteErrore(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
		return
	}

	common.WriteOK(ctx, parsedResp.Data)
}

var GetTokenRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyiiorw&namespace=opendoc%2Fsaas_api", "OAuth2.0鉴权方式调用接口说明")
	}

	//var respModelDesc = func() string {
	//	builder := markdown.Builder{}
	//	builder.Table(
	//		[][]string{
	//			[]string{"参数名称", "类型", "必选", "参数说明"},
	//			[]string{"token", "string", "是", "授权码注意：120分钟失效，请在expiresIn参数的有效截止时间失效前重新获取token，建议提前5分钟重新获取。"},
	//			[]string{"expiresIn", "string", "是", "有效截止时间（毫秒）"},
	//			[]string{"refreshToken", "string", "是", "刷新授权码"},
	//		}, []markdown.TableAlignment{
	//			markdown.AlignLeft,
	//			markdown.AlignCenter,
	//			markdown.AlignCenter,
	//			markdown.AlignLeft,
	//		},
	//	)
	//
	//	return builder.String()
	//}

	r := router.New(
		&GetTokenRequest{},
		router.Summary("获取鉴权Token。注意：前端无需单独调用此接口。"),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: GetTokenResponseData{},
			},
			"400": router.ResponseItem{
				Model: common.ErrorResp{},
			},
		}),
	)

	return r
}
