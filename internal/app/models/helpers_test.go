package models_test

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestHelper_RandomString(t *testing.T) {

	var set = make(map[string]int)

	for i := 0; i < 100; i++ {
		set[models.RandomString(32)]++
	}

	for _, n := range set {
		assert.Equal(t, 1, n)
	}
}
