package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email			string    `json:"email,omitempty"`
	Username  string    `json:"userName,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) ValidateFields() error {

	if user.Name == "" && user.Password == "" && user.Username == ""  && user.Email == "" {
		return errors.New("all fields are required")
	}

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	if user.Username == "" {
		return errors.New("username is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	user.cleanWhiteSpaces()

	return nil

}

func (user *User) cleanWhiteSpaces() {

	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)

}
