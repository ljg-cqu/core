package template_

import (
	"fmt"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// https://github.com/pdfcpu/pdfcpu/blob/master/pkg/api/test/merge_test.go
// https://github.com/phpdave11/gofpdi

type mergeTemplateResult struct {
	joinName string
	fileBody []byte
	//TemplateId string //`json:"templateId" binding:"required" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板ID."`
	//TemplateName   string `json:"fileName" binding:"required" default:"商贷通用收入证明.pdf" description:"文件名称，必须带扩展名."`
}

// mergeTemplates merge template files that found from esign platform
// todo: add pdf limit?
func mergeTemplates(templateIds ...string) (result *mergeTemplateResult, errObj *common.RespObj) {
	// esign authentication
	//oauth, err := token.GetOauthInfo()
	//if err != nil {
	//	return nil, &common.RespObj{
	//		Code: http.StatusNetworkAuthenticationRequired,
	//		Msg:  fmt.Sprintf("got an error when try to get authentication info:%v", err),
	//	}
	//}

	if len(templateIds) < 2 {
		return nil, &common.RespObj{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf(" not enough file number to merge"),
		}
	}

	// get files from esign
	//var files []*os.File
	var filenames []string
	var tempDir = "./merge_temp_dir/"
	for _, templateId := range templateIds {
		if templateId == "" {
			return nil, &common.RespObj{
				Code: http.StatusInternalServerError,
				Msg:  fmt.Sprintf("file id should not be empty string"),
			}
		}

		// get template details, including save its content with download url
		details, errObj := GetTemplDetails(templateId)
		if errObj != nil {
			return nil, errObj
		}

		//parsedResp := MergeThenUploadFilesResponse{}
		restyResp, err := common.Client.R().
			SetHeaders(map[string]string{
				//"X-Tsign-Open-App-Id": oauth.AppId,
				//"X-Tsign-Open-Token":  oauth.Token,
				//"Content-Type":        oauth.ContentType,
			}).SetOutput(tempDir + details.TemplateName).Get(details.DownloadUrl)

		if restyResp.RawResponse.StatusCode != http.StatusOK {
			return nil, &common.RespObj{
				Code: restyResp.RawResponse.StatusCode,
				Msg:  restyResp.RawResponse.Status,
			}
		}
		if err != nil {
			return nil, &common.RespObj{
				Code: http.StatusInternalServerError,
				Msg:  fmt.Sprintf("failed to download file %q:%v", details.TemplateName, err),
			}
		}
		//
		//file, err := os.Open("./" + details.TemplateName)
		//if err != nil {
		//	return nil, &common.RespObj{
		//		Code: http.StatusInternalServerError,
		//		Msg:  fmt.Sprintf("failed to open file %q:%v", details.TemplateName, err),
		//	}
		//}

		//errObj = common.ErrObjFromEsignRequest(restyResp.RawResponse, err, &common.EsignError{Code: parsedResp.Code, Msg: parsedResp.Msg})
		//if errObj != nil {
		//	return
		//}

		//filename := details.TemplateName
		//file, err := os.CreateTemp("", filename)
		//if err != nil {
		//	return nil, &common.RespObj{
		//		Code: http.StatusInternalServerError,
		//		Msg:  fmt.Sprintf("failed to create temp file %q:%v", filename, err),
		//	}
		//}
		//
		//_, err = file.Write(restyResp.Body())
		//if err != nil {
		//	return nil, &common.RespObj{
		//		Code: http.StatusInternalServerError,
		//		Msg:  fmt.Sprintf("failed to write temp file %q:%v", filename, err),
		//	}
		//}
		//
		//files = append(files, file)

		filenames = append(filenames, details.TemplateName)
	}
	//defer func() {
	//	for _, file := range files {
	//		file.Close()
	//	}
	//}()

	// todo: delete temp files
	defer func() {

	}()

	// join file names open it for combination
	var joinName string
	var filepaths []string

	//for _, file := range files {
	//	joinName += strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + "."
	//}
	//joinName = strings.TrimSuffix(joinName, ".")

	for _, filename := range filenames {
		filepaths = append(filepaths, tempDir+filename)
		joinName += strings.TrimSuffix(filename, filepath.Ext(filename)) + "."
	}
	joinName = strings.TrimSuffix(joinName, ".")

	// merge file
	//var readSeekers []io.ReadSeeker
	//for _, f := range files {
	//	readSeekers = append(readSeekers, f)
	//}
	//
	//buf := &bytes.Buffer{}
	//if err := api.Merge(readSeekers, buf, nil); err != nil {
	//	return nil, &common.RespObj{
	//		Code: http.StatusInternalServerError,
	//		Msg:  fmt.Sprintf("failed to merge file %q:%v", joinName, err),
	//	}
	//}

	//var filepaths string
	//for _,

	if err := api.MergeCreateFile(filepaths, tempDir+joinName, nil); err != nil {
		return nil, &common.RespObj{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("failed to merge file %q:%v", joinName, err),
		}
	}

	joinFileBytes, err := ioutil.ReadFile(tempDir + joinName)
	if err != nil {
		return nil, &common.RespObj{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("failed to read joined file %q:%v", joinName, err),
		}
	}

	if err := os.RemoveAll(tempDir); err != nil {
		return nil, &common.RespObj{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("failed to remove temp merged files %q:%v", joinName, err),
		}
	}

	//return &mergeTemplateResult{joinName: joinName, fileBody: buf.Bytes()}, nil
	return &mergeTemplateResult{joinName: joinName, fileBody: joinFileBytes}, nil
}
