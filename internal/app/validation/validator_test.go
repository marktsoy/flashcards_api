package validation_test

import (
	"testing"

	"github.com/marktsoy/flashcards_api/internal/app/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidation_MinLength(t *testing.T) {
	type testcase struct {
		name         string
		value        string
		minlen       int
		excpected    bool
		excpectedMsg string
	}
	tc := []testcase{
		{
			name:         "firstline",
			value:        "my length is ok",
			minlen:       5,
			excpected:    true,
			excpectedMsg: "",
		},
		{
			name:         "firstline",
			value:        "my length is not ok",
			minlen:       100,
			excpected:    false,
			excpectedMsg: "firstline is too short, accepted length: 100",
		},
	}

	for _, c := range tc {
		passed, errMsg := validation.MinLen(c.name, c.value, c.minlen)
		assert.Equal(t, c.excpected, passed)
		assert.Equal(t, c.excpectedMsg, errMsg)
	}

	for _, c := range tc {
		passed, errMsg := validation.MinLen(c.name, c.value, c.minlen, "optional error")
		assert.Equal(t, c.excpected, passed)
		if !passed {
			assert.Equal(t, "optional error", errMsg)
		}
	}
}
