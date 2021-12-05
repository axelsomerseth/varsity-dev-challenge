package models

type Follower struct {
	UserID     uint `validate:"required"`
	FollowerID uint `validate:"required"`
}

func (Follower) TableName() string {
	return "FOLLOWER"
}
