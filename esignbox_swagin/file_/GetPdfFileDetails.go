package file_

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/token"
	"github.com/long2ice/swagin/router"
	"github.com/wI2L/fizz/markdown"
	"net/http"
)

const (
	EsignSandBoxQueryPdfFileDetailsUrl = "/v1/files/{fileId}"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyingmd&namespace=opendoc%2Fsaas_api

type GetPdfFileDetailsRequest struct {
	FileId string `uri:"fileId" binding:"required" default:"ede1fa4504954c29ad210637c15f42cf" description:"文件ID"`
}

type GetPdfFileDetailsResponse struct {
	Code int                           `json:"code" binding:"required:" description:"业务码，0表示成功"`
	Msg  string                        `json:"message" description:"信息"`
	Data GetPdfFileDetailsResponseData `json:"data" description:"业务信息"`
}

type GetPdfFileDetailsResponseData struct {
	FileId   string `json:"fileId" default:"ede1fa4504954c29ad210637c15f42cf" description:"文件ID"`
	FileName string `json:"name" default:"商贷通用收入证明.pdf" description:"文件名称"`
	Size     string `json:"size" description:"文件大小，单位byte"`
	Status   int    `json:"status" description:"文件上传状态: 0-文件未上传；1-文件上传中 ；2-文件上传已完成,；3-文件上传失败 ；4-文件等待转pdf ；5-文件已转换pdf ；6-加水印中；7-加水印完毕；8-文件转换中；9-文件转换失败"`

	DownloadUrl   string `json:"downloadUrl" description:"PDF文件下载链接，有效期60分钟"`
	PdfTotalPages int    `json:"pdfTotalPages" description:"pdf文件总页数,仅当文件类型为pdf时有值"`
}

func (req *GetPdfFileDetailsRequest) Handler(ctx *gin.Context) {
	//parsedResp := GetPdfFileDetailsResponse{}
	//oauth, err := token.GetOauthInfo()
	//if err != nil {
	//	common.RespErrf(ctx, 400, fmt.Sprintf("got an error when try to get authentication info:%w", err))
	//	return
	//}
	//
	//restyResp, err := common.Client.R().SetHeaders(map[string]string{
	//	"X-Tsign-Open-App-Id": oauth.AppId,
	//	"X-Tsign-Open-Token":  oauth.Token,
	//	"Content-Type":        oauth.ContentType,
	//}).SetResult(&parsedResp).Get("/v1/files/" + req.FileId)
	//
	//if common.RespErre(ctx, restyResp.RawResponse, err, &common.EsignError{Code: parsedResp.Code, Msg: parsedResp.Msg}) {
	//	return
	//}
	//common.RespSucc(ctx, parsedResp.Data)

	detail, errObj := GetPdfFileDetails(req.FileId)
	if errObj != nil {
		common.RespErrObj(ctx, errObj)
		return
	}

	common.RespSucc(ctx, detail)
}

func GetPdfFileDetails(fileId string) (data *GetPdfFileDetailsResponseData, errObj *common.RespObj) {
	oauth, err := token.GetOauthInfo()
	if err != nil {
		return nil, &common.RespObj{
			Code: http.StatusNetworkAuthenticationRequired,
			Msg:  fmt.Sprintf("got an error when try to get authentication info:%v", err),
		}
	}

	parsedResp := GetPdfFileDetailsResponse{}
	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
		"Content-Type":        oauth.ContentType,
	}).SetResult(&parsedResp).Get("/v1/files/" + fileId)

	errObj = common.ErrObjFromEsignRequest(restyResp.RawResponse, err, &common.EsignError{Code: parsedResp.Code, Msg: parsedResp.Msg})
	if errObj != nil {
		return
	}

	return &parsedResp.Data, nil
}

var GetPdfFileDetailsRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyingmd&namespace=opendoc%2Fsaas_api", "查询PDF文件详情.")
	}

	r := router.New(
		&GetPdfFileDetailsRequest{},
		router.Summary("获取PDF文件详情."),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: struct {
					Code int                           `json:"code" binding:"required" default:"0"`
					Msg  string                        `json:"msg" binding:"required" default:"ok"`
					Data GetPdfFileDetailsResponseData `json:"data"`
				}{},
			},
		}),
	)

	return r
}
