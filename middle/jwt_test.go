package middle

import (
	"github.com/appleboy/gofight/v2"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/auth"
	"github.com/ljg-cqu/core/resp"
	"github.com/ljg-cqu/core/utils"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func TestValidAccessToken(t *testing.T) {
	var userId, userName = "admin001", "Zealy"
	tkpair, _ := auth.CreateTokenPair(userId, userName)

	e := gin.Default()
	e.Use(ValidAccessToken())
	e.POST("/todo", func(c *gin.Context) {
		v, ok := c.Get(CtxAccessTokenDetailKey)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}
		tokenDetail, ok := v.(*auth.AccessTokenDetail)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}

		resp.Succ(c, tokenDetail)
	})

	r := gofight.New()
	r.POST("/todo").
		SetDebug(true).
		SetHeader(gofight.H{auth.AccessTokenKey: tkpair.AccessTokenStr}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)
			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Equal(t, resp.StatusOK, respObj.Code, respObj.Msg)
			require.Nil(t, err, respObj)
			var tokenDetail auth.AccessTokenDetail

			err = json.UnmarshalFromString(respObj.Data, &tokenDetail)
			require.Nil(t, err)
			require.Equal(t, tkpair.UserId, tokenDetail.UserId)
			require.Equal(t, tkpair.UserName, tokenDetail.UserName)
			require.Equal(t, tkpair.AccessTokenStr, tokenDetail.AccessTokenStr)
			utils.PrintlnAsJson("access tokenDetail details:", tokenDetail)
		})
}

func TestValidRefreshTokenToken(t *testing.T) {
	var userId, userName = "admin001", "Zealy"
	tkpair, _ := auth.CreateTokenPair(userId, userName)

	e := gin.Default()
	e.Use(ValidRefreshToken())
	e.POST("/refresh", func(c *gin.Context) {
		v, ok := c.Get(CtxRefreshTokenDetailKey)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}
		tokenDetail, ok := v.(*auth.RefreshTokenDetail)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}

		resp.Succ(c, tokenDetail)
	})

	r := gofight.New()
	r.POST("/refresh").
		SetDebug(true).
		SetHeader(gofight.H{auth.RefreshTokenKey: tkpair.RefreshTokenStr}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)
			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Equal(t, resp.StatusOK, respObj.Code, respObj.Msg)
			require.Nil(t, err, respObj)
			var tokenDetail auth.RefreshTokenDetail

			err = json.UnmarshalFromString(respObj.Data, &tokenDetail)
			require.Nil(t, err)
			require.Equal(t, tkpair.UserId, tokenDetail.UserId)
			require.Equal(t, tkpair.UserName, tokenDetail.UserName)
			require.Equal(t, tkpair.RefreshTokenStr, tokenDetail.RefreshTokenStr)
			utils.PrintlnAsJson("access tokenDetail details:", tokenDetail)
		})
}

func TestValidAccessTokenTimeout(t *testing.T) {
	auth.AccessTimeout = time.Nanosecond
	var userId, userName = "admin001", "Zealy"
	tkpair, _ := auth.CreateTokenPair(userId, userName)

	time.Sleep(time.Second * 2)

	e := gin.Default()
	e.Use(ValidAccessToken())
	e.POST("/todo", func(c *gin.Context) {
		v, ok := c.Get(CtxAccessTokenDetailKey)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}
		tokenDetail, ok := v.(*auth.AccessTokenDetail)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}

		resp.Succ(c, tokenDetail)
	})

	r := gofight.New()
	r.POST("/todo").
		SetDebug(true).
		SetHeader(gofight.H{auth.AccessTokenKey: tkpair.AccessTokenStr}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)
			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Equal(t, resp.StatusOK, respObj.Code, respObj.Msg)
			require.Nil(t, err, respObj)
			var tokenDetail auth.AccessTokenDetail

			err = json.UnmarshalFromString(respObj.Data, &tokenDetail)
			require.Nil(t, err)
			require.Equal(t, tkpair.UserId, tokenDetail.UserId)
			require.Equal(t, tkpair.UserName, tokenDetail.UserName)
			require.Equal(t, tkpair.AccessTokenStr, tokenDetail.AccessTokenStr)
			utils.PrintlnAsJson("access tokenDetail details:", tokenDetail)
		})
}

func TestValidRefreshTokenTokenTimeout(t *testing.T) {
	auth.RefreshTimeout = time.Nanosecond
	var userId, userName = "admin001", "Zealy"
	tkpair, _ := auth.CreateTokenPair(userId, userName)

	time.Sleep(time.Second * 2)

	e := gin.Default()
	e.Use(ValidRefreshToken())
	e.POST("/refresh", func(c *gin.Context) {
		v, ok := c.Get(CtxRefreshTokenDetailKey)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}
		tokenDetail, ok := v.(*auth.RefreshTokenDetail)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}

		resp.Succ(c, tokenDetail)
	})

	r := gofight.New()
	r.POST("/refresh").
		SetDebug(true).
		SetHeader(gofight.H{auth.RefreshTokenKey: tkpair.RefreshTokenStr}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)
			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Equal(t, resp.StatusOK, respObj.Code, respObj.Msg)
			require.Nil(t, err, respObj)
			var tokenDetail auth.RefreshTokenDetail

			err = json.UnmarshalFromString(respObj.Data, &tokenDetail)
			require.Nil(t, err)
			require.Equal(t, tkpair.UserId, tokenDetail.UserId)
			require.Equal(t, tkpair.UserName, tokenDetail.UserName)
			require.Equal(t, tkpair.RefreshTokenStr, tokenDetail.RefreshTokenStr)
			utils.PrintlnAsJson("access tokenDetail details:", tokenDetail)
		})
}
