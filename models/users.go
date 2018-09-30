package models

import "github.com/go-ozzo/ozzo-validation"

type User struct {
	UUID     string
	UserName string
	Password string
	Email    string
}

type identity interface {
	GetID() string
	GetName() string
}

func (m User) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.UserName, validation.Required, validation.Length(0, 120)),
		validation.Field(&m.Email, validation.Required),
	)
}

func (m User) GetID() string {
	return m.UUID
}

func (m User) GetName() string {
	return m.UserName
}
