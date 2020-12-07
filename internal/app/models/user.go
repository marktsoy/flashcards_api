package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User Model
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

func (u *User) Creating() (*User, error) {
	s, err := encryptPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.EncryptedPassword = s
	return u, nil
}

func encryptPassword(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
