package models

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func encryptPassword(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Create new math/rand.Rand struct that will generate random int
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomString generates random string of length passed
// Uses charset and seedRand
func RandomString(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
