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
	EsignSandBoxAddFillControl = "/v1/docTemplates/{templateId}/components"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fzv9rch&namespace=opendoc%2Fsaas_api

type AddFillControlRequest struct {
	TemplateId string `uri:"templateId" binding:"required" json:"templateId" description:"模板id" default:"c4d4fe1b48184ba28982f68bf2c7bf25"`

	Body []StructComponent `form:"structComponents" binding:"required" json:"structComponents" description:"添加填写控件请求参数列表"`
}

type StructComponent struct {
	ID      string  `form:"id" json:"id" default:"cp-459470" description:"输入项组件id，使用时可用id填充，为空时表示添加，不为空时表示修改"`
	Key     string  `form:"key" json:"key" default:"cp-459470" description:"模板下输入项组件唯一标识，使用模板时也可用根据key值填充"`
	Type    int     `form:"type" binding:"required, oneof=1,2,3,8,11" json:"type" default:"1" example:"1" enum:"1,2,3,8,11" description:"输入项组件类型，1-单行文本，2-数字，3-日期，8-多行文本，11-图片，不支持修改"`
	Context Context `form:"context" json:"context" binding:"required"  type:"object" description:"输入项组件上下文信息，包含了名称，填充格式，样式以及坐标"`
}

type Context struct {
	Label    string `form:"label" json:"label" default:"甲方代表" binding:"required" description:"输入项组件显示名称"`
	Required bool   `form:"required" json:"required" default:"true" description:"是否必填，默认true"`
	Limit    string `form:"limit" json:"limit" example:"yyy-MM-dd" example:"#00" description:"输入项组件type=2,type=3时填充格式校验规则;数字格式如：#,#00.0# 日期格式如： yyyy-MM-dd"`
	Style    Style  `form:"style" json:"style" binding:"required" description:"输入项组件样式"`
	Pos      Pos    `form:"pos" json:"pos" binding:"required"  description:"输入项组件坐标"`
}

type Style struct {
	Width     float32 `form:"width" json:"width" format:"float" binding:"required" default:"100.11" description:"输入项组件宽度"`
	Height    float32 `form:"height" json:"height" format:"float" binding:"required" example:"50.05" default:"50.05" description:"输入项组件高"`
	Font      int     `form:"font" json:"font" default:"1" enum:"1,2,4,5" description:"填充字体,默认1，1-宋体，2-新宋体，4-黑体，5-楷体"`
	FontSize  float32 `form:"fontSize" json:"fontSize" format:"float" example:"12" default:"12" description:"填充字体大小,默认12"`
	TextColor string  `form:"textColor" json:"textColor" example:"#000000" default:"#000000" description:"字体颜色，默认#000000黑色"`
}

type Pos struct {
	Page int     `form:"page" json:"page" binding:"required" required:"true" example:"1" default:"1" description:"页码"`
	X    float32 `form:"x" json:"x" binding:"required" required:"true" format:"float" example:"20.2" default:"20" description:"x轴坐标，左下角为原点"`
	Y    float32 `form:"y" json:"y" binding:"required" required:"true" format:"float" example:"100.1" default:"100" description:"y轴坐标，左下角为原点"`
}

type AddFillControlResponse struct {
	Code int      `json:"code" binding:"required" required:"true" description:"业务码，0表示成功"`
	Msg  string   `json:"message" description:"信息"`
	Data []string `json:"data" description:"添加/编辑的输入项组件id列表"`
}

type AddFillControlResponseData struct {
	ComponentIdList []string `json:"componentIdList" example:"a73bd123857445df97b3d45ae33891ed" description:"添加/编辑的输入项组件id列表"`
}

func (req *AddFillControlRequest) Handler(ctx *gin.Context) {
	parsedResp := AddFillControlResponse{}
	oauth, err := token.GetOauthInfo()
	if err != nil {
		common.WriteErrore(ctx, fmt.Errorf("got an error when try to get authentication info:%w", err), 0, "", 0, "")
		return
	}

	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
		"Content-Type":        oauth.ContentType,
	}).SetBody(&req).
		SetResult(&parsedResp).Post("/v1/docTemplates/" + req.TemplateId + "/components")

	if err != nil {
		common.WriteErrorf(ctx, 400, "got an error when try to delete fill control, error:%v", err)
		return
	}

	if parsedResp.Code != 0 {
		common.WriteErrorf(ctx, 400, "got an error when try to delete fill control, code:%v, message:%v", parsedResp.Code, parsedResp.Msg)
		return
	}

	_ = restyResp

	common.WriteOK(ctx, AddFillControlResponseData{parsedResp.Data})
}

var AddFillControlRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("官方文档：")
		return builder.String() + builder.Link(" https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fzv9rch&namespace=opendoc%2Fsaas_api", "添加或修改填写控件")
	}

	r := router.New(
		&AddFillControlRequest{},
		router.Summary("添加或修改填写控件。"),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: AddFillControlResponseData{},
			},
			"400": router.ResponseItem{
				Model: common.ErrorResp{},
			},
		}),
	)
	return r
}
