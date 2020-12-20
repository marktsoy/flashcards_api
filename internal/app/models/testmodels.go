package models

import (
	"strconv"
	"testing"
)

// TestUser ..
func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Email:    "example@test.com",
		Password: "some password",
	}
}

// TestDeck ..
func TestDeck(t *testing.T) *Deck {
	t.Helper()
	return &Deck{
		Name: "First collection",
	}
}

// TestCard ...
func TestCard(t *testing.T, i int) *Card {
	t.Helper()
	return &Card{
		Question: "Question " + strconv.Itoa(i),
		Answer:   "Question " + strconv.Itoa(i),
	}
}
