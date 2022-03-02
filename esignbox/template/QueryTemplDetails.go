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
	EsignSandBoxQueryTemplDetailsUrl = "/v1/docTemplates/{templateId}"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fviygk4&namespace=opendoc%2Fsaas_api

type QueryTemplDetailsRequest struct {
	TemplateId string `json:"templateId" required:"true" path:"templateId" example:"c4d4fe1b48184ba28982f68bf2c7bf25" doc:"模板ID, 该参数需放在请求地址里面"`
}

type QueryTemplDetailsResponse struct {
	Code int                           `json:"code" required:"true" doc:"业务码，0表示成功"`
	Msg  string                        `json:"message" doc:"信息"`
	Data QueryTemplDetailsResponseData `json:"data" doc:"业务信息"`
}

type QueryTemplDetailsResponseData struct {
	TemplateId       string             `json:"templateId" example:"c4d4fe1b48184ba28982f68bf2c7bf25" doc:"模板ID"`
	TemplateName     string             `json:"templateName" required:"true" example:"商贷通用收入证明.pdf" doc:"模板名称"`
	TemplateType     int                `json:"templateType" enum:"3" doc:"固定值 3"`
	DownloadUrl      string             `json:"DownloadUrl" doc:"模板文件下载链接，有效期60分钟。"`
	FileSize         int64              `json:"fileSize" doc:"模板文件大小"`
	CreateTime       int64              `json:"createTime" doc:"创建时间，Unix时间戳（毫秒级）"`
	UpdatedTime      int64              `json:"updatedTime" required:"true" doc:"更新时间，Unix时间戳（毫秒级）"`
	StructComponents []_StructComponent `json:"structComponents" type:"array" dock:"文件模板中的填写控件列表"`
}

type _StructComponent struct {
	ID        string   `json:"id" example:"cp-459470" doc:"输入项组件id，使用时可用id填充，为空时表示添加，不为空时表示修改"`
	Key       string   `json:"key" example:"cp-459470" doc:"模板下输入项组件唯一标识，使用模板时也可用根据key值填充"`
	Type      int      `json:"type" example:"1" default:"1" required:"true" enum:"1,2,3,8,11" doc:"输入项组件类型，1-单行文本，2-数字，3-日期，8-多行文本，11-图片，不支持修改"`
	Context   _Context `json:"context" required:"true" type:"object" doc:"输入项组件上下文信息，包含了名称，填充格式，样式以及坐标"`
	RefId     string   `json:"refId" doc:"关联签署方的ID，开发者可忽略此字段"`
	AllowEdit bool     `json:"allowEdit" doc:"是否允许编辑，用于管控页面是否能对该控件进行修改；线下制作的模板（带表单域）此值为false,其余为true"`
}

type _Context struct {
	Context
	Options string `json:"options" doc:"控件结构中的子元素，用于选择控件，如：单选，多选，下拉选项"`
	Version int    `json:"version" doc:"控件版本，开发者可忽略此字段"`
	Ext     string `json:"ext" enum:"imgType,page,signRequirements,qiFeng,units,signDatePos,fillLengthLimit" doc:"扩展字段，用于支持一些扩展功能支持以下字段：imgType：图片类型；page：页码，用于签署区，单页签署区为独立数字；骑缝章时用于标明骑缝章跨度，如：1-5；signRequirements：签署要求，逗号分隔 1-企业章 2-经办人 3-法定代表人章；qiFeng：是否骑缝章签署，true - 骑缝章，false - 非骑缝章；units：暂无用途；signDatePos：签署区日期，当签署区设置了签署时间时，用于时间的坐标位置，如{\"x\":30.8,\"y\":724.01,\"page\":1}；fillLengthLimit：控件填充的限制长度（限制填充的字数），用于单行/多行文本。"`
}

func QueryTemplDetails(client *resty.Client) func(ctx huma.Context, req QueryTemplDetailsRequest) {
	return func(ctx huma.Context, req QueryTemplDetailsRequest) {
		parsedResp := QueryTemplDetailsResponse{}
		oauth, err := token.GetOauthInfo(client)
		if err != nil {
			ctx.WriteError(http.StatusBadRequest, "got an error when try to get authentication info", err)
			return
		}

		restyResp, err := client.R().SetHeaders(map[string]string{
			"X-Tsign-Open-App-Id": oauth.AppId,
			"X-Tsign-Open-Token":  oauth.Token,
			"Content-Type":        oauth.ContentType,
		}).SetResult(&parsedResp).Get("/v1/docTemplates/" + req.TemplateId)

		if common.WriteError(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
			return
		}

		// TODO: save file to db

		ctx.WriteModel(http.StatusOK, parsedResp.Data)
	}
}

func RunQueryTemplDetails(r *huma.Resource, client *resty.Client) {
	r.Get("get-doc-templ-url", "查询模板文件详情",
		responses.BadRequest(),
		responses.OK().Model(QueryTemplDetailsResponseData{}),
	).Run(QueryTemplDetails(client))
}
