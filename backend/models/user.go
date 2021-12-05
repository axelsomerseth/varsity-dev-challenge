package models

type User struct {
	Model
	Username string `validate:"required,max=20"`
	Email    string `validate:"required,max=255"`
	Subject  string `validate:"required"`
	Picture  string `validate:"url"`
}

func (User) TableName() string {
	return "USER"
}
