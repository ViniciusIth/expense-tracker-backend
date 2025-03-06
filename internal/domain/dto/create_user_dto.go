package dto

import "errors"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ErrNameIsRequired     = errors.New("name is required")
	ErrEmailIsRequired    = errors.New("email is required")
	ErrPasswordIsRequired = errors.New("password is required")
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters")
)

func (req *CreateUserRequest) Validate() error {
	if req.Name == "" {
		return ErrNameIsRequired
	}
	if req.Email == "" {
		return ErrEmailIsRequired
	}
	if req.Password == "" {
		return ErrPasswordIsRequired
	}
	if len(req.Password) < 8 {
		return ErrPasswordTooShort
	}
	return nil
}
