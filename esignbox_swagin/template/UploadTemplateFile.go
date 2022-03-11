package template

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/template/models/models"
	"github.com/ljg-cqu/core/esignbox_swagin/token"
	"github.com/ljg-cqu/core/utils"
	"github.com/long2ice/swagin/router"
	"github.com/wI2L/fizz/markdown"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fawgyis&namespace=opendoc%2Fsaas_api

type UploadTemplFileRequest struct {
	//File *multipart.FileHeader `form:"file" binding:"required" description:"文件名称必须带扩展名:.pdf"`
}

type UploadTemplFileResponse struct {
	Code int    `json:"errCode" required:"true" description:"业务码，0表示成功"`
	Msg  string `json:"message" required:"true" description:"信息"`
}

type UploadTemplFileResponseData struct {
	TemplateId string `json:"templateId" binding:"required" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板ID."`
	FileName   string `json:"fileName" binding:"required" default:"商贷通用收入证明.pdf" description:"文件名称，必须带扩展名."`
}

// todo: log strategy improvement

func (req *UploadTemplFileRequest) Handler(ctx *gin.Context) {
	var supported_File_Formats = map[string]struct{}{
		".pdf":  struct{}{},
		".docx": struct{}{},
		".doc":  struct{}{},
		".rtf":  struct{}{},
		".xlsx": struct{}{},
		".xls":  struct{}{},
		".pptx": struct{}{},
		".ppt":  struct{}{},
		".wps":  struct{}{},
		".et":   struct{}{},
		".dps":  struct{}{},
		".jpeg": struct{}{},
		".jpg":  struct{}{},
		".png":  struct{}{},
		".bmp":  struct{}{},
		".tiff": struct{}{},
		".tif":  struct{}{},
		".gif":  struct{}{},
		".html": struct{}{},
		".htm":  struct{}{},
	}

	// check context availability before work
	if err := common.PgxPool.Ping(context.Background()); err != nil {
		common.RespErrf(ctx, http.StatusInternalServerError, "database is down! service refused:%v", err)
		return
	}

	parsedResp := UploadTemplFileResponse{}
	oauth, err := token.GetOauthInfo()
	if err != nil {
		common.RespErrf(ctx, http.StatusInternalServerError, "got an error for esign authentication:%v", err)
		return
	}

	// read file stream
	fileH, err := ctx.FormFile("file")
	if err != nil {
		common.RespErrf(ctx, http.StatusBadRequest, "failed to receive file:%v", err)
		return
	}

	if fileH == nil {
		common.RespErrf(ctx, http.StatusBadRequest, "failed to receive file: file is nil")
		return
	}

	fileName := fileH.Filename
	fileSize := fileH.Size

	ext := filepath.Ext(fileName)
	if _, ok := supported_File_Formats[ext]; !ok {
		common.RespErrf(ctx, http.StatusBadRequest, "bad file extension %q!", ext)
		return
	}

	file, err := fileH.Open()
	defer file.Close()
	if err != nil {
		common.RespErrf(ctx, http.StatusInternalServerError, "got and error when open file %q:%v", fileName, err)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		common.RespErrf(ctx, http.StatusInternalServerError, "failed to read %q file body:%v", fileName, err)
		return
	}

	// require url and id for template to upload
	contentMD5 := utils.MD5B64(fileBytes)
	contentType := "application/pdf"
	if ext != ".pdf" {
		contentType = "application/octet-stream"
	}

	uploadUrlAndId, err := getTemplUploadUrl(fileName, contentMD5, contentType)
	if err != nil {
		common.RespErrf(ctx, http.StatusInternalServerError, "got and error when request template upload url for %q:%v", fileName, err)
		return
	}

	// perform template upload chores
	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
	}).SetHeaders(map[string]string{
		"Content-MD5":  contentMD5,
		"Content-Type": contentType,
	}).
		SetBody(fileBytes).
		SetResult(&parsedResp).Put(uploadUrlAndId.UploadUrl)

	_ = restyResp

	if err != nil {
		common.RespErrf(ctx, http.StatusBadRequest, "failed to upload file %q to esign, error:%v", fileName, err)
		return
	}

	if parsedResp.Code != 0 {
		common.RespErrf(ctx, http.StatusBadRequest, "failed to upload file %q to esign, error code:%v, error message:%v", fileName, parsedResp.Code, parsedResp.Msg)
		return
	}

	// issue a task to sync template upload status
	// TODO:

	_, err = models.New(common.PgxPool).CreateTemplate(context.Background(), &models.CreateTemplateParams{
		TemplateID:   uploadUrlAndId.TemplateId,
		TemplateName: fileName,
		FileSize:     fileSize,
		FileBody:     fileBytes,
	})
	if err != nil {
		common.RespErrf(ctx, http.StatusInternalServerError, "failed to create template %q in database for template upload:%v", fileName, err)
		return
	}

	common.RespSucc(ctx, UploadTemplFileResponseData{TemplateId: uploadUrlAndId.TemplateId, FileName: fileName})
}

var UploadPDFTemplFileRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fawgyis&namespace=opendoc%2Fsaas_api", "模板文件上传.")
	}

	r := router.New(
		&UploadTemplFileRequest{},
		router.Summary("模板文件上传. 注意：因UI界面限制，请用Postman、curl或其他工具，通过表单上传模板文件。"),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.ContentType(binding.MIMEMultipartPOSTForm),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: struct {
					Code int                         `json:"code" binding:"required" default:"0"`
					Msg  string                      `json:"msg" binding:"required" default:"ok"`
					Data UploadTemplFileResponseData `json:"data"`
				}{},
			},
		}),
	)

	return r
}
