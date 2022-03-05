package middleware

import (
	"github.com/appleboy/gofight/v2"
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/logger"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogrus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Recovery(), Logrus(logger.WithFormatter(&logrus.TextFormatter{ForceColors: true})))

	// pingpong
	r.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("pong"))
	})

	f := gofight.New()
	f.GET("/ping").SetDebug(true).Run(r, func(res gofight.HTTPResponse, req gofight.HTTPRequest) {
		data := res.Body.String()
		require.Equal(t, "pong", data)
	})

	// Check log in /tmp/abfpaas-log/
}

func TestLogrus2(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Recovery(), Logrus(logger.WithFormatter(&logrus.TextFormatter{ForceColors: true})))

	// pingpong
	r.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("pong"))
	})

	r.Run("127.0.0.1:8080")

	// Run http://127.0.0.1:8080/ping and see what happen
}
