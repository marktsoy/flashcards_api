package sqlstore_test

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
	"github.com/marktsoy/flashcards_api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUser_Create(t *testing.T) {
	s, clear := sqlstore.TestStore(t, databaseURL)
	defer clear("users")
	u := models.TestUser(t)
	err := s.User().Create(models.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUser_FindByEmail(t *testing.T) {
	s, clear := sqlstore.TestStore(t, databaseURL)
	defer clear("users")
	u := models.TestUser(t)

	s.User().Create(u)
	user, err := s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, user)

}

func TestUser_FindByEmailNotFound(t *testing.T) {
	s, clear := sqlstore.TestStore(t, databaseURL)
	defer clear("users")

	user, err := s.User().FindByEmail("notfoundtest@example.com")

	assert.Error(t, err, store.ErrRecordNotFound)
	assert.Nil(t, user)
}
