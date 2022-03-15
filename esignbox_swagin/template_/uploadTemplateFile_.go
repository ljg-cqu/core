package template_

import (
	"fmt"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/token"
	"net/http"
)

type uploadFileResponse struct {
	Code int    `json:"errCode" required:"true" description:"业务码，0表示成功"`
	Msg  string `json:"message" required:"true" description:"信息"`
}

func uploadTemplateFile(body []byte, contentMD5, contentType, uploadUrl string) (errObj *common.RespObj) {
	oauth, err := token.GetOauthInfo()
	if err != nil {
		return &common.RespObj{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("got an error for esign authentication:%v", err),
		}
	}

	parsedResp := uploadFileResponse{}
	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
	}).SetHeaders(map[string]string{
		"Content-MD5":  contentMD5,
		"Content-Type": contentType,
	}).SetBody(body).SetResult(&parsedResp).Put(uploadUrl)

	return common.ErrObjFromEsignRequest(restyResp.RawResponse, err, &common.EsignError{
		Code: parsedResp.Code,
		Msg:  parsedResp.Msg,
	})
}
