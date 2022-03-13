package file_

import (
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/utils"
	"github.com/long2ice/swagin/router"
	"net/http"
	"strings"
)

//const (
//	EsignSandBoxQueryPdfFileDetailsUrl = "/v1/files/{fileId}"
//)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyingmd&namespace=opendoc%2Fsaas_api

type MergeThenUploadFilesRequest struct {
	FileIds string `uri:"fileIds" binding:"required" default:"ede1fa4504954c29ad210637c15f42cf,ede1fa4504954c29ad210637c15f42cf" description:"用于合并的PDF文件ID，以逗号按顺序连续排列"`
}

type MergeThenUploadFilesResponse struct {
	Code int                              `json:"code" binding:"required:" description:"业务码，0表示成功"`
	Msg  string                           `json:"message" description:"信息"`
	Data MergeThenUploadFilesResponseData `json:"data" description:"业务信息"`
}

type MergeThenUploadFilesResponseData struct {
	FileId   string `json:"fileId" default:"ede1fa4504954c29ad210637c15f42cf" description:"文件ID"`
	FileName string `json:"name" default:"商贷通用收入证明.pdf" description:"文件名称"`
	Size     string `json:"size" description:"文件大小，单位byte"`
	//Status   int    `json:"status" description:"文件上传状态: 0-文件未上传；1-文件上传中 ；2-文件上传已完成,；3-文件上传失败 ；4-文件等待转pdf ；5-文件已转换pdf ；6-加水印中；7-加水印完毕；8-文件转换中；9-文件转换失败"`

	//DownloadUrl   string `json:"downloadUrl" description:"PDF文件下载链接，有效期60分钟"`
	//PdfTotalPages int    `json:"pdfTotalPages" description:"pdf文件总页数,仅当文件类型为pdf时有值"`
}

func (req *MergeThenUploadFilesRequest) Handler(ctx *gin.Context) {
	//parsedResp := MergeThenUploadFilesResponse{}
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

	// validate input file ids
	fileids_ := strings.Split(req.FileIds, ",")
	if len(fileids_) < 2 {
		common.RespErr(ctx, http.StatusInternalServerError, "There must be at least two files for merge.")
		return
	}
	var fileids []string
	for _, fileid := range fileids_ {
		fileids = append(fileids, strings.TrimSpace(fileid))
	}

	mergeResult, errObj := mergeFiles(fileids...)
	if errObj != nil {
		common.RespErrObj(ctx, errObj)
		return
	}

	_ = mergeResult

	// require url and id for template to upload
	contentMD5 := utils.MD5B64(mergeResult.fileBody)
	contentType := "application/pdf"
	fileSize := int64(len(mergeResult.fileBody))
	//if ext != ".pdf" {
	//	contentType = "application/octet-stream"
	//}

	uploadUrlAndId, errObj := getUploadUrl(mergeResult.joinName, contentMD5, contentType, fileSize)
	if errObj != nil {
		common.RespErrObj(ctx, errObj)
		return
	}

	// perform file upload chores

	if errObj := uploadFile(mergeResult.fileBody, contentMD5, contentType, uploadUrlAndId.UploadUrl); errObj != nil {
		common.RespErrObj(ctx, errObj)
		return
	}

	//parsedResp := uploadFileResponse{}
	//restyResp, err := common.Client.R().SetHeaders(map[string]string{
	//	"X-Tsign-Open-App-Id": oauth.AppId,
	//	"X-Tsign-Open-Token":  oauth.Token,
	//}).SetHeaders(map[string]string{
	//	"Content-MD5":  contentMD5,
	//	"Content-Type": contentType,
	//}).
	//	SetBody(fileBytes).
	//	SetResult(&parsedResp).Put(uploadUrlAndId.UploadUrl)
	//
	//if common.RespErre(ctx, restyResp.RawResponse, err, &common.EsignError{parsedResp.Code, parsedResp.Msg}) {
	//	return
	//}

	//_, err = models.New(common.PgxPool).CreateContractFile(context.Background(), &models.CreateContractFileParams{
	//	FileID:    uploadUrlAndId.FileId,
	//	FileName:  fileName,
	//	CreatorID: gofakeit.UUID(),
	//	FileSize:  fileSize,
	//	FileBody:  fileBytes,
	//})
	//if err != nil {
	//	common.RespErrf(ctx, http.StatusInternalServerError, "failed to create template %q in database for template upload:%v", fileName, err)
	//	return
	//}

	common.RespSucc(ctx, MergeThenUploadFilesResponseData{FileId: uploadUrlAndId.FileId, FileName: mergeResult.joinName})
}

var MergeThenUploadFilesRequestH = func() *router.Router {
	//var apiDesc = func() string {
	//	builder := markdown.Builder{}
	//	builder.P("官方文档：")
	//	return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyingmd&namespace=opendoc%2Fsaas_api", "合并PDF文件.")
	//}

	r := router.New(
		&MergeThenUploadFilesRequest{},
		router.Summary("合并PDF文件."),
		//router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: struct {
					Code int                              `json:"code" binding:"required" default:"0"`
					Msg  string                           `json:"msg" binding:"required" default:"ok"`
					Data MergeThenUploadFilesResponseData `json:"data"`
				}{},
			},
		}),
	)

	return r
}
