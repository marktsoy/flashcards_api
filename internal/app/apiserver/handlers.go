package apiserver

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/marktsoy/flashcards_api/internal/app/models"
)

func TestIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		io.WriteString(ctx.Writer, "Testing is ok")
	}
}

func (s *server) createUser() gin.HandlerFunc {
	type req struct {
		Email                string `json:"email" binding:"required"`
		Password             string `json:"password" binding:"required"`
		PasswordConfirmation string `json:"password_confirmation" binding:"required"`
	}
	return func(c *gin.Context) {

		r := &req{}

		if err := c.ShouldBind(r); err != nil {
			c.AbortWithStatusJSON(422, gin.H{
				"error": "Invalid request",
			})
			return
		}

		u := &models.User{
			Email:    r.Email,
			Password: r.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": "Could not save the user",
			})
			return
		}

		c.JSON(201, u)
		return
	}
}
