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
