package common

import (
	"github.com/danielgtaylor/huma"
	"github.com/pkg/errors"
	"net/http"
)

func WriteError(ctx huma.Context, respErr error, respStatusCode int, respStatus string, businessCode int, businessMsg string) bool {
	if respErr != nil {
		ctx.WriteError(http.StatusBadRequest, "got an error when request esign", respErr)
		return true
	}

	if respStatusCode != http.StatusOK {
		err := errors.Errorf("response status code:%v, response status:%v", respStatusCode, respStatus)
		ctx.WriteError(http.StatusBadRequest, "got an error when request esign", err)
		return true
	}

	if businessCode != 0 {
		err := errors.Errorf("esgin business code:%v, esign business message:%v", businessCode, businessMsg)
		ctx.WriteError(http.StatusBadRequest, "got an error when request esign", err)
		return true
	}

	return false
}
