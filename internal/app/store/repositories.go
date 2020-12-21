package store

import "github.com/marktsoy/flashcards_api/internal/app/models"

// UserRepository base
type UserRepository interface {
	Create(u *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindAll() ([]*models.User, error)
}

// DeckRepository base
type DeckRepository interface {
	Create(*models.Deck) error
	FindByID(id string) (*models.Deck, error)
	Update(*models.Deck) error
	Delete(*models.Deck) error
}

// CardRepository ...
type CardRepository interface {
	Create(*models.Card) error
	FindAllByDeck(*models.Deck) ([]*models.Card, error)
	FindByID(id int) (*models.Card, error)
	Update(*models.Card) error
	Delete(*models.Card) error
}
