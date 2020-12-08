package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	_ "github.com/lib/pq" // Include postgress driver
)

// TestStore ...
func TestStore(t *testing.T, databaseURL string) (*SQLStore, func(...string)) {
	t.Helper() //
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	s := New(db)

	return s, func(tablenames ...string) {
		if len(tablenames) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE  %v  CASCADE", strings.Join(tablenames, ", "))); err != nil {
				t.Fatal(err)
			}
			fmt.Printf("Cleared %v", strings.Join(tablenames, ", "))
		}
		s.db.Close()
	}
}
