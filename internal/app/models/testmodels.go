package models

import "testing"

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
