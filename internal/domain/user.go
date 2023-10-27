package domain

import (
	"github.com/go-playground/validator"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

type User struct {
	Login string `json:"name" validate:"required,gte=2"`
	Age   int    `json:"age" validate:"required,gte=1"`
}

type UserOutput struct {
	ID int `json:"id"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
