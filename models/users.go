package models

import "github.com/go-ozzo/ozzo-validation"

type User struct {
	UUID     string
	username string
	password string
	email    string
}

type identity interface {
	GetID() string
	GetName() string
}

func (m User) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.username, validation.Required, validation.Length(0, 120)),
		validation.Field(&m.email, validation.Required),
	)
}

func (m User) GetID() string {
	return m.UUID
}

func (m User) GetName() string {
	return m.username
}
