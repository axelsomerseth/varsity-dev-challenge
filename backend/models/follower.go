package models

type Follower struct {
	UserID     uint `validate:"required"`
	FollowerID uint `validate:"required"`
}
