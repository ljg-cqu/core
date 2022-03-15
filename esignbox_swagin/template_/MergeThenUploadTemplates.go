package template_

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/models/models"
	"github.com/ljg-cqu/core/utils"
	"github.com/long2ice/swagin/router"
	"github.com/spf13/cast"
	"net/http"

	"strings"
)

//const (
//	EsignSandBoxQueryPdfFileDetailsUrl = "/v1/files/{fileId}"
//)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyingmd&namespace=opendoc%2Fsaas_api

type MergeThenUploadTemplatesRequest struct {
	TemplateIds string `uri:"templateIds" default:"ede1fa4504954c29ad210637c15f42cf,ede1fa4504954c29ad210637c15f42cf" description:"用于合并的PDF模板ID，以逗号按顺序连续排列"`
}

type MergeThenUploadTemplatesResponse struct {
	Code int                                  `json:"code" binding:"required:" description:"业务码，0表示成功"`
	Msg  string                               `json:"message" description:"信息"`
	Data MergeThenUploadTemplatesResponseData `json:"data" description:"业务信息"`
}

type MergeThenUploadTemplatesResponseData struct {
	TemplateId   string `json:"templateId" default:"ede1fa4504954c29ad210637c15f42cf" description:"文件ID"`
	TemplateName string `json:"name" default:"商贷通用收入证明.pdf" description:"文件名称"`
	Size         string `json:"size" description:"文件大小，单位byte"`
	//Status   int    `json:"status" description:"文件上传状态: 0-文件未上传；1-文件上传中 ；2-文件上传已完成,；3-文件上传失败 ；4-文件等待转pdf ；5-文件已转换pdf ；6-加水印中；7-加水印完毕；8-文件转换中；9-文件转换失败"`

	//DownloadUrl   string `json:"downloadUrl" description:"PDF文件下载链接，有效期60分钟"`
	//PdfTotalPages int    `json:"pdfTotalPages" description:"pdf文件总页数,仅当文件类型为pdf时有值"`
}

func (req *MergeThenUploadTemplatesRequest) Handler(ctx *gin.Context) {
	// validate and clean input template ids
	templateids_ := strings.Split(req.TemplateIds, ",")
	if len(templateids_) < 2 {
		common.RespErrf(ctx, http.StatusBadRequest, "There must be at least two files for merge.")
		return
	}
	var temlateids []string
	for _, fileid := range templateids_ {
		temlateids = append(temlateids, strings.TrimSpace(fileid))
	}
	result, errObj := MergeThenUploadTemplates(templateids_)
	if errObj != nil {
		common.RespErrObj(ctx, errObj)
		return
	}

	// todo: return more info
	common.RespSucc(ctx, result)
}

func MergeThenUploadTemplates(templateIds []string) (data *MergeThenUploadTemplatesResponseData, errObj *common.RespObj) {
	// merge templates relying on esign storage
	mergeResult, errObj := mergeTemplates(templateIds...)
	if errObj != nil {
		return nil, errObj
	}

	// require url and id for template to upload
	contentMD5 := utils.MD5B64(mergeResult.fileBody)
	contentType := "application/pdf"
	fileSize := int64(len(mergeResult.fileBody))
	//if ext != ".pdf" {	// currently only deal with pdf
	//	contentType = "application/octet-stream"
	//}

	uploadUrlAndId, errObj := getTemplUploadUrl(mergeResult.joinName, contentMD5, contentType)
	if errObj != nil {
		return nil, errObj
	}

	// perform file upload chores

	if errObj := uploadTemplateFile(mergeResult.fileBody, contentMD5, contentType, uploadUrlAndId.UploadUrl); errObj != nil {
		return nil, errObj
	}

	// save upload results meta data
	createTemplateParams := models.CreateTemplateParams{
		TemplateID:   uploadUrlAndId.TemplateId,
		TemplateName: mergeResult.joinName,
		DocType:      models.DocType("0-合同"),
		CreatorID:    gofakeit.UUID(), // todo: use true id
		FileSize:     fileSize,
		FileBody:     mergeResult.fileBody,
	}

	_, err := models.New(common.PgxPool).CreateTemplate(context.Background(), &createTemplateParams)
	if err != nil {
		return nil, &common.RespObj{
			Code: http.StatusInternalServerError, Msg: fmt.Sprintf("failed to save merged templates %q:%v", mergeResult.joinName, err),
		}
	}

	// todo: return more info
	return &MergeThenUploadTemplatesResponseData{
		TemplateId:   uploadUrlAndId.TemplateId,
		TemplateName: mergeResult.joinName,
		Size:         cast.ToString(fileSize), // todo: to unify the responded size data format
	}, nil
}

var MergeThenUploadTemplatesRequestH = func() *router.Router {
	//var apiDesc = func() string {
	//	builder := markdown.Builder{}
	//	builder.P("官方文档：")
	//	return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyingmd&namespace=opendoc%2Fsaas_api", "合并PDF文件.")
	//}

	r := router.New(
		&MergeThenUploadTemplatesRequest{},
		router.Summary("合并PDF模板."),
		//router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: struct {
					Code int                                  `json:"code" binding:"required" default:"0"`
					Msg  string                               `json:"msg" binding:"required" default:"ok"`
					Data MergeThenUploadTemplatesResponseData `json:"data"`
				}{},
			},
		}),
	)

	return r
}
