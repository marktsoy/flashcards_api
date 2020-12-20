package teststore_test

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestCardRepository_CreateAndFindByDeck(t *testing.T) {
	store := teststore.New()
	cards := []*models.Card{
		{Question: "Is it ok 1?", DeckID: "DECK_ID", Answer: "I guess so"},
		{Question: "Is it ok 2?", DeckID: "DECK_ID", Answer: "I guess So so"},
		{Question: "Is it ok 3?", DeckID: "DECK_ID", Answer: "I guess So so so"},
		{Question: "Is it ok 1?", DeckID: "DECK_ID_1", Answer: "I guess so"},
		{Question: "Is it ok 2?", DeckID: "DECK_ID_1", Answer: "I guess So so"},
	}
	for _, m := range cards {
		err := store.Card().Create(m)
		assert.NoError(t, err)
		assert.NotEmpty(t, m.ID)
	}
	var rows []*models.Card
	rows, _ = store.Card().FindAllByDeck(&models.Deck{ID: "DECK_ID"})
	assert.Equal(t, 3, len(rows))
	rows, _ = store.Card().FindAllByDeck(&models.Deck{ID: "DECK_ID_1"})
	assert.Equal(t, 2, len(rows))

}
