package template

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/token"
	"github.com/ljg-cqu/core/utils"
	"github.com/long2ice/swagin/router"
	"github.com/wI2L/fizz/markdown"
	"io/ioutil"
	"strings"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fawgyis&namespace=opendoc%2Fsaas_api

type UploadPDFTemplFileRequest struct {
	//ContentMd5     string `header:"Content-MD5" example:"1Qinq+/TLV3UZGbzifvmjw==" description:"与【步骤1 获取文件上传地址】Body体中contentMd5值一致"`
	//ContentType string `header:"Content-Type" default:"multipart/form-data"`
	//TemplUploadUrl string `query:"templUploadUrl" binding:"required" description:"模板文件上传地址" example:"http://esignoss.esign.cn/1111564182/8759b09d-4b3e-426f-8ae1-03b1fca3daf6/%E5%95%86%E8%B4%B7%E9%80%9A%E7%94%A8%E6%94%B6%E5%85%A5%E8%AF%81%E6%98%8E.pdf?Expires=1646216465&OSSAccessKeyId=STS.NTfJkmomE8v8hAUvB1NZPHVsJ&Signature=GKEO9kXnX0LagjVE%2BM6vYGT2%2BeA%3D&callback-var=eyJ4OmZpbGVfa2V5IjoiJDk2YzNkYzExLTYxMGYtNDY0Ny05Zjg2LTE5YWNiNmEwYzJmMiQ5ODk3MTE1MzUifQ%3D%3D%0A&callback=eyJjYWxsYmFja1VybCI6Imh0dHA6Ly9maWxlc3lzdGVtLXNtbC50c2lnbi5jbi9maWxlLXN5c3RlbS9jYWxsYmFjay9hbGlvc3MiLCJjYWxsYmFja0JvZHkiOiAie1wibWltZVR5cGVcIjoke21pbWVUeXBlfSxcInNpemVcIjogJHtzaXplfSxcImJ1Y2tldFwiOiAke2J1Y2tldH0sXCJvYmplY3RcIjogJHtvYmplY3R9LFwiZXRhZ1wiOiAke2V0YWd9LFwiZmlsZV9rZXlcIjoke3g6ZmlsZV9rZXl9fSIsImNhbGxiYWNrQm9keVR5cGUiOiAiYXBwbGljYXRpb24vanNvbiJ9%0A&security-token=CAIS%2BAF1q6Ft5B2yfSjIr5fTAdHZgrJkj7TTamfkkkJkQtV8p5PYqDz2IHtKdXRvBu8Xs%2F4wnmxX7f4YlqB6T55OSAmcNZEob27lS5%2FmMeT7oMWQweEurv%2FMQBqyaXPS2MvVfJ%2BOLrf0ceusbFbpjzJ6xaCAGxypQ12iN%2B%2Fm6%2FNgdc9FHHPPD1x8CcxROxFppeIDKHLVLozNCBPxhXfKB0ca0WgVy0EHsPnvm5DNs0uH1AKjkbRM9r6ceMb0M5NeW75kSMqw0eBMca7M7TVd8RAi9t0t1%2FIVpGiY4YDAWQYLv0rda7DOltFiMkpla7MmXqlft%2BhzcgeQY0pc%2FRqAAUwxVoW0nfk4w92%2Bm0ikUU43mcQR3XgrTF2geBGPg5GvqQVuUaZvPdzQIODgCG2tcEvUuMmKU%2FqbGzuj7bWEOrdrIfV7UXwSAa9NdJ5okZmmQvoKAmnCWiocgK1rG%2B9ItCGc5f5JAiYH%2F0UrxI8HKE6FUFT5oWLvgngYiKEm%2FDAk"`

	//File *multipart.FileHeader `form:"file" binding:"required" description:"文件名称必须带扩展名:.pdf"`
}

type UploadTemplFileResponse struct {
	Code int    `json:"errCode" required:"true" doc:"业务码，0表示成功"`
	Msg  string `json:"message" doc:"信息"`
}

type UploadTemplFileResponseData struct {
	TemplateId string `json:"templateId" example:"c4d4fe1b48184ba28982f68bf2c7bf25" doc:"模板ID"`
	FileName   string `json:"fileName" binding:"required" default:"商贷通用收入证明.pdf" description:"文件名称，必须带扩展名:.pdf."`
	//UploadUrl  string `json:"uploadUrl" doc:"文件上传地址，链接有效期60分钟。" example:"http://esignoss.esign.cn/1111564182/8759b09d-4b3e-426f-8ae1-03b1fca3daf6/%E5%95%86%E8%B4%B7%E9%80%9A%E7%94%A8%E6%94%B6%E5%85%A5%E8%AF%81%E6%98%8E.pdf?Expires=1646216465&OSSAccessKeyId=STS.NTfJkmomE8v8hAUvB1NZPHVsJ&Signature=GKEO9kXnX0LagjVE%2BM6vYGT2%2BeA%3D&callback-var=eyJ4OmZpbGVfa2V5IjoiJDk2YzNkYzExLTYxMGYtNDY0Ny05Zjg2LTE5YWNiNmEwYzJmMiQ5ODk3MTE1MzUifQ%3D%3D%0A&callback=eyJjYWxsYmFja1VybCI6Imh0dHA6Ly9maWxlc3lzdGVtLXNtbC50c2lnbi5jbi9maWxlLXN5c3RlbS9jYWxsYmFjay9hbGlvc3MiLCJjYWxsYmFja0JvZHkiOiAie1wibWltZVR5cGVcIjoke21pbWVUeXBlfSxcInNpemVcIjogJHtzaXplfSxcImJ1Y2tldFwiOiAke2J1Y2tldH0sXCJvYmplY3RcIjogJHtvYmplY3R9LFwiZXRhZ1wiOiAke2V0YWd9LFwiZmlsZV9rZXlcIjoke3g6ZmlsZV9rZXl9fSIsImNhbGxiYWNrQm9keVR5cGUiOiAiYXBwbGljYXRpb24vanNvbiJ9%0A&security-token=CAIS%2BAF1q6Ft5B2yfSjIr5fTAdHZgrJkj7TTamfkkkJkQtV8p5PYqDz2IHtKdXRvBu8Xs%2F4wnmxX7f4YlqB6T55OSAmcNZEob27lS5%2FmMeT7oMWQweEurv%2FMQBqyaXPS2MvVfJ%2BOLrf0ceusbFbpjzJ6xaCAGxypQ12iN%2B%2Fm6%2FNgdc9FHHPPD1x8CcxROxFppeIDKHLVLozNCBPxhXfKB0ca0WgVy0EHsPnvm5DNs0uH1AKjkbRM9r6ceMb0M5NeW75kSMqw0eBMca7M7TVd8RAi9t0t1%2FIVpGiY4YDAWQYLv0rda7DOltFiMkpla7MmXqlft%2BhzcgeQY0pc%2FRqAAUwxVoW0nfk4w92%2Bm0ikUU43mcQR3XgrTF2geBGPg5GvqQVuUaZvPdzQIODgCG2tcEvUuMmKU%2FqbGzuj7bWEOrdrIfV7UXwSAa9NdJ5okZmmQvoKAmnCWiocgK1rG%2B9ItCGc5f5JAiYH%2F0UrxI8HKE6FUFT5oWLvgngYiKEm%2FDAk"`
}

// TODO: save file to db
// TODO： avoid upload duplicate pdf template.

func (req *UploadPDFTemplFileRequest) Handler(ctx *gin.Context) {
	// check if database work normall todo:
	// use db as backup, but read from esign as newly.
	parsedResp := UploadTemplFileResponse{}
	oauth, err := token.GetOauthInfo()
	if err != nil {
		common.WriteError(ctx, 400, fmt.Sprintf("got an error when try to get authentication info:%w", err))
		return
	}

	fileH, _ := ctx.FormFile("file")
	fileName := fileH.Filename
	if !strings.HasSuffix(fileName, ".pdf") {
		common.WriteError(ctx, 400, "文件名称格式错误。文件名称必须带扩展名:.pdf")
		return
	}

	//fileName := req.File.Filename

	//file, _ := req.File.Open()
	file, err := fileH.Open()
	defer file.Close()
	if err != nil {
		common.WriteError(ctx, 400, fmt.Sprintf("got and error when open file:%v", err))
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		common.WriteError(ctx, 400, fmt.Sprintf("failed to read file body:%v", err))
		return
	}

	contentMD5 := utils.MD5B64(fileBytes)

	uploadUrlId, err := getPDFTemplUploadUrl(fileName, contentMD5)
	if err != nil {
		common.WriteError(ctx, 400, fmt.Sprintf("got and error when request pdf template upload url:%v", err))
		return
	}

	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
	}).SetHeaders(map[string]string{
		"Content-MD5":  contentMD5,
		"Content-Type": "application/pdf",
	}).
		SetBody(fileBytes).
		SetResult(&parsedResp).Put(uploadUrlId.UploadUrl)

	if common.WriteErrorE(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
		return
	}

	common.WriteOK(ctx, UploadTemplFileResponseData{TemplateId: uploadUrlId.TemplateId, FileName: fileName})
}

var UploadPDFTemplFileRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fawgyis&namespace=opendoc%2Fsaas_api", "PDF模板文件上传.")
	}

	r := router.New(
		&UploadPDFTemplFileRequest{},
		router.Summary("PDF模板文件上传. 注意：因UI界面限制，请用Postman、curl或其他工具，通过表单上传模板文件。"),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.ContentType(binding.MIMEMultipartPOSTForm),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: UploadTemplFileResponseData{},
			},
			"400": router.ResponseItem{
				Model: common.ErrorResp{},
			},
		}),
	)

	return r
}
