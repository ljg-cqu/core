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
	EsignSandBoxQueryTemplUploadStatus = "/v1/docTemplates/{templateId}/getBaseInfo"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fmmrbd6&namespace=opendoc%2Fsaas_api

type QueryTemplUploadStatusRequest struct {
	TemplateId string `uri:"templateId" binding:"required" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板ID."`
}

type QueryTemplUploadStatusResponse struct {
	Code int                                `json:"errCode" binding:"required" description:"业务码，0表示成功"`
	Msg  string                             `json:"message" description:"信息"`
	Data QueryTemplUploadStatusResponseData `json:"data" description:"业务信息"`
}

type QueryTemplUploadStatusResponseData struct {
	TemplateId         string `json:"templateId" binding:"required" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板ID"`
	TemplateName       string `json:"templateName" binding:"required" default:"商贷通用收入证明.pdf" description:"模板名称"`
	TemplateFileStatus int    `json:"templateFileStatus" validate:"oneof=0,1,2,3" binding:"required" enum:"0,1,2,3" description:"模板文件上传状态。0-未上传;1-未转换成PDF; 2-已上传成功; 3-已转换成PDF"` // todo: check enum
	FileSize           int64  `json:"fileSize" format:"int64" binding:"required" description:"模板文件大小,单位byte"`                                                                   // todo: check format
	CreateTime         int64  `json:"createTime" binding:"required" description:"创建时间，Unix时间戳（毫秒级）"`
	UpdatedTime        int64  `json:"updatedTime" binding:"required" description:"更新时间，Unix时间戳（毫秒级）"`

	//UploadUrl string `json:"uploadUrl" doc:"文件上传地址，链接有效期60分钟。" example:"http://esignoss.esign.cn/1111564182/8759b09d-4b3e-426f-8ae1-03b1fca3daf6/%E5%95%86%E8%B4%B7%E9%80%9A%E7%94%A8%E6%94%B6%E5%85%A5%E8%AF%81%E6%98%8E.pdf?Expires=1646216465&OSSAccessKeyId=STS.NTfJkmomE8v8hAUvB1NZPHVsJ&Signature=GKEO9kXnX0LagjVE%2BM6vYGT2%2BeA%3D&callback-var=eyJ4OmZpbGVfa2V5IjoiJDk2YzNkYzExLTYxMGYtNDY0Ny05Zjg2LTE5YWNiNmEwYzJmMiQ5ODk3MTE1MzUifQ%3D%3D%0A&callback=eyJjYWxsYmFja1VybCI6Imh0dHA6Ly9maWxlc3lzdGVtLXNtbC50c2lnbi5jbi9maWxlLXN5c3RlbS9jYWxsYmFjay9hbGlvc3MiLCJjYWxsYmFja0JvZHkiOiAie1wibWltZVR5cGVcIjoke21pbWVUeXBlfSxcInNpemVcIjogJHtzaXplfSxcImJ1Y2tldFwiOiAke2J1Y2tldH0sXCJvYmplY3RcIjogJHtvYmplY3R9LFwiZXRhZ1wiOiAke2V0YWd9LFwiZmlsZV9rZXlcIjoke3g6ZmlsZV9rZXl9fSIsImNhbGxiYWNrQm9keVR5cGUiOiAiYXBwbGljYXRpb24vanNvbiJ9%0A&security-token=CAIS%2BAF1q6Ft5B2yfSjIr5fTAdHZgrJkj7TTamfkkkJkQtV8p5PYqDz2IHtKdXRvBu8Xs%2F4wnmxX7f4YlqB6T55OSAmcNZEob27lS5%2FmMeT7oMWQweEurv%2FMQBqyaXPS2MvVfJ%2BOLrf0ceusbFbpjzJ6xaCAGxypQ12iN%2B%2Fm6%2FNgdc9FHHPPD1x8CcxROxFppeIDKHLVLozNCBPxhXfKB0ca0WgVy0EHsPnvm5DNs0uH1AKjkbRM9r6ceMb0M5NeW75kSMqw0eBMca7M7TVd8RAi9t0t1%2FIVpGiY4YDAWQYLv0rda7DOltFiMkpla7MmXqlft%2BhzcgeQY0pc%2FRqAAUwxVoW0nfk4w92%2Bm0ikUU43mcQR3XgrTF2geBGPg5GvqQVuUaZvPdzQIODgCG2tcEvUuMmKU%2FqbGzuj7bWEOrdrIfV7UXwSAa9NdJ5okZmmQvoKAmnCWiocgK1rG%2B9ItCGc5f5JAiYH%2F0UrxI8HKE6FUFT5oWLvgngYiKEm%2FDAk"`
}

// TODO: save file to db

func (req *QueryTemplUploadStatusRequest) Handler(ctx *gin.Context) {
	parsedResp := QueryTemplUploadStatusResponse{}
	oauth, err := token.GetOauthInfo()
	if err != nil {
		common.WriteError(ctx, 400, fmt.Sprintf("got an error when try to get authentication info:%w", err))
		return
	}

	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
		"Content-Type":        oauth.ContentType,
	}).SetResult(&parsedResp).Get("/v1/docTemplates/" + req.TemplateId + "/getBaseInfo")

	if common.WriteErrorE(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
		return
	}

	common.WriteOK(ctx, parsedResp.Data)
}

var QueryTemplUploadStatusRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fmmrbd6&namespace=opendoc%2Fsaas_api", "查询模板文件上传状态.")
	}

	r := router.New(
		&QueryTemplUploadStatusRequest{},
		router.Summary("查询模板文件上传状态."),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: QueryTemplUploadStatusResponseData{},
			},
			"400": router.ResponseItem{
				Model: common.ErrorResp{},
			},
		}),
	)

	return r
}
