package domain

type UserID uint

type User struct {
	UserID UserID
}

func (u *User) Validate() error {
	return nil
}
