package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marktsoy/flashcards_api/internal/app/store"
	"go.uber.org/zap"
)

// APIServer ...
type APIServer struct {
	config *Config
	router *gin.Engine
	logger *zap.SugaredLogger
	store  *store.Store
}

// New - creates default instance of API server
func New(conf *Config) *APIServer {
	return &APIServer{
		config: conf,
	}
}

// Start Starts server
func (s *APIServer) Start() error {
	if err := s.configureStore(); err != nil {
		panic(err)
	}
	if err := s.configureLogger(); err != nil {
		panic(err)
	}
	s.configureRouter()
	s.logger.Debugf("Starting API server at address: %v", s.config.BindAddr)

	return s.router.Run(s.config.BindAddr)
}

func (s *APIServer) configureLogger() error {
	loggerType := s.config.LoggerType
	var cfg zap.Config
	if loggerType == "prod" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	cfg.OutputPaths = []string{
		"stdout",
		s.config.LogPath,
	}
	logger, err := cfg.Build()
	if err == nil {
		s.logger = logger.Sugar()
	}
	return err
}

func (s *APIServer) configureRouter() {
	s.router = gin.New()
	s.router.Use(LogMiddleware(s.logger))
	test := s.router.Group("/test")
	{
		test.GET("/", TestIndex())
	}
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

// LogMiddleware instance a Logger middleware with config.
func LogMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {

	logger.Info("Using Zap")
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
