package template

import (
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
	"github.com/go-resty/resty/v2"
	"github.com/ljg-cqu/core/esignbox/common"
	"github.com/ljg-cqu/core/esignbox/token"
	"net/http"
)

const (
	EsignSandBoxDeleteFillControlUrl = "/v1/docTemplates/{templateId}/components/{ids}"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fgy0qlv&namespace=opendoc%2Fsaas_api

type DeleteFillControlRequest struct {
	TemplateId string `json:"templateId" required:"true" example:"c4d4fe1b48184ba28982f68bf2c7bf25" path:"templateId" doc:"模板id"`
	IDs        string `json:"ids" require:"true" path:"ids" doc:"输入项组件ID集合，英文逗号分隔"`
}

type DeleteFillControlResponse struct {
	Code int    `json:"code" required:"true" doc:"业务码，0表示成功"`
	Msg  string `json:"message" doc:"信息"`
}

func DeleteFillControl(client *resty.Client) func(ctx huma.Context, req DeleteFillControlRequest) {
	return func(ctx huma.Context, req DeleteFillControlRequest) {
		parsedResp := DeleteFillControlResponse{}
		oauth, err := token.GetOauthInfo(client)
		if err != nil {
			ctx.WriteError(http.StatusBadRequest, "got an error when try to get authentication info", err)
			return
		}

		restyResp, err := client.R().SetHeaders(map[string]string{
			"X-Tsign-Open-App-Id": oauth.AppId,
			"X-Tsign-Open-Token":  oauth.Token,
			"Content-Type":        oauth.ContentType,
		}).SetBody(&req).
			SetResult(&parsedResp).Delete("/v1/docTemplates/" + req.TemplateId + "/components/" + req.IDs)

		if common.WriteError(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
			return
		}

		ctx.WriteHeader(http.StatusOK)
	}
}

func RunDeleteFillControl(r *huma.Resource, client *resty.Client) {
	r.Post("DeleteFillControl", "删除填写控件",
		responses.BadRequest(),
		responses.OK().Model(DeleteFillControlResponse{}),
	).Run(DeleteFillControl(client))
}
