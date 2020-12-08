package teststore_test

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
	"github.com/marktsoy/flashcards_api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestDeck_Create(t *testing.T) {
	s := teststore.New()
	u := models.TestUser(t)
	s.User().Create(u)
	deck := models.TestDeck(t)
	deck.BindUser(u)
	err := s.Deck().Create(deck)
	assert.NoError(t, err)
	assert.NotNil(t, deck)
	assert.NotNil(t, u)
}

func TestDeckRep_Implements(t *testing.T) {
	assert.Implements(t, (*store.DeckRepository)(nil), &teststore.DeckRepository{})
}
