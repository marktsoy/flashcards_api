package sqlstore_test

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
	"github.com/marktsoy/flashcards_api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestDeck_Create(t *testing.T) {
	s, clear := sqlstore.TestStore(t, databaseURL)
	defer clear("users", "decks")

	u := models.TestUser(t)
	err := s.User().Create(u)
	if err != nil {
		t.Fatal(err)
	}

	d := models.TestDeck(t)
	d.BindUser(u)
	err = s.Deck().Create(d)
	assert.NoError(t, err)
	assert.Len(t, d.ID, 32)
}
func TestDeck_FindByID(t *testing.T) {
	s, clear := sqlstore.TestStore(t, databaseURL)
	defer clear("users", "decks")

	u := models.TestUser(t)
	err := s.User().Create(u)
	if err != nil {
		t.Fatal(err)
	}
	d := models.TestDeck(t)
	d.BindUser(u)
	err = s.Deck().Create(d)
	assert.NoError(t, err)
	assert.Len(t, d.ID, 32)

	var retrievedDeck *models.Deck
	retrievedDeck, err = s.Deck().FindByID(d.ID)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedDeck)
	assert.Equal(t, retrievedDeck.UserID, u.ID)

	retrievedDeck, err = s.Deck().FindByID("Unknown ID")
	assert.Error(t, err)
	assert.Equal(t, err, store.ErrRecordNotFound)
	assert.Nil(t, retrievedDeck)

}
