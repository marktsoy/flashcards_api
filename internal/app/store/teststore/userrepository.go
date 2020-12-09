package teststore

import (
	"errors"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
)

// UserRepository implementation for teststore package
type UserRepository struct {
	users map[string]*models.User
}

// Create Implementation
func (r *UserRepository) Create(u *models.User) error {
	if err := u.Creating(); err != nil {
		return err
	}
	if _, ok := r.users[u.Email]; ok {
		return errors.New("User exists")
	}
	r.users[u.Email] = u
	u.ID = len(r.users)
	return nil
}

// FindByEmail implementation
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u, found := r.users[email]
	if !found {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
