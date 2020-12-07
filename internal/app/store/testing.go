package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper() //
	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tablenames ...string) {
		if len(tablenames) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE  %v  CASCADE", strings.Join(tablenames, ", "))); err != nil {
				t.Fatal(err)
			}
			fmt.Printf("Cleared %v", strings.Join(tablenames, ", "))
		}
		s.Close()
	}
}
