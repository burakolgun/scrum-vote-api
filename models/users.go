package models

import "github.com/go-ozzo/ozzo-validation"

type User struct {
	UUID     string `json:"uuid" form:"-"`
	UserName string `json:"username" form:"username"`
	Password string	`json:"password" form:"password"`
	Email    string	`json:"email" form:"email"`
}

type identity interface {
	GetID() string
	GetName() string
}

func (m User) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.UserName, validation.Required, validation.Length(0, 120)),
	)
}

func (m User) GetID() string {
	return m.UUID
}

func (m User) GetName() string {
	return m.UserName
}
