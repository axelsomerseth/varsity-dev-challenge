package models

type Users struct {
	Model
	Username string `validate:"required,max=20"`
	Password string `validate:"required,max=255"`
	Email    string `validate:"required,max=255"`
}
