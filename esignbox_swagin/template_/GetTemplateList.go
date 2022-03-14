package template_

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/models/models"
	"github.com/long2ice/swagin/router"
	"net/http"
)

const (
//EsignSandBoxGetTemplInfoPath = "/v1/flow-templates/basic-info?pageNum=XX&pageSize=XX"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fuseb10&namespace=opendoc%2Fsaas_api

type GetTemplateListRequest struct {
	//PageNum  int `query:"pageNum" description:"查询页码，不传默认第1页，起始值：1"`
	//PageSize int `query:"pageSize" description:"分页大小，不传默认10条，最大值：100"`
	DocType string `uri:"docType" default:"0-合同" description:"文档类型：0-合同, 1-协议, 2-订单. 若不指定，将返回所有模板列表"`
}

//type GetTemplListResponse struct {
//	Code int                        `json:"code" required:"true" description:"业务码，0表示成功"`
//	Msg  string                     `json:"message,omitempty" description:"信息"`
//	Data []GetTemplInfoResponseData `json:"data,omitempty" type:"object" description:"业务信息"`
//}
//
//type GetTemplInfoResponseData struct {
//	FlowTemplateBasicInfos FlowTemplateBasicInfos `json:"flowTemplateBasicInfos" type:"array" required:"false" description:"流程模板基本信息列表"`
//	Total                  int                    `json:"total" required:"false" description:"数据总条数"`
//}
//
//type FlowTemplateBasicInfos struct {
//	FlowTemplateId   string       `json:"flowTemplateId" required:"false" description:"流程模板编号"`
//	FlowTemplateName string       `json:"flowTemplateName" required:"false" description:"流程模板名称"`
//	DocTemplates     DocTemplates `json:"docTemplates" type:"array" required:"false" description:"文件模板信息"`
//}
//
//type DocTemplates struct {
//	DocTemplateId   string `json:"docTemplateId" required:"false" description:"文件模板id注意：该参数为接口中使用的模板id"`
//	DocTemplateName string `json:"docTemplateName" required:"false" description:"文件模板名称"`
//}

//type ContractTemplate struct {
//	TemplID          []string `json:"templID" description:"模板ID"`
//	TemplName        []string `json:"templName" description:"模板名称"`
//	TemplType        int32    `json:"templType" description:"模板类型，固定值 3"`
//	FileType         []string `json:"fileType" description:"文件类型"`
//	CreateTime       int64    `json:"createTime" description:"创建时间，Unix时间戳（毫秒级）"`
//	UpdateTime       int64    `json:"updateTime" description:"更新时间，Unix时间戳（毫秒级）"`
//	FileSize         int64    `json:"fileSize" description:"模板文件大小, 单位byte"`
//	FileBytes        []byte   `json:"fileBytes" description:"模板二进制文件内容"`
//	StructComponents []byte   `json:"structComponents" description:"文件模板中的填写控件列表"`
//}

type GetTemplateListResponseData_ struct {
	TemplateId   string `json:"templateId" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板ID"`
	TemplateName string `json:"templateName" binding:"required" default:"商贷通用收入证明.pdf" description:"模板名称"`
	DownloadUrl  string `json:"DownloadUrl" binding:"required" description:"模板文件下载链接，有效期60分钟。"`
	CreateTime   int64  `json:"createTime" binding:"required" description:"创建时间，Unix时间戳（毫秒级）"`
	CreatorID    string `json:"creatorID" binding:"required" description:"创建者账户ID"`
	CreatorName  string `json:"creatorName" binding:"required" description:"创建者姓名"`
	DocType      string `json:"docType" default:"0-合同" description:"文档类型：0-合同, 1-协议, 2-订单."`
}

// Query template list from database

func (req *GetTemplateListRequest) Handler(ctx *gin.Context) {
	queries := models.New(common.PgxPool)
	var contractTemplIds []string
	var err error
	// todo: return type from db
	if req.DocType == "" {
		contractTemplIds, err = queries.ListContractTemplateIds(context.Background())
		if err != nil {
			common.RespErrf(ctx, http.StatusInternalServerError, "failed to query contract template ids from db, error:%v", err)
			return
		}

		if len(contractTemplIds) == 0 {
			common.RespSucc(ctx, []GetTemplateListResponseData_{})
			return
		}
	} else {
		contractTemplIds, err = queries.ListContractTemplateIdsByDocType(context.Background(), models.DocType(req.DocType))
		if err != nil {
			common.RespErrf(ctx, http.StatusInternalServerError, "failed to query contract template ids from db, error:%v", err)
			return
		}

		if len(contractTemplIds) == 0 {
			common.RespSucc(ctx, []GetTemplateListResponseData_{})
			return
		}
	}

	var templDetailsList []GetTemplateListResponseData_
	for _, contractTemplId := range contractTemplIds {
		if contractTemplId == "" {
			continue
		}
		detail, errObj := GetTemplDetails(contractTemplId)
		if errObj != nil {
			common.RespErrObj(ctx, errObj)
			return
		}
		templDetailsList = append(templDetailsList, GetTemplateListResponseData_{
			TemplateId:   detail.TemplateId,
			TemplateName: detail.TemplateName,
			DownloadUrl:  detail.DownloadUrl,
			CreateTime:   detail.CreateTime,
			CreatorID:    detail.CreatorID,
			CreatorName:  detail.CreatorName,
			DocType:      req.DocType,
		})
	}

	// todo: return account-id and name ...
	common.RespSucc(ctx, templDetailsList)
}

var GetTemplListRequestH = func() *router.Router {
	r := router.New(
		&GetTemplateListRequest{},
		router.Summary("获取模板文件列表. docType文档类型：0-合同, 1-协议, 2-订单. 若不指定，将返回所有模板列表"),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: struct {
					Code int                            `json:"code" binding:"required" default:"0"`
					Msg  string                         `json:"msg" binding:"required" default:"ok"`
					Data []GetTemplateListResponseData_ `json:"data"`
				}{},
			},
		}),
	)

	return r
}
