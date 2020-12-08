package sqlstore

import (
	"database/sql"

	"github.com/marktsoy/flashcards_api/internal/app/store"
)

// SQLStore ...
type SQLStore struct {
	db             *sql.DB
	userRepository *UserRepository
	deckRepository *DeckRepository
}

// New ...
func New(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

/**
* Declare Repositories
**/

// User ...
func (s *SQLStore) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}
	return s.userRepository
}

// Deck ...
func (s *SQLStore) Deck() store.DeckRepository {
	if s.deckRepository == nil {
		s.deckRepository = &DeckRepository{
			store: s,
		}
	}
	return s.deckRepository
}
