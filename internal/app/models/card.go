package models

import "time"

// Card ...
type Card struct {
	ID        int       `json:"id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
	DeckID    string    `json:"deck_id"`
}

func (c *Card) Creating() {
	c.CreatedAt = time.Now()
}
