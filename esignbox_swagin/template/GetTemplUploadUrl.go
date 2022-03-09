package template

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/token"
	"github.com/long2ice/swagin/router"
	"github.com/pkg/errors"
	"github.com/wI2L/fizz/markdown"
)

const (
	EsignSandBoxUploadTemplUrl = "/v1/docTemplates/createByUploadUrl"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fawgyis&namespace=opendoc%2Fsaas_api
//
//type GetTemplUploadUrlRequest struct {
//	Body TemplBasicInfo `description:"doc格式模板文件，如 商贷通用收入证明.doc"`
//}

type GetTemplUploadUrlRequest struct {
	FileName string `form:"fileName" json:"fileName" binding:"required" default:"商贷通用收入证明.pdf" description:"文件名称，必须带扩展名:.pdf."`

	ContentMd5 string `form:"contentMd5" json:"contentMd5" binding:"required" default:"1Qinq+/TLV3UZGbzifvmjw==" description:"模板文件md5值，先计算文件md5值，在对该md5值进行base64编码,可使用E签宝官网工具进行计算.https://smlopen.esign.cn/tools/file-md5"`

	ContentType string `json:"-" binding:"required" enum:"application/pdf" description:"目标文件的MIME类型.注意，后面文件流上传的Content-Type参数要和这里一致，不然就会有403的报错"` // todo:check enum
	Convert2Pdf bool   `json:"-" enum:"false" description:"是否需要转成pdf"`
}

type GetTemplUploadUrlResponse struct {
	Code int                           `json:"code" doc:"业务码，0表示成功"`
	Msg  string                        `json:"message" doc:"信息"`
	Data GetTemplUploadUrlResponseData `json:"data" doc:"业务信息"`
}

type GetTemplUploadUrlResponseData struct {
	TemplateId string `json:"templateId" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板ID"`
	UploadUrl  string `json:"uploadUrl" description:"文件上传地址，链接有效期60分钟。"`
}

/*
{
  "templateId": "c4d4fe1b48184ba28982f68bf2c7bf25",
  "uploadUrl": "http://esignoss.esign.cn/1111564182/8759b09d-4b3e-426f-8ae1-03b1fca3daf6/%E5%95%86%E8%B4%B7%E9%80%9A%E7%94%A8%E6%94%B6%E5%85%A5%E8%AF%81%E6%98%8E.pdf?Expires=1646216465&OSSAccessKeyId=STS.NTfJkmomE8v8hAUvB1NZPHVsJ&Signature=GKEO9kXnX0LagjVE%2BM6vYGT2%2BeA%3D&callback-var=eyJ4OmZpbGVfa2V5IjoiJDk2YzNkYzExLTYxMGYtNDY0Ny05Zjg2LTE5YWNiNmEwYzJmMiQ5ODk3MTE1MzUifQ%3D%3D%0A&callback=eyJjYWxsYmFja1VybCI6Imh0dHA6Ly9maWxlc3lzdGVtLXNtbC50c2lnbi5jbi9maWxlLXN5c3RlbS9jYWxsYmFjay9hbGlvc3MiLCJjYWxsYmFja0JvZHkiOiAie1wibWltZVR5cGVcIjoke21pbWVUeXBlfSxcInNpemVcIjogJHtzaXplfSxcImJ1Y2tldFwiOiAke2J1Y2tldH0sXCJvYmplY3RcIjogJHtvYmplY3R9LFwiZXRhZ1wiOiAke2V0YWd9LFwiZmlsZV9rZXlcIjoke3g6ZmlsZV9rZXl9fSIsImNhbGxiYWNrQm9keVR5cGUiOiAiYXBwbGljYXRpb24vanNvbiJ9%0A&security-token=CAIS%2BAF1q6Ft5B2yfSjIr5fTAdHZgrJkj7TTamfkkkJkQtV8p5PYqDz2IHtKdXRvBu8Xs%2F4wnmxX7f4YlqB6T55OSAmcNZEob27lS5%2FmMeT7oMWQweEurv%2FMQBqyaXPS2MvVfJ%2BOLrf0ceusbFbpjzJ6xaCAGxypQ12iN%2B%2Fm6%2FNgdc9FHHPPD1x8CcxROxFppeIDKHLVLozNCBPxhXfKB0ca0WgVy0EHsPnvm5DNs0uH1AKjkbRM9r6ceMb0M5NeW75kSMqw0eBMca7M7TVd8RAi9t0t1%2FIVpGiY4YDAWQYLv0rda7DOltFiMkpla7MmXqlft%2BhzcgeQY0pc%2FRqAAUwxVoW0nfk4w92%2Bm0ikUU43mcQR3XgrTF2geBGPg5GvqQVuUaZvPdzQIODgCG2tcEvUuMmKU%2FqbGzuj7bWEOrdrIfV7UXwSAa9NdJ5okZmmQvoKAmnCWiocgK1rG%2B9ItCGc5f5JAiYH%2F0UrxI8HKE6FUFT5oWLvgngYiKEm%2FDAk"
}
*/

func (req *GetTemplUploadUrlRequest) Handler(ctx *gin.Context) {
	parsedResp := GetTemplUploadUrlResponse{}
	oauth, err := token.GetOauthInfo()
	if err != nil {
		common.WriteError(ctx, 400, fmt.Sprintf("got an error when try to get authentication info:%w", err))
		return
	}
	//
	//req.Body.ContentType = "application/octet-stream"
	//req.Body.Convert2Pdf = true

	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
		"Content-Type":        oauth.ContentType,
	}).SetBody(&req).
		SetResult(&parsedResp).Post(EsignSandBoxUploadTemplUrl)

	if common.WriteErrore(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
		return
	}

	common.WriteOK(ctx, GetTemplUploadUrlResponseData{parsedResp.Data.TemplateId, parsedResp.Data.UploadUrl})

	// TODO: save file to db
}

var GetTemplUploadUrlRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("请求参数说明")
		builder.Table(
			[][]string{
				[]string{"参数名称", "类型", "必选", "参数说明", "举例"},
				[]string{"fileName", "string", "是", "文件名称，必须带扩展名:.pdf", "商贷通用收入证明.pdf"},
				[]string{"contentMd5", "string", "是", "有效截止时间（毫秒）", "1Qinq+/TLV3UZGbzifvmjw=="},
			}, []markdown.TableAlignment{
				markdown.AlignLeft,
				markdown.AlignCenter,
			},
		)
		builder.P("模板文件md5值，先计算文件md5值，在对该md5值进行base64编码,可使用E签宝官网工具进行计算.https://smlopen.esign.cn/tools/file-md5")

		builder.P("官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fawgyis&namespace=opendoc%2Fsaas_api", "获取模板文件上传地址.")
	}

	//var respModelDesc = func() string {
	//	builder := markdown.Builder{}
	//	builder.Table(
	//		[][]string{
	//			[]string{"参数名称", "类型", "必选", "参数说明", "举例"},
	//			[]string{"fileName", "string", "是", "文件名称，必须带扩展名:.pdf", "商贷通用收入证明.pdf"},
	//			[]string{"contentMd5", "string", "是", "有效截止时间（毫秒）", "1Qinq+/TLV3UZGbzifvmjw=="},
	//		}, []markdown.TableAlignment{
	//			markdown.AlignLeft,
	//			markdown.AlignCenter,
	//		},
	//	)
	//	builder.P("模板文件md5值，先计算文件md5值，在对该md5值进行base64编码,可使用E签宝官网工具进行计算.https://smlopen.esign.cn/tools/file-md5")
	//
	//	return builder.String()
	//}

	r := router.New(
		&GetTemplUploadUrlRequest{},
		router.Summary("获取模板文件上传地址."),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.ContentType(binding.MIMEJSON),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: GetTemplUploadUrlResponseData{},
				//Description: respModelDesc(),
			},
			"400": router.ResponseItem{
				Model: common.ErrorResp{},
			},
		}),
	)

	return r
}

func getTemplUploadUrl(fileName, contentMD5 string, contentType string) (*GetTemplUploadUrlResponseData, error) {
	type getPDFTemplUploadUrlReq struct {
		FileName    string `json:"fileName"`
		ContentMd5  string `json:"contentMd5"`
		ContentType string `json:"contentType"`
		Convert2Pdf bool   `json:"convert2Pdf"`
	}

	parsedResp := GetTemplUploadUrlResponse{}
	oauth, err := token.GetOauthInfo()
	if err != nil {
		return nil, err
	}

	convert2Pdf := false
	if contentType != "application/pdf" {
		convert2Pdf = true
	}

	req := getPDFTemplUploadUrlReq{
		FileName:    fileName,
		ContentMd5:  contentMD5,
		ContentType: contentType,
		Convert2Pdf: convert2Pdf,
	}

	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
		"Content-Type":        oauth.ContentType,
	}).SetBody(&req).
		SetResult(&parsedResp).Post(EsignSandBoxUploadTemplUrl)

	_ = restyResp

	if err != nil {
		return nil, err
	}

	if parsedResp.Code != 0 {
		return nil, errors.Errorf("error code:%v, error message:%v", parsedResp.Code, parsedResp.Msg)
	}

	return &GetTemplUploadUrlResponseData{parsedResp.Data.TemplateId, parsedResp.Data.UploadUrl}, nil
}
