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
}
