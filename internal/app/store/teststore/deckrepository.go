package teststore

import (
	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
)

// DeckRepository ...
type DeckRepository struct {
	decks map[string]*models.Deck
}

// Create ...
func (r *DeckRepository) Create(d *models.Deck) error {
	d.Creating()
	r.decks[d.ID] = d
	return nil
}

// FindByID ...
func (r *DeckRepository) FindByID(id string) (*models.Deck, error) {
	d, ok := r.decks[id]

	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return d, nil
}

// Update ...
func (r *DeckRepository) Update(d *models.Deck) error {
	if _, ok := r.decks[d.ID]; !ok {
		return store.ErrRecordNotFound
	}
	r.decks[d.ID] = d
	return nil
}

// Delete ...
func (r *DeckRepository) Delete(d *models.Deck) error {
	if _, ok := r.decks[d.ID]; !ok {
		return store.ErrRecordNotFound
	}
	delete(r.decks, d.ID)
	return nil
}
