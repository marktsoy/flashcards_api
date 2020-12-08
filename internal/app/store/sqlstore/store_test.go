package sqlstore_test

import (
	"os"
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/store"
	"github.com/marktsoy/flashcards_api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

var (
	databaseURL string
)

func TestMain(t *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=flashcards_test sslmode=disable"
	}

	os.Exit(t.Run())
}

func TestSQLStore_ImplementsStore(t *testing.T) {
	assert.Implements(t, (*store.Store)(nil), &sqlstore.SQLStore{})

}
