package template_

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgtype"
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/ljg-cqu/core/esignbox_swagin/models/models"
	"github.com/ljg-cqu/core/esignbox_swagin/token"
	"github.com/long2ice/swagin/router"
	"github.com/wI2L/fizz/markdown"
	"net/http"
)

const (
	EsignSandBoxFillTemplateContentUrl = "/v1/files/createByTemplate"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fsiipw3&namespace=opendoc%2Fsaas_api

//type FillTemplateContentRequest struct {
//	Body FillTemplateContentRequestBody `json:"structComponents" required:"true" type:"array" doc:"添加填写控件请求参数"`
//}

//type FillTemplateContentRequest struct {
//	TemplateId       string       `form:"templateId" binding:"required" json:"templateId" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板id"`
//	TemplateName         string       `form:"name" binding:"required" json:"name" description:"文件名称，文件名称（必须带上文件扩展名，不然会导致后续发起流程校验过不去.该字段的文件后缀名称和真实的文件后缀需要一致" default:"商贷通用收入证明.pdf"`
//	StrictCheck      bool         `form:"strictCheck" json:"strictCheck" description:"开启simpleFormFields为空校验，默认false，传 false：允许simpleFormFields为空，此时模板中所有待填写字段均为空值；传 true： 当模板中存在必填字段时，不允许simpleFormFields为空，否则会报错" default:"false"`
//	SimpleFormFields []FieldValue `form:"simpleFormFields" binding:"required" json:"simpleFormFields" description:"输入项填充内容，key:value 传入；可使用输入项组件id+填充内容，也可使用输入项组件key+填充内容方式填充.注意：E签宝官网获取的模板id，在通过模板创建文件的时候只支持输入项组件id+填充内容.如何进行图片控件填充:https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fzwlv9n&namespace=opendoc%2Fsaas_api"` // todo:check object type
//}

type FillTemplateContentRequest struct {
	DocType          string       `uri:"docType" default:"0-合同" description:"文档类型：0-合同, 1-协议, 2-订单."`
	TemplateId       string       `form:"templateId" json:"templateId" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"必须。模板id"`
	FileName         string       `form:"name" json:"name" description:"必须。文件名称，文件名称（必须带上文件扩展名，不然会导致后续发起流程校验过不去.该字段的文件后缀名称和真实的文件后缀需要一致" default:"商贷通用收入证明.pdf"`
	StrictCheck      bool         `form:"strictCheck" json:"strictCheck" description:"必须。开启simpleFormFields为空校验，默认false，传 false：允许simpleFormFields为空，此时模板中所有待填写字段均为空值；传 true： 当模板中存在必填字段时，不允许simpleFormFields为空，否则会报错" default:"false"`
	SimpleFormFields []FieldValue `form:"simpleFormFields" json:"simpleFormFields" description:"必须。输入项填充内容，key:value 传入；可使用输入项组件id+填充内容，也可使用输入项组件key+填充内容方式填充.注意：E签宝官网获取的模板id，在通过模板创建文件的时候只支持输入项组件id+填充内容.如何进行图片控件填充:https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fzwlv9n&namespace=opendoc%2Fsaas_api"` // todo:check object type
}

//type FillTemplateContentRequest struct {
//	TemplateId       string       `binding:"required" json:"templateId" default:"c4d4fe1b48184ba28982f68bf2c7bf25" description:"模板id"`
//	TemplateName         string       `binding:"required" json:"name" description:"文件名称，文件名称（必须带上文件扩展名，不然会导致后续发起流程校验过不去.该字段的文件后缀名称和真实的文件后缀需要一致" default:"商贷通用收入证明.pdf"`
//	StrictCheck      bool         `json:"strictCheck" description:"开启simpleFormFields为空校验，默认false，传 false：允许simpleFormFields为空，此时模板中所有待填写字段均为空值；传 true： 当模板中存在必填字段时，不允许simpleFormFields为空，否则会报错" default:"false"`
//	SimpleFormFields []FieldValue `binding:"required" json:"simpleFormFields" description:"输入项填充内容，key:value 传入；可使用输入项组件id+填充内容，也可使用输入项组件key+填充内容方式填充.注意：E签宝官网获取的模板id，在通过模板创建文件的时候只支持输入项组件id+填充内容.如何进行图片控件填充:https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fzwlv9n&namespace=opendoc%2Fsaas_api"` // todo:check object type
//}

type FieldValue struct {
	Field string `form:"field" json:"field" binding:"required" default:"name"`
	Value string `form:"value" json:"value" binding:"required" default:"zealy"`
}

type FillTemplateContentRequest_ struct {
	DocType          string            `uri:"docType"`
	TemplateId       string            `json:"templateId"`
	FileName         string            `json:"name"`
	StrictCheck      bool              `json:"strictCheck"`
	SimpleFormFields map[string]string `json:"simpleFormFields"`
}

/*
Template ID:c4d4fe1b48184ba28982f68bf2c7bf25
{
"staff_name": "罗继高",
"id_num": "522731198905083798",
"from_year": "2021",
"from_month": "11",
"industry": "软件开发",
"job": "区块链工程师",
"fixed_salary": "1000",
"average_salary": "1200",
"total_salary": "12000",
"company_nature": "民营企业",
"company_address": "北京市朝阳区酒仙桥东路10号",
"company_leader": "邓琦",
"company_tel": "10086",
"prove_year": "2022",
"prove_month": "03",
"prove_date": "02"
}
*/

/*
{
  "name": "合同模板-商贷通用收入证明-0310.pdf",
  "simpleFormFields": [
  	{"field": "staff_name","value": "罗继高"},
	{"field": "id_num","value": "522731198905083798"},
	{"field": "from_year","value": "2021"},
	{"field": "from_month","value": "11"},
	{"field": "industry","value": "软件开发"},
	{"field": "job","value": "区块链工程师"},
	{"field": "fixed_salary","value": "1000"},
	{"field": "average_salary","value": "1200"},
	{"field": "total_salary","value": "12000"},
	{"field": "company_nature","value": "民营企业"},
	{"field": "company_address","value": "北京市朝阳区酒仙桥东路10号"},
	{"field": "company_leader","value": "邓琦"},
	{"field": "company_tel","value": "10086"},
	{"field": "prove_year","value": "2021"},
	{"field": "prove_month","value": "03"},
	{"field": "prove_date","value": "02"}
  ],
  "strictCheck": false,
  "templateId": "b34875c5c5f24757a39375954fae1dcf"
}
*/

type FillTemplateContentResponse struct {
	Code int                             `json:"code" required:"true" description:"业务码，0表示成功"`
	Msg  string                          `json:"message" description:"信息"`
	Data FillTemplateContentResponseData `json:"data" type:"object" description:"业务信息"`
}

type FillTemplateContentResponseData struct {
	FileId      string `json:"fileId" description:"文件ID"`
	FileName    string `json:"fileName" description:"文件名称"`
	DownloadUrl string `json:"downloadUrl" description:"模板文件下载链接，有效期60分钟"`
}

func (req *FillTemplateContentRequest) Handler(ctx *gin.Context) {

	var fieldValues = make(map[string]string)
	for _, fieldValue := range req.SimpleFormFields {
		fieldValues[fieldValue.Field] = fieldValue.Value
	}

	_req := FillTemplateContentRequest_{}
	_req.DocType = req.DocType
	_req.TemplateId = req.TemplateId
	_req.StrictCheck = req.StrictCheck
	_req.FileName = req.FileName
	_req.SimpleFormFields = fieldValues

	result, errObj := FillTemplateContent(&_req)
	if errObj != nil {
		common.RespErrObj(ctx, errObj)
		return
	}

	common.RespSucc(ctx, result)
}

func FillTemplateContent(params *FillTemplateContentRequest_) (data *FillTemplateContentResponseData, errObj *common.RespObj) {
	oauth, err := token.GetOauthInfo()
	if err != nil {
		return nil, &common.RespObj{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("got an error for esign authentication:%v", err),
		}
	}

	parsedResp := FillTemplateContentResponse{}
	restyResp, err := common.Client.R().SetHeaders(map[string]string{
		"X-Tsign-Open-App-Id": oauth.AppId,
		"X-Tsign-Open-Token":  oauth.Token,
		"Content-Type":        oauth.ContentType,
	}).SetBody(&params).
		SetResult(&parsedResp).Post(EsignSandBoxFillTemplateContentUrl)

	if errObj := common.ErrObjFromEsignRequest(restyResp.RawResponse, err, &common.EsignError{
		Code: parsedResp.Code,
		Msg:  parsedResp.Msg,
	}); errObj != nil {
		return nil, errObj
	}

	bytes, err := json.Marshal(params.SimpleFormFields)
	if err != nil {
		return nil, &common.RespObj{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("failed to marshal simple form fields for %q:%v", params.FileName, err),
		}
	}

	// todo: norm doc type...
	_, err = models.New(common.PgxPool).CreateFile(context.Background(), &models.CreateFileParams{
		FileID:           parsedResp.Data.FileId,
		FileName:         parsedResp.Data.FileName,
		DocType:          models.DocType(params.DocType),
		TemplateID:       params.TemplateId,
		ParentFileIds:    []string{},      // todo
		CreatorID:        gofakeit.UUID(), // todo: use true account
		SimpleFormFields: pgtype.JSONB{Bytes: bytes, Status: pgtype.Present},
		FileSize:         0,        // todo
		FileBody:         []byte{}, // todo
	})
	if err != nil {
		return nil, &common.RespObj{
			Code: http.StatusInternalServerError,
			Msg:  fmt.Sprintf("failed to create contract file %q in database for template fill:%v", params.FileName, err),
		}
	}

	return &parsedResp.Data, nil
}

var FillTemplateContentRequestH = func() *router.Router {
	var apiDesc = func() string {
		builder := markdown.Builder{}
		//		builder.P("请求参数说明")
		//		builder.Table(
		//			[][]string{
		//				[]string{"参数名称", "类型", "必选", "参数说明", "举例"},
		//				[]string{"fileName", "string", "是", "文件名称，必须带扩展名:.pdf", "商贷通用收入证明.pdf"},
		//				[]string{"templateId", "string", "是", "模板id", "1Qinq+/TLV3UZGbzifvmjw=="},
		//				[]string{"strictCheck", "bool", "否", "开启simpleFormFields为空校验，默认false", "false"},
		//				[]string{"simpleFormFields", "object", "是", "输入项填充内容,key:value 传入.", "\"simpleFormFields\":{\"name\":\"测试甲方\",\"yifang\":\"测试乙方\"},"},
		//			}, []markdown.TableAlignment{
		//				markdown.AlignLeft,
		//				markdown.AlignCenter,
		//				markdown.AlignCenter,
		//				markdown.AlignLeft,
		//				markdown.AlignLeft,
		//			},
		//		)
		//		builder.P("请求示例")
		//
		//		builder.P(`
		//{
		//    "name":"商贷通用收入证明.pdf",
		//    "simpleFormFields":{
		//        "name":"测试甲方",
		//        "yifang":"测试乙方"
		//    },
		//    "templateId":"bd4f8b3fc02047a9be661d164eceb288"
		//}
		//`)

		builder.P("官方文档：")
		return builder.String() + builder.Link("https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fsiipw3&namespace=opendoc%2Fsaas_api", "填充内容生成PDF")
	}

	//req := FillTemplateContentRequest{}
	//req.SimpleFormFields = make(map[string]string, 0)

	r := router.New(
		//&req,
		&FillTemplateContentRequest{},
		router.Summary("填充内容生成PDF。docType文档类型：0-合同, 1-协议, 2-订单."),
		router.Description(apiDesc()),
		//router.Security(&security.Basic{}),
		router.Responses(router.Response{
			"200": router.ResponseItem{
				Model: struct {
					Code int                             `json:"code" binding:"required" default:"0"`
					Msg  string                          `json:"msg" binding:"required" default:"ok"`
					Data FillTemplateContentResponseData `json:"data"`
				}{},
			},
		}),
	)

	return r
}
