package teststore_test

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/store"
	"github.com/marktsoy/flashcards_api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestTeststore_ImplementsStore(t *testing.T) {

	assert.Implements(t, (*store.Store)(nil), &teststore.Teststore{})
}
