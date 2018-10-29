package models

import (
	"../database"
	"fmt"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/gommon/log"
)

type User struct {
	UUID     string `json:"uuid" form:"-"`
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
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

func (m User) Store() (User, error) {
	DB, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()
	err = DB.Ping()

	if err != nil {
		return User{}, err
	}
	fmt.Println("Connection Opened")
	fmt.Println("1")

	return User{
		"1",
		"2",
		"3",
		"4",
	}, nil
}
