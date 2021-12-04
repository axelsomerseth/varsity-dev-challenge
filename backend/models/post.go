package models

import "time"

// Post is a microbloggin post
type Post struct {
	Model
	UserID    uint   `validate:"required"`
	Post      string `validate:"required,max=140"` // tweet
	CreatedAt time.Time
}
