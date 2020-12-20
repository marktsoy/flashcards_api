package teststore

import "github.com/marktsoy/flashcards_api/internal/app/models"

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
