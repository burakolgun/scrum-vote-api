package models

import "github.com/go-ozzo/ozzo-validation"

type Company struct {
	UUID        string
	OwnerUUID   string
	Name        string
	Description string
}

func (m Company) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
		validation.Field(&m.Description, validation.Required, validation.Length(0, 120)),
	)
}
