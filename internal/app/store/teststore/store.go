package teststore

/**
* teststore is mock store to be used in api testing
**/

import (
	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
)

// Teststore ...
type Teststore struct {
	userRepository *UserRepository
	deckRepository *DeckRepository
	cardRepository *CardRepository
}

// New Teststore ...
func New() *Teststore {
	return &Teststore{}
}

// User implementation
func (t *Teststore) User() store.UserRepository {
	if t.userRepository == nil {
		t.userRepository = &UserRepository{
			users: make(map[string]*models.User),
		}
	}
	return t.userRepository
}

// Deck ...
func (t *Teststore) Deck() store.DeckRepository {
	if t.deckRepository == nil {
		t.deckRepository = &DeckRepository{
			decks: make(map[string]*models.Deck),
		}
	}
	return t.deckRepository
}

// Card ...
func (t *Teststore) Card() store.CardRepository {
	if t.cardRepository == nil {
		t.cardRepository = &CardRepository{
			cards: make(map[int]*models.Card, 0),
		}
	}
	return t.cardRepository
}
