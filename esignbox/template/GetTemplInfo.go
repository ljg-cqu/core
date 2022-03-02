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
	EsignSandBoxGetTemplInfoPath = "/v1/flow-templates/basic-info?pageNum=XX&pageSize=XX"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fuseb10&namespace=opendoc%2Fsaas_api

type GetTempInfoRequest struct {
	PageNum  int `json:"pageNum" query:"pageNum" required:"false" doc:"查询页码，不传默认第1页，起始值：1"`
	PageSize int `json:"pageSize" query:"pageSize" required:"false" doc:"分页大小，不传默认10条，最大值：100"`
}

type GetTempInfoResponse struct {
	Code int                     `json:"code" required:"true" doc:"业务码，0表示成功"`
	Msg  string                  `json:"message,omitempty" doc:"信息"`
	Data GetTempInfoResponseData `json:"data,omitempty" type:"object" doc:"业务信息"`
}

type GetTempInfoResponseData struct {
	FlowTemplateBasicInfos FlowTemplateBasicInfos `json:"flowTemplateBasicInfos" type:"array" required:"false" doc:"流程模板基本信息列表"`
	Total                  int                    `json:"total" required:"false" doc:"数据总条数"`
}

type FlowTemplateBasicInfos struct {
	FlowTemplateId   string       `json:"flowTemplateId" required:"false" doc:"流程模板编号"`
	FlowTemplateName string       `json:"flowTemplateName" required:"false" doc:"流程模板名称"`
	DocTemplates     DocTemplates `json:"docTemplates" type:"array" required:"false" doc:"文件模板信息"`
}

type DocTemplates struct {
	DocTemplateId   string `json:"docTemplateId" required:"false" doc:"文件模板id注意：该参数为接口中使用的模板id"`
	DocTemplateName string `json:"docTemplateName" required:"false" doc:"文件模板名称"`
}

// TODO: deal with Not Found reality.

func GetTemplInfo(client *resty.Client) func(ctx huma.Context, req GetTempInfoRequest) {
	return func(ctx huma.Context, req GetTempInfoRequest) {
		parsedResp := GetTempInfoResponse{}

		oauth, err := token.GetOauthInfo(client)
		if err != nil {
			ctx.WriteError(http.StatusBadRequest, "got an error when try to get authentication info", err)
			return
		}

		restyResp, err := client.R().SetHeaders(map[string]string{
			"X-Tsign-Open-App-Id": oauth.AppId,
			"X-Tsign-Open-Token":  oauth.Token,
			"Content-Type":        oauth.ContentType,
		}).SetResult(&parsedResp).Get(EsignSandBoxGetTemplInfoPath)

		if common.WriteError(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
			return
		}

		ctx.WriteModel(http.StatusOK, parsedResp.Data)
	}
}

func RunGetTemplInfo(r *huma.Resource, client *resty.Client) {
	r.Get("GetTemplInfo", "查询e签宝官网模板信息。通过此接口查询调用的appid所对应e签宝官网企业主体下的模板基础信息。",
		//responses.NotFound(),
		responses.BadRequest(),
		responses.OK().Model(GetTempInfoResponseData{}),
	).Run(GetTemplInfo(client))
}
