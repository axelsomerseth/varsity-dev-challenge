package models

type User struct {
	Model
	Username string `json:"username" validate:"required,max=20"`
	Email    string `json:"email" validate:"required,max=255"`
	Subject  string `json:"subject" validate:"required"`
	Picture  string `json:"picture" validate:"url"`
}

func (User) TableName() string {
	return "USER"
}
