package models

// User Model
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"-"`
	EncryptedPassword string `json:"-"`
}

// Creating ...
func (u *User) Creating() error {
	s, err := encryptPassword(u.Password)
	if err != nil {
		return err
	}
	u.EncryptedPassword = s
	return nil
}
