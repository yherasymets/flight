package models

type Article struct {
	Id                     uint16
	Title, Anons, FullText string
	UserId                 uint16
	User                   User
}
