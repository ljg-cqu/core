package openchain

import (
	"encoding/hex"
	"github.com/ljg-cqu/core/_errors"
	"github.com/ljg-cqu/core/utils"
	"time"
)

type HandshakeRequest struct {
	AccessId string `json:"accessId"`
	Time     string `json:"time"`
	Secret   string `json:"secret"` // signature
}

type HandshakeResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Token   string `json:"data"`
}

func (c *Client) HandshakeIfNeeded() _errors.Error {
	if c.CurrentToken != "" {
		if time.Now().Add(time.Second * 10).Before(c.TokenExpireAt) {
			return nil
		}
	}

	var req = &HandshakeRequest{
		AccessId: c.AccessId,
		Time:     utils.NowTimestamp13(),
	}

	signature, err := utils.SignRS256(c.AccessKey, []byte(req.AccessId+req.Time))
	if err != nil {
		return err.WithMsgf("failed to handshake:%v", err).WithTag(_errors.ErrTagHttpHandshak)
	}

	req.Secret = hex.EncodeToString(signature)

	var res HandshakeResponse
	_, err_ := c.R().
		SetBody(req).
		SetResult(&res).
		Post(c.HandshakeUrl)

	if err_ != nil {
		return _errors.NewWithMsgf("failed to handshake:%v", err).WithTag(_errors.ErrTagHttpHandshak)
	}

	c.CurrentToken = res.Token
	c.TokenExpireAt = time.Now().Add(c.TokenTimeoutInMinutes_)

	c.HandshakeNum++

	return nil
}
