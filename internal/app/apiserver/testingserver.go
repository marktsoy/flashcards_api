package apiserver

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/marktsoy/flashcards_api/internal/app/store/teststore"
)

// TestConfig ...
func TestConfig(t *testing.T) *Config {
	t.Helper()
	gin.SetMode(gin.ReleaseMode)
	return &Config{
		BindAddr:   "localhost:3333",
		LoggerType: "",
		LogPath:    "",
	}
}

// TestServer returns nonloggin,mocked store server
func TestServer(t *testing.T) *server {
	t.Helper()
	store := teststore.New()
	config := TestConfig(t)

	s := newServer(config, store)
	return s
}
