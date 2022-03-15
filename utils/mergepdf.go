package utils

//
//import (
//	"bytes"
//	"fmt"
//	"github.com/ljg-cqu/core/esignbox_swagin/common"
//	"github.com/pdfcpu/pdfcpu/pkg/api"
//	"io"
//	"net/http"
//	"os"
//	"path/filepath"
//	"strings"
//)
//
//type mergeFileResult struct {
//	joinName string
//	fileBody []byte
//}
//
//// mergeFiles merge files that found from esign platform
//// todo: add pdf limit?
//func mergeFiles(fileIds ...string) (result *mergeFileResult, errObj *common.RespObj) {
//	// esign authentication
//	oauth, err := token.GetOauthInfo()
//	if err != nil {
//		return nil, &common.RespObj{
//			Code: http.StatusNetworkAuthenticationRequired,
//			Msg:  fmt.Sprintf("got an error when try to get authentication info:%v", err),
//		}
//	}
//
//	if len(fileIds) < 2 {
//		return nil, &common.RespObj{
//			Code: http.StatusInternalServerError,
//			Msg:  fmt.Sprintf(" not enough file number to merge"),
//		}
//	}
//
//	// get files from esign
//	var files []*os.File
//	for _, fileId := range fileIds {
//		if fileId == "" {
//			return nil, &common.RespObj{
//				Code: http.StatusInternalServerError,
//				Msg:  fmt.Sprintf("file id should not be empty string"),
//			}
//		}
//		parsedResp := MergeThenUploadFilesResponse{}
//		restyResp, err := common.Client.R().SetHeaders(map[string]string{
//			"X-Tsign-Open-App-Id": oauth.AppId,
//			"X-Tsign-Open-Token":  oauth.Token,
//			"Content-Type":        oauth.ContentType,
//		}).Get("/v1/files/" + fileId)
//
//		errObj = common.ErrObjFromEsignRequest(restyResp.RawResponse, err, &common.EsignError{Code: parsedResp.Code, Msg: parsedResp.Msg})
//		if errObj != nil {
//			return
//		}
//
//		filename := parsedResp.Data.TemplateName
//		file, err := os.CreateTemp("", filename)
//		if err != nil {
//			return nil, &common.RespObj{
//				Code: http.StatusInternalServerError,
//				Msg:  fmt.Sprintf("failed to create temp file %q:%v", filename, err),
//			}
//		}
//
//		_, err = file.Write(restyResp.Body())
//		if err != nil {
//			return nil, &common.RespObj{
//				Code: http.StatusInternalServerError,
//				Msg:  fmt.Sprintf("failed to write temp file %q:%v", filename, err),
//			}
//		}
//
//		files = append(files, file)
//	}
//	defer func() {
//		for _, file := range files {
//			file.Close()
//		}
//	}()
//
//	// join file names open it for combination
//	var joinName string
//	for _, file := range files {
//		joinName += strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + "."
//	}
//	joinName = strings.TrimSuffix(joinName, ".")
//
//	// merge file
//	var readSeekers []io.ReadSeeker
//	for _, f := range files {
//		readSeekers = append(readSeekers, f)
//	}
//
//	buf := &bytes.Buffer{}
//	if err := api.Merge(readSeekers, buf, nil); err != nil {
//		return nil, &common.RespObj{
//			Code: http.StatusInternalServerError,
//			Msg:  fmt.Sprintf("failed to merge file %q:%v", joinName, err),
//		}
//	}
//	return &mergeFileResult{joinName: joinName, fileBody: buf.Bytes()}, nil
//}
