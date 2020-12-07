package models

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// Md5 ...
func Md5(items ...string) string {
	h := md5.New()
	for _, s := range items {
		io.WriteString(h, s)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func encryptPassword(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
