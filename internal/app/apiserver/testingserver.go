package apiserver

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/store/teststore"
)

func TestConfig(t *testing.T) *Config {
	t.Helper()
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
