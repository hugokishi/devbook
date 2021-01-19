package models

import (
	"errors"
	"strings"
	"time"
)

// User - Model for user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare - Call methods to format and validate recieved users
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()
	return nil
}

// validate - Validate user
func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("The name is required")
	}
	if user.Nick == "" {
		return errors.New("The nickname is required")
	}
	if user.Email == "" {
		return errors.New("The email is required")
	}
	if user.Password == "" {
		return errors.New("The password is required")
	}
	return nil
}

// Format - Format user (remove spaces)
func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
