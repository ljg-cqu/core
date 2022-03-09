package template

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/token"
	"github.com/long2ice/swagin/router"
	"github.com/wI2L/fizz/markdown"
)

const (
	EsignSandBoxDeleteFillControlUrl = "/v1/docTemplates/{templateId}/components/{ids}"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fgy0qlv&namespace=opendoc%2Fsaas_api

type DeleteFillControlRequest struct {
	TemplateId string `uri:"templateId" binding:"required" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板id"`
	IDs        string `uri:"ids"  binding:"required"  description:"输入项组件ID集合，英文逗号分隔"`
}

type _DeleteFillControlResponse struct {
	Code int    `json:"code" required:"true" doc:"业务码，0表示成功"`
	Msg  string `json:"message" doc:"信息"`
}

type DeleteFillControlResponse struct {
	TemplateId string `json:"templateId" binding:"required" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板id"`
	IDs        string `json:"ids"  binding:"required"  description:"输入项组件ID集合，英文逗号分隔"`
}

func (req *DeleteFillControlRequest) Handler(ctx *gin.Context) {
	parsedResp := _DeleteFillControlResponse{}
	oauth, err := token.GetOauthInfo()
	if err != nil {
		common.WriteErrorE(ctx, fmt.Errorf("got an error when try to get authentication info:%w", err), 0, "", 0, "")
		return
	}

	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
		"Content-Type":        oauth.ContentType,
	}).SetBody(&req).
		SetResult(&parsedResp).Delete("/v1/docTemplates/" + req.TemplateId + "/components/" + req.IDs)
	if err != nil {
		common.WriteErrorF(ctx, 400, "got an error when try to delete fill control, error:%v", err)
		return
	}

	if parsedResp.Code != 0 {
		common.WriteErrorF(ctx, 400, "got an error when try to delete fill control, code:%v, message:%v", parsedResp.Code, parsedResp.Msg)
		return
	}

	_ = restyResp

	common.WriteOK(ctx, DeleteFillControlResponse{req.TemplateId, req.IDs})
}

var DeleteFillControlRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("官方文档：")
		return builder.String() + builder.Link(" https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fgy0qlv&namespace=opendoc%2Fsaas_api", "删除填写控件")
	}

	r := router.New(
		&DeleteFillControlRequest{},
		router.Summary("删除填写控件。"),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: DeleteFillControlResponse{},
			},
			"400": router.ResponseItem{
				Model: common.ErrorResp{},
			},
		}),
	)
	return r
}
