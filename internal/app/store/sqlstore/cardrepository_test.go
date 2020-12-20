package sqlstore_test

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func makeTestDeck(t *testing.T, store *sqlstore.SQLStore) *models.Deck {
	t.Helper()
	u := models.TestUser(t)
	err := store.User().Create(u)
	if err != nil {
		t.Fatal(err)
	}

	d := models.TestDeck(t)
	d.BindUser(u)
	err = store.Deck().Create(d)
	return d
}

func TestCard_Create(t *testing.T) {
	s, clear := sqlstore.TestStore(t, databaseURL)
	defer clear("users", "decks", "cards")

	iterationCount := 10
	d := makeTestDeck(t, s)
	for i := 0; i < iterationCount; i++ {
		card := models.TestCard(t, i)
		card.DeckID = d.ID
		err := s.Card().Create(card)
		assert.NoError(t, err)
		assert.NotNil(t, card.CreatedAt)
		assert.NotNil(t, card.ID)
	}
}
func TestCard_FindByDeck(t *testing.T) {
	s, clear := sqlstore.TestStore(t, databaseURL)
	defer clear("users", "decks", "cards")

	iterationCount := 10
	d := makeTestDeck(t, s)
	createdIDs := make([]int, 0)
	for i := 0; i < iterationCount; i++ {
		card := models.TestCard(t, i)
		card.DeckID = d.ID
		s.Card().Create(card)
		createdIDs = append(createdIDs, card.ID)
	}

	cards, err := s.Card().FindAllByDeck(d)

	assert.NoError(t, err)
	assert.Len(t, cards, len(createdIDs))
	for _, c := range cards {
		assert.Contains(t, createdIDs, c.ID)
	}

	// No deck
	cards, err = s.Card().FindAllByDeck(models.TestDeck(t))
	assert.NoError(t, err)
	assert.Len(t, cards, 0)

}
