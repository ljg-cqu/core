package middle

import (
	"github.com/gin-gonic/gin"
	"github.com/ljg-cqu/core/logger"
	ginlogrus "github.com/toorop/gin-logrus"
)

func Logrus(options ...logger.Option) gin.HandlerFunc {
	return ginlogrus.Logger(logger.New(options...))
}
