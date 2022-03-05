package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/auth"
	"github.com/ljg-cqu/core/resp"
)

// Keys to be set in gin.Context
const (
	CtxAccessJwtDetailKey  = "AccessJwtDetailKey"
	CtxRefreshJwtDetailKey = "RefreshJwtDetailKey"
)

func ValidAccessJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenDetail, err := auth.VerifyAccessJwt(c.Request)
		if err != nil {
			resp.Err(c, resp.ErrUnauthorized.AppendMsg(err.Error()))
			c.Abort()
			return
		}
		c.Set(CtxAccessJwtDetailKey, *tokenDetail)
		c.Next()
	}
}

func ValidRefreshJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenDetail, err := auth.VerifyRefreshJwt(c.Request)
		if err != nil {
			resp.Err(c, resp.ErrUnauthorized.AppendMsg(err.Error()))
			c.Abort()
			return
		}
		c.Set(CtxRefreshJwtDetailKey, *tokenDetail)
		c.Next()
	}
}
