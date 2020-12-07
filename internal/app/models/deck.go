package models

import (
	"strconv"
	"time"
)

// Deck ...
type Deck struct {
	ID     string
	UserID int
	Name   string

	user *User
}

// Creating - Before creating hook
func (d *Deck) Creating() {
	if d.user != nil && d.UserID == 0 {
		d.UserID = d.user.ID
	}
	if d.ID == "" {
		d.ID = Md5(time.Now().String(), d.Name, strconv.Itoa(d.UserID))
	}
}

// BindUser to Deck
func (d *Deck) BindUser(user *User) {
	d.user = user
}
