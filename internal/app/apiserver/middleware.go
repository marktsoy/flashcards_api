package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LogMiddleware instance a Logger middleware with config.
func LogMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	if logger == nil {
		return func(c *gin.Context) {
			c.Next()
		}
	}
	return func(c *gin.Context) {
		// Start timer
		method := c.Request.Method
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// Process request
		c.Next()
		status := c.Writer.Status()
		if http.StatusOK <= status && status < 400 {
			logger.Infof("%v\t %v\t %v \t %v", method, path, raw, status)
		} else {
			logger.Warnf("%v\t %v\t %v \t %v", method, path, raw, status)
		}

	}
}
