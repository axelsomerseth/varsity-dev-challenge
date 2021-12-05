package models

// Model is the base model definition.
type Model struct {
	ID uint `json:"id" gorm:"primarykey"`
}
