package template

//import (
//	"context"
//	"github.com/gin-gonic/gin"
//	"github.com/ljg-cqu/core/esignbox_swagin/common"
//	"github.com/ljg-cqu/core/esignbox_swagin/template/models/models"
//	"github.com/long2ice/swagin/router"
//)
//
//const (
//	EsignSandBoxGetTemplInfoPath = "/v1/flow-templates/basic-info?pageNum=XX&pageSize=XX"
//)
//
//// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fuseb10&namespace=opendoc%2Fsaas_api
//
//type GetTemplListRequest struct{}
//
////type GetTemplListResponse struct {
////	Code int                        `json:"code" required:"true" description:"业务码，0表示成功"`
////	Msg  string                     `json:"message,omitempty" description:"信息"`
////	Data []GetTemplInfoResponseData `json:"data,omitempty" type:"object" description:"业务信息"`
////}
////
////type GetTemplInfoResponseData struct {
////	FlowTemplateBasicInfos FlowTemplateBasicInfos `json:"flowTemplateBasicInfos" type:"array" required:"false" description:"流程模板基本信息列表"`
////	Total                  int                    `json:"total" required:"false" description:"数据总条数"`
////}
////
////type FlowTemplateBasicInfos struct {
////	FlowTemplateId   string       `json:"flowTemplateId" required:"false" description:"流程模板编号"`
////	FlowTemplateName string       `json:"flowTemplateName" required:"false" description:"流程模板名称"`
////	DocTemplates     DocTemplates `json:"docTemplates" type:"array" required:"false" description:"文件模板信息"`
////}
////
////type DocTemplates struct {
////	DocTemplateId   string `json:"docTemplateId" required:"false" description:"文件模板id注意：该参数为接口中使用的模板id"`
////	DocTemplateName string `json:"docTemplateName" required:"false" description:"文件模板名称"`
////}
//
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
//
//// Query template list from database
//
//func (req *GetTemplListRequest) Handler(ctx *gin.Context) {
//	queries := models.New(common.PgxPool)
//	contractTempls, err := queries.ListTemplates(context.Background())
//	if err != nil {
//		common.WriteErrorF(ctx, 500, "failed to query contract templates from db, error:%v", err)
//		return
//	}
//
//	if len(contractTempls) == 0 {
//		common.WriteError(ctx, 400, "Not Found")
//		return
//	}
//
//	var _contractTempls []ContractTemplate
//	for _, e := range contractTempls {
//		var t ContractTemplate
//		t.TemplID = e.TemplID
//		t.TemplName = e.TemplName
//		t.CreateTime = e.CreateTime
//		t.UpdateTime = e.UpdateTime
//		t.FileSize = e.FileSize
//		t.FileBytes = e.FileBytes
//
//		ff, err := e.FormFields.MarshalJSON()
//		if err != nil {
//			common.WriteErrorF(ctx, 500, "failed to marshal json, error:%v", err)
//			return
//		}
//		t.FormFields = ff
//
//		_contractTempls = append(_contractTempls, t)
//	}
//
//	common.WriteOK(ctx, _contractTempls)
//}
//
//var GetTemplListRequestH = func() *router.Router {
//	r := router.New(
//		&GetTemplListRequest{},
//		router.Summary("获取模板列表"),
//		//router.Security(&security.Basic{}),
//		router.Responses(router.Response{
//			"200": router.ResponseItem{
//				Model: []ContractTemplate{},
//			},
//			"400": router.ResponseItem{
//				Model: common.ErrorResp{},
//			},
//			"500": router.ResponseItem{
//				Model: common.ErrorResp{},
//			},
//		}),
//	)
//
//	return r
//}
