package middleware

import (
	"crypto/subtle"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuth(realm string, creds map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if !ok {
			basicAuthFailed(c.Writer, realm)
			return
		}

		credPass, credUserOk := creds[user]
		if !credUserOk || subtle.ConstantTimeCompare([]byte(pass), []byte(credPass)) != 1 {
			basicAuthFailed(c.Writer, realm)
			return
		}
		c.Next()
	}
}

func basicAuthFailed(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	w.WriteHeader(http.StatusUnauthorized)
}
