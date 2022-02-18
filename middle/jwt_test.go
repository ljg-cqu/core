package middle

import (
	"fmt"
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

func TestValidAccessJwt(t *testing.T) {
	var userId, userName = "admin001", "Zealy"
	tkpair, _ := auth.CreateJwtPair(userId, userName)

	e := gin.Default()
	e.Use(ValidAccessJwt())
	e.POST("/todo", func(c *gin.Context) {
		v, ok := c.Get(CtxAccessJwtDetailKey)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}
		tokenDetail, ok := v.(auth.AccessJwtDetail)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}

		resp.Succ(c, tokenDetail)
	})

	r := gofight.New()
	r.POST("/todo").
		SetDebug(true).
		SetHeader(gofight.H{auth.AccessJwtKey: tkpair.AccessJwt}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)
			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Equal(t, resp.StatusOK, respObj.Code, respObj.Msg)
			require.Nil(t, err, respObj)
			var tokenDetail auth.AccessJwtDetail

			err = json.UnmarshalFromString(respObj.Data, &tokenDetail)
			require.Nil(t, err)
			require.Equal(t, tkpair.UserId, tokenDetail.UserId)
			require.Equal(t, tkpair.UserName, tokenDetail.UserName)
			require.Equal(t, tkpair.AccessJwt, tokenDetail.AccessJwt)
			utils.PrintlnAsJson("access tokenDetail details:", tokenDetail)
		})
}

func TestValidRefreshJwt(t *testing.T) {
	var userId, userName = "admin001", "Zealy"
	tkpair, _ := auth.CreateJwtPair(userId, userName)

	e := gin.Default()
	e.Use(ValidRefreshJwt())
	e.POST("/refresh", func(c *gin.Context) {
		v, ok := c.Get(CtxRefreshJwtDetailKey)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}
		tokenDetail, ok := v.(auth.RefreshJwtDetail)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}

		resp.Succ(c, tokenDetail)
	})

	r := gofight.New()
	r.POST("/refresh").
		SetDebug(true).
		SetHeader(gofight.H{auth.RefreshJwtKey: tkpair.RefreshJwt}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)
			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Equal(t, resp.StatusOK, respObj.Code, respObj.Msg)
			require.Nil(t, err, respObj)
			var tokenDetail auth.RefreshJwtDetail

			err = json.UnmarshalFromString(respObj.Data, &tokenDetail)
			require.Nil(t, err)
			require.Equal(t, tkpair.UserId, tokenDetail.UserId)
			require.Equal(t, tkpair.UserName, tokenDetail.UserName)
			require.Equal(t, tkpair.RefreshJwt, tokenDetail.RefreshJwt)
			utils.PrintlnAsJson("access tokenDetail details:", tokenDetail)
		})
}

func TestValidAccessJwtTimeout(t *testing.T) {
	auth.AccessTimeout = time.Nanosecond
	var userId, userName = "admin001", "Zealy"
	tkpair, _ := auth.CreateJwtPair(userId, userName)

	time.Sleep(time.Second * 2)

	e := gin.Default()
	e.Use(ValidAccessJwt())
	e.POST("/todo", func(c *gin.Context) {
		v, ok := c.Get(CtxAccessJwtDetailKey)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}
		tokenDetail, ok := v.(*auth.AccessJwtDetail)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}

		resp.Succ(c, tokenDetail)
	})

	r := gofight.New()
	r.POST("/todo").
		SetDebug(true).
		SetHeader(gofight.H{auth.AccessJwtKey: tkpair.AccessJwt}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)
			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Nil(t, err)

			require.Equal(t, resp.StatusUnauthorized, respObj.Code)
			fmt.Println(respObj.Msg)
		})
}

func TestValidRefreshJwtTimeout(t *testing.T) {
	auth.RefreshTimeout = time.Nanosecond
	var userId, userName = "admin001", "Zealy"
	tkpair, _ := auth.CreateJwtPair(userId, userName)

	time.Sleep(time.Second * 2)

	e := gin.Default()
	e.Use(ValidRefreshJwt())
	e.POST("/refresh", func(c *gin.Context) {
		v, ok := c.Get(CtxRefreshJwtDetailKey)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}
		tokenDetail, ok := v.(*auth.RefreshJwtDetail)
		if !ok {
			resp.Err(c, resp.ErrUnauthorized)
			return
		}

		resp.Succ(c, tokenDetail)
	})

	r := gofight.New()
	r.POST("/refresh").
		SetDebug(true).
		SetHeader(gofight.H{auth.RefreshJwtKey: tkpair.RefreshJwt}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)
			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Nil(t, err)

			require.Equal(t, resp.StatusUnauthorized, respObj.Code)
			fmt.Println(respObj.Msg)
		})
}
