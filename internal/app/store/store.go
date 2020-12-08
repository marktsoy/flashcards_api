package store

// Store Interface
type Store interface {
	User() UserRepository
	Deck() DeckRepository
}
