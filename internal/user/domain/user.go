package domain

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/babyhando/order-service/pkg/conv"
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
	Password  string
	Phone     Phone
}

func (u *User) Validate() error {
	if !u.Phone.IsValid() {
		return errors.New("phone is not valid")
	}
	return nil
}

func (u *User) PasswordIsCorrect(pass string) bool {
	return NewPassword(pass) == u.Password
}

func NewPassword(pass string) string {
	h := sha256.New()
	h.Write(conv.ToBytes(pass))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

type UserFilter struct {
	ID    UserID
	Phone string
}

func (f *UserFilter) IsValid() bool {
	f.Phone = strings.TrimSpace(f.Phone)
	return f.ID > 0 || len(f.Phone) > 0
}
