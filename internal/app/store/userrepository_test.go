package store_test

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUser_Create(t *testing.T) {
	s, clear := store.TestStore(t, databaseURL)
	defer clear("users")
	u, err := s.User().Create(models.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUser_FindByEmail(t *testing.T) {
	
}
