package file_

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/utils"
	"github.com/long2ice/swagin/router"
	"github.com/wI2L/fizz/markdown"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fgcu36n&namespace=opendoc%2Fsaas_api

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

type UploadFileRequest struct {
	DocType string `uri:"docType" default:"0-合同" description:"文档类型：0-合同, 1-协议, 2-订单.`
	//File *multipart.FileHeader `form:"file" binding:"required" description:"文件名称必须带扩展名:.pdf"`
}

type UploadFileResponseData struct {
	FileId   string `json:"FileId" binding:"required" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"文件ID."`
	FileName string `json:"fileName" binding:"required" default:"商贷通用收入证明.pdf" description:"文件名称，必须带扩展名."`
}

// todo: log strategy improvement

func (req *UploadFileRequest) Handler(ctx *gin.Context) {
	//// check context availability before workTemplateId
	//if err := common.PgxPool.Ping(context.Background()); err != nil {
	//	common.RespErrf(ctx, http.StatusInternalServerError, "database is down! service refused:%v", err)
	//	return
	//}
	//
	//oauth, err := token.GetOauthInfo()
	//if err != nil {
	//	common.RespErrf(ctx, http.StatusInternalServerError, "got an error for esign authentication:%v", err)
	//	return
	//}

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

	uploadUrlAndId, errObj := getUploadUrl(fileName, contentMD5, contentType, fileSize)
	if errObj != nil {
		common.RespErrObj(ctx, errObj)
		return
	}

	// perform file upload chores

	if errObj := uploadFile(fileBytes, contentMD5, contentType, uploadUrlAndId.UploadUrl); errObj != nil {
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
	if err != nil {
		common.RespErrf(ctx, http.StatusInternalServerError, "failed to create template %q in database for template upload:%v", fileName, err)
		return
	}

	common.RespSucc(ctx, UploadFileResponseData{FileId: uploadUrlAndId.FileId, FileName: fileName})
}

var UploadFileRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fawgyis&namespace=opendoc%2Fsaas_api", "模板文件上传.")
	}

	r := router.New(
		&UploadFileRequest{},
		router.Summary("模板文件上传. docType文档类型：0-合同, 1-协议, 2-订单. 注意：因UI界面限制，请用Postman、curl或其他工具，通过表单上传模板文件。"),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.ContentType(binding.MIMEMultipartPOSTForm),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: struct {
					Code int                    `json:"code" binding:"required" default:"0"`
					Msg  string                 `json:"msg" binding:"required" default:"ok"`
					Data UploadFileResponseData `json:"data"`
				}{},
			},
		}),
	)

	return r
}
