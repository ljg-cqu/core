package auth

import (
	"fmt"
	"github.com/appleboy/gofight/v2"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/utils"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestJWTParseAndVerifyAcessJwt(t *testing.T) {
	var userId, userName = "admin001", "Zealy"
	tkpair, err := CreateJwtPair(userId, userName)
	require.Nil(t, err)
	require.Equal(t, userId, tkpair.UserId)
	require.Equal(t, userName, tkpair.UserName)
	utils.PrintlnAsJson("token pair struct:", tkpair)

	e := gin.New()
	e.POST("/", func(c *gin.Context) {
		c.Header(AccessJwtKey, c.GetHeader(AccessJwtKey))
		c.Header(RefreshJwtKey, c.GetHeader(RefreshJwtKey))
		c.JSON(http.StatusOK, nil)
	})
	r := gofight.New()
	r.POST("/").
		SetDebug(true).
		SetHeader(gofight.H{AccessJwtKey: tkpair.AccessJwt}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			accessDetail, err := VerifyAccessJwt(req)
			require.Nil(t, err)
			require.Equal(t, accessDetail.UserId, tkpair.UserId)
			require.Equal(t, accessDetail.UserName, tkpair.UserName)
			require.Equal(t, accessDetail.AccessJwt, tkpair.AccessJwt)
			require.Equal(t, accessDetail.AccessJwtUuid, tkpair.AccessJwtUuid)

			utils.PrintlnAsJson("access token details:", accessDetail)
		})
}

func TestJWTParseAndVerifyRefreshJwt(t *testing.T) {
	var userId, userName = "admin001", "Zealy"
	tkpair, err := CreateJwtPair(userId, userName)
	require.Nil(t, err)
	require.Equal(t, userId, tkpair.UserId)
	require.Equal(t, userName, tkpair.UserName)
	utils.PrintlnAsJson("token pair struct:", tkpair)

	e := gin.New()
	e.POST("/", func(c *gin.Context) {
		c.Header(AccessJwtKey, c.GetHeader(AccessJwtKey))
		c.Header(RefreshJwtKey, c.GetHeader(RefreshJwtKey))
		c.JSON(http.StatusOK, nil)
	})
	r := gofight.New()
	r.POST("/").
		SetDebug(true).
		SetHeader(gofight.H{RefreshJwtKey: tkpair.RefreshJwt}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			refreshDetail, err := VerifyRefreshJwt(req)
			require.Nil(t, err)
			require.Equal(t, refreshDetail.UserId, tkpair.UserId)
			require.Equal(t, refreshDetail.UserName, tkpair.UserName)
			require.Equal(t, refreshDetail.RefreshJwt, tkpair.RefreshJwt)
			require.Equal(t, refreshDetail.RefreshJwtUuid, tkpair.RefreshJwtUuid)

			utils.PrintlnAsJson("refresh token details:", refreshDetail)
		})
}

func TestJWTParseAndVerifyAcessJwtTimeout(t *testing.T) {
	AccessTimeout = time.Nanosecond

	var userId, userName = "admin001", "Zealy"
	tkpair, _ := CreateJwtPair(userId, userName)

	time.Sleep(time.Second * 2)

	e := gin.New()
	e.POST("/", func(c *gin.Context) {
		c.Header(AccessJwtKey, c.GetHeader(AccessJwtKey))
		c.JSON(http.StatusOK, nil)
	})
	r := gofight.New()
	r.POST("/").
		SetDebug(true).
		SetHeader(gofight.H{AccessJwtKey: tkpair.AccessJwt}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			_, err := VerifyAccessJwt(req)
			require.NotNil(t, err)
			fmt.Println(err)
		})
}

func TestJWTParseAndVerifyFreshJwtTimeout(t *testing.T) {
	RefreshTimeout = time.Nanosecond

	var userId, userName = "admin001", "Zealy"
	tkpair, _ := CreateJwtPair(userId, userName)

	time.Sleep(time.Second * 2)

	e := gin.New()
	e.POST("/", func(c *gin.Context) {
		c.Header(RefreshJwtKey, c.GetHeader(RefreshJwtKey))
		c.JSON(http.StatusOK, nil)
	})
	r := gofight.New()
	r.POST("/").
		SetDebug(true).
		SetHeader(gofight.H{RefreshJwtKey: tkpair.RefreshJwt}).
		Run(e, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
			_, err := VerifyRefreshJwt(req)
			require.NotNil(t, err)
			fmt.Println(err)
		})
}
