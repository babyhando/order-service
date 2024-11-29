package domain

import (
	"errors"
	"time"
)

type (
	UserID uint
	Phone  string
)

func (p Phone) IsValid() bool {
	// todo regex
	return true
}

type User struct {
	ID        UserID
	CreatedAt time.Time
	DeletedAt time.Time
	FirstName string
	LastName  string
	Phone     Phone
}

func (u *User) Validate() error {
	if !u.Phone.IsValid() {
		return errors.New("phone is not valid")
	}
	return nil
}
