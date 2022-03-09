package openchain

import (
	"github.com/go-resty/resty/v2"
	"github.com/ljg-cqu/core/_errors"
)

type QueryLastBlockReq struct {
	BizId    string `json:"bizid"`
	Method   Method `json:"method"`
	AccessId string `json:"accessId"`
	Token    string `json:"token"`
}

type QueryLastBlockRes struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Data    string `json:"data"`
}

func (c *Client) QueryLastBlock() (string, error) {
	var req = &QueryLastBlockReq{
		BizId:    c.BizId,
		Method:   MethodQueryLastBlock,
		AccessId: c.AccessId,
		Token:    c.CurrentToken,
	}

	//c.SetDoNotParseResponse(true)

	var res QueryLastBlockRes
	resp, err := c.R().SetBody(req).SetResult(&res).Post(c.QueryUrl)
	if err != nil {
		if v, ok := err.(*resty.ResponseError); ok {
			_ = v
			_ = ok
		}
		return "", _errors.NewWithMsgf("failed to query account:%v", err)
	}
	_ = resp
	return res.Data, nil
}
