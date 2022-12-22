package models

import "time"

type Follow struct {
	Id          int
	AuthId      int
	FollowingId int
	CreatedAt   time.Time
}

type FollowWithAuth struct {
	Id    int
	Email string
}
