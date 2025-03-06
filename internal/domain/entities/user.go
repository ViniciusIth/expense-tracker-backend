package entities

import (
	"errors"
	"net/mail"
)

type User struct {
	ID    string
	Name  string
	Email string
}

var (
    ErrIDIsRequired         = errors.New("id is required")
	ErrNameIsRequired       = errors.New("name is required")
	ErrEmailIsRequired      = errors.New("email is required")
	ErrEmailFormatIsInvalid = errors.New("email format is invalid")
)

func (user *User) Validate() error {
    if user.ID == "" {
        return ErrIDIsRequired
    }
	if user.Name == "" {
		return ErrNameIsRequired
	}
	if user.Email == "" {
		return ErrEmailIsRequired
	}
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return ErrEmailFormatIsInvalid
	}
	return nil
}
