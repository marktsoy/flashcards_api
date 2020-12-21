package teststore

import (
	"errors"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
)

// CardRepository ...
type CardRepository struct {
	cards map[int]*models.Card
}

// Create ...
func (rep *CardRepository) Create(card *models.Card) error {
	card.Creating()
	card.ID = len(rep.cards) + 1
	rep.cards[card.ID] = card
	return nil
}

// FindAllByDeck ...
func (rep *CardRepository) FindAllByDeck(deck *models.Deck) ([]*models.Card, error) {
	records := make([]*models.Card, 0)
	for _, card := range rep.cards {
		if card.DeckID != deck.ID {
			continue
		}
		records = append(records, card)
	}
	return records, nil
}

// FindByID ...
func (rep *CardRepository) FindByID(id int) (*models.Card, error) {
	card, ok := rep.cards[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return card, nil
}

// Update ...
func (rep *CardRepository) Update(card *models.Card) error {
	_, ok := rep.cards[card.ID]
	if !ok {
		return store.ErrRecordNotFound
	}
	rep.cards[card.ID] = card
	return nil
}

// Delete ...
func (rep *CardRepository) Delete(card *models.Card) error {
	_, ok := rep.cards[card.ID]
	if !ok {
		return errors.New("Nothing was deleted")
	}
	delete(rep.cards, card.ID)
	return nil
}
