package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/auth"
	"github.com/ljg-cqu/core/middle"
	"github.com/ljg-cqu/core/resp"
)

type LoginReq struct {
	LoginName string `json:"login_name"`
	Password  string `json:"password"`
}

type LoginRes struct {
	AccessJwt  string `json:"access_jwt"`
	RefreshJwt string `json:"refresh_jwt"`
}

func Login(c *gin.Context) {
	var loginReq LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		resp.Err(c, resp.ErrUnauthorized.AppendMsg(err.Error()))
		return
	}

	// validate name and password params format ...

	// Get user id by login name
	var userId = "cli0001"

	// TODO: store token in Redis, or PostgreSQL?

	tkp, err := auth.CreateJwtPair(userId, loginReq.LoginName)
	if err != nil {
		resp.Err(c, resp.ErrInternalServerError.AppendMsg(err.Error()))
		return
	}

	resp.Succ(c, &LoginRes{tkp.AccessJwt, tkp.RefreshJwt})
}

func Logout(c *gin.Context) {
	resp.Succ(c, nil)
}

type RefreshReq struct {
	RefreshJwt string `json:"refresh_jwt"`
}

type RefreshRes struct {
	AccessJwt  string `json:"access_jwt"`
	RefreshJwt string `json:"refresh_jwt"`
}

// Refresh generates refresh JWT for later access
// Note: this handler must be used with middle.ValidRefreshJwt
func Refresh(c *gin.Context) {
	v, _ := c.Get(middle.CtxRefreshJwtDetailKey)
	tokenDetail := v.(auth.RefreshJwtDetail)
	tkp, err := auth.CreateJwtPair(tokenDetail.UserId, tokenDetail.UserName)
	if err != nil {
		resp.Err(c, resp.ErrInternalServerError.AppendMsg(err.Error()))
		return
	}

	resp.Succ(c, &RefreshRes{tkp.AccessJwt, tkp.RefreshJwt})
}
