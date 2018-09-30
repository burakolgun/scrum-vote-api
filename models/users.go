package models

import "github.com/go-ozzo/ozzo-validation"

type User struct {
	uuid     string
	username string
	password string
	email    string
}

func (m User) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.username, validation.Required, validation.Length(0, 120)),
		validation.Field(&m.email, validation.Required),
	)
}