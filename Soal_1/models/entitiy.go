package models

import (
	"github.com/go-playground/validator/v10"
)

type UserData struct {
	ID       uint
	Name     string `gorm:"not null" validate:"required"`
	Email    string `gorm:"not null;unique" validate:"required,email"`
	Password string `gorm:"not null" validate:"required"`
}

func (UserData) TableName() string {
	return "users"
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *UserData) Validate() error {
	return validate.Struct(u)
}
