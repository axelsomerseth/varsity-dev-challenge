package models

import "net/url"

type User struct {
	Model
	Username string `validate:"required,max=20"`
	// Password string `validate:"required,max=255"`
	Email   string  `validate:"required,max=255"`
	Subject string  `validate:"required"`
	Picture url.URL `validate:"url"`
}

func (User) TableName() string {
	return "USER"
}
