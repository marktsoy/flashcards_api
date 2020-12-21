package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/marktsoy/flashcards_api/internal/app/store"
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

// FindByID ...
func (r *CardRepository) FindByID(id int) (*models.Card, error) {
	db := r.store.db
	card := &models.Card{
		Deck: &models.Deck{},
	}

	err := db.QueryRow(`
		SELECT 
			c.id, c.question, c.answer,c.deck_id,c.created_at , decks.id, decks.user_id, decks.name 
		FROM cards c
		INNER JOIN decks
			ON decks.id=c.deck_id 
		WHERE c.id=$1 `,
		id,
	).Scan(&card.ID, &card.Question, &card.Answer, &card.DeckID, &card.CreatedAt, &card.Deck.ID, &card.Deck.UserID, &card.Deck.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return card, nil
}

// Update ...
func (r *CardRepository) Update(card *models.Card) error {
	db := r.store.db

	sqlStatement := "UPDATE cards SET question=$2 , answer=$3 where decks.id = $1 "
	res, err := db.Exec(sqlStatement, card.ID, card.Question, card.Answer)
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

// Delete ....
func (r *CardRepository) Delete(card *models.Card) error {
	db := r.store.db

	sqlStatement := "DELETE FROM cards WHERE cards.id = $1"
	res, err := db.Exec(sqlStatement, card.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Nothing was deleted")
	}
	return nil
}
