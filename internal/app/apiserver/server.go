package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marktsoy/flashcards_api/internal/app/store"
	"go.uber.org/zap"
)

/**
*
* Hide implementation of the server from outer world
*
 */

type server struct {
	config *Config
	router *gin.Engine
	logger *zap.SugaredLogger
	store  store.Store
}

// New - creates default instance of API server
func newServer(conf *Config, st store.Store) *server {
	s := &server{
		config: conf,
		store:  st,
	}
	if err := s.configureLogger(); err != nil {
		panic(err)
	}
	s.configureRouter()
	return s
}

// Start Starts server
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureLogger() error {
	loggerType := s.config.LoggerType
	if loggerType == "" {
		return nil
	}
	var cfg zap.Config
	if loggerType == "prod" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	outputs := []string{
		"stdout",
	}
	if s.config.LogPath != "" {
		outputs = append(outputs, s.config.LogPath)
	}
	cfg.OutputPaths = outputs
	logger, err := cfg.Build()
	if err == nil {
		s.logger = logger.Sugar()
	}
	return err
}

/**
* Endpoint declaration goes here
*
 */
func (s *server) configureRouter() {
	s.router = gin.New()
	s.router.Use(LogMiddleware(s.logger))
	// s.router.Use(gin.Recovery())
	userHandlers := s.router.Group("/users")
	{
		userHandlers.POST("/", s.createUser())
	}
}
