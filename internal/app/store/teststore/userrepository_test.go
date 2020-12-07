package teststore_test

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
	"github.com/marktsoy/flashcards_api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUser_Create(t *testing.T) {
	s := teststore.New()
	u := models.TestUser(t)
	err := s.User().Create(models.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUser_FindByEmail(t *testing.T) {
	s := teststore.New()
	u := models.TestUser(t)

	s.User().Create(u)
	user, err := s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, user)

}

func TestUser_FindByEmailNotFound(t *testing.T) {
	s := teststore.New()

	user, err := s.User().FindByEmail("notfoundtest@example.com")

	assert.Error(t, err, store.ErrRecordNotFound)
	assert.Nil(t, user)
}
