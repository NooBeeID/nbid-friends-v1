package models

type Auth struct {
	Id       int // will be auto generated
	Email    string
	Password string
	ImgUrl   string
}
