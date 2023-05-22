package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func StructuredLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		sugar := logger.Sugar()
		start := time.Now()
		c.Next()
		stop := time.Now()
		latency := stop.Sub(start)
		if c.Writer.Status() >= 500 {
			sugar.Errorw("error", c.Errors.ByType(gin.ErrorTypePrivate).String(),
				"IP", c.ClientIP(),
				"method", c.Request.Method,
				"status", c.Writer.Status(),
				"size", c.Writer.Size(),
				"path", c.Request.URL.Path,
				"latency", latency)
		} else {
			sugar.Infow("request",
				"IP", c.ClientIP(),
				"method", c.Request.Method,
				"status", c.Writer.Status(),
				"size", c.Writer.Size(),
				"path", c.Request.URL.Path,
				"latency", latency)
		}
	}

}
