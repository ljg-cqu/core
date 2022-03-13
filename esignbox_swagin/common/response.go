package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type ErrorResp struct {
	Code int    `json:"code" binding:"required" default:"400" description:"错误码"`
	Msg  string `json:"msg" binding:"required" default:"请求发生错误" description:"错误信息"`
}

type EsignError struct {
	Code int
	Msg  string
}

func WriteErrore(ctx *gin.Context, rawResp *http.Response, respErr error, eError *EsignError) bool {
	if rawResp.StatusCode != http.StatusOK {
		err := errors.Errorf("response status code:%v, response status:%v", rawResp.StatusCode, rawResp.Status)
		ctx.JSON(http.StatusBadRequest, ErrorResp{Code: http.StatusBadRequest, Msg: "got an error when request esign:%v" + err.Error()})
		return true
	}

	if respErr != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResp{Code: http.StatusBadRequest, Msg: "got an error when request esign:%v" + respErr.Error()})
		return true
	}

	if eError.Code != 0 {
		err := errors.Errorf("esgin business code:%v, esign business message:%v", eError.Code, eError.Msg)
		ctx.JSON(http.StatusBadRequest, ErrorResp{Code: http.StatusBadRequest, Msg: "got an error when request esign:%v" + err.Error()})
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

func WriteErrorf(ctx *gin.Context, code int, format string, values ...interface{}) {
	ctx.JSON(http.StatusBadRequest, ErrorResp{
		Code: code,
		Msg:  fmt.Sprintf(format, values),
	})
}

func WriteOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

// ----------------

type RespObj struct {
	Code int         `json:"code" binding:"required"`
	Msg  string      `json:"msg" bindding:"required"`
	Data interface{} `json:"data"`
}

func RespSucc(c *gin.Context, data interface{}) {
	resp(c, &RespObj{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func RespErre(ctx *gin.Context, rawResp *http.Response, respErr error, eError *EsignError) bool {
	if rawResp.StatusCode != http.StatusOK {
		RespErrf(ctx, rawResp.StatusCode, "got an error when request esign:%v", rawResp.Status)
		return true
	}

	if respErr != nil {
		RespErrf(ctx, http.StatusBadRequest, "got an error when request esign:%v", respErr)
		return true
	}

	if eError.Code != 0 {
		RespErrf(ctx, http.StatusBadRequest, "esgin business code:%v, esign business message:%v", eError.Code, eError.Msg)
		return true
	}

	return false
}

func RespErr(c *gin.Context, code int, msg string) {
	resp(c, &RespObj{
		Code: code,
		Msg:  msg,
	})
}

func RespErrObj(c *gin.Context, errObj *RespObj) {
	resp(c, errObj)
}

func RespErrf(c *gin.Context, code int, format string, values ...interface{}) {
	resp(c, &RespObj{
		Code: code,
		Msg:  fmt.Sprintf(format, values),
	})
}

func resp(c *gin.Context, respObj *RespObj) {
	c.JSON(http.StatusOK, respObj)
}

// ----------------

func ErrObjFromEsignRequest(rawResp *http.Response, respErr error, eError *EsignError) (errObj *RespObj) {
	if rawResp.StatusCode != http.StatusOK {
		return &RespObj{
			Code: rawResp.StatusCode,
			Msg:  rawResp.Status,
		}
	}

	if respErr != nil {
		return &RespObj{
			Code: http.StatusBadRequest,
			Msg:  respErr.Error(),
		}
	}

	if eError.Code != 0 {
		return &RespObj{
			Code: eError.Code,
			Msg:  eError.Msg,
		}
	}

	return nil
}
