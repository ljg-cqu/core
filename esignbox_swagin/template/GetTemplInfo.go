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
	EsignSandBoxGetTemplInfoPath = "/v1/flow-templates/basic-info?pageNum=XX&pageSize=XX"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fuseb10&namespace=opendoc%2Fsaas_api

type GetTemplInfoRequest struct {
	PageNum  int `query:"pageNum" description:"查询页码，不传默认第1页，起始值：1"`
	PageSize int `query:"pageSize" description:"分页大小，不传默认10条，最大值：100"`
}

type GetTempInfoResponse struct {
	Code int                      `json:"code" required:"true" description:"业务码，0表示成功"`
	Msg  string                   `json:"message,omitempty" description:"信息"`
	Data GetTemplInfoResponseData `json:"data,omitempty" type:"object" description:"业务信息"`
}

type GetTemplInfoResponseData struct {
	FlowTemplateBasicInfos FlowTemplateBasicInfos `json:"flowTemplateBasicInfos" type:"array" required:"false" description:"流程模板基本信息列表"`
	Total                  int                    `json:"total" required:"false" description:"数据总条数"`
}

type FlowTemplateBasicInfos struct {
	FlowTemplateId   string       `json:"flowTemplateId" required:"false" description:"流程模板编号"`
	FlowTemplateName string       `json:"flowTemplateName" required:"false" description:"流程模板名称"`
	DocTemplates     DocTemplates `json:"docTemplates" type:"array" required:"false" description:"文件模板信息"`
}

type DocTemplates struct {
	DocTemplateId   string `json:"docTemplateId" required:"false" description:"文件模板id注意：该参数为接口中使用的模板id"`
	DocTemplateName string `json:"docTemplateName" required:"false" description:"文件模板名称"`
}

// TODO: deal with Not Found reality.

func (req *GetTemplInfoRequest) Handler(ctx *gin.Context) {
	parsedResp := GetTempInfoResponse{}

	oauth, err := token.GetOauthInfo()
	if err != nil {
		common.WriteErrorE(ctx, fmt.Errorf("got an error when try to get authentication info:%w", err), 0, "", 0, "")
		return
	}

	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
		"Content-Type":        oauth.ContentType,
	}).SetResult(&parsedResp).Get(EsignSandBoxGetTemplInfoPath)

	if common.WriteErrorE(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
		return
	}

	common.WriteOK(ctx, parsedResp.Data)
}

var GetTemplInfoRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("通过此接口查询调用的appid所对应e签宝官网企业主体下的模板基础信息。官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fuseb10&namespace=opendoc%2Fsaas_api", "查询e签宝官网模板信息")
	}

	r := router.New(
		&GetTemplInfoRequest{},
		router.Summary("查询e签宝官网模板信息。注意：此接口仅能查询官网企业主体下的模板信息。"),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: GetTemplInfoResponseData{},
			},
			"400": router.ResponseItem{
				Model: common.ErrorResp{},
			},
		}),
	)

	return r
}
