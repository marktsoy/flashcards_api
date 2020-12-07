package models

// User Model
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
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
