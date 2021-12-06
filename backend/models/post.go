package models

import "time"

// Post is a microbloggin post
type Post struct {
	Model
	UserID    uint      `json:"userID" validate:"required"`
	Post      string    `json:"post" validate:"required,max=140"` // tweet
	CreatedAt time.Time `json:"createdAt"`
}

func (Post) TableName() string {
	return "POST"
}
