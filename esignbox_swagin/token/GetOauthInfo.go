package token

import (
	"github.com/ljg-cqu/core/esignbox_swagin/common"
	"github.com/spf13/cast"
	"time"
)

const (
	AppId       = "7438902503"                       // 应用id，需在e签宝开放平台创建
	Secret      = "018df25ada48ea6b0204a3b03d92ab95" // 应用密钥，不可泄露
	GrantType   = "client_credentials"               // 授权类型，固定值
	ContentType = "application/json; charset=UTF-8"
)

const (
	EsignSandBoxGetTokenPath = "/v1/oauth2/access_token"
)

var oauthInfo *OauthInfo

type OauthInfo struct {
	Token        string
	ExpiresAt    time.Time
	RefreshToken string

	AppId       string
	GrantType   string
	ContentType string
}

func GetOauthInfo() (*OauthInfo, error) {
	if oauthInfo == nil {
		oauthInfo = &OauthInfo{}
		oauthInfo.AppId = AppId
		oauthInfo.GrantType = GrantType
		oauthInfo.ContentType = ContentType

		res, err := _getToken()
		if err != nil {
			return nil, err
		}

		oauthInfo.Token = res.Data.Token
		oauthInfo.RefreshToken = res.Data.RefreshToken
		expireIn := cast.ToInt(res.Data.ExpiresIn)
		oauthInfo.ExpiresAt = time.Now().Add(time.Millisecond * time.Duration(expireIn))

		return oauthInfo, nil
	}

	// TODO: consider using refresh token
	if time.Now().Add(time.Second * 60).After(oauthInfo.ExpiresAt) {
		res, err := _getToken()
		if err != nil {
			return nil, err
		}

		oauthInfo.Token = res.Data.Token
		oauthInfo.RefreshToken = res.Data.RefreshToken
		expireIn := cast.ToInt(res.Data.ExpiresIn)
		oauthInfo.ExpiresAt = time.Now().Add(time.Millisecond * time.Duration(expireIn))

		return oauthInfo, nil
	}

	return oauthInfo, nil
}

func _getToken() (*GetTokenResponse, error) {
	res := GetTokenResponse{}

	_, err := common.Client.R().SetQueryParams(map[string]string{
		"appId":     AppId,
		"secret":    Secret,
		"grantType": GrantType,
	}).SetResult(&res).Get(EsignSandBoxGetTokenPath)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
