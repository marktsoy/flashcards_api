package apiserver

import "github.com/gin-gonic/gin"

/**
* Endpoint declaration goes here
*
 */
func (s *server) configureRouter() {
	s.router = gin.New()
	s.router.Use(LogMiddleware(s.logger))
	s.router.Use(OnlyJSON)
	authMiddleware := BasicAuth(s.store)
	userHandlers := s.router.Group("/users")
	{
		userHandlers.POST("/", s.createUser())
		userHandlers.GET("/", authMiddleware, s.listUsers())
		userHandlers.GET("/me", authMiddleware, s.me())
	}

	deckHandlers := s.router.Group("/decks")
	deckHandlers.Use(authMiddleware)
	{
		deckHandlers.POST("/", s.createDeck())
		deckHandlers.PUT("/:id", s.updateDeck())
		deckHandlers.GET("/:id", s.getDeck())
		deckHandlers.DELETE("/:id", s.destroyDeck())
	}

	cardHandlers := s.router.Group("/cards")
	cardHandlers.Use(authMiddleware)
	{
		cardHandlers.GET("/:deck", s.getCards())
		cardHandlers.POST("/", s.createCard())
		cardHandlers.PUT("/:id", s.updateCard())
		cardHandlers.DELETE("/:id", s.destroyCard())
	}

}
