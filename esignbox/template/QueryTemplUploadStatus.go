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
	EsignSandBoxQueryTemplUploadStatus = "/v1/docTemplates/{templateId}/getBaseInfo"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fmmrbd6&namespace=opendoc%2Fsaas_api

type QueryTemplUploadStatusRequest struct {
	TemplateId string `json:"templateId" required:"true" path:"templateId" example:"c4d4fe1b48184ba28982f68bf2c7bf25" doc:"模板ID"`
}

type QueryTemplUploadStatusResponse struct {
	Code int                                `json:"errCode" required:"true" doc:"业务码，0表示成功"`
	Msg  string                             `json:"message" doc:"信息"`
	Data QueryTemplUploadStatusResponseData `json:"data" doc:"业务信息"`
}

type QueryTemplUploadStatusResponseData struct {
	TemplateId         string `json:"templateId" required:"true" example:"c4d4fe1b48184ba28982f68bf2c7bf25" doc:"模板ID"`
	TemplateName       string `json:"templateName" required:"true" example:"商贷通用收入证明.pdf" doc:"模板名称"`
	TemplateFileStatus int    `json:"templateFileStatus" enum:"0,1,2,3" required:"true" doc:"模板文件上传状态。0-未上传;1-未转换成PDF; 2-已上传成功; 3-已转换成PDF"`
	FileSize           int64  `json:"fileSize" required:"true" format:"int64" doc:"模板文件大小"`
	CreateTime         int64  `json:"createTime" required:"true" doc:"创建时间，Unix时间戳（毫秒级）"`
	UpdatedTime        int64  `json:"updatedTime" required:"true" doc:"更新时间，Unix时间戳（毫秒级）"`

	//UploadUrl string `json:"uploadUrl" doc:"文件上传地址，链接有效期60分钟。" example:"http://esignoss.esign.cn/1111564182/8759b09d-4b3e-426f-8ae1-03b1fca3daf6/%E5%95%86%E8%B4%B7%E9%80%9A%E7%94%A8%E6%94%B6%E5%85%A5%E8%AF%81%E6%98%8E.pdf?Expires=1646216465&OSSAccessKeyId=STS.NTfJkmomE8v8hAUvB1NZPHVsJ&Signature=GKEO9kXnX0LagjVE%2BM6vYGT2%2BeA%3D&callback-var=eyJ4OmZpbGVfa2V5IjoiJDk2YzNkYzExLTYxMGYtNDY0Ny05Zjg2LTE5YWNiNmEwYzJmMiQ5ODk3MTE1MzUifQ%3D%3D%0A&callback=eyJjYWxsYmFja1VybCI6Imh0dHA6Ly9maWxlc3lzdGVtLXNtbC50c2lnbi5jbi9maWxlLXN5c3RlbS9jYWxsYmFjay9hbGlvc3MiLCJjYWxsYmFja0JvZHkiOiAie1wibWltZVR5cGVcIjoke21pbWVUeXBlfSxcInNpemVcIjogJHtzaXplfSxcImJ1Y2tldFwiOiAke2J1Y2tldH0sXCJvYmplY3RcIjogJHtvYmplY3R9LFwiZXRhZ1wiOiAke2V0YWd9LFwiZmlsZV9rZXlcIjoke3g6ZmlsZV9rZXl9fSIsImNhbGxiYWNrQm9keVR5cGUiOiAiYXBwbGljYXRpb24vanNvbiJ9%0A&security-token=CAIS%2BAF1q6Ft5B2yfSjIr5fTAdHZgrJkj7TTamfkkkJkQtV8p5PYqDz2IHtKdXRvBu8Xs%2F4wnmxX7f4YlqB6T55OSAmcNZEob27lS5%2FmMeT7oMWQweEurv%2FMQBqyaXPS2MvVfJ%2BOLrf0ceusbFbpjzJ6xaCAGxypQ12iN%2B%2Fm6%2FNgdc9FHHPPD1x8CcxROxFppeIDKHLVLozNCBPxhXfKB0ca0WgVy0EHsPnvm5DNs0uH1AKjkbRM9r6ceMb0M5NeW75kSMqw0eBMca7M7TVd8RAi9t0t1%2FIVpGiY4YDAWQYLv0rda7DOltFiMkpla7MmXqlft%2BhzcgeQY0pc%2FRqAAUwxVoW0nfk4w92%2Bm0ikUU43mcQR3XgrTF2geBGPg5GvqQVuUaZvPdzQIODgCG2tcEvUuMmKU%2FqbGzuj7bWEOrdrIfV7UXwSAa9NdJ5okZmmQvoKAmnCWiocgK1rG%2B9ItCGc5f5JAiYH%2F0UrxI8HKE6FUFT5oWLvgngYiKEm%2FDAk"`
}

// TODO: save file to db

func QueryTemplUploadStatus(client *resty.Client) func(ctx huma.Context, req QueryTemplUploadStatusRequest) {
	return func(ctx huma.Context, req QueryTemplUploadStatusRequest) {
		parsedResp := QueryTemplUploadStatusResponse{}
		oauth, err := token.GetOauthInfo(client)
		if err != nil {
			ctx.WriteError(http.StatusBadRequest, "got an error when try to get authentication info", err)
			return
		}

		restyResp, err := client.R().SetHeaders(map[string]string{
			"X-Tsign-Open-App-Id": oauth.AppId,
			"X-Tsign-Open-Token":  oauth.Token,
			"Content-Type":        oauth.ContentType,
		}).SetResult(&parsedResp).Get("/v1/docTemplates/" + req.TemplateId + "/getBaseInfo")

		if common.WriteError(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
			return
		}

		ctx.WriteModel(http.StatusOK, parsedResp.Data)
	}
}

func RunQueryTemplUploadStatus(r *huma.Resource, client *resty.Client) {
	r.Get("QueryTemplUploadStatus", "查询模板文件上传状态",
		responses.BadRequest(),
		responses.OK().Model(QueryTemplUploadStatusResponseData{}),
	).Run(QueryTemplUploadStatus(client))
}
