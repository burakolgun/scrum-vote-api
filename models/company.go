package models

import "github.com/go-ozzo/ozzo-validation"

type Company struct {
	uuid        string
	ownerUUID   string
	name        string
	description string
}

func (m Company) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.name, validation.Required, validation.Length(0, 120)),
		validation.Field(&m.description, validation.Required, validation.Length(0, 120)),
	)
}
