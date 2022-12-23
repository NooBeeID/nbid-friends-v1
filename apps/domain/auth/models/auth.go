package models

type Auth struct {
	Id       int // will be auto generated
	Email    string
	Password string
	ImgUrl   string
}

type SearchAuth struct {
	Id          int
	Email       string
	ImgUrl      string
	FollowingId *int
	AuthId      *int
}
