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
	EsignSandBoxQueryPdfFileDetailsUrl = "/v1/files/{fileId}"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyingmd&namespace=opendoc%2Fsaas_api

type QueryPdfFileDetailsRequest struct {
	FileId string `json:"fileId" required:"true" example:"ede1fa4504954c29ad210637c15f42cf" path:"fileId" doc:"文件ID"`
}

type QueryPdfFileDetailsResponse struct {
	Code int                             `json:"code" required:"true" doc:"业务码，0表示成功"`
	Msg  string                          `json:"message" doc:"信息"`
	Data QueryPdfFileDetailsResponseData `json:"data" type:"object" doc:"业务信息"`
}

type QueryPdfFileDetailsResponseData struct {
	FileId   string `json:"fileId" example:"ede1fa4504954c29ad210637c15f42cf" doc:"文件ID"`
	FileName string `json:"name" example:"商贷通用收入证明.pdf" doc:"文件名称"`
	Size     string `json:"size" doce:"文件大小，单位byte"`
	Status   int    `json:"status" doc:"文件上传状态: 0-文件未上传；1-文件上传中 ；2-文件上传已完成,；3-文件上传失败 ；4-文件等待转pdf ；5-文件已转换pdf ；6-加水印中；7-加水印完毕；8-文件转换中；9-文件转换失败"`

	DownloadUrl   string `json:"downloadUrl" doc:"模板文件下载链接，有效期60分钟" example:"https://esignoss.esign.cn/1111564182/f33c1c35-e47e-451d-b9b9-0855d682ea79/%E5%95%86%E8%B4%B7%E9%80%9A%E7%94%A8%E6%94%B6%E5%85%A5%E8%AF%81%E6%98%8E.pdf?Expires=1646233196&OSSAccessKeyId=LTAI4G23YViiKnxTC28ygQzF&Signature=fpguS0x08KbyJ3xOfPaujARK010%3D"`
	PdfTotalPages int    `json:"pdfTotalPages" doc:"pdf文件总页数,仅当文件类型为pdf时有值"`
}

func QueryPdfFileDetails(client *resty.Client) func(ctx huma.Context, req QueryPdfFileDetailsRequest) {
	return func(ctx huma.Context, req QueryPdfFileDetailsRequest) {
		parsedResp := QueryPdfFileDetailsResponse{}
		oauth, err := token.GetOauthInfo(client)
		if err != nil {
			ctx.WriteError(http.StatusBadRequest, "got an error when try to get authentication info", err)
			return
		}

		restyResp, err := client.R().SetHeaders(map[string]string{
			"X-Tsign-Open-App-Id": oauth.AppId,
			"X-Tsign-Open-Token":  oauth.Token,
			"Content-Type":        oauth.ContentType,
		}).SetResult(&parsedResp).Get("/v1/files/" + req.FileId)

		if common.WriteError(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
			return
		}

		ctx.WriteModel(http.StatusOK, parsedResp.Data)
	}
}

func RunQueryPdfFileDetails(r *huma.Resource, client *resty.Client) {
	r.Post("QueryPdfFileDetails", "查询PDF文件详情",
		responses.BadRequest(),
		responses.OK().Model(QueryPdfFileDetailsResponseData{}),
	).Run(QueryPdfFileDetails(client))
}
