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

	authProtectedRoute := s.router.Group("/test")
	{
		authProtectedRoute.POST("/", TestIndex())
		authProtectedRoute.GET("/", TestIndex())
	}
}