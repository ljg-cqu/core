package auth

import (
	"github.com/appleboy/gofight/v2"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/utils"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestJWTParseAndVerifyAcessToken(t *testing.T) {
	var userId, userName = "admin001", "Zealy"
	tkpair, err := CreateTokenPair(userId, userName)
	require.Nil(t, err)
	require.Equal(t, userId, tkpair.UserId)
	require.Equal(t, userName, tkpair.UserName)
	utils.PrintlnAsJson("token pair struct:", tkpair)

	e := gin.New()
	e.POST("/", func(c *gin.Context) {
		c.Header(AccessTokenKey, c.GetHeader(AccessTokenKey))
		c.Header(RefreshTokenKey, c.GetHeader(RefreshTokenKey))
		c.JSON(http.StatusOK, nil)
	})
	r := gofight.New()
	r.POST("/").
		SetDebug(true).
		SetHeader(gofight.H{AccessTokenKey: tkpair.AccessTokenStr}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			accessDetail, err := VerifyAccessToken(req)
			require.Nil(t, err)
			require.Equal(t, accessDetail.UserId, tkpair.UserId)
			require.Equal(t, accessDetail.UserName, tkpair.UserName)
			require.Equal(t, accessDetail.AccessTokenStr, tkpair.AccessTokenStr)
			utils.PrintlnAsJson("access token details:", accessDetail)
		})
}

func TestJWTParseAndVerifyRefreshToken(t *testing.T) {
	var userId, userName = "admin001", "Zealy"
	tkpair, err := CreateTokenPair(userId, userName)
	require.Nil(t, err)
	require.Equal(t, userId, tkpair.UserId)
	require.Equal(t, userName, tkpair.UserName)
	utils.PrintlnAsJson("token pair struct:", tkpair)

	e := gin.New()
	e.POST("/", func(c *gin.Context) {
		c.Header(AccessTokenKey, c.GetHeader(AccessTokenKey))
		c.Header(RefreshTokenKey, c.GetHeader(RefreshTokenKey))
		c.JSON(http.StatusOK, nil)
	})
	r := gofight.New()
	r.POST("/").
		SetDebug(true).
		SetHeader(gofight.H{RefreshTokenKey: tkpair.RefreshTokenStr}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			refreshDetail, err := VerifyRefreshToken(req)
			require.Nil(t, err)
			require.Equal(t, refreshDetail.UserId, tkpair.UserId)
			require.Equal(t, refreshDetail.UserName, tkpair.UserName)
			require.Equal(t, refreshDetail.RefreshTokenStr, tkpair.RefreshTokenStr)
			utils.PrintlnAsJson("refresh token details:", refreshDetail)
		})
}

func TestJWTParseAndVerifyAcessTokenTimeout(t *testing.T) {
	accessTimeout = time.Nanosecond
	var userId, userName = "admin001", "Zealy"
	tkpair, err := CreateTokenPair(userId, userName)
	require.Nil(t, err)
	require.Equal(t, userId, tkpair.UserId)
	require.Equal(t, userName, tkpair.UserName)
	utils.PrintlnAsJson("token pair struct:", tkpair)

	time.Sleep(time.Second * 2)

	e := gin.New()
	e.POST("/", func(c *gin.Context) {
		c.Header(AccessTokenKey, c.GetHeader(AccessTokenKey))
		c.Header(RefreshTokenKey, c.GetHeader(RefreshTokenKey))
		c.JSON(http.StatusOK, nil)
	})
	r := gofight.New()
	r.POST("/").
		SetDebug(true).
		SetHeader(gofight.H{AccessTokenKey: tkpair.AccessTokenStr}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			accessDetail, err := VerifyAccessToken(req)
			require.Nil(t, err, err)
			require.Equal(t, accessDetail.UserId, tkpair.UserId)
			require.Equal(t, accessDetail.UserName, tkpair.UserName)
			require.Equal(t, accessDetail.AccessTokenStr, tkpair.AccessTokenStr)
			utils.PrintlnAsJson("access token details:", accessDetail)
		})
}
