package template

import (
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
	"github.com/go-resty/resty/v2"
	"github.com/ljg-cqu/core/esignbox/common"
	"github.com/ljg-cqu/core/esignbox/token"
	"net/http"
)

const (
	EsignSandBoxFillTemplateContentUrl = "/v1/files/createByTemplate"
)

// https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fsiipw3&namespace=opendoc%2Fsaas_api

type FillTemplateContentRequest struct {
	Body FillTemplateContentRequestBody `json:"structComponents" required:"true" type:"array" doc:"添加填写控件请求参数"`
}

type FillTemplateContentRequestBody struct {
	FileName         string            `json:"name" example:"商贷通用收入证明.pdf" required:"true" doc:"文件名称，文件名称（必须带上文件扩展名，不然会导致后续发起流程校验过不去 示例：合同.pdf.该字段的文件后缀名称和真实的文件后缀需要一致"`
	TemplateId       string            `json:"templateId" required:"true" example:"c4d4fe1b48184ba28982f68bf2c7bf25" path:"templateId" doc:"模板id"`
	StrictCheck      bool              `json:"strictCheck" doc:"开启simpleFormFields为空校验，默认false，传 false：允许simpleFormFields为空，此时模板中所有待填写字段均为空值；传 true： 当模板中存在必填字段时，不允许simpleFormFields为空，否则会报错"`
	SimpleFormFields map[string]string `json:"simpleFormFields" required:"true" type:"object" doc:"输入项填充内容，key:value 传入；可使用输入项组件id+填充内容，也可使用输入项组件key+填充内容方式填充.注意：E签宝官网获取的模板id，在通过模板创建文件的时候只支持输入项组件id+填充内容.如何进行图片控件填充:https://open.esign.cn/doc/detail?id=opendoc%2Fsaas_api%2Fzwlv9n&namespace=opendoc%2Fsaas_api"`
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

type FillTemplateContentResponse struct {
	Code int                             `json:"code" required:"true" doc:"业务码，0表示成功"`
	Msg  string                          `json:"message" doc:"信息"`
	Data FillTemplateContentResponseData `json:"data" type:"object" doc:"业务信息"`
}

type FillTemplateContentResponseData struct {
	FileId      string `json:"fileId" doc:"文件ID"`
	FileName    string `json:"fileName" doc:"文件名称"`
	DownloadUrl string `json:"downUrl" doc:"模板文件下载链接，有效期60分钟" example:"https://esignoss.esign.cn/1111564182/f33c1c35-e47e-451d-b9b9-0855d682ea79/%E5%95%86%E8%B4%B7%E9%80%9A%E7%94%A8%E6%94%B6%E5%85%A5%E8%AF%81%E6%98%8E.pdf?Expires=1646233196&OSSAccessKeyId=LTAI4G23YViiKnxTC28ygQzF&Signature=fpguS0x08KbyJ3xOfPaujARK010%3D"`
}

/*
{
   "code": 0,
   "message": "成功",
   "data": {
      "fileId": "ede1fa4504954c29ad210637c15f42cf",
      "fileName": "商贷通用收入证明.pdf",
      "downloadUrl": "https://esignoss.esign.cn/1111564182/f33c1c35-e47e-451d-b9b9-0855d682ea79/%E5%95%86%E8%B4%B7%E9%80%9A%E7%94%A8%E6%94%B6%E5%85%A5%E8%AF%81%E6%98%8E.pdf?Expires=1646233196&OSSAccessKeyId=LTAI4G23YViiKnxTC28ygQzF&Signature=fpguS0x08KbyJ3xOfPaujARK010%3D"
   }
}

{
   "code": 0,
   "message": "成功",
   "data": {
      "fileId": "3d35e0bf3ce94a48a28dec72d9d7487e",
      "fileName": "商贷通用收入证明.pdf",
      "downloadUrl": "https://esignoss.esign.cn/1111564182/66023c56-d607-4e1c-b640-7bfcbac8e125/%E5%95%86%E8%B4%B7%E9%80%9A%E7%94%A8%E6%94%B6%E5%85%A5%E8%AF%81%E6%98%8E.pdf?Expires=1646233451&OSSAccessKeyId=LTAI4G23YViiKnxTC28ygQzF&Signature=5Em9r04LkArUfK8nzymo71dmpaA%3D"
   }
}


*/

func FillTemplateContent(client *resty.Client) func(ctx huma.Context, req FillTemplateContentRequest) {
	return func(ctx huma.Context, req FillTemplateContentRequest) {
		parsedResp := FillTemplateContentResponse{}
		oauth, err := token.GetOauthInfo(client)
		if err != nil {
			ctx.WriteError(http.StatusBadRequest, "got an error when try to get authentication info", err)
			return
		}

		restyResp, err := client.R().SetHeaders(map[string]string{
			"X-Tsign-Open-App-Id": oauth.AppId,
			"X-Tsign-Open-Token":  oauth.Token,
			"Content-Type":        oauth.ContentType,
		}).SetBody(&req.Body).
			SetResult(&parsedResp.Data).Post(EsignSandBoxFillTemplateContentUrl)

		if common.WriteError(ctx, err, restyResp.RawResponse.StatusCode, restyResp.RawResponse.Status, parsedResp.Code, parsedResp.Msg) {
			return
		}

		ctx.WriteModel(http.StatusOK, parsedResp.Data)
	}
}

func RunFillTemplateContent(r *huma.Resource, client *resty.Client) {
	r.Post("FillTemplateContent", "填充内容生成PDF",
		responses.BadRequest(),
		responses.OK().Model(FillTemplateContentResponseData{}),
	).Run(FillTemplateContent(client))
}
