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

func TestUser_FindAll(t *testing.T) {
	s, clear := sqlstore.TestStore(t, databaseURL)
	defer clear("users")
	testUsers := []*models.User{
		{Email: "my@me.at", Password: "qwerty123123"},
		{Email: "mysecondUser@me.at", Password: "qwerty123123"},
		{Email: "my3rdUser@me.at", Password: "qwerty123123"},
	}
	for _, u := range testUsers {
		s.User().Create(u)
	}
	users, err := s.User().FindAll()

	assert.NoError(t, err)
	assert.IsType(t, []*models.User{}, users)
	assert.Equal(t, len(users), len(testUsers))
}
