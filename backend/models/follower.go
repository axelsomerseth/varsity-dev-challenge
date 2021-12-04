package models

type Followers struct {
	UserID      uint `validate:"required"`
	FollowingID uint `validate:"required"`
}
