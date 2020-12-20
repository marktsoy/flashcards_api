package sqlstore

import (
	"github.com/marktsoy/flashcards_api/internal/app/models"
)

// CardRepository ...
type CardRepository struct {
	store *SQLStore
}

// Create ...
func (r *CardRepository) Create(c *models.Card) error {
	db := r.store.db

	c.Creating()
	err := db.QueryRow(
		"INSERT INTO cards(question,answer,created_at,deck_id) VALUES ($1,$2,$3,$4) RETURNING id",
		c.Question, c.Answer, c.CreatedAt, c.DeckID,
	).Scan(&c.ID)

	if err != nil {
		return err
	}
	return nil
}

// FindAllByDeck ...
func (r *CardRepository) FindAllByDeck(d *models.Deck) ([]*models.Card, error) {
	db := r.store.db
	records := make([]*models.Card, 0)
	rows, err := db.Query("SELECT id,question,answer,deck_id,created_at FROM cards  WHERE deck_id=$1", d.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := &models.Card{}
		rows.Scan(
			&c.ID, &c.Question, &c.Answer, &c.DeckID, &c.CreatedAt,
		)
		records = append(records, c)
	}
	return records, nil
}
