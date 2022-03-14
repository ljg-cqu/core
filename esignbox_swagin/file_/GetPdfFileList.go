package file_

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/models/models"
	"github.com/long2ice/swagin/router"
	"github.com/wI2L/fizz/markdown"
	"net/http"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyingmd&namespace=opendoc%2Fsaas_api

type GetPdfFileDetailsListRequest struct {
	//FileId string `uri:"fileId" binding:"required" default:"ede1fa4504954c29ad210637c15f42cf" description:"文件ID"`
	DocType string `uri:"docType" default:"0-合同" description:"文档类型：0-合同, 1-协议, 2-订单. 若不指定，将返回所有模板列表"`
}

func (req *GetPdfFileDetailsListRequest) Handler(ctx *gin.Context) {
	queries := models.New(common.PgxPool)

	var contractFileIds []string
	var err error

	if req.DocType == "" {
		contractFileIds, err = queries.ListContractFileIds(context.Background())
		if err != nil {
			common.RespErrf(ctx, http.StatusInternalServerError, "failed to query contract file ids from db, error:%v", err)
			return
		}

		if len(contractFileIds) == 0 {
			common.RespErrf(ctx, http.StatusNotFound, "Not Found")
			common.RespSucc(ctx, []GetPdfFileDetailsResponseData{})
			return
		}
	} else {
		contractFileIds, err = queries.ListContractFileIdsByDocType(context.Background(), models.DocType(req.DocType))
		if err != nil {
			common.RespErrf(ctx, http.StatusInternalServerError, "failed to query contract file ids from db, error:%v", err)
			return
		}

		if len(contractFileIds) == 0 {
			common.RespErrf(ctx, http.StatusNotFound, "Not Found")
			common.RespSucc(ctx, []GetPdfFileDetailsResponseData{})
			return
		}
	}

	var fileDetailsList []GetPdfFileDetailsResponseData
	for _, contractFileId := range contractFileIds {
		if contractFileId == "" {
			continue
		}
		detail, errObj := GetPdfFileDetails(contractFileId)
		if errObj != nil {
			common.RespErrObj(ctx, errObj)
			return
		}
		fileDetailsList = append(fileDetailsList, *detail)
	}

	// todo: return account-id and name ...
	common.RespSucc(ctx, fileDetailsList)
}

var GetPdfFileDetailsListRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		builder.P("官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fyingmd&namespace=opendoc%2Fsaas_api", "查询PDF文件详情.")
	}

	r := router.New(
		&GetPdfFileDetailsListRequest{},
		router.Summary("获取PDF文件列表.docType文档类型：0-合同, 1-协议, 2-订单. 若不指定，将返回所有模板列表"),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: struct {
					Code int                             `json:"code" binding:"required" default:"0"`
					Msg  string                          `json:"msg" binding:"required" default:"ok"`
					Data []GetPdfFileDetailsResponseData `json:"data"`
				}{},
			},
		}),
	)

	return r
}
