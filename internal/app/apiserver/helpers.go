package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marktsoy/flashcards_api/internal/app/models"
)

var (
	userAuthKey = "auth"
)

// Auth helpers
func setUser(c *gin.Context, user *models.User) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[userAuthKey] = user
}

func getUser(c *gin.Context) *models.User {
	return c.Keys[userAuthKey].(*models.User)
}

func authErr(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "Unauthorized",
	})
}

// END ----- Auth helpers
