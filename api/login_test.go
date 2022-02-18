package api

import (
	"encoding/json"
	"github.com/appleboy/gofight/v2"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/auth"
	"github.com/ljg-cqu/core/middle"
	"github.com/ljg-cqu/core/resp"
	"github.com/ljg-cqu/core/utils"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	e := gin.Default()
	e.POST("/login", Login)

	var user = LoginReq{
		LoginName: "Zealy",
		Password:  "11111188",
	}

	f := gofight.New()
	f.POST("/login").
		SetDebug(true).
		SetJSONInterface(&user).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)

			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Nil(t, err, respObj)
			require.Equal(t, resp.StatusOK, respObj.Code, respObj.Msg)

			var jwtPair RefreshRes

			err = utils.Json.UnmarshalFromString(respObj.Data, &jwtPair)
			require.Nil(t, err)

			utils.PrintlnAsJson("new access jwt:", jwtPair.AccessJwt)
			utils.PrintlnAsJson("new refresh jwt:", jwtPair.RefreshJwt)
		})
}

func TestLogout(t *testing.T) {
	e := gin.Default()
	e.POST("/logout", Logout)

	r := gofight.New()
	r.POST("/logout").
		SetDebug(true).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)

			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Nil(t, err, respObj)
			require.Equal(t, resp.StatusOK, respObj.Code)
		})
}

func TestRefresh(t *testing.T) {
	var userId, userName = "admin001", "Zealy"
	pastTkpair, err := auth.CreateJwtPair(userId, userName)
	require.Nil(t, err)
	utils.PrintlnAsJson("past refresh jwt:", pastTkpair.RefreshJwt)

	e := gin.Default()
	e.Use(middle.ValidRefreshJwt())
	e.POST("/refresh", Refresh)

	r := gofight.New()
	r.POST("/refresh").
		SetDebug(true).
		SetHeader(gofight.H{auth.RefreshJwtKey: pastTkpair.RefreshJwt}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)

			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Nil(t, err, respObj)
			require.Equal(t, resp.StatusOK, respObj.Code, respObj.Msg)

			var jwtPair RefreshRes

			err = utils.Json.UnmarshalFromString(respObj.Data, &jwtPair)
			require.Nil(t, err)
			require.NotEqual(t, jwtPair.AccessJwt, pastTkpair.AccessJwt)
			require.NotEqual(t, jwtPair.RefreshJwt, pastTkpair.RefreshJwt)

			utils.PrintlnAsJson("new access jwt:", jwtPair.AccessJwt)
			utils.PrintlnAsJson("new refresh jwt:", jwtPair.RefreshJwt)
		})
}

func TestRefreshTimeout(t *testing.T) {
	auth.RefreshTimeout = time.Nanosecond

	var userId, userName = "admin001", "Zealy"
	pastTkpair, err := auth.CreateJwtPair(userId, userName)
	require.Nil(t, err)
	utils.PrintlnAsJson("past refresh jwt:", pastTkpair.RefreshJwt)

	time.Sleep(time.Second * 2)

	e := gin.Default()
	e.Use(middle.ValidRefreshJwt())
	e.POST("/refresh", Refresh)

	r := gofight.New()
	r.POST("/refresh").
		SetDebug(true).
		SetHeader(gofight.H{auth.RefreshJwtKey: pastTkpair.RefreshJwt}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			require.Equal(t, http.StatusOK, res.Code)

			var respObj resp.RespObj
			err := json.Unmarshal(res.Body.Bytes(), &respObj)
			require.Nil(t, err, respObj)
			require.Equal(t, resp.StatusUnauthorized, respObj.Code)
		})
}
