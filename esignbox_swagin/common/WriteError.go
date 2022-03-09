package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type ErrorResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func WriteErrorE(ctx *gin.Context, respErr error, respStatusCode int, respStatus string, businessCode int, businessMsg string) bool {
	if respErr != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResp{Code: http.StatusBadRequest, Msg: "got an error when request esign:" + respErr.Error()})
		return true
	}

	if respStatusCode != http.StatusOK {
		err := errors.Errorf("response status code:%v, response status:%v", respStatusCode, respStatus)
		ctx.JSON(http.StatusBadRequest, ErrorResp{Code: http.StatusBadRequest, Msg: "got an error when request esign:" + err.Error()})
		return true
	}

	if businessCode != 0 {
		err := errors.Errorf("esgin business code:%v, esign business message:%v", businessCode, businessMsg)
		ctx.JSON(http.StatusBadRequest, ErrorResp{Code: http.StatusBadRequest, Msg: "got an error when request esign:" + err.Error()})
		return true
	}

	return false
}

func WriteError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusBadRequest, ErrorResp{
		Code: code,
		Msg:  msg,
	})
}

func WriteErrorF(ctx *gin.Context, code int, format string, values ...interface{}) {
	ctx.JSON(http.StatusBadRequest, ErrorResp{
		Code: code,
		Msg:  fmt.Sprintf(format, values),
	})
}

func WriteOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
