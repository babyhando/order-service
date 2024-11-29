package domain

import "time"

type (
	UserID uint
	Phone  string
)

func PhoneFromString(p string) (Phone, error) {
	// todo : phone regex check
	return Phone(p), nil
}

type User struct {
	ID        UserID
	CreatedAt time.Time
	DeletedAt time.Time
	FirstName string
	LastName  string
	Phone     Phone
}
