package template_

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/token"
	"github.com/long2ice/swagin/router"
	"github.com/wI2L/fizz/markdown"
	"net/http"
)

const (
	EsignSandBoxQueryTemplDetailsUrl = "/v1/docTemplates/{templateId}"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fviygk4&namespace=opendoc%2Fsaas_api

type GetTemplDetailsRequest struct {
	TemplateId string `uri:"templateId" binding:"required" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板ID, 该参数需放在请求地址里面"`
}

type GetTemplDetailsResponse struct {
	Code int                         `json:"code" binding:"required"  description:"业务码，0表示成功"`
	Msg  string                      `json:"message" description:"信息"`
	Data GetTemplDetailsResponseData `json:"data" description:"业务信息"`
}

type GetTemplDetailsResponseData struct {
	TemplateId       string             `json:"templateId" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板ID"`
	TemplateName     string             `json:"templateName" binding:"required" default:"商贷通用收入证明.pdf" description:"模板名称"`
	TemplateType     int                `json:"templateType" enum:"3" description:"固定值 3"` // todo:check enum
	DownloadUrl      string             `json:"DownloadUrl" description:"模板文件下载链接，有效期60分钟。"`
	FileSize         int64              `json:"fileSize" description:"模板文件大小"`
	CreateTime       int64              `json:"createTime" description:"创建时间，Unix时间戳（毫秒级）"`
	UpdatedTime      int64              `json:"updatedTime" binding:"required" description:"更新时间，Unix时间戳（毫秒级）"`
	StructComponents []_StructComponent `json:"structComponents" type:"array" description:"文件模板中的填写控件列表"` // todo: check type
	CreatorID        string             `json:"creatorID"`
	CreatorName      string             `json:"creatorName"`
}

type _StructComponent struct {
	ID        string   `json:"id" default:"cp-459470" description:"输入项组件id，使用时可用id填充，为空时表示添加，不为空时表示修改"`
	Key       string   `json:"key" default:"cp-459470" description:"模板下输入项组件唯一标识，使用模板时也可用根据key值填充"`
	Type      int      `json:"type" default:"1" binding:"required" enum:"1,2,3,8,11" description:"输入项组件类型，1-单行文本，2-数字，3-日期，8-多行文本，11-图片，不支持修改"` // todo: check enum
	Context   _Context `json:"context" binding:"required" type:"object" description:"输入项组件上下文信息，包含了名称，填充格式，样式以及坐标"`                             // todo:chek type
	RefId     string   `json:"-" description:"关联签署方的ID，开发者可忽略此字段"`
	AllowEdit bool     `json:"allowEdit" description:"是否允许编辑，用于管控页面是否能对该控件进行修改；线下制作的模板（带表单域）此值为false,其余为true"`
}

type _Context struct {
	Context
	Options string `json:"options" description:"控件结构中的子元素，用于选择控件，如：单选，多选，下拉选项"`
	Version int    `json:"-" description:"控件版本，开发者可忽略此字段"` // todo: check enum
	Ext     string `json:"ext" enum:"imgType,page,signRequirements,qiFeng,units,signDatePos,fillLengthLimit" description:"扩展字段，用于支持一些扩展功能支持以下字段：imgType：图片类型；page：页码，用于签署区，单页签署区为独立数字；骑缝章时用于标明骑缝章跨度，如：1-5；signRequirements：签署要求，逗号分隔 1-企业章 2-经办人 3-法定代表人章；qiFeng：是否骑缝章签署，true - 骑缝章，false - 非骑缝章；units：暂无用途；signDatePos：签署区日期，当签署区设置了签署时间时，用于时间的坐标位置，如{\"x\":30.8,\"y\":724.01,\"page\":1}；fillLengthLimit：控件填充的限制长度（限制填充的字数），用于单行/多行文本。"`
}

func (req *GetTemplDetailsRequest) Handler(ctx *gin.Context) {
	data, errObj := GetTemplDetails(req.TemplateId)
	if errObj != nil {
		common.RespErrObj(ctx, errObj)
		return
	}

	// todo: return account-id and name ...
	common.RespSucc(ctx, data)
}

var GetTemplDetailsRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fviygk4&namespace=opendoc%2Fsaas_api", "查询模板文件详情.")
	}

	r := router.New(
		&GetTemplDetailsRequest{},
		router.Summary("获取模板文件详情."),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: struct {
					Code int                         `json:"code" binding:"required" default:"0"`
					Msg  string                      `json:"msg" binding:"required" default:"ok"`
					Data GetTemplDetailsResponseData `json:"data"`
				}{},
			},
		}),
	)

	return r
}

func GetTemplDetails(templeId string) (data *GetTemplDetailsResponseData, errObj *common.RespObj) {
	oauth, err := token.GetOauthInfo()
	if err != nil {
		return nil, &common.RespObj{
			Code: http.StatusNetworkAuthenticationRequired,
			Msg:  fmt.Sprintf("got an error when try to get authentication info:%v", err),
		}
	}

	parsedResp := GetTemplDetailsResponse{}
	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
		"Content-Type":        oauth.ContentType,
	}).SetResult(&parsedResp).Get("/v1/docTemplates/" + templeId)

	// TODO: use true creator info
	_data := parsedResp.Data
	_data.CreatorID = gofakeit.UUID()
	_data.CreatorName = gofakeit.Name()

	errObj = common.ErrObjFromEsignRequest(restyResp.RawResponse, err, &common.EsignError{Code: parsedResp.Code, Msg: parsedResp.Msg})
	if errObj != nil {
		return
	}

	return &_data, nil
}
