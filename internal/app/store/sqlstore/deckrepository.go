package sqlstore

import (
	"database/sql"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
)

// DeckRepository ...
type DeckRepository struct {
	store *SQLStore
}

// Create ...
func (r *DeckRepository) Create(d *models.Deck) error {
	db := r.store.db

	d.Creating()
	err := db.QueryRow(
		"INSERT INTO decks(id,name,user_id) VALUES ($1,$2,$3) RETURNING id",
		d.ID, d.Name, d.UserID,
	).Scan(&d.ID)

	if err != nil {
		return err
	}
	return nil
}

// FindByID ...
func (r *DeckRepository) FindByID(id string) (*models.Deck, error) {
	db := r.store.db

	d := &models.Deck{}
	err := db.QueryRow("SELECT id,name,user_id FROM decks WHERE id=$1", id).Scan(&d.ID, &d.Name, &d.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return d, nil
}
