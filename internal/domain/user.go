package domain

import (
	"errors"

	"github.com/go-playground/validator"
)

var (
	validate        *validator.Validate
	ErrUserNotFound = errors.New("user with such credentials not found")
)

func init() {
	validate = validator.New()
}

type User struct {
	Login string `json:"name" validate:"required,gte=2"`
	Age   int    `json:"age" validate:"required,gte=1"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
