package file_

import (
	"fmt"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/token"
	"net/http"
)

const (
	EsignSandBoxUploadUrl = "/v1/files/getUploadUrl"
)

type getUploadUrlResponse struct {
	Code int                      `json:"code" doc:"业务码，0表示成功"`
	Msg  string                   `json:"message" doc:"信息"`
	Data getUploadUrlResponseData `json:"data" doc:"业务信息"`
}

type getUploadUrlResponseData struct {
	FileId    string `json:"fileId" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板ID"`
	UploadUrl string `json:"uploadUrl" description:"文件上传地址，链接有效期60分钟。"`
}

func getUploadUrl(fileName, contentMD5 string, contentType string, fileSize int64) (data *getUploadUrlResponseData, errObj *common.RespObj) {
	type getPDFUploadUrlReq struct {
		FileName    string `json:"fileName"`
		ContentMd5  string `json:"contentMd5"`
		ContentType string `json:"contentType"`
		FileSize    int64  `json:"fileSize"`
		Convert2Pdf bool   `json:"convert2Pdf"`
	}

	parsedResp := getUploadUrlResponse{}
	oauth, err := token.GetOauthInfo()
	if err != nil {
		return nil, &common.RespObj{
			Code: http.StatusNetworkAuthenticationRequired,
			Msg:  fmt.Sprintf("got an error when try to get authentication info:%v", err),
		}
	}

	convert2Pdf := false
	if contentType != "application/pdf" {
		convert2Pdf = true
	}

	req := getPDFUploadUrlReq{
		FileName:    fileName,
		ContentMd5:  contentMD5,
		ContentType: contentType,
		FileSize:    fileSize,
		Convert2Pdf: convert2Pdf,
	}

	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
		"Content-Type":        oauth.ContentType,
	}).SetBody(&req).
		SetResult(&parsedResp).Post(EsignSandBoxUploadUrl)

	errObj = common.ErrObjFromEsignRequest(restyResp.RawResponse, err, &common.EsignError{Code: parsedResp.Code, Msg: parsedResp.Msg})
	if errObj != nil {
		return
	}

	return &getUploadUrlResponseData{parsedResp.Data.FileId, parsedResp.Data.UploadUrl}, nil
}
