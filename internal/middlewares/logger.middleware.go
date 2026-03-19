package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		c.Next()
		logger.Info("Incoming Request",
			zap.String("method", method),
			zap.String("path", path),
			zap.Duration("duration", time.Since(start)),
			zap.Int("status", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
		)
	}
}