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

// Update ...
func (r *DeckRepository) Update(d *models.Deck) error {
	db := r.store.db

	sqlStatement := "UPDATE decks SET name=$2 where decks.id = $1 "
	res, err := db.Exec(sqlStatement, d.ID, d.Name)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return store.ErrRecordNotFound
	}
	return nil
}

// Delete ...
func (r *DeckRepository) Delete(d *models.Deck) error {
	db := r.store.db
	if d.ID == "" {
		return store.ErrRecordNotFound
	}
	sqlStatement := "DELETE FROM decks WHERE decks.id = $1 "
	res, err := db.Exec(sqlStatement, d.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return store.ErrRecordNotFound
	}
	return nil
}
