package openchain

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/ljg-cqu/core/errors"
	"github.com/ljg-cqu/core/utils"
	"strconv"
)

type CreateAccountReq struct {
	OrderId         string `json:"orderId"`
	BizId           string `json:"bizid"`
	Account         string `json:"account"`
	MyKmsKeyId      string `json:"mykmsKeyId"`
	NewAccountId    string `json:"newAccountId"`
	NewAccountKmsId string `json:"newAccountKmsId"`
	Method          Method `json:"method"`
	AccessId        string `json:"accessId"`
	Token           string `json:"token"`
	TenantId        string `json:"tenantid"`
	GasLimit        string `json:"gas"`
}

func (c *Client) CreateAccount(id string) (*NewAccount, errors.Error) {
	var acc = &NewAccount{
		Id:    id,
		KmsId: utils.Uuid(),
	}

	var req = &CreateAccountReq{
		OrderId:         utils.Uuid(),
		BizId:           c.BizId,
		Account:         c.Account,
		MyKmsKeyId:      c.KmsKeyId,
		NewAccountId:    acc.Id,
		NewAccountKmsId: acc.KmsId,
		Method:          MethodTenantCreateAccount,
		AccessId:        c.AccessId,
		Token:           c.CurrentToken,
		TenantId:        c.TenantId,
		GasLimit:        strconv.Itoa(int(c.GasLimit)),
	}

	var res Response
	resp, err := c.R().SetBody(req).SetResult(&res).Post(c.TransactUrl)
	if err != nil {
		return nil, errors.NewWithMsgf("failed to create account:%v", err).WithTag(errors.ErrTagHttpRequest)
	}

	_ = resp

	acc.PublicKey = string(res.Data)

	acc.CreatedAt = Timestamp(utils.NowTimestamp13())

	return acc, nil
}

type QueryAccountReq struct {
	BizId      string `json:"bizid"`
	Method     Method `json:"method"`
	RequestStr string `json:"requestStr"`
	AccessId   string `json:"accessId"`
	Token      string `json:"token"`
}

// ----------

type QueryAccountStr struct {
	QueryAccount string `json:"queryAccount"`
}

// TODO: insight of open chain error code

func (c *Client) QueryAccount(account string) (*Account, errors.Error) {
	acc, _ := json.Marshal(&QueryAccountStr{QueryAccount: account})
	var req = &QueryAccountReq{
		BizId:      c.BizId,
		Method:     MethodQueryAccount,
		RequestStr: string(acc),
		AccessId:   c.AccessId,
		Token:      c.CurrentToken,
	}

	var res Response
	resp, err := c.R().SetBody(req).SetResult(&res).Post(c.QueryUrl)
	if err != nil {
		if v, ok := err.(*resty.ResponseError); ok {
			_ = v
			_ = ok
		}
		return nil, errors.NewWithMsgf("failed to query account:%v", err)
	}
	_ = resp
	return res.ParseAccount() // TODO;
}
