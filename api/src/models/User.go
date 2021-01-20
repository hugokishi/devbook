package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}
	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

// validate - Validate user
func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("The name is required")
	}
	if user.Nick == "" {
		return errors.New("The nickname is required")
	}
	if user.Email == "" {
		return errors.New("The email is required")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email is in an invalid format")
	}
	if step == "register" && user.Password == "" {
		return errors.New("The password is required")
	}
	return nil
}

// Format - Format user (remove spaces)
func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}
