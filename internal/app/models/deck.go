package models

// Deck ...
type Deck struct {
	ID     string `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`

	user *User
}

// Creating - Before creating hook
func (d *Deck) Creating() {
	if d.user != nil && d.UserID == 0 {
		d.UserID = d.user.ID
	}
	if d.ID == "" {
		d.ID = RandomString(32)
	}
}

// BindUser to Deck
func (d *Deck) BindUser(user *User) {
	d.user = user
}
