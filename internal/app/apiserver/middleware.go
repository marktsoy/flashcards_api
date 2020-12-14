package apiserver

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/marktsoy/flashcards_api/internal/app/store"
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

// OnlyJSON ...
func OnlyJSON(c *gin.Context) {
	if contentType := c.Request.Header.Get("Content-Type"); contentType != "application/json" {
		c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
			"error": "Unsupported content type",
		})
	}
}

// BasicAuth ...
func BasicAuth(store store.Store) gin.HandlerFunc {
	repository := store.User()
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		h := strings.Split(authHeader, " ")
		if h[0] != "Basic" || len(h) != 2 {
			authErr(c)
			return
		}
		cred, err := base64.StdEncoding.DecodeString(h[1])
		if err != nil {
			authErr(c)
			return
		}
		auth := strings.Split(string(cred), ":")
		if len(auth) != 2 {
			authErr(c)
			return
		}
		user, err := repository.FindByEmail(auth[0])
		if err != nil {
			authErr(c)
			return
		}
		if err := user.CheckPassword(auth[1]); err != nil {
			authErr(c)
			return
		}

		setUser(c, user)

		c.Next()
	}
}
