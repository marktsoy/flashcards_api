package apiserver

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/validation"
)

// TestIndex ...
func TestIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		io.WriteString(ctx.Writer, "Testing is ok")
	}
}

// Get current user
func (s *server) me() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := getUser(c)
		c.JSON(200, user)
	}
}

func (s *server) listUsers() gin.HandlerFunc {
	repo := s.store.User()
	return func(c *gin.Context) {
		users, err := repo.FindAll()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, users)
	}
}

func (s *server) createUser() gin.HandlerFunc {
	type req struct {
		Email                string `json:"email" `
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	}
	return func(c *gin.Context) {

		r := &req{}

		if err := c.ShouldBind(r); err != nil {
			c.AbortWithStatusJSON(422, gin.H{
				"error": "Invalid request",
			})
			return
		}

		valErrors := gin.H{}

		if ok, msg := validation.Email("Email", r.Email); !ok {
			valErrors["Email"] = msg
		}
		if ok, msg := validation.MinLen("Password", r.Password, 8); !ok {
			valErrors["Password"] = msg
		}
		if r.PasswordConfirmation != r.Password {
			valErrors["PasswordConfirmation"] = "Passwords does not match"
		}

		if len(valErrors) > 0 {
			c.AbortWithStatusJSON(422, gin.H{
				"error": valErrors,
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

func (s *server) createDeck() gin.HandlerFunc {
	type req struct {
		Name string `json:"name" `
	}
	return func(c *gin.Context) {
		user := getUser(c)
		r := &req{}

		if err := c.ShouldBind(r); err != nil {
			c.AbortWithStatusJSON(422, gin.H{
				"error": "Invalid request",
			})
			return
		}
		valErrors := gin.H{}
		if ok, msg := validation.MinLen("Name", r.Name, 1, "Name is required"); !ok {
			valErrors["Name"] = msg
		}
		if len(valErrors) > 0 {
			c.AbortWithStatusJSON(422, gin.H{
				"error": valErrors,
			})
			return
		}

		d := &models.Deck{
			Name: r.Name,
		}
		d.BindUser(user)

		if err := s.store.Deck().Create(d); err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": "Could not create deck",
			})
			return
		}

		c.JSON(201, d)
		return
	}
}

func (s *server) updateDeck() gin.HandlerFunc {
	type req struct {
		Name string `json:"name" `
	}
	return func(c *gin.Context) {
		user := getUser(c)
		id := c.Param("id")
		r := &req{}
		d, err := s.store.Deck().FindByID(id)
		if err != nil {
			c.AbortWithStatusJSON(422, gin.H{
				"error": "Deck:" + id + " was not found",
			})
			return
		}
		if d.UserID != user.ID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not authorized",
			})
			return
		}
		if err := c.ShouldBind(r); err != nil {
			c.AbortWithStatusJSON(422, gin.H{
				"error": "Invalid request",
			})
			return
		}
		valErrors := gin.H{}
		if ok, msg := validation.MinLen("Name", r.Name, 1, "Name is required"); !ok {
			valErrors["Name"] = msg
			c.AbortWithStatusJSON(422, gin.H{
				"error": valErrors,
			})
			return
		}
		d.Name = r.Name

		if err := s.store.Deck().Update(d); err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": "Could not update deck",
			})
			return
		}

		c.JSON(200, d)
		return
	}
}
