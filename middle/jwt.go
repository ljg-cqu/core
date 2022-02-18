package middle

import (
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/auth"
	"github.com/ljg-cqu/core/resp"
)

// Keys to be set in gin.Context
const (
	CtxAccessTokenDetailKey  = "AccessTokenDetail"
	CtxRefreshTokenDetailKey = "RefreshTokenDetail"
)

func ValidAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenDetail, err := auth.VerifyAccessToken(c.Request)
		if err != nil {
			resp.Err(c, resp.ErrUnauthorized.AppendMsg(err.Error()))
			c.Abort()
			return
		}
		c.Set("AccessTokenDetail", tokenDetail)
		c.Next()
	}
}

func ValidRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenDetail, err := auth.VerifyRefreshToken(c.Request)
		if err != nil {
			resp.Err(c, resp.ErrUnauthorized.AppendMsg(err.Error()))
			c.Abort()
			return
		}
		c.Set("RefreshTokenDetail", tokenDetail)
		c.Next()
	}
}
