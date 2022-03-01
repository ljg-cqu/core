package main

var MyOauthInfo = Oauth{
	AppId:       "7438902503",
	Token:       "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJnSWQiOiIyNGM4YTFhOWU3YWM0MTAwOTcxZTJhNTcyMmZkZGE0NCIsImFwcElkIjoiNzQzODkwMjUwMyIsIm9JZCI6IjY1MmNlOWIxNzNhMDQ5M2FhNzE3MzM4OGVhYjg2NzNiIiwidGltZXN0YW1wIjoxNjQ2MDk5NjYxMzIyfQ.7K6B_nw8uua0GUUxPlQtdqQfDRzYGOevkYDN9t3hykA",
	ContentType: "application/json; charset=UTF-8",
}

type Oauth struct {
	AppId       string `json:"appId" required:"true" header:"X-Tsign-Open-App-Id" default:"7438902503" example:"7438902503" doc:"应用id，需在e签宝开放平台创建"`
	Token       string `json:"token" required:"true" header:"X-Tsign-Open-App-Id" default:"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJnSWQiOiIyNGM4YTFhOWU3YWM0MTAwOTcxZTJhNTcyMmZkZGE0NCIsImFwcElkIjoiNzQzODkwMjUwMyIsIm9JZCI6IjY1MmNlOWIxNzNhMDQ5M2FhNzE3MzM4OGVhYjg2NzNiIiwidGltZXN0YW1wIjoxNjQ2MDk5NjYxMzIyfQ.7K6B_nw8uua0GUUxPlQtdqQfDRzYGOevkYDN9t3hykA" example:"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJnSWQiOiIyNGM4YTFhOWU3YWM0MTAwOTcxZTJhNTcyMmZkZGE0NCIsImFwcElkIjoiNzQzODkwMjUwMyIsIm9JZCI6IjY1MmNlOWIxNzNhMDQ5M2FhNzE3MzM4OGVhYjg2NzNiIiwidGltZXN0YW1wIjoxNjQ2MDk5NjYxMzIyfQ.7K6B_nw8uua0GUUxPlQtdqQfDRzYGOevkYDN9t3hykA" doc:"通过获取鉴权Token接口返回。授权码注意：120分钟失效，请在expiresIn参数的有效截止时间失效前重新获取token，建议提前5分钟重新获取。"`
	ContentType string `json:"contentType" required:"true" header:"Content-Type" default:"application/json; charset=UTF-8" example:"application/json; charset=UTF-8"`
}

type GetRequest struct {
	AppId     string `json:"appId" required:"true" query:"appId" default:"7438902503" example:"7438902503" doc:"应用id，需在e签宝开放平台创建"`
	Secret    string `json:"secret" required:"true" query:"secret" default:"7438902503" example:"018df25ada48ea6b0204a3b03d92ab95" doc:"应用密钥，不可泄露"`
	GrantType string `json:"grantType" required:"true" query:"grantType" default:"client_credentials" enum:"client_credentials" doc:"授权类型，固定值: client_credentials"`
}

type GetResponse struct {
	Code int             `json:"code" required:"true" doc:"业务码，0表示成功"`
	Msg  string          `json:"message,omitempty" doc:"信息"`
	Data GetResponseData `json:"data,omitempty" doc:"业务信息"`
}

type GetResponseData struct {
	Token        string `json:"token" doc:"授权码注意：120分钟失效，请在expiresIn参数的有效截止时间失效前重新获取token，建议提前5分钟重新获取。"`
	ExpiresIn    string `json:"expiresIn" doc:" 有效截止时间（毫秒）"`
	RefreshToken string `json:"refreshToken" doc:"刷新授权码"`
}
