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
	EsignSandBoxAddFillControl = "/v1/docTemplates/{templateId}/components"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fzv9rch&namespace=opendoc%2Fsaas_api

type AddFillControlRequest struct {
	TemplateId string `json:"templateId" required:"true" example:"c4d4fe1b48184ba28982f68bf2c7bf25" path:"templateId" doc:"模板id"`

	Body []StructComponent `json:"structComponents" required:"true" type:"array" doc:"添加填写控件请求参数"`
}

type StructComponent struct {
	ID      string  `json:"id" example:"cp-459470" doc:"输入项组件id，使用时可用id填充，为空时表示添加，不为空时表示修改"`
	Key     string  `json:"key" example:"cp-459470" doc:"模板下输入项组件唯一标识，使用模板时也可用根据key值填充"`
	Type    int     `json:"type" example:"1" default:"1" required:"true" enum:"1,2,3,8,11" doc:"输入项组件类型，1-单行文本，2-数字，3-日期，8-多行文本，11-图片，不支持修改"`
	Context Context `json:"context" required:"true" type:"object" doc:"输入项组件上下文信息，包含了名称，填充格式，样式以及坐标"`
}

type Context struct {
	Label    string `json:"label" example:"甲方代表" required:"true" doc:"输入项组件显示名称"`
	Required bool   `json:"required" default:"true" doc:"是否必填，默认true"`
	Limit    string `json:"limit" example:"yyy-MM-dd" example:"#00" doc:"输入项组件type=2,type=3时填充格式校验规则;数字格式如：#,#00.0# 日期格式如： yyyy-MM-dd"`
	Style    Style  `json:"style" required:"true" doc:"输入项组件样式"`
	Pos      Pos    `json:"pos" required:"true" doc:"输入项组件坐标"`
}

type Style struct {
	Width     float32 `json:"width" format:"float" required:"true"  example:"100.11" default:"100" doc:"输入项组件宽度"`
	Height    float32 `json:"height" format:"float" required:"true" example:"50.05" default:"50" doc:"输入项组件高"`
	Font      int     `json:"font" default:"1" enum:"1,2,4,5" doc:"填充字体,默认1，1-宋体，2-新宋体，4-黑体，5-楷体"`
	FontSize  float32 `json:"fontSize" format:"float" example:"12" default:"12" doc:"填充字体大小,默认12"`
	TextColor string  `json:"textColor" example:"#000000" default:"#000000" doc:"字体颜色，默认#000000黑色"`
}

type Pos struct {
	Page int     `json:"page" required:"true" example:"1" default:"1" doc:"页码"`
	X    float32 `json:"x" required:"true" format:"float" example:"20.2" default:"20" doc:"x轴坐标，左下角为原点"`
	Y    float32 `json:"y" required:"true" format:"float" example:"100.1" default:"100" doc:"y轴坐标，左下角为原点"`
}

type AddFillControlResponse struct {
	Code int      `json:"code" required:"true" doc:"业务码，0表示成功"`
	Msg  string   `json:"message" doc:"信息"`
	Data []string `json:"data" doc:"添加/编辑的输入项组件id列表"`
}

type AddFillControlResponseData struct {
	ComponentIdList []string `json:"componentIdList" example:"a73bd123857445df97b3d45ae33891ed" doc:"添加/编辑的输入项组件id列表"`
}

func AddFillControl(client *resty.Client) func(ctx huma.Context, req AddFillControlRequest) {
	return func(ctx huma.Context, req AddFillControlRequest) {
		parsedResp := AddFillControlResponse{}
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
			SetResult(&parsedResp).Post("/v1/docTemplates/" + req.TemplateId + "/components")

		if common.WriteError(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
			return
		}

		ctx.WriteModel(http.StatusOK, AddFillControlResponseData{parsedResp.Data})
	}
}

func RunAddFillControl(r *huma.Resource, client *resty.Client) {
	r.Post("AddFillControl", "添加或修改填写控件",
		responses.BadRequest(),
		responses.OK().Model(AddFillControlResponseData{}),
	).Run(AddFillControl(client))
}
